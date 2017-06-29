package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	// optparse for golang
	ipaddress := flag.String("ip", "1.1.1.1", "this is a string")
	numbPtr := flag.Int("number", 666, "Yay numbers")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("ipaddress:", *ipaddress)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	fmt.Println("object", reflect.TypeOf(flag.Args()))cd ../../

	if *boolPtr == true {
		fmt.Println("it's true !")
	}

}
