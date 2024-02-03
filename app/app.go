package app

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, startPort, endPort int) {
	for port := startPort; port <= endPort; port++ {
		if scanPort(host, port) {
			fmt.Printf("Port %d is open\n", port)
		} else {
			fmt.Printf("Port %d is closed\n", port)
		}
	}
}

func scanPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false
	}
	err = conn.Close()
	if err != nil {
		return false
	}
	return true
}
