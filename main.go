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
	records.GetRecords(args[0], args[1])
}
