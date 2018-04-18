package main

import (
	"fmt"
)

import "github.com/morganwu277/go-iptables/iptables"

func printRules(ipt *iptables.IPTables) {
	ipt.ApplyCmdType(iptables.CmdIPTables)
	rules, _ := ipt.List("filter", "INPUT")
	for _, r := range rules {
		fmt.Printf("\t %v\n", r)
	}
	fmt.Println("===========================")
}
func main() {
	ipt, _ := iptables.New()
	printRules(ipt)


	fmt.Println("Adding unique new port 22 from 192.0.2.0/24")
	ipt.AppendUnique("filter", "INPUT", "-p", "udp", "-s", "192.0.2.0/24", "--dport", "22", "-j", "ACCEPT")
	printRules(ipt)

	path := "/tmp/rules.v4"
	fmt.Printf("Save iptables to %v\n", path)
	ipt.ApplyCmdType(iptables.CmdIPTablesSave)
	err := ipt.Save(path)
	if err != nil {
		fmt.Printf("Error! %v", err)
	}
	printRules(ipt)

	path = "/tmp/rules.v4"
	fmt.Printf("Restore iptables from %v\n", path)
	ipt.ApplyCmdType(iptables.CmdIPTablesRestore)
	err = ipt.Restore(path)
	if err != nil {
		fmt.Printf("Error! %v", err)
	}
	printRules(ipt)

	fmt.Println("Deleting unique new port 22 from 192.0.2.0/24")
	ipt.Delete("filter", "INPUT", "-p", "udp", "-s", "192.0.2.0/24", "--dport", "22", "-j", "ACCEPT")
	printRules(ipt)

}
