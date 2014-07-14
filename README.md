goiostat
========
[![Build Status](https://travis-ci.org/CapillarySoftware/goiostat.png)](https://travis-ci.org/CapillarySoftware/goiostat)

Implementation of iostat in go that allows you to send data over zeromq with protobuffers.

Currently only support linux 2.6 kernel.

<h2>install directions</h2>
<pre><code>
export GOPATH=$HOME/code/goiostat
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
sh bin/install_dependencies.sh

go install src/github.com/CapillarySoftware/goiostat/iostat.go

if you want zeromq output make sure to install zmq3 from http://zeromq.org/intro:get-the-software
on mac just use brew:
brew install homebrew/versions/zeromq32

</code></pre>
