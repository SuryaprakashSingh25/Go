package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello\nGO!"
	message1 := "Hello\tGO!"
	message2 := "Hello\rGO!"
	rawMessage := `Hello\nGO!`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message varaible is:", len(message))

	fmt.Println("First Character in message is:", message[0])

	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  //A has an ASCII value of 65
	str2 := "Banana" //B has an ASCII value of 66
	str3 := "app"    //a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	fmt.Println(ch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

}
