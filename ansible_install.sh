#!/bin/bash
#Don't change this to $home it will fail in ansible
export GOPATH=/home/capillaryDeploy/go
export GOROOT=/usr/local/go
export GOBIN=/home/capillaryDeploy/go/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
go get github.com/tools/godep
go install github.com/tools/godep
godep restore
go install github.com/CapillarySoftware/goiostat
