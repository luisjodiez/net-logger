package adapters

import (
	"net"
	"time"
)

type TCPPinger struct{}

func (p *TCPPinger) Ping(target string) string {
	conn, err := net.DialTimeout("tcp", target, time.Second)
	if err == nil {
		conn.Close()
		return "ok"
	}
	return "ko"
}
