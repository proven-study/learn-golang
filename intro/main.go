package main

import "fmt"

func oldMain() {
	var x float64 = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	// Late Initialization
	// var num int
	// var amount float32
	// var isValid bool
	// var name string
	// num = 20
	// amount = 49.99
	// isValid = true
	// name = "Bappy"
	// fmt.Println(num, amount, isValid, name)

	// Declare and Initialize Variables without DataType
	var num = 20
	var amount = 49.99
	var isValid = true
	var name = "Bappy"
	fmt.Println(num, amount, isValid, name)
	//Using the shorthand syntax
	y := 42
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y) //y is of type int
	// Declare Multiple Variables
	num1, num2 := 20, 30          //int variables
	amount, name = 49.99, "Bappy" // float and string variables
	fmt.Println(num1, num2, amount, name)

	var_1, var_2 := 1, "hi" //declare var_1 int and var_2 string
	fmt.Println(var_1, var_2)
	var_3, var_2 := 2, "hello" //var_3 is declared, but var_2 is just reassigned
	fmt.Println(var_3, var_2)
}

// Declaring Global variable
var scalar string = "Welcome to scalar"

func main() {
	// Printing Global variable in main
	fmt.Println(scalar + " from main")
	//function call
	displayGreeting()

	const myConst = 10
	fmt.Println(myConst)
	// cannot assign a new value to a constant
	// myConst = 20 // this will cause a compile-time error
}

func displayGreeting() {
	// Printing Global variable from a random function
	fmt.Println(scalar + " from function")
}
