package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// a place to store and call osx_facts

func osxFacts() {
	// basic facts switcher
	getOsxHostname()
	getOsxFQDN()
	getOsxSerialNumber()
	getOsxUptime()
}

func getOsxHostname() {
	system_hostname, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	munched_hostname := strings.Split(string(system_hostname), ".")
	fmt.Println(munched_hostname[0])
}

func getOsxFQDN() {
	hostname_fqdn, err := exec.Command("uname", "-n").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(hostname_fqdn))
}

func getOsxSerialNumber() {
	serial_number, err := exec.Command("system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		fmt.Println(err)
	}
	parsed_serial_array := strings.Split(string(serial_number), "\n")
	parsed_serial_number := strings.Split(string(parsed_serial_array[15]), " ")
	fmt.Println(parsed_serial_number[9])
	//fmt.Println(reflect.TypeOf(parsed_serial_number))
	// figure out why my chopped up string field seems tp startw with a new line.
}

func getOsxUptime() {
	sys_uptime, err := exec.Command("uptime").Output()
	if err != nil {
		fmt.Println(err)
	}
	// need to split stuff up or flip to clock secs up.
	//fmt.Println(string(sys_uptime))
	parsed_uptime_array := strings.Split(string(sys_uptime), " ")
	fmt.Printf("%s days\n", parsed_uptime_array[3])


}

func getOsxInterfaces() {
	fmt.Println("Not yet implemented.")

}

func getOsxIpAddress() {
	fmt.Println("Not yet implemented.")

}

func getOsxMacAddress() {
	fmt.Println("Not yet implemented.")

}

func getOsxOsVersion() {
	fmt.Println("Not yet implemented.")

}

func getOsxOsMajorVersion() {
	fmt.Println("Not yet implemented.")

}

func getOsxOsMinorVersion() {
	fmt.Println("Not yet implemented.")

}

func getOsxAllSwap() {
	fmt.Println("Not yet implemented.")

}

func getOsxUsedSwap() {
	fmt.Println("Not yet implemented.")

}

func getOsxAllMem() {
	fmt.Println("Not yet implemented.")

}

func getOsxUsedMem() {
	fmt.Println("Not yet implemented.")

}

func getOsxProcs() {
	fmt.Println("Not yet implemented.")

}

func isOsxVirtual() {
	fmt.Println("Not yet implemented.")

}

func getOsxDomain() {
	fmt.Println("Not yet implemented.")

}
