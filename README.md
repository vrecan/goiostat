goiostat
========
[![Build Status](https://travis-ci.org/CapillarySoftware/goiostat.svg?branch=master)](https://travis-ci.org/CapillarySoftware/goiostat)

Implementation of iostat in go that allows you to send data over zeromq with protobuffers or json.

Currently only support linux 2.6 kernel.

<h2>install directions</h2>
```
export GOPATH=$HOME/go
export GOBIN=$HOME/bin
export PATH=$PATH:$GOBIN
go get github.com/tools/godep
go get code.google.com/p/go.tools/cmd/cover
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go install github.com/onsi/ginkgo/ginkgo
go install github.com/onsi/gomega
export PATH=$PATH:$HOME/gopath/bin
$GOBIN/godep restore
go install github.com/CapillarySoftware/goiostat
$GOBIN/ginkgo -r -cover -race //unit tests 
```
<p>
 by default it only compiles with log output. If you want to compile with
 ZeroMQ support you need to run
 </p>
 ```
go build -tags zmq 
 ```
 or
 ```
go install -tags zmq
 ```
<p>
 If you want nanomsg support you can also compile with 
 </p>
 ```
go install -tags nano
 ```

 if you want to build with both nano and go you can do the following
```
go install -tags 'nano zmq'
```
<p>
make sure to install zmq3 from http://zeromq.org/intro:get-the-software
on mac just use brew:
brew install homebrew/versions/zeromq32
Also install nanomsg version 0.4
brew install nanomsg 
</p>
