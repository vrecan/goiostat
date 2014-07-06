export GOPATH=/home/travis/gopath
export PATH=$PATH:$GOPATH/bin 
export GOBIN=$GOPATH/bin
ginkgo -cover -r 
