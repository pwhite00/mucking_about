package main

import (
	"flag"
	"fmt"
	"os/exec"
)



func main() {
	myflag := flag.String("fact", "null", "Specify a fact to query.")
	flag.Parse()

	fmt.Printf("Flag: %s\n", *myflag)
	time_stamp, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time Stamp: %s", time_stamp)
}
