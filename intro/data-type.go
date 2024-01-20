package main

import "fmt"

func main() {
	// basicType()
	// integers()
	floats()
}

func basicType() {
	str := "Hello, world!"
	num1 := 42
	num2 := 3.14159
	num3 := 1234567890
	boolean := true
	char := 'A'
	// %T: print type of the value
	fmt.Printf("%T\n", str)
	// %v: print default format for each type
	fmt.Printf("%v\n", str)
	// String format specifiers
	fmt.Printf("%s\n", str)         // %s: print string
	fmt.Printf("%q\n", str)         // %q: print quoted string
	fmt.Printf("%x\n", []byte(str)) // %x: print hex encoding of bytes
	fmt.Printf("%X\n", []byte(str)) // %X: print uppercase hex encoding of bytes
	// Integer format specifiers
	fmt.Printf("%d\n", num1) // %d: print decimal integer
	fmt.Printf("%b\n", num1) // %b: print binary integer
	fmt.Printf("%o\n", num1) // %o: print octal integer
	fmt.Printf("%x\n", num1) // %x: print hex encoding of integer
	fmt.Printf("%X\n", num1) // %X: print uppercase hex encoding of integer
	fmt.Printf("%c\n", char) // %c: print character
	// Floating-point format specifiers
	fmt.Printf("%f\n", num2) // %f: print floating-point number
	fmt.Printf("%e\n", num2) // %e: print scientific notation of floating-point       number
	fmt.Printf("%E\n", num2) // %E: print scientific notation of floating-point number with uppercase E
	fmt.Printf("%g\n", num2) // %g: print floating-point number in decimal or scientific notation, depending on the value
	fmt.Printf("%G\n", num2) // %G: print floating-point number in decimal or scientific notation, depending on the value with uppercase E
	// Width and precision
	fmt.Printf("|%5d|\n", num1)    // %5d: print decimal integer with minimum width of 5 characters
	fmt.Printf("|%-5d|\n", num1)   // %-5d: print decimal integer with minimum width of 5 characters, left-justified
	fmt.Printf("|%5.2f|\n", num2)  // %5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point
	fmt.Printf("|%-5.2f|\n", num2) // %-5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point, left-justified
	// Boolean format specifiers
	fmt.Printf("%t\n", boolean) // %t: print boolean value
	// Pointer format specifier
	fmt.Printf("%p\n", &num3) // %p: print pointer address
}

func integers() {
	// Signed integers
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807
	// Unsigned integers
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615
	// Print the values
	fmt.Println("Signed Integers:")
	fmt.Printf("int8: %d\n", a)
	fmt.Printf("int16: %d\n", b)
	fmt.Printf("int32: %d\n", c)
	fmt.Printf("int64: %d\n", d)
	fmt.Println("Unsigned Integers:")
	fmt.Printf("uint8: %d\n", e)
	fmt.Printf("uint16: %d\n", f)
	fmt.Printf("uint32: %d\n", g)
	fmt.Printf("uint64: %d\n", h)
}

func floats() {
	// float32 example
	var myFloat32 float32 = 3.14
	fmt.Printf("myFloat32 = %f, type = %T\n", myFloat32, myFloat32)
	// float64 example
	var myFloat64 float64 = 3.141592653589793238462643383279502
	fmt.Printf("myFloat64 = %f, type = %T\n", myFloat64, myFloat64)

}
