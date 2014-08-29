#!/bin/bash
go get github.com/tools/godep
go install github.com/tools/godep
$GOBIN/godep restore
go install github.com/CapillarySoftware/goiostat
