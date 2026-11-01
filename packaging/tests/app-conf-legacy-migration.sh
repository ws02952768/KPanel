#!/usr/bin/env bash
set -euo pipefail

# Host-side harness; every host-filesystem mutation by the tests stays inside
# disposable containers. The fixture is the exact published RC3 updater.
cd "${1:?repository root required}"
test "$(uname -s)" = Linux
test "$(uname -m)" = x86_64
mkdir -p dist/test-fixtures
git show fd7033717f6f47d66277da760e7d255bd85d3c25:packaging/kejilion-app/kpanel.conf \
  > dist/test-fixtures/rc3-kpanel.conf
printf '%s\n' '6516e8173e6c8688529b56cf049d47b0288b6c229251d64251bd1c8bdff2670d  dist/test-fixtures/rc3-kpanel.conf' | sha256sum -c -
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath \
  -o dist/test-fixtures/migration-agent ./cmd/kejilion-agent
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -c \
  -o dist/test-fixtures/migration.test ./cmd/kejilion-agent
image=bash:5.2.37-alpine3.22@sha256:3bee76a96d86d5d2d5efc7c1c570e5a7c95db22348a26944e0e546fa174e3324
docker run --rm -e KPANEL_MIGRATION_ROOTFS_TEST=1 -e KPANEL_TEST_SOURCE_ROOT=/src \
  -v "$PWD:/src:ro" "$image" bash -ec \
  'ln -s /usr/local/bin/bash /bin/bash; exec /src/dist/test-fixtures/migration.test -test.run "^TestSharedScriptMigrationLegacyRootfs$" -test.v'
docker run --rm -e KPANEL_APP_CONF_TEST_ROOTFS=1 -e KPANEL_REQUIRE_LEGACY_MIGRATION=1 \
  -v "$PWD:/src:ro" "$image" bash /src/packaging/tests/app-conf-lifecycle.sh /src
