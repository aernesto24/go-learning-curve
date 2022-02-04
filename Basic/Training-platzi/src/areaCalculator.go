package main

import (
	"fmt"
	"math"
)

//scan function is part of fmt if you will enter a string you will need bufio

func main() {
	//THis code is to calculate the are of a sqaure, a rectangle, a circle, and a trapezoid

	//Get it the values of the variables from the user
	var base int
	fmt.Println("What is the value of the base? (THis value will also be use for the trapezoid)")
	fmt.Scan(&base)

	var height int
	fmt.Println("What is the value of the height?(THis value will also be use for the trapezoid)")
	fmt.Scan(&height)

	var shortside int
	fmt.Println("What is the value of the shortside of the trapezoid?")
	fmt.Scan(&shortside)

	var radius int
	fmt.Println("What is the value of the radius of your circle?")
	fmt.Scan(&radius)

	//constants
	const pi = 3.14

	//Calculate the area of a SQUARE
	squareArea := base * base
	fmt.Println("The area for a square with base value of:", base, "is:", squareArea)

	//Calculate the area of a Rectangle
	rectangleArea := base * height
	fmt.Println("The area for a square with base value of:", base, "and a height of", height, "is:", rectangleArea)

	//Calculate the area of a trapezoid
	trapezoidArea := ((shortside + base) / 2) * height
	fmt.Println("THe trapezoid area for values corresponding to base, height, short upper side", base, height, shortside, "is: ", trapezoidArea)

	//calculate the area of the circle
	floatRadius := float64(radius) //we need to first convert the int to a float so we can multiply it with another float
	circleArea := pi * math.Pow(floatRadius, 2)
	fmt.Println("THe area of a circle with radius", radius, "is:", circleArea)

}
