package main

import (
	"fmt"
	"os/exec"
)

// a place to debian_store

func debianFacts() {
	// basic facts switcher
	getDebianHostname()
}

func getDebianHostname() {
	system_hostname, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	munched_hostname := string(system_hostname)
	fmt.Println(munched_hostname)
}

func getDebianFQDN() {
	fmt.Println("Not yet implemented.")

}

func getDebianSerialNumber() {
	fmt.Println("Not yet implemented.")

}

func getDebianUptime() {
	fmt.Println("Not yet implemented.")

}

func getDebianInterfaces() {
	fmt.Println("Not yet implemented.")

}

func getDebianIpAddress() {
	fmt.Println("Not yet implemented.")

}

func getDebianMacAddress() {
	fmt.Println("Not yet implemented.")

}

func getDebianOsVersion() {
	fmt.Println("Not yet implemented.")

}

func getDebianOsMajorVersion() {
	fmt.Println("Not yet implemented.")

}

func getDebianOsMinorVersion() {
	fmt.Println("Not yet implemented.")

}

func getDebianAllSwap() {
	fmt.Println("Not yet implemented.")

}

func getDebianUsedSwap() {
	fmt.Println("Not yet implemented.")

}

func getDebianAllMem() {
	fmt.Println("Not yet implemented.")
}

func getDebianUsedMem() {
	fmt.Println("Not yet implemented.")

}

func getDebianProcs() {
	fmt.Println("Not yet implemented.")

}

func isDebianVirtual() {
	fmt.Println("Not yet implemented.")

}

func getDebianDomain() {
	fmt.Println("Not yet implemented.")

}
