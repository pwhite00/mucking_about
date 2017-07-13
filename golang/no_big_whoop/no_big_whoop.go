package main

import (
	"fmt"
)


// note that you do NOT define the 2nd file in import statements.

func main() {
	fmt.Println("Message from func main in main.go.")

	fmt.Println("How can I run a function defined in a different file in the same package.")

	aMessageForYou()

}
