package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

func getRecords(args string, url string) {
	switch args {
	case "ns":
		arg, err := net.LookupNS(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range arg {
			fmt.Println(arg[i].Host)
		}
	case "ip":
		arg, err := net.LookupIP(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range arg {
			fmt.Println(arg[i])
		}
	case "cname":
		arg, err := net.LookupCNAME(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(arg)
	case "mx":
		arg, err := net.LookupMX(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range arg {
			fmt.Println(arg[i].Host, arg[i].Pref)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		help()
		return
	}
	getRecords(args[0], args[1])
}
