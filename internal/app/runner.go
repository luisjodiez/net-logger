package app

import (
	"context"
	"net"
	"strings"
	"time"
	"net-logger/internal/core"
)

type Runner struct {
	PingerICMP core.Pinger
	PingerTCP core.Pinger
	Results core.ResultsRepository
}

func (r *Runner) ProbeLoop(ctx context.Context, target string, connType string) {
	host := target
	port := ""
	if h, p, err := net.SplitHostPort(target); err == nil && p != "" {
		host = h
		port = p
	}
	interval := time.Second
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var status string
			if connType == "icmp" {
				status = r.PingerICMP.Ping(host)
			} else {
				targetAddr := net.JoinHostPort(host, port)
				status = r.PingerTCP.Ping(targetAddr)
			}
			now := time.Now()
			r.Results.Add(core.Result{
				Target: target,
				ConnType: connType,
				Timestamp: now.Unix(),
				Datetime: now.Format(time.RFC3339),
				Status: status,
			})
			time.Sleep(interval)
		}
	}
}

func DetectConnType(target string) string {
	if _, _, err := net.SplitHostPort(target); err == nil {
		return "tcp"
	}
	if strings.Contains(target, ":") {
		return "tcp"
	}
	return "icmp"
}
