package integration

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNetlogger_FullCycle(t *testing.T) {
	targets := []string{"127.0.0.1", "google.com:443", "localhost:22", "invalid.invalid"}
	args := append([]string{"run", "./cmd/netlogger", "--for", "2s"}, targets...)
	cmd := exec.Command("go", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("netlogger run failed: %v\nOutput: %s", err, string(output))
	}

	// Check that a report is generated for each target
	for _, tgt := range targets {
		fname := "report_" + strings.ReplaceAll(strings.ReplaceAll(tgt, ":", "_"), "/", "_") + ".md"
		if _, err := os.Stat(fname); err != nil {
			t.Errorf("expected report file %s to be created", fname)
		} else {
			// Check report content for expected table header
			data, _ := os.ReadFile(fname)
			if !strings.Contains(string(data), "| Target | Type | Timestamp | Datetime | Status |") {
				t.Errorf("report %s missing table header", fname)
			}
			// Clean up
			_ = os.Remove(fname)
		}
	}

	// Check output for summary lines
	for _, tgt := range targets {
		fname := "report_" + strings.ReplaceAll(strings.ReplaceAll(tgt, ":", "_"), "/", "_") + ".md"
		if !strings.Contains(string(output), fname) {
			t.Errorf("expected output to mention %s", fname)
		}
	}
}
