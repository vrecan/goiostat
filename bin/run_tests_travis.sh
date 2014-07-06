#export GOPATH:/home/travis/gopath/src/github.com/CapillarySoftware/goiostat
#export PATH=$PATH:$GOPATH/bin 
#export GOBIN=$GOPATH/bin
#go get github.com/dustin/go-humanize
#go get github.com/onsi/gomega
#go get github.com/onsi/ginkgo
ginkgo -cover -r 
