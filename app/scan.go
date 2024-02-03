package app

import (
	"fmt"
	"github.com/briandowns/spinner"
	. "github.com/tikarammardi/portscanner/utils"
	"net"
	"time"
)

func ScanPorts(host string, startPort, endPort int) {

	for port := startPort; port <= endPort; port++ {
		if scanPort(host, port) {
			fmt.Printf("%sPort Number %d is open\n", RedColor, port)
		} else {
			fmt.Printf("%sPort Number %d is closed\n", GreenColor, port)
		}
	}
}

func scanPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	showSpinner(host)
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

func showSpinner(host string) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Scanning port on %s", host)
	s.Color("red")
	s.Start()
	time.Sleep(time.Second)
	s.Stop()
}
