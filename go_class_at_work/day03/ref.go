package main

import (
	"fmt"
	"strings"
)

type damkot struct {
	A           int  // defined as an integer
	Routers     []string  // defined as an array of strings
	Datacenters map[string]string // defined as a hash of strings referncing other strings.
}

func main() {
	// define damkot as d and fire up the instance
	var d damkot
	d = damkot{
		A:           42, // integer
		Routers:     []string{"BBR1.was", "BBR2.was", "BBR1.chi", "BBR2.chi"}, // list of strings
		Datacenters: make(map[string]string), // hash of strings
	}
	// this is a for loop that iterates through d.Routers, it chops up the router name by '.' and grabs field
	// 1 to get datanceter name. and it maps those values as a key value pair in the map(hash) d.Datacenters.
	for _, value := range d.Routers[:] {  // the _ means discard the value (kind of like self from python).
		tokens := strings.Split(value, ".") // chop up the value variables by '.'
		d.Datacenters[value] = tokens[1]  // this maps datacenter by routername => dc
	}

	fmt.Println(d.Datacenters)
	//fmt.Printf("%v \n", d)
	change(d)
	//fmt.Printf("%v \n", d)
}
func change(d damkot) {
	d.A = 553
	//fmt.Printf("%v \n", d)
}
