package main

import (
	"fmt"
	"os/exec"
)

// a place to store ubuntu_facts
func ubuntuFacts() {
	// basic facts switcher
	getUbuntuHostname()
}

func getUbuntuHostname() {
	system_hostname, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	munched_hostname := string(system_hostname)
	fmt.Println(munched_hostname)
}

func getUbuntuFQDN() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuSerialNumber() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuUptime() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuInterfaces() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuIpAddress() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuMacAddress() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuOsVersion() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuOsMajorVersion() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuOsMinorVersion() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuAllSwap() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuUsedSwap() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuAllMem() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuUsedMem() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuProcs() {
	fmt.Println("Not yet implemented.")

}

func isUbuntuVirtual() {
	fmt.Println("Not yet implemented.")

}

func getUbuntuDomain() {
	fmt.Println("Not yet implemented.")

}
