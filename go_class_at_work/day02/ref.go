package main

import (
	"fmt"
)

type damkot struct {
	A int
}

func main() {
	var d *damkot
	d = &damkot{A: 42}
	fmt.Printf("%v \n", d)
	change(d)
	fmt.Printf("%v \n", d)
}
func change(d *damkot) {
	d.A = 553
	fmt.Printf("%v \n", d)
}

