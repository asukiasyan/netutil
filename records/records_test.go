package records

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNS(t *testing.T) {
	url := "google.com"
	got := GetNS(url)
	expected := []string{"ns1.google.com.", "ns2.google.com.", "ns3.google.com.", "ns4.google.com."}
	if reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGetIP(t *testing.T) {
	url := "google.com"
	got := GetIP(url)
	expected := []string{"2a00:1450:4001:828::200e", "142.250.186.78"}
	if reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGetCNAME(t *testing.T) {
	cname := "google.com."
	if GetCNAME("google.com") != cname {
		fmt.Printf("CNAME of google.com doesn't match, expected %v, got %v", cname, GetCNAME("google.com"))
	}
}

func TestGetMX(t *testing.T) {
	url := "google.com"
	got := GetMX(url)
	expected := []string{"10", "smtp.google.com."}
	if reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
