package main

// in this exmaple we set a to = "G" at a package level.
// then echo it , then set it as "O" at a function level and call it
// then call it from the package level again.
// we then reset it's value and call it and see it now equals "T"

// what see is that by resetting var a to := "O" it can change the variable value but only in the function.
// This is because when we st it as := we treat it as a new variable. and so it only impacts that function.
// if we reset it using = it assigns a new global value and chganges it after.

var a = "G"

func main() {
	n()
	m()
	n()
	o()
	n()

}

func n() {
	print(a)
}

func m() {
	a := "O"
	print(a)
}

func o() {
	a = "T"
	print(a)
}
