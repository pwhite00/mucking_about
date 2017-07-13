package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	gather_basic_info()
	time_info()
}

func time_info() {
	mytime, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", mytime)

	uptime, err := exec.Command("uptime").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Uptime: %s\n", uptime)
}

func gather_basic_info() {

	my_host, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("Hostname: %s\n", my_host)

}
