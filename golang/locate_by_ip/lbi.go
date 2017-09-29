package main

import (
	"flag"
	"fmt"
	"net/http"
)

type details struct {
	IPAddress   string
	CountryCode string
	City        string
	RegionCode  string
	ZipCode     string
	TimeZone    string
	Latitude    string
	Longitude   string
}

func main() {
	ip := flag.String("ip", "none", "Define an IPv4 address or leave it blank and use the discovered IP of your current host.")
	flag.Parse()

	if *ip == "none" {
		url := "https://freegeoip.net/json/"
		call_api(url)
	} else {
		url := "https://freegeoip.net/json/" + *ipaddress
		call_api(url)
	}
}

func call_api(url string) {

}
