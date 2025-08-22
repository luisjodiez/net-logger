package core

import "testing"

type DummyPinger struct{
	Resp string
}

func (d *DummyPinger) Ping(target string) string {
	return d.Resp
}

func TestSafeResults_AddAndAll(t *testing.T) {
	r := &SafeResults{}
	res := Result{"host", "icmp", 123, "2025-01-01T00:00:00Z", "ok"}
	r.Add(res)
	all := r.All()
	if len(all) != 1 || all[0] != res {
		t.Errorf("expected one result, got %v", all)
	}
}

func TestDummyPinger(t *testing.T) {
	p := &DummyPinger{"ok"}
	if p.Ping("host") != "ok" {
		t.Error("expected ok")
	}
}
