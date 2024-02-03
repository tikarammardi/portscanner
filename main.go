package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/tikarammardi/portscanner/app"
	. "github.com/tikarammardi/portscanner/utils"
)

func main() {
	var (
		host      string
		startPort int
		endPort   int
	)

	pflag.StringVar(&host, "host", "", "Specify the host to scan for open ports.")
	pflag.IntVar(&startPort, "start", StartPort, "Specify the start port for scanning.")
	pflag.IntVar(&endPort, "end", EndPort, "Specify the end port for scanning.")
	pflag.Parse()

	host = app.GetHost(host)
	startPort, endPort = app.GetPorts(startPort, endPort)

	fmt.Printf("Scanning ports on %s from %d to %d\n", host, startPort, endPort)

	app.ScanPorts(host, startPort, endPort)
}
