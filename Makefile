GOPATH:=`pwd`/vendor:$(GOPATH)
GOPATH:=`pwd`/vendor/src/github.com/miekg/dns:$(GOPATH)
VERSION:=0.1
HARDWARE=$(shell uname -m)

build: vendor
	GOPATH=$(GOPATH) go build

test: vendor

release:

clean:

vendor:
	sh vendor.sh
