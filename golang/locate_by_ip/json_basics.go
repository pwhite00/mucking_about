package main

import (
	"encoding/json" // needed for json
	"fmt"
	"io/ioutil" // added the following packages to work with teh api
	"log"
	"net/http"
	"time"
)

type people struct {
	// make the variables from the people struct exportable
	Number  int    `json:"number"` // json.number corisponds the top level "number" field in the json.
	Outcome string `json:"message"`
}

func main() {
	// broke the exmaples out into seperate functions and call them both from main
	local_json()
	api_json()
}

func local_json() {
	// text is a locally contained JSON used in this exmaple to allow easyier json learning.
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	// json.Unmarshall works with type byte[] instead of string. So []byte(text) allows us to import teh string and use it as a []byte
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)
	fmt.Println(people1.Outcome)
}

func api_json() {
	// use an api to get the json and then do somthing with it.
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, //maximum 2 sec
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(people1.Number)
	fmt.Println(people1)
}
