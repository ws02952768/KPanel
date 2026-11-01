package sites

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// Exercise production executable admission as well as the runtime commands.
func TestProxyRuntimeUsesTrustedDockerExecutable(t *testing.T) {
	if os.Geteuid() != 0 {
		t.Skip("trusted host executables require root ownership")
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "docker")
	if err := os.WriteFile(path, []byte("#!/bin/sh\ncase \"$*\" in\n'inspect --format {{.State.Running}} nginx') echo true;;\n'exec nginx nginx -t') exit 0;;\n*) exit 99;;\nesac\n"), 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	if err := verifyProxyRuntime(context.Background(), systemRecipeJobRunner{}); err != nil {
		t.Fatal(err)
	}
}
