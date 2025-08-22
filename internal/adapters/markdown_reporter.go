package adapters

import (
	"fmt"
	"os"
	"strings"
	"net-logger/internal/core"
)

type MarkdownReporter struct{}

func (r *MarkdownReporter) Report(results []core.Result) error {
	targetMap := make(map[string][]core.Result)
	for _, res := range results {
		targetMap[res.Target] = append(targetMap[res.Target], res)
	}
	for target, res := range targetMap {
		reportFile := fmt.Sprintf("report_%s.md", sanitizeFilename(target))
		f, err := os.Create(reportFile)
		if err != nil {
			return err
		}
		total := len(res)
		oks := 0
		for _, r := range res {
			if r.Status == "ok" {
				oks++
			}
		}
		failures := total - oks
		first, last := "", ""
		if total > 0 {
			first = res[0].Datetime
			last = res[total-1].Datetime
		}
		fmt.Fprintf(f, "# Report for %s\n\n", target)
		fmt.Fprintf(f, "**Total probes:** %d  |  **OK:** %d  |  **KO:** %d  |  **First:** %s  |  **Last:** %s\n\n", total, oks, failures, first, last)
		fmt.Fprintln(f, "| Target | Type | Timestamp | Datetime | Status |")
		fmt.Fprintln(f, "|--------|------|-----------|----------|--------|")
		for _, r := range res {
			fmt.Fprintf(f, "| %s | %s | %d | %s | %s |\n", r.Target, r.ConnType, r.Timestamp, r.Datetime, r.Status)
		}
		f.Close()
	}
	return nil
}

func sanitizeFilename(s string) string {
	replacer := strings.NewReplacer(":", "_", "/", "_", "\\", "_")
	return replacer.Replace(s)
}
