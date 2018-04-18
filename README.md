# go-iptables

[![GoDoc](https://godoc.org/github.com/coreos/go-iptables/iptables?status.svg)](https://godoc.org/github.com/coreos/go-iptables/iptables)
[![Build Status](https://travis-ci.org/coreos/go-iptables.png?branch=master)](https://travis-ci.org/coreos/go-iptables)

Go bindings for iptables utility.

In-kernel netfilter does not have a good userspace API. The tables are manipulated via setsockopt that sets/replaces the entire table. Changes to existing table need to be resolved by userspace code which is difficult and error-prone. Netfilter developers heavily advocate using iptables utlity for programmatic manipulation.

go-iptables wraps invocation of iptables utility with functions to append and delete rules; create, clear and delete chains.

## Test
```bash
$ SUDO_PERMITTED=true ./test
Building go-iptables...
Checking gofmt...
Running tests...
ok  	github.com/morganwu277/go-iptables/iptables	(cached)	coverage: 74.1% of statements
Success
```