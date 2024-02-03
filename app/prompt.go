package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetPorts(defaultStart, defaultEnd int) (int, int) {
	startPort, endPort := defaultStart, defaultEnd
	if askForCustomRange() {
		startPort = getPort("Enter start port:", startPort)
		endPort = getPort("Enter end port:", endPort)
		for startPort > endPort {
			fmt.Println("Start port cannot be greater than end port. Please enter again.")
			startPort = getPort("Enter start port:", startPort)
			endPort = getPort("Enter end port:", endPort)
		}
	}
	return startPort, endPort
}

func getPort(prompt string, defaultPort int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(prompt)
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if port, err := strconv.Atoi(input); err == nil && port > 0 && port <= 65535 {
				return port // Valid port
			} else {
				fmt.Println("Invalid port. Please enter a valid port number.")
			}
		} else {
			fmt.Printf("No input detected. Using default port: %d\n", defaultPort)
			return defaultPort // Use default if no input is given
		}
	}
}

func askForCustomRange() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to enter a custom ports range? (yes/no)")
	if scanner.Scan() {
		response := strings.ToLower(strings.TrimSpace(scanner.Text()))
		return response == "yes" || response == "y"
	}
	return false
}

func GetHost(defaultHost string) string {
	if defaultHost != "" {
		return defaultHost
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please specify a host to scan:")
	if scanner.Scan() {
		host := strings.TrimSpace(scanner.Text())
		if host != "" {
			return host
		}
	}
	fmt.Println("No host specified. Exiting.")
	os.Exit(1)
	return "" // This return is never reached but necessary for compilation.
}
