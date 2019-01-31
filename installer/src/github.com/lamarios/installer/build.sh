#!/bin/sh

GOPATH="$(pwd)/../../../../dotfiles/installer"

VERSION=$1
ARCH=$(go env GOARCH)
OS=$(go env GOOS)

echo "Building installer $VERSION  os:$OS arch:$ARCH"


go clean
go build -ldflags="-s -w -X main.VERSION=$VERSION" -o install
#Compressing

chmod +x install
mv install ../../../../../install


#tar -czf unifi-throughput-$VERSION-$OS-$ARCH.tar.gz