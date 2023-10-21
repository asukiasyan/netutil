package records

import (
	"fmt"
	"log"
	"net"
)

func GetRecords(args string, url string) {
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
