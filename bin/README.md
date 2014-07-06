goiostat
========

Implementation of iostat in go. 

Currently only support linux 2.6 kernel.

<h2>install directions</h2>
<pre><code>export GOPATH=$HOME/goiostat
export PATH=$PATH:$GOPATH/bin 
export GOBIN=$GOPATH/bin
sh bin/install_dependencies.sh

go install src/github.com/CapillarySoftware/goiostat/iostat.go
</code></pre>