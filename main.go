package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/tikarammardi/portscanner/app"
	"os"
	"strings"
)

func main() {
	var host string
	pflag.StringVar(&host, "host", "", "Specify the host to scan for open ports.")
	fmt.Println("Before parsing", host)
	pflag.Parse()

	fmt.Println("After parsing", host)

	if host == "" {
		fmt.Println("Please specify a host to scan.")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			host = strings.TrimSpace(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input:", err)
			os.Exit(1)
		}

		if host == "" {
			fmt.Println("No host specified.")
			os.Exit(1)
		}
	}

	startPort := 1
	endPort := 65535

	fmt.Println("Scanning ports on", host, "from", startPort, "to", endPort)
	app.ScanPorts(host, startPort, endPort)

}
