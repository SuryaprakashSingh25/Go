package intermediate

import "fmt"

func main() {

	// --- General Formatting Verbs
	// %v Prints the value in the default format
	// %#v Prints the value in Go-syntax format
	// %T Prints the type of the value
	// %% Prints the % sign

	i := 15.5
	string := "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	// --- Integer formatting Verbs
	// %b Base 2
	// %d Base 10
	// %o Base 8
	// %O Base 8 with prefix 0o
	// %+d Base 10 with sign
	// %x Base 16, lowercase
	// %X Base 16, uppercase
	// %#x Base 16 with prefix 0x
	// %4d Minimum width of 4 characters, padded with spaces
	//%-4d Minimum width of 4 characters, padded with spaces, left-aligned
	//%04d Minimum width of 4 characters, padded with zeros

	int := 18
	fmt.Printf("%b\n", int)   // Base 2
	fmt.Printf("%d\n", int)   // Base 10
	fmt.Printf("%o\n", int)   // Base 8
	fmt.Printf("%O\n", int)   // Base 8 with prefix 0o
	fmt.Printf("%+d\n", int)  // Base 10 with sign
	fmt.Printf("%x\n", int)   // Base 16, lowercase
	fmt.Printf("%X\n", int)   // Base 16, uppercase
	fmt.Printf("%#x\n", int)  // Base 16 with prefix 0x
	fmt.Printf("%4d\n", int)  // Minimum width of 4 characters, padded with spaces
	fmt.Printf("%-4d\n", int) // Minimum width of 4 characters, padded with spaces, left-aligned
	fmt.Printf("%04d\n", int) // Minimum width of 4 characters, padded with zeros

	// --- String formatting Verbs
	// %s Prints the value as plain string
	// %q Prints the value as a double-quoted string
	// %8s Prints the value as a string with a minimum width of 8 characters, padded with spaces
	// %-8s Prints the value as a string with a minimum width of 8 characters, padded with spaces, left-aligned
	// %x Prints the value as hex dump of byte values
	// % x Prints the value as hex dump of byte values with spaces between bytes

	txt := "World"

	fmt.Printf("%s\n", txt)   // Prints the value as plain string
	fmt.Printf("%q\n", txt)   // Prints the value as a double-quoted string
	fmt.Printf("%8s\n", txt)  // Prints the value as a string with a minimum width of 8 characters, padded with spaces
	fmt.Printf("%-8s\n", txt) // Prints the value as a string with a minimum width of 8 characters, padded with spaces, left-aligned
	fmt.Printf("%x\n", txt)   // Prints the value as hex dump of byte values
	fmt.Printf("% x\n", txt)  // Prints the value as hex dump of byte values with spaces between bytes

	// --- Boolean formatting Verbs
	// %t Prints the value as true or false

	t := true
	f := false
	fmt.Printf("%t\n", t) // Prints the value as true
	fmt.Printf("%t\n", f) // Prints the value as false

	// --- Floating-point formatting Verbs
	// %e Scientific notation (e.g., 1.234567e+10)
	// %f Decimal notation (e.g., 123456789.123456789)
	// %.2f Decimal notation with 2 decimal places (e.g., 123456789.12)
	// %6.2f Decimal notation with 6 characters wide and 2 decimal places (e.g., " 123456.12")
	// %g Exponent as needed, no trailing zeros (e.g., 123456789.123456789)

	flt := 9.18

	fmt.Printf("%e\n", flt)    // Scientific notation
	fmt.Printf("%f\n", flt)    // Decimal notation
	fmt.Printf("%.2f\n", flt)  // Decimal notation with 2 decimal places
	fmt.Printf("%6.2f\n", flt) // Decimal notation with 6 characters wide and 2 decimal places
	fmt.Printf("%g\n", flt)    // Exponent as needed, no trailing zeros

}
