package main

import (
	"fmt"
	"sort"
	//	"strconv"
)

// constant dog = rex
const dog = "Rex"

var pkg_variable int = 4

func main() {
	fmt.Println("Test Program with some random functions to see how things work.")

	//run_constant()

	//run_booleans()

	//run_scope()

	//convert_f_to_c()

	//convert_f_to_m()

	//control_structures()

	//data_store_stypes()
	//find_smallest_in_array()
	//my_message := "You foolish English KNIGIGHT"
	//run_a_function(my_message)

	//scores := []float64{98, 93, 79, 82, 83, 62, 100, 56}
	//more_scores := []float64{5, 10, 15, 20}
	// you can all a function in a function
	//fmt.Println(average(scores))
	//fmt.Println(add_stuff(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, int(average(more_scores)), int(average(scores))))
	//fmt.Println(add_stuff(1,2,3,4,5,6,7,8,9,10))

	// a function can be a variable
	//add_more_stuff := add_stuff(1, 2, 3, 4)
	//fmt.Println(add_more_stuff)

	// a function can call itself
	//fmt.Println(factorial(2))

	// the defer statement scdeules a function to be called after a diferent function completes
	//defer defer_second()
	//defer_first()

	// defer and panic
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("PANIC")

	// write a sum function for an array
	sum_numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The sum of the numbers in array.", sum_numbers, "=", sum_test(sum_numbers))

	// half an inetger and return it with true/flse based on even/odd.
	fmt.Println(half_integer(4))
	fmt.Println(half_integer(3))
	fmt.Println(half_integer(2))
	fmt.Println(half_integer(1))

}

func run_constant() {
	// print a constant defined at package level.
	fmt.Println("Running Constant method.")
	fmt.Println("My Constant is \"", dog, "\". He is a dog.")

}

func run_booleans() {
	// print some boolean values
	fmt.Println("Running Boolean method.")
	fmt.Println("When you have true && true you get = ", true && true)
	fmt.Println("When you have true && false you get = ", true && false)
	fmt.Println("When you have true || false you get = ", true || false)
	fmt.Println("When you have false || true you get = ", false || true)
	fmt.Println("When you have true || true you get =", true || true)
	fmt.Println("When you have !true you get = ", !true)
}

func run_scope() {
	// show some exmaples of variable scope.
	fmt.Println("We can take variables at package level and use them in a function. Like this [", pkg_variable, "]")
	var function_variable int = 5
	fmt.Println("You can also define a variable inside a method and use it. [", function_variable, "]")
	fmt.Println("Sadly or not depending on you a method variable is not usable outsde it's scope at say package level.")
}

func convert_f_to_c() {
	// tutorial to change over fahrenheit to celcius ( C = (F - 32) * 5/9 )
	starting_temp := 75
	fmt.Printf("Starting Temp is %d degress Fahrenheit.\n", starting_temp)
	celsius_temp := ((starting_temp - 32) * 5 / 9)
	fmt.Printf("If the it's %d Fahrenheit it's also %d Celsius.\n", starting_temp, celsius_temp)
}

func convert_f_to_m() {
	//tutorial to change feet to meters (1 ft = 0.3048 m)
	starting_feet := float64(25)
	fmt.Printf("Starting with %f feet.\n", starting_feet)
	ending_meters := (starting_feet * 0.3048)
	fmt.Printf("Starting with %F feet. you end up with %F meters\n", float64(starting_feet), ending_meters)
}

func control_structures() {
	// tutorial for control stuctrues like if , case etc...

	// repeat a few lines.
	fmt.Println("this")
	fmt.Println("would")
	fmt.Println("work")
	fmt.Printf("\nBut wow no.\nBut try this.")
	fmt.Println(`1
	2
	3
	4
	5`)

	// for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// if statement
	thing := 4
	if thing == 4 {
		fmt.Println("Thing is equal to ", thing, "True.")
	} else if thing == 76 {
		fmt.Println("Something has gone very wrong", thing, " is not even close to the number...")

	} else {
		fmt.Println("Thing is not equal to ", thing, "Like the thought FALSE.")
	}
}

func data_store_stypes() {
	// tutorial for arrays, slices and maps
	var array_one [5]float64
	fmt.Println("Printed after creation but before objects were defined.", array_one)
	//array_one[0] already == 0 beacuse of the wy it was created.
	array_one[1] = 56
	array_one[2] = 99
	array_one[3] = 77
	array_one[4] = 100

	fmt.Println("Now that objects are defined we acan see that we have values.", array_one)

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += array_one[i]
	}
	fmt.Println(total / float64(len(array_one))) // converts array_one length from a int to a floatt64 so it can be a decimal.

	// slices - a segemt of an array
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	fmt.Println(slice2[2])
	fmt.Println(slice2[2:4])

	// maps (aka hashes)
	y := make(map[string]int)
	y["ten"] = 10
	y["eleven"] = 11
	fmt.Println(y)
	fmt.Println(y["ten"])

	// remove y["eleven"]
	delete(y, "eleven")
	fmt.Println(y)

	colors := map[string]map[string]string{
		"blk": map[string]string{
			"name": "Black",
			"type": "Primary",
		},
		"pink": map[string]string{
			"name": "Pink",
			"type": "bright",
		},
	}
	fmt.Println(colors)
	fmt.Println(colors["blk"]["type"])

}

func find_smallest_in_array() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	fmt.Println(x)

	sort.Ints(x)
	lowest := x[0]
	biggest := x[len(x)-1]
	fmt.Println("The smallest is", lowest, ". But the largest is", biggest, ".")
	fmt.Print("and yes i totally pulled in the sort function to avoid some long comparative loop looking for big / small.")
}

func run_a_function(message string) {
	// give the function a string and blamo it says stuff.
	lame_response := "I know you are but what am I?"
	fmt.Println("You Said: [", message, "].")
	fmt.Println("To that I say... ", lame_response, "?")

	// you can call other functions from inside functions
	find_smallest_in_array()
}

func add_stuff(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func average(xs []float64) float64 {
	// find the average of an array of float64 values.
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func defer_first() {
	fmt.Println("first")
}

func defer_second() {
	fmt.Println("second")
	fmt.Println("Defer is useful when you are say opening a file and need it closed at the end.")
	fmt.Println("You can just toss a \"defer f.close()\" and it'll close the file when the writing function exits.")
}

func sum_test(some_numbers []int) int {
	// take an array and add them all togetther.
	sum_total := 0
	for _, n := range some_numbers {
		sum_total += n
	}
	return sum_total
}

func half_integer(starting_number int) (int, bool) {
	// take an integer , cut it in half and return the nwe number as well as true for even false for odd.
	fmt.Println("The Starting Number is :", starting_number)
	var mybool bool = starting_number%2 == 0
	var myval int = starting_number % 2
	return myval, mybool
}
