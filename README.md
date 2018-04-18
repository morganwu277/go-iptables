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

## Sample Usage
Pls review `example` folder
And next is the corresponding output when executing `./build_run.sh`:
```bash
	 -P INPUT ACCEPT
===========================
Adding unique new port 22 from 192.0.2.0/24
	 -P INPUT ACCEPT
	 -A INPUT -s 192.0.2.0/24 -p udp -m udp --dport 22 -j ACCEPT
===========================
Save iptables to /tmp/rules.v4
	 -P INPUT ACCEPT
	 -A INPUT -s 192.0.2.0/24 -p udp -m udp --dport 22 -j ACCEPT
===========================
Restore iptables from /tmp/rules.v4
	 -P INPUT ACCEPT
	 -A INPUT -s 192.0.2.0/24 -p udp -m udp --dport 22 -j ACCEPT
===========================
Deleting unique new port 22 from 192.0.2.0/24
	 -P INPUT ACCEPT
===========================
```

