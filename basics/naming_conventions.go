package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// eg. CalculateArea
	// Structs, interfaces, enums

	// snake_case
	// eg. user_id

	// UPPERCASE
	// Use case is constants

	// mixedCase
	// eg. javaScript

	const MAXENTRIES = 5

	var employeeId = 1001
	fmt.Println("Emplyee Id: ", employeeId)
}
