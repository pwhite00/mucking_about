package main

import (
	"fmt"
	"strconv"
)

type Human interface {
	// an Interface is a structured way to create a process for taking in data in variables forms and
	// processing it in a set way.  { come up with a better wording for this. }
	GetBaseballCard() *BaseballCard
}

type BaseballCard struct {
	// Struct called BaseballCard that keeps a String called Vitals.
	Vitals string
}

type Damkot struct {
	// Struct called Damkot that keeps a String called Name and an Integer called Age.
	Name string
	Age  int
}

type Eileen struct {
	// Struct called Eileen that keeps a String for FirstName , a String for LastName and a Bool value for Alive.
	FirstName string
	LastName  string
	Alive     bool
}

func (e *Eileen) GetBaseballCard() *BaseballCard {
	// using the e instance of pointer to Eileen return some values as &BaseballCard
	return &BaseballCard{e.FirstName + "" + e.LastName}
}

func (d *Damkot) GetAge() int {
	// run the GetAge method on the d instance of Damkot
	return d.Age
}

func (d *Damkot) GetBaseballCard() *BaseballCard {
	// this runs the GetBaseballCard method using d *Damkot and a pointer to struct BaseballCard
	// it then returns an instance of
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
