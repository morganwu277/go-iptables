#!/bin/bash

ORG_PATH="github.com/morganwu277"
REPO_PATH="${ORG_PATH}/go-iptables"

if [ ! -h gopath/src/${REPO_PATH} ]; then
	mkdir -p gopath/src/${ORG_PATH}
	ln -s ../../../../.. gopath/src/${REPO_PATH} || exit 255
fi

export GOBIN=${PWD}/bin
export GOPATH=${PWD}/gopath

eval $(go env)

go build -v
rm -rf gopath
sudo ./example