//You Must Define Package on first line.
package Day_3

import (
	"fmt"
	"os"
)

func main() {
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
}

func openrouterdb(filename string) *os.File { return nil }