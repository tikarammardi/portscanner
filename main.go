package main

import (
	"fmt"
	"portscanner/app"
)

func main() {
	host := "localhost"
	startPort := 1
	endPort := 65535

	fmt.Println("Scanning ports on", host, "from", startPort, "to", endPort)
	app.ScanPorts(host, startPort, endPort)

}
