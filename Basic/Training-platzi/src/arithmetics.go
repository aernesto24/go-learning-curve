package main

import "fmt"

func main() {

	//Here we are goint to work withs ome arithmetic

	//Declare some variables
	x := 10
	y := 50

	//sum
	result := x + y
	fmt.Println(x, "+", y, "equals to:", result)

	//substract
	result = x - y
	fmt.Println(x, "-", y, "equals to:", result)

	//multiplication
	result = x * y
	fmt.Println(x, "*", y, "equals to:", result)

	//substract
	result = y / x
	fmt.Println(y, "/", x, "equals to:", result)

	//module
	result = x % y
	fmt.Println("module from > ", x, "/", y, "equals to:", result)
	fmt.Println()

	//INcrement
	fmt.Println("Previous value of X is:", x)
	x++
	fmt.Println("X after incrementing is:", x)

	//Decrement
	fmt.Println()
	fmt.Println("Now we are going to decremet X")
	x--
	fmt.Println("NOw x value is:", x)
}
