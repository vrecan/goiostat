goiostat
========
[![Build Status](https://travis-ci.org/CapillarySoftware/goiostat.png)](https://travis-ci.org/CapillarySoftware/goiostat)

Implementation of iostat in go that allows you to send data over zeromq with protobuffers or json.

Currently only support linux 2.6 kernel.

<h2>install directions</h2>
<pre>
<code>
* go get github.com/tools/godep
* go get code.google.com/p/go.tools/cmd/cover
* go get github.com/onsi/ginkgo/ginkgo
* go get github.com/onsi/gomega
* go install github.com/onsi/ginkgo/ginkgo
* go install github.com/onsi/gomega
* export PATH=$PATH:$HOME/gopath/bin
* $HOME/gopath/bin/godep restore
* go install github.com/CapillarySoftware/goiostat
* $HOME/gopath/bin/ginkgo -cover -r --race //unit tests


if you want zeromq output make sure to install zmq3 from http://zeromq.org/intro:get-the-software
on mac just use brew:
brew install homebrew/versions/zeromq32


</code></pre>
