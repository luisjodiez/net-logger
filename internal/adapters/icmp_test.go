package adapters

import (
	"testing"
)

type fakePinger struct{
	resp string
}

func (f *fakePinger) Ping(target string) string {
	return f.resp
}

func TestICMPPinger_Ping(t *testing.T) {
	p := &ICMPPinger{}
	// This test only checks that the method returns a string ("ok" or "ko")
	res := p.Ping("127.0.0.1")
	if res != "ok" && res != "ko" {
		t.Errorf("expected ok or ko, got %s", res)
	}
}
