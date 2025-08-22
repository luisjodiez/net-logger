package app

import (
	"context"
	"testing"
	"time"
	"net-logger/internal/core"
)

type dummyPinger struct{ resp string }
func (d *dummyPinger) Ping(target string) string { return d.resp }

type dummyResults struct{
	added []core.Result
}
func (d *dummyResults) Add(r core.Result) { d.added = append(d.added, r) }
func (d *dummyResults) All() []core.Result { return d.added }

func TestRunner_ProbeLoop_ICMP(t *testing.T) {
	dr := &dummyResults{}
	r := &Runner{
		PingerICMP: &dummyPinger{"ok"},
		PingerTCP: &dummyPinger{"ko"},
		Results: dr,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	r.ProbeLoop(ctx, "host", "icmp")
	if len(dr.added) == 0 {
		t.Error("expected at least one result")
	}
}

func TestRunner_ProbeLoop_TCP(t *testing.T) {
	dr := &dummyResults{}
	r := &Runner{
		PingerICMP: &dummyPinger{"ko"},
		PingerTCP: &dummyPinger{"ok"},
		Results: dr,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	r.ProbeLoop(ctx, "localhost:80", "tcp")
	if len(dr.added) == 0 {
		t.Error("expected at least one result")
	}
}

func TestDetectConnType(t *testing.T) {
	if DetectConnType("host:80") != "tcp" {
		t.Error("expected tcp")
	}
	if DetectConnType("host") != "icmp" {
		t.Error("expected icmp")
	}
}
