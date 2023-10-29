package records

import (
	"fmt"
	"log"
	"net"
)

func GetRecords(args string, url string) {
	switch args {
	case "ns":
		ns, err := net.LookupNS(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range ns {
			fmt.Println(ns[i].Host)
		}
	case "ip":
		ip, err := net.LookupIP(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range ip {
			fmt.Println(ip[i])
		}
	case "cname":
		cname, err := net.LookupCNAME(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(cname)
	case "mx":
		mx, err := net.LookupMX(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := range mx {
			fmt.Println(mx[i].Host, mx[i].Pref)
		}
	default:
		fmt.Println("\nIvalid argument, valid arguments are: ns, ip, cname, mx\n")
	}
}
