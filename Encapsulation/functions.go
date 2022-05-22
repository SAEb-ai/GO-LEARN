//Go program to implemnt encapsulation using functions
package main

//Importing the Format Package
import (
	"fmt"
	"strings"
)

//Entry point of Go File
func main() {
	//Perform two tasks
	printName("Md Shahbaz Alam")
	validateEmail("test@gmail.com")
}

// Prints Name
func printName(name string) {
	fmt.Println("My Name is", name)
}

//Validates Email
func validateEmail(email string) {
	//Check whether email is valid or not
	var isEmailValid = strings.Contains(email, "@");
	if isEmailValid {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is invalid")
	}
}