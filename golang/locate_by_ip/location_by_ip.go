package main

// writing this to get used to calling a web api with golang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ip_url string = "https://freegeoip.net/json/"

func main() {
	// run functions from here
	call_api()
	//parse_output(dumped_data)
}

func call_api() {
	// call the API and get some JSON.
	resp, err := http.Get(ip_url)
	if err != nil {
		log.Fatal(err)
	}
	data_dump, err := ioutil.ReadAll(resp.Body)

	var data_raw map[string]interface{}
	data, err := json.Unmarshal(data_dump, &data_raw)

	fmt.Printf(data)

}

func parse_output() {
	// take the json from call_api and do stuff with it.
	fmt.Println("Not Yet Implemented.")
}
