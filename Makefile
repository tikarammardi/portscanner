# Makefile for building and distributing a Go application

APP_NAME := portscanner
DIST_DIR := dist

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean

# Targets for building on different platforms
.PHONY: all build linux darwin windows clean package

all: clean linux darwin windows

build: all

linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(DIST_DIR)/$(APP_NAME)-linux-amd64

darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64

windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe

clean:
	$(GOCLEAN)
	rm -rf $(DIST_DIR)/*

# Placeholder for packaging and distribution steps
package:
