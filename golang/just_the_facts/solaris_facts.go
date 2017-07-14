package main

import (
	"fmt"
	"os/exec"
)

// a place to store solaris_facts
func solarisFacts() {
	// basic facts switcher
	getSolarisHostname()
}

func getSolarisHostname() {
	system_hostname, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	munched_hostname := string(system_hostname)
	fmt.Println(munched_hostname)
}

func getSolarisFQDN() {
	fmt.Println("Not yet implemented.")

}

func getSolarisSerialNumber() {
	fmt.Println("Not yet implemented.")

}

func getSolarisUptime() {
	fmt.Println("Not yet implemented.")

}

func getSolarisInterfaces() {
	fmt.Println("Not yet implemented.")

}

func getSolarisIpAddress() {
	fmt.Println("Not yet implemented.")

}

func getSolarisMacAddress() {
	fmt.Println("Not yet implemented.")

}

func getSolarisOsVersion() {
	fmt.Println("Not yet implemented.")

}

func getSolarisOsMajorVersion() {
	fmt.Println("Not yet implemented.")

}

func getSolarisOsMinorVersion() {
	fmt.Println("Not yet implemented.")

}

func getSolarisAllSwap() {
	fmt.Println("Not yet implemented.")

}

func getSolarisUsedSwap() {
	fmt.Println("Not yet implemented.")

}

func getSolarisAllMem() {
	fmt.Println("Not yet implemented.")

}

func getSolarisUsedMem() {
	fmt.Println("Not yet implemented.")

}

func getSolarisProcs() {
	fmt.Println("Not yet implemented.")

}

func isSolarisVirtual() {
	fmt.Println("Not yet implemented.")

}

func getSolarisDomain() {
	fmt.Println("Not yet implemented.")

}
