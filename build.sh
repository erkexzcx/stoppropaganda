#!/bin/bash

VERSION=0.0.24
BINARY_NAME=stoppropaganda

# Remove old binaries (if any)
mkdir -p ./dist
rm -rf ./dist/*

env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_linux_i386" cmd/stoppropaganda/main.go             # Linux i386
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_linux_x86_64" cmd/stoppropaganda/main.go         # Linux 64bit
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_linux_arm" cmd/stoppropaganda/main.go      # Linux armv5/armel/arm (it also works on armv6)
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_linux_armhf" cmd/stoppropaganda/main.go    # Linux armv7/armhf
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_linux_aarch64" cmd/stoppropaganda/main.go        # Linux armv8/aarch64
env CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_freebsd_x86_64" cmd/stoppropaganda/main.go     # FreeBSD 64bit
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_darwin_x86_64" cmd/stoppropaganda/main.go       # Darwin 64bit
env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_darwin_aarch64" cmd/stoppropaganda/main.go      # Darwin armv8/aarch64
env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_windows_i386.exe" cmd/stoppropaganda/main.go     # Windows 32bit
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "dist/${BINARY_NAME}_v${VERSION}_windows_x86_64.exe" cmd/stoppropaganda/main.go # Windows 64bit
