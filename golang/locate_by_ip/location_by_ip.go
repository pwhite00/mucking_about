package main

// writing this to get used to calling a web api with golang
// woohoo first real golang script. -pwhite 6212017
// version 1.0

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"     // only uncomment for tests using debug section in main.
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
	// this is flag; it works like optparse in other language note; ip == -ip flag, none == default value, usage well usage
	// call help using -h or -help. note calling help will exit with 2.
	ipaddress := flag.String("ip", "none", "Enter an IP address to trace or defaults to current IP in use.")
	flag.Parse()

	// This is the debugging section to work out if conditions as well as using strings stored in flag.
	//fmt.Print("ip:", *ipaddress, "\n")
	//if *ipaddress == "none" {
	//	fmt.Println("it's nothing really.")
	//	fmt.Println(reflect.TypeOf(*ipaddress))
	//	fmt.Println("https://freegeoip.net/json/")
	//} else {
	//	fmt.Println("I'm pretty sure it's something actually.")
	//	fmt.Println(reflect.TypeOf(*ipaddress))
	//	fmt.Println("https://freegeoip.net/json/" + *ipaddress)
	//	thing := "https://freegeoip.net/json/" + *ipaddress
	//	fmt.Println(thing)
	//}

	// this if statement checks if an IP was manualy passed in. It will change the URL used to
	// allow the script to target the passed in IP vs the current one the node is using.
	if *ipaddress == "none" {
		url := "https://freegeoip.net/json/"
		call_api(url)
	} else {
		url := "https://freegeoip.net/json/" + *ipaddress
		call_api(url)
	}
}

func call_api(url string) {
	// call freegeoip.net/json/ and get data on location

	// define the http.Client and set a timeout
	locationClient := http.Client{
		Timeout: time.Second * 2, // 2 second max
	}

	// make the http get call to the api and check if errors occur, report them if them do.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// define a keyvalue pair - find out more about this one.
	req.Header.Set("User-Agent", "location-data")

	// check the response from the api and repo errors
	res, getErr := locationClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	// read the output from the api and report errors
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// do some json munching of the data and report errors
	details1 := details{}
	jsonErr := json.Unmarshal(body, &details1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// write the json output. normally in a more interesting way.
	fmt.Println("Ipaddress", details1.Ipaddress)
	fmt.Println(details1.City, ",", details1.Region_name, ",", details1.Country_name)
	fmt.Println("Latitude:", details1.Latitude, "Longitude:", details1.Longitude)

}
