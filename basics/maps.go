package basics

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[keyType]valueType

	// mapVariable := make(map[keyType]valueType)

	// using a Map literal
	// mapVariable=map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["Rohit"] = 45
	myMap["Surya"] = 6
	fmt.Println(myMap)
	fmt.Println(myMap["Rohit"])
	fmt.Println(myMap["Suraj"])

	myMap["Surya"] = 21
	fmt.Println(myMap)

	delete(myMap, "Surya")
	fmt.Println(myMap)

	myMap["Bumrah"] = 11
	myMap["Hardik"] = 33
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	value, unknownValue := myMap["Rohit"]
	fmt.Println(value)
	fmt.Println(unknownValue)

	_, ok := myMap["Rohit"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("Both the maps are equal")
	}

	for i, v := range myMap2 {
		fmt.Println(i, v)
	}

	for _, v := range myMap2 {
		fmt.Println(v)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val := myMap4["Rohit"]
	fmt.Println("Value:", val)

	// myMap4["Rohit"] = "Sharma"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["Rohit"] = "Sharma"
	fmt.Println(myMap4)

	fmt.Println("Length of map is:", len(myMap4))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

}
