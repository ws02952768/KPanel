package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

var errSharedScriptMigrationBusy = errors.New("KPanel lifecycle is still running")

// Older installed updaters finish using their own lifecycle before installing
// the target lifecycle. Let that transaction release its lock before migrating;
// neither failed upgrades nor an Agent health probe may change the SSH script.
type sharedScriptMigration struct {
	ready    atomic.Bool
	execute  func(context.Context) error
	interval time.Duration
}

func newSharedScriptMigration(enabled bool) *sharedScriptMigration {
	m := &sharedScriptMigration{execute: migrateSharedScript, interval: time.Second}
	_, managedErr := os.Lstat("/home/docker/kpanel/.managed-by-kejilion-app")
	_, privateErr := os.Lstat("/home/docker/kpanel/bin/kejilion.sh")
	m.ready.Store(!enabled || errors.Is(managedErr, os.ErrNotExist) || errors.Is(privateErr, os.ErrNotExist))
	return m
}

func (m *sharedScriptMigration) Ready() bool { return m.ready.Load() }

func (m *sharedScriptMigration) Run(ctx context.Context) {
	for !m.Ready() {
		if ctx.Err() != nil {
			return
		}
		err := m.execute(ctx)
		if err == nil {
			m.ready.Store(true)
			return
		}
		if !errors.Is(err, errSharedScriptMigrationBusy) {
			if ctx.Err() == nil {
				slog.Error("shared host script migration failed; host mutations remain blocked", "error", err)
			}
			return
		}
		timer := time.NewTimer(m.interval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
		}
	}
}

func migrateSharedScript(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, "/bin/bash", "-c", sharedScriptMigrationCommand)
	command.Env = []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin", "LANG=C.UTF-8"}
	if err := command.Run(); err != nil {
		var exit *exec.ExitError
		if errors.As(err, &exit) && exit.ExitCode() == 75 {
			return errSharedScriptMigrationBusy
		}
		return fmt.Errorf("migrate packaged shared script: %w", err)
	}
	return nil
}

// Only the root-owned lifecycle extracted from the pinned image is sourced.
// No request, environment variable or remote source can select these paths.
const sharedScriptMigrationCommand = `set -eu
app=/home/docker/kpanel
private="$app/bin/kejilion.sh"
life="$app/bin/kpanel.conf"
[ "$(id -u)" = 0 ] || exit 77
[ -e "$private" ] || [ -L "$private" ] || exit 0
trusted_file() {
  [ -f "$1" ] && [ ! -L "$1" ] && [ "$(stat -c '%u' "$1")" = 0 ] || return 1
  local mode
  mode="$(stat -c '%a' "$1")" || return 1
  (( (8#$mode & 8#022) == 0 ))
}
[ -d /home ] && [ ! -L /home ] && [ "$(stat -c '%u' /home)" = 0 ] || exit 77
mode="$(stat -c '%a' /home)"
(( (8#$mode & 8#022) == 0 )) || exit 77
[ "$(stat -c '%u' /home/docker)" = 0 ] && [ ! -L "$app" ] || exit 77
if [ ! -L /home/docker ]; then
  mode="$(stat -c '%a' /home/docker)"
  (( (8#$mode & 8#022) == 0 )) || exit 77
fi
resolved="$(readlink -f "$app")" || exit 77
[ -d "$resolved" ] || exit 77
parent="$resolved/bin"
while :; do
  [ -d "$parent" ] && [ ! -L "$parent" ] && [ "$(stat -c '%u' "$parent")" = 0 ] || exit 77
  mode="$(stat -c '%a' "$parent")" || exit 77
  (( (8#$mode & 8#022) == 0 )) || exit 77
  [ "$parent" != / ] || break
  parent="$(dirname "$parent")"
done
trusted_file "$app/.managed-by-kejilion-app" && trusted_file "$life" && trusted_file "$private" || exit 77
lock=/run/lock/kejilion-kpanel-lifecycle.lock
for parent in /run /run/lock; do
  [ -d "$parent" ] && [ ! -L "$parent" ] && [ "$(stat -c '%u' "$parent")" = 0 ] || exit 77
  mode="$(stat -c '%a' "$parent")"
  if (( (8#$mode & 8#022) != 0 )); then
    [ "$parent" = /run/lock ] && (( (8#$mode & 8#1000) != 0 )) || exit 77
  fi
done
[ ! -L "$lock" ] || exit 77
if [ -e "$lock" ]; then trusted_file "$lock" || exit 77; fi
umask 077
exec 9>>"$lock"
[ -f /proc/$$/fd/9 ] && [ "$(stat -Lc '%u' /proc/$$/fd/9)" = 0 ] && [ ! -L "$lock" ] || exit 77
mode="$(stat -Lc '%a' /proc/$$/fd/9)"
(( (8#$mode & 8#022) == 0 )) || exit 77
flock -n 9 || exit 75
[ ! -e "$app/update-state/transaction" ] || exit 75
trusted_file "$life" && trusted_file "$private" || exit 77
grep -Fx '# KPanel shared script migration: v1' "$life" >/dev/null || exit 77
docker_app_plus() { :; }
entry() {
  . "$life"
  local KJ_KPANEL_LIFECYCLE_LOCK_HELD=1
  kpanel_migrate_legacy_shared_script
}
entry
`
