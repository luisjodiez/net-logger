package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func TestNetloggerICMPAndTCP(t *testing.T) {
	cmd := exec.Command("go", "run", "./cmd/netlogger", "--for", "2s", "127.0.0.1", "google.com:443")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("netlogger run failed: %v\nOutput: %s", err, string(output))
	}
	if !strings.Contains(string(output), "Report written to report_127.0.0.1.md") {
		t.Error("ICMP report not generated")
	}
	if !strings.Contains(string(output), "Report written to report_google.com_443.md") {
		t.Error("TCP report not generated")
	}
}
