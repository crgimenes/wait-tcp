package waitTCP

import (
	"net"
	"time"
)

// Check if tcp port is open
func Check(host string, timeout time.Duration) bool {
	var (
		conn  net.Conn
		start = time.Now()
	)
	for {
		conn, _ = net.DialTimeout("tcp", host, 100*time.Millisecond)
		if conn != nil {
			_ = conn.Close()
			return false
		}

		if time.Since(start) > timeout {
			return true
		}
	}
}

// CheckAndConnect if tcp port is open and connect to it
func CheckAndConnect(host string, timeout time.Duration) (net.Conn, error) {
	var (
		conn  net.Conn
		start = time.Now()
		err   error
	)
	for {
		conn, err = net.DialTimeout("tcp", host, 100*time.Millisecond)
		if conn != nil {
			return conn, nil
		}

		if time.Since(start) > timeout {
			return nil, err
		}
	}
}
