#!/bin/sh
set -e
apk add --update go build-base git mercurial ca-certificates
mkdir -p /go/src/github.com/clevertechru
cp -r /src /go/src/github.com/clevertechru/logspout
cd /go/src/github.com/clevertechru/logspout
export GOPATH=/go
go get github.com/Masterminds/glide && $GOPATH/bin/glide install
go get golang.org/x/crypto/ssh/terminal
go build -ldflags "-X main.Version=$1" -o /bin/logspout
apk del go git mercurial build-base
rm -rf /go /var/cache/apk/* /root/.glide

# backwards compatibility
ln -fs /tmp/docker.sock /var/run/docker.sock
