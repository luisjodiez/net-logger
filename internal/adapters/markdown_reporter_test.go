package adapters

import (
	"os"
	"testing"
	"net-logger/internal/core"
)

func TestMarkdownReporter_Report(t *testing.T) {
	reporter := &MarkdownReporter{}
	results := []core.Result{
		{Target: "test", ConnType: "icmp", Timestamp: 1, Datetime: "2025-01-01T00:00:00Z", Status: "ok"},
	}
	err := reporter.Report(results)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, err := os.Stat("report_test.md"); err != nil {
		t.Errorf("expected report file to be created, got %v", err)
	}
	_ = os.Remove("report_test.md")
}
