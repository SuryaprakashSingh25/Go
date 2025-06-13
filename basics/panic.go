package basics

import "fmt"

func main() {

	// panic(interface{})
	//Example of valid input
	process(10)

	//Example of invalid input
	process(-3)

}

func process(input int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non negative number")
		//not be called after panic
		//fmt.Println("After panic")
		//defer fmt.Println("Defer 3")
	}
	fmt.Println("Processing Input", input)
}
