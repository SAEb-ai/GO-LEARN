//Go program to check whether email is valid or not.
package main

//Importing the Format Package
import (
	"fmt"
	"strings"
)

//Entry point of Go File
func main() {

	var email = "test@gmail.com"
	//Check whether email is valid or not
	var isEmailValid = strings.Contains(email, "@");
	if isEmailValid {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is invalid")
	}
	
}

