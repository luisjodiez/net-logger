package adapters

import (
	"testing"
)

func TestTCPPinger_Ping(t *testing.T) {
	p := &TCPPinger{}
	// This test only checks that the method returns a string ("ok" or "ko")
	res := p.Ping("localhost:80")
	if res != "ok" && res != "ko" {
		t.Errorf("expected ok or ko, got %s", res)
	}
}
