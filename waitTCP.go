package waitTCP

import (
	"net"
	"time"
)

// Check if tcp port is open
func Check(host string, timeout time.Duration) (ret bool, err error) {
	for {
		var timeout time.Duration
		if timeout == 0 {
			timeout = time.Duration(1) * time.Second
		}
		var conn net.Conn
		conn, err = net.DialTimeout("tcp", host, timeout)
		if err != nil {
			ret = true
			err = nil
			return
		}
		if conn != nil {
			err = conn.Close()
			return
		}
	}
}
