package records

import (
	"log"
	"net"
	"strconv"
)

// Returns pointer to slice of NS struct - []*NS
func GetNS(url string) []string {
	ns, err := net.LookupNS(url)
	if err != nil {
		log.Fatal(err)
	}
	var n []string
	for _, v := range ns {
		n = append(n, v.Host)
	}
	return n
}

// Returns slice of IP struct - []IP
func GetIP(url string) []string {

	ip, err := net.LookupIP(url)
	if err != nil {
		log.Fatal(err)
	}
	var i []string
	for _, v := range ip {
		i = append(i, v.String())
	}
	return i
}

// Returns string
func GetCNAME(url string) string {
	cname, err := net.LookupCNAME(url)
	if err != nil {
		log.Fatal(err)
	}
	return cname
}

// Returns pointer to slice of MX struct - []*MX
func GetMX(url string) []string {
	mx, err := net.LookupMX(url)
	if err != nil {
		log.Fatal(err)
	}
	var m []string
	for _, v := range mx {
		m = append(m, v.Host, strconv.Itoa(int(v.Pref)))
	}
	return m
}
