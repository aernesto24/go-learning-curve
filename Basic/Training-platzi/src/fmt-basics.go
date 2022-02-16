//This code will show us more functionalities of the fmt command

package main

import "fmt"

func main() {
	//Declaring variables
	helloMessage := "Hello"
	worldMessage := "world"

	//println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf to add a function
	name := "Ernesto"
	certifications := 20

	fmt.Printf("%s tiene mas de %d certificacones\n", name, certifications)
	//in case we do not know what type of variable we are inserting in the print we can use %v
	fmt.Printf("%v tiene mas de %v certificacones\n", name, certifications)

	//Sprintf > does not print into dfe conse but you can use it to assing it to a variable
	certMessage := fmt.Sprintf("%s tiene mas de %d certificaciones", name, certifications)
	fmt.Println(certMessage)

	//print thre type of data
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("certifications: %T \n", certifications)
}
