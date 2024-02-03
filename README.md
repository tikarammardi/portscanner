
# Port Scanner in Golang
Introduction
This is a simple and efficient Port Scanner written in the Go programming language (Golang). A Port Scanner is a tool that helps you discover open ports on a target system. It can be a useful utility for network administrators and security professionals to identify potential vulnerabilities in a network.

Features
Fast and lightweight
Concurrent scanning for improved speed
Range-based or single-port scanning
Easy-to-use command-line interface

## Requirements
- Go 1.16 or later
## Installation
1. Make sure you have Go installed on your system. You can download it from https://golang.org/dl/.

2. Clone this repository or download the source code:

```bash
 git clone https://github.com/tikarammardi/portscanner.git
````
1. Navigate to the project directory:

```bash
cd portscanner
```
1. Build the project:

```bash
go build
```
## Usage

```bash
# Run the project
./portscanner --host <host> --start <start> --end <end>
```

Options:
- `--host`: The target host to scan (required)
- `--start`: The start port (default: 1)
- `--end`: The end port (default: 65535)

Example:
```bash
./portscanner --host example.com --start 80 --end 443
#or
./portscanner --host=example.com --start=80 --end=443
```


## You can also use the Makefile to build the project:
```bash
make build
```
## Clean the project
```bash
make clean
```


