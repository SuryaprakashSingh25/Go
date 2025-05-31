package basics

import "fmt"

var middleName = "Cane"

func main() {
	var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"

	//Default Values
	//Numeric Types: 0
	//Boolean: false
	//String: ""
	//Pointers,slices,maps,functions and structs: nil

	// ----Scope
	fmt.Println(firstName)
	fmt.Println(middleName)
}

func printname() {
	firstName := "Michael"
	fmt.Println((firstName))
}
