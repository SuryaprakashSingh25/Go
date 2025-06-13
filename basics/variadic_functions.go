package basics

import "fmt"

func main() {

	// ... Ellipsis
	//func functionName(param1 type1,param2 type2,param3 ...type3) returnType{
	// function body
	//}

	//fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
	// statement, total := sum("The sum is:", 1, 2, 3)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence, total := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence, "Total:", total)

}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
