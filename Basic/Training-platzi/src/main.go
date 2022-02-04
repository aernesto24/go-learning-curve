package main

import "fmt"

func main() {
	// Declaring constant values
	const pi float64 = 3.14
	const pi2 = 3.1415

	// Using soome print statements
	fmt.Println("HEllo World!!")
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println()

	//Declaring int variables
	base := 12
	var height int = 14
	//var area int

	//Assigning the result of multip to a variable
	// you need to use := as this variables has not been declared before
	rectangleArea := base * height
	squareArea := base * base

	fmt.Println("rectangle area:", rectangleArea)
	fmt.Println("square area:", squareArea)
	fmt.Println()

	//Zero values - GO dooes not use null or none it assing a default values to variables
	var inta int
	var float64b float64
	var stringc string
	var boold bool

	fmt.Println("the int default value is:", inta)
	fmt.Println("the float64 default value is:", float64b)
	fmt.Println("the string default value is:", stringc)
	fmt.Println("the bool default value is:", boold)
}
