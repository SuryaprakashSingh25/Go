package basics

import "fmt"

func main() {

	//var arrayName [size]elementType
	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// String Arrays
	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array: ", fruits)

	// fmt.Println("Third Element: ", fruits[2])

	// Copying one array to another
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int
	copiedArray = &originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	// fmt.Println("Copied Array: ", *copiedArray)

	// Ways to iterate over array
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index,", i, ": ", numbers[i])
	// }

	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// Underscore is blank identifier, used to store unused values
	// for _, value := range numbers {
	// 	fmt.Printf("Value: %d\n", value)
	// }

	// Length of an array
	// fmt.Println("The length of numbers array is: ", len(numbers))

	// Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}
	// fmt.Println("Array 1 is equal to Array 2? ", array1 == array2)

	// Matrix
	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(matrix)

	// Underscore example using a function
	// a, _ := someFunction()
	// fmt.Println(a)
	// // fmt.Println(b)
}

// func someFunction() (int, int) {
// 	return 1, 2
// }
