package integration

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestNetlogger_Interrupt(t *testing.T) {
	// Start the process
	cmd := exec.Command("go", "run", "./cmd/netlogger", "--for", "10s", "127.0.0.1")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start netlogger: %v", err)
	}
	// Give it a moment to start
	time.Sleep(1 * time.Second)
	// Send interrupt
	_ = cmd.Process.Signal(os.Interrupt)
	// Wait for it to exit
	err := cmd.Wait()
	if err != nil && !strings.Contains(err.Error(), "signal: interrupt") {
		t.Errorf("expected interrupt exit, got: %v", err)
	}
	// Check report file
	fname := "report_127.0.0.1.md"
	if _, err := os.Stat(fname); err != nil {
		t.Errorf("expected report file %s to be created after interrupt", fname)
	} else {
		_ = os.Remove(fname)
	}
}
