//You Must Define Package on first line.
package main

import (
	"fmt"
	"os"
)

func main() {
	// use os.Args to get commandline options that are single word options (:ie script {install|help|unistall})
	// use flags to get optparse style options (ie: script {-h} {-install full})


	fmt.Printf("line elements: %d\n", len(os.Args))
	var alen int
	var i int
	alen = len(os.Args)
	if alen == 2 {
		for i = 0; i < 5; i++ {
			fmt.Printf("hello %d %s \n", i, os.Args[1])
		}
	} else {
		fmt.Printf("provide single entry\n%d \n", i)
	}

	fmt.Print(os.Args, "\n")
	do_a_thing := os.Args[1]
	fmt.Printf("%s \n", do_a_thing)
}
