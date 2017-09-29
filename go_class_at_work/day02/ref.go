package main

import (
	"fmt"
)

type damkot struct {
	// a struct is a sort of format map. different from a map in that it's not a hash.
	// it defines what a object should have in it by type and identifier.
	// as in below variable "A" will be an integer when an instance of struct 'damkot'
	// also damkot can only be called in this file because it is not capitolized and therefore
	// not global.
	A int
}

func main() {
	// the next 2 lines define a variable "d" as a instance of damkot and then instantiate it.
	var d *damkot
	d = &damkot{A: 42} // here we set d.A to 42
	fmt.Printf("%v \n", d)
	// by then calling the change function we get to see that we have changed d.A from 42 to 553.
	change(d)
	// likewise we can see our change carried through and has reset the value of d.A beyind the change function as well.
	fmt.Printf("%v \n", d)
}
func change(d *damkot) {
	d.A = 553
	fmt.Printf("%v \n", d)
}
