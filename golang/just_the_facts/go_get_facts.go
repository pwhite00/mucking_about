package main

import (
	"flag"
	"fmt"
	"os/exec"
	//"reflect"
	"runtime"
)

func main() {
	myflag := flag.String("fact", "null", "Specify a fact to query.")
	verbose := flag.Bool("verbose", false, "Run in verbose Mode.")
	flag.Parse()

	fmt.Printf("Flag: %s\n", *myflag)
	time_stamp, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Verbose mode is %t\n", *verbose)

	fmt.Printf("Time Stamp: %s", time_stamp)

	// determine os and assign that a variable.
	// check uname
	// if uname == Darwin it's OSX
	// if uname == Linux check fpr /etc/lsb-release or /etc/redhat-release
	// if uname == SunOS it's Solaris of some sort
	// if uname != any of those complain and exit.

	if err != nil {
		fmt.Println(err)
	}

	// use a case to manage figuring this stuff out.
	// pass fact and verbose flags into os specific fact functions.
	// make each OS it's own set of facts.
	switch os := runtime.GOOS; os {
	case "darwin":
		// OSX
		fmt.Print("Operating System Detected " + os + "\n")
		fmt.Println("a")
		osxFacts()
	case "linux":
		// Linux (debian, ubuntu,redhat,centos)
		fmt.Print("Operating System Detected " + os + "\n")
		fmt.Println("b")
	case "sunOS":
		// Solaris
		fmt.Print("Operating System Detected " + os + "\n")
		fmt.Println("c")
	default:
		// anything else
		fmt.Print("Operating System Detected " + os + "\n")
		fmt.Println("d")
	}

}
