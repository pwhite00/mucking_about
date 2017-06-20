package main

// writing this to get used to calling a web api with golang

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type details struct {
	Ipaddress    string  `json:"ip"`
	Country_code string  `json:country_code`
	Country_name string  `json:"country_name"`
	Region_code  string  `json:"region_code"`
	Region_name  string  `json:"region_name"`
	City         string  `json:"city"`
	Zip_code     string  `json:"zip_code"`
	Time_zone    string  `json:"time_zone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Metro_Code   int     `json:"metro_code"`
}

func main() {
	ipaddress := flag.String("ip", "none", "Enter an IP address to trace or defaults to current IP in use.")
	flag.Parse()

	fmt.Print("ip:", *ipaddress, "\n")
	if *ipaddress == "none" {
		fmt.Println("it's nothing really.")
	} else {
		fmt.Println("I'm pretty sure it's comething actually.")
	}
	if *ipaddress == "none" {
		url := "https://freegeoip.net/json/"
		call_api(url)
	} else {
		combined_info := strings.Join("https://freegeoip.net/json/", *ipaddress)
		url := combined_info
		call_api(url)
	}
}

func call_api(url string) {
	// call freegeoip.net/json/ and get data on location
	locationClient := http.Client{
		Timeout: time.Second * 2, // 2 second max
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "location-data")

	res, getErr := locationClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	details1 := details{}
	jsonErr := json.Unmarshal(body, &details1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(details1)

}
