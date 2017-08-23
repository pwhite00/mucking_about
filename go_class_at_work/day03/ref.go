package main

import (
	"fmt"
	"strings"
)

type damkot struct {
	A           int
	Routers     []string
	Datacenters map[string]string
}

func main() {
	var d damkot
	d = damkot{
		A:           42,
		Routers:     []string{"BBR1.was", "BBR2.was", "BBR1.chi", "BBR2.chi"},
		Datacenters: make(map[string]string),
	}
	for _, value := range d.Routers[:] {
		tokens := strings.Split(value, ".")
		d.Datacenters[value] = tokens[1]
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
