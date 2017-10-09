package main

// todo:
// create an interface for structured data returns
// maybe create some structs to make it easier.

// todo:
// add the ability for script to use 2 flags and not just -a or one f the others.
// -a should override the others so no duplication is performed.
// missing -a we should get a boolean to continue , if only one exists we should call report_output
// after just one test.

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pid_lookup := flag.Int("p", 0, "Define the PID to search for.")
	get_memory_usage := flag.Bool("m", false, "Report memory usage for PID.")
	get_cpu_usage := flag.Bool("c", false, "Report cpu usage for PID.")
	get_sys_load := flag.Bool("l", false, "Display systemload for PID")
	get_all_outputs := flag.Bool("a", false, "Display output from all checks.")
	flag.Parse()

	// convert flag opts to usage variables
	pid_lookup_value := *pid_lookup
	get_all_outputs_value := *get_all_outputs
	get_memory_usage_value := *get_memory_usage
	get_cpu_usage_value := *get_cpu_usage
	get_sys_load_value := *get_sys_load


	// run all checks if -a
	if get_all_outputs_value {
		check_memory_usage(pid_lookup_value)
		check_cpu_usage(pid_lookup_value)
		check_sys_load(pid_lookup_value)
		report_output()
	}

	// run cpu check if -c
	if get_cpu_usage_value {
		check_cpu_usage(pid_lookup_value)
		report_output()
	}

	// run memory check if -m
	if get_memory_usage_value {
		check_memory_usage(pid_lookup_value)
		report_output()

	}

	// run system load check if -l
	if get_sys_load_value {
		check_sys_load(pid_lookup_value)
		report_output()
	}

}

func check_memory_usage(pid int) {
	// take pid and find memory usage info
	fmt.Println("not implemented yet.", pid)
}

func check_cpu_usage(pid int) {
	// take pid and find cpu usgae info
	fmt.Println("not implemented yet.", pid)
}

func check_sys_load(pid int) {
	// determine the pids impact on system load.
	fmt.Println("not implemented yet.", pid)
}

func report_output() {
	// figure out a good round up output manager using a struct and or interface.
	fmt.Println("not implemented yet. [Exiting].")
	os.Exit(0)
}
