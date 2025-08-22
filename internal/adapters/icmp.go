package adapters

import (
	"github.com/prometheus-community/pro-bing"
)

type ICMPPinger struct{}

func (p *ICMPPinger) Ping(target string) string {
	pinger, err := probing.NewPinger(target)
	if err != nil {
		return "ko"
	}
	pinger.Count = 1
	pinger.Timeout = 2 * 1e9 // 2s
	pinger.SetPrivileged(false)
	err = pinger.Run()
	if err != nil {
		return "ko"
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		return "ok"
	}
	return "ko"
}
