package basics

import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Deferred value:", i)
	defer fmt.Println("Deferred statement executed")
	defer fmt.Println("2nd Deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i", i)
}
