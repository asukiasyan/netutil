package main

import (
	"fmt"
	"os"

	"github.com/asukiasyan/netutil/records"
)

func help() {
	fmt.Println(`
NAME:
   	netutil - Network utility to lookup various records for given host

USAGE:
	netutil command [command options] [url...]

COMMANDS:
	  ns	Lookup NameServer for given host
	  ip	Lookup the IP addresses for given host
	  cname	Lookup CNAME records for given host
	  mx	Lookup MX records for given host

GLOBAL OPTIONS:
	  --help, -h  show help
	`)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		help()
		return
	}
	switch args[0] {
	case "ns":
		ns := records.GetNS(args[1])
		fmt.Println(ns)
	case "ip":
		fmt.Println(records.GetIP(args[1]))
	case "cname":
		fmt.Println(records.GetCNAME(args[1]))
	case "mx":
		fmt.Println(records.GetMX(args[1]))
	default:
		fmt.Println("Invalid argument, valid args are: ns, ip, cname, mx")
	}
}
