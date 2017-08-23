package main

import (
	"fmt"
	"strconv"
)

type Human interface {
	GetBaseballCard() *BaseballCard
}

type BaseballCard struct {
	Vitals string
}

type Damkot struct {
	Name string
	Age  int
}

type Eileen struct {
	FirstName string
	LastName  string
	Alive     bool
}

func (e *Eileen) GetBaseballCard() *BaseballCard {
	return &BaseballCard{e.FirstName + "" + e.LastName}
}

func (d *Damkot) GetAge() int {
	return d.Age
}

func (d *Damkot) GetBaseballCard() *BaseballCard {
	return &BaseballCard{d.Name + strconv.Itoa(d.Age)}
}

func main() {
	d := &Damkot{"Peter", 43}
	e := &Eileen{"Eileen", "Watson", true}
	//fmt.Println(d.GetAge())
	//fmt.Println(d.GetBaseballCard())
	humans := []Human{
		d,
		e,
	}

	for _, human := range humans {
		fmt.Println(human.GetBaseballCard())
	}

}
