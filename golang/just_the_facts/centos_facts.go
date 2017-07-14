package main

import (
	"fmt"
	"os/exec"
)

// a place to store centos_facts
func centosFacts() {
	// basic facts switcher
	getCentosHostname()
}

func getCentosHostname() {
	system_hostname, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	munched_hostname := string(system_hostname)
	fmt.Println(munched_hostname)
}

func getCentosFQDN() {
	fmt.Println("Not yet implemented.")

}

func getCentosSerialNumber() {
	fmt.Println("Not yet implemented.")

}

func getCentosUptime() {
	fmt.Println("Not yet implemented.")

}

func getCentosInterfaces() {
	fmt.Println("Not yet implemented.")

}

func getCentosIpAddress() {
	fmt.Println("Not yet implemented.")

}

func getCentosMacAddress() {
	fmt.Println("Not yet implemented.")

}

func getCentosOsVersion() {
	fmt.Println("Not yet implemented.")

}

func getCentosOsMajorVersion() {
	fmt.Println("Not yet implemented.")

}

func getCentosOsMinorVersion() {
	fmt.Println("Not yet implemented.")

}

func getCentosAllSwap() {
	fmt.Println("Not yet implemented.")

}

func getCentosUsedSwap() {
	fmt.Println("Not yet implemented.")

}

func getCentosAllMem() {
	fmt.Println("Not yet implemented.")

}

func getCentosUsedMem() {
	fmt.Println("Not yet implemented.")

}

func getCentosProcs() {
	fmt.Println("Not yet implemented.")

}

func isCentosVirtual() {
	fmt.Println("Not yet implemented.")

}

func getCentosDomain() {
	fmt.Println("Not yet implemented.")

}
