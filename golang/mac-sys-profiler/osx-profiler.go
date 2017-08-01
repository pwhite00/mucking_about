package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

func main() {
	// just a script to use golang to wrapper osX /usr/sbin/system_profiler
	modeFlag := flag.String("mode", "null", "Specify a system_profiler mode to query.")
	flag.Parse()

	fmt.Printf("Collecting system_profiler info using mode = [%s]\n", *modeFlag)
	run_query(string(*modeFlag))
}

func run_query(search string) {
	fmt.Println("not yet implemented.(query)")
	if search == "null" {
		fmt.Printf("\tERROR:\tNo mode selected. Retry command with -help for assistance.\n")
		os.Exit(1)
	}

	var queryXml = []byte{}
	//queryXml, err := exec.Command("/usr/sbin/system_profiler", "-xml", search).Output()
	queryXml, err := exec.Command("/usr/sbin/system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(queryXml))
	//fmt.Print(queryXml)
	//convert_xml_to_json(queryXml)
}

func convert_xml_to_json(rawData byte) {
	fmt.Println("not yet implemented.(convert)")
}
