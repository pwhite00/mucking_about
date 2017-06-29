package main

// writing this to work on reading and executing system commands and parsing options.

import (
	"fmt"

	"flag"
	"os"
)

func main() {
	// run some things and do some stuff
	// getopts ?   // maybe struct ?
	batch_mode := flag.String("batch", "none", "Run in batch mode by feeding a file to teh script")
	input_file := flag.String("in-file", "none", "filename to change.")
	media_tv := flag.Bool("tv", false, "Is the media a tv show ?")
	resolution := flag.String("res", "none", "Video resolution.")
	year := flag.String("year", "none", "What year is the video file from.")
	flag.Parse()

	fmt.Println(flag.Args()

	if *batch_mode != "none" {
		process_batch(*batch_mode)
	}

	if *input_file != "none" {
		process_single_file(*input_file, *media_tv, *resolution, *year)
	}

	fmt.Println("wtf... you forget somthing ?")
	os.Exit(1)
}

func process_single_file(infile string, mediaTv bool, resolution string, year string) {
	// process a single file.
	fmt.Println(infile)
	fmt.Println(resolution)
	fmt.Println(year)
	if mediaTv {
		fmt.Println("it's true !")
	}
	os.Exit(0)
}

func process_batch(batch_file string) {
	// eat a batch file . much it up and process a bunch of files.
	fmt.Println("Process a bunch of files !")
	os.Exit(0)
}