goiostat
========
[![Build Status](https://travis-ci.org/CapillarySoftware/goiostat.png)](https://travis-ci.org/CapillarySoftware/goiostat)

Implementation of iostat in go. 

Currently only support linux 2.6 kernel.

<h2>install directions</h2>
<pre><code>
export GOPATH=$HOME/goiostat
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBINk
sh bin/install_dependencies.sh

go install src/github.com/CapillarySoftware/goiostat/iostat.go


</code></pre>
