#!/bin/bash
source ~/.bash_profile
go get github.com/tools/godep
go install github.com/tools/godep
godep restore
go install github.com/CapillarySoftware/goiostat
