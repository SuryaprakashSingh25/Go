package basics

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const q float64 = 22 / 7
	fmt.Println(q) // 3

	const p float64 = 22 / 7.0
	fmt.Println(p) //3.142

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 //int64 max value
	fmt.Println(maxInt)

	maxInt += 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 //uint64 max value
	fmt.Println(uMaxInt)

	uMaxInt += 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFLoat float64 = 1.0e-323
	fmt.Println(smallFLoat)

	smallFLoat /= math.MaxFloat64
	fmt.Println(smallFLoat)
}
