package intermediate

import "fmt"

func main() {

	// Printing Functions in Go
	// fmt.Print("Hello ")

	// fmt.Println("Hello ")

	// name := "John"
	// fmt.Printf("Name:%s\n", name)

	// //Formatting Functions
	// s := fmt.Sprint("Hello", "World", 123, 456)
	// fmt.Println(s)

	// s = fmt.Sprintln("Hello", "World", 123, 456)
	// fmt.Print(s)

	// sf := fmt.Sprintf("Hello %s, your score is %d", "Alice", 95)
	// fmt.Println(sf)

	// Scanning Functions
	// var name string
	// var age int
	// fmt.Print("Enter your name and age: ")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error Formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is not valid, must be 18 or older", age)
	}
	return nil
}
