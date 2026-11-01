package sites

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type proxyRuntimeRunner struct {
	running               string
	inspectErr, configErr error
	calls                 []string
}

func (r *proxyRuntimeRunner) LookPath(name string) (string, error) { return name, nil }
func (r *proxyRuntimeRunner) Run(_ context.Context, name string, args ...string) ([]byte, error) {
	r.calls = append(r.calls, name+" "+strings.Join(args, " "))
	if args[0] == "inspect" {
		return []byte(r.running), r.inspectErr
	}
	return nil, r.configErr
}

func TestProxyRuntimeRejectsFilesWithoutRunningValidNginx(t *testing.T) {
	for _, tc := range []struct {
		name, running         string
		inspectErr, configErr error
		wantOK                bool
		calls                 int
	}{
		{name: "missing", inspectErr: errors.New("no container"), calls: 1},
		{name: "stopped", running: "false", calls: 1},
		{name: "invalid config", running: "true", configErr: errors.New("invalid config"), calls: 2},
		{name: "ready", running: "true\n", wantOK: true, calls: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := &proxyRuntimeRunner{running: tc.running, inspectErr: tc.inspectErr, configErr: tc.configErr}
			err := verifyProxyRuntime(context.Background(), r)
			if (err == nil) != tc.wantOK || len(r.calls) != tc.calls {
				t.Fatalf("err=%v calls=%v", err, r.calls)
			}
		})
	}
}

func TestProxyInvocationKeepsDirectAccessAndRequiresPairedScript(t *testing.T) {
	invocation := proxyInvocation("panel.example.com", "127.0.0.1", "8080")
	if !containsString(invocation.environment, "KJ_WEB_PROXY_KEEP_DIRECT=1") ||
		!containsString(invocation.required, `KPANEL_WEB_PROXY_ACCESS_PROTOCOL_VERSION="1"`) {
		t.Fatalf("unpaired proxy invocation: %+v", invocation)
	}
}
