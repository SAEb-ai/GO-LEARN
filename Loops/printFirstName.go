//Go program to print the firstnames from the list of names present in the array
package main

//Importing the Format Package
import (
	"fmt"
	"strings"
)

//Entry point of Go File
func main() {
	
	//Declaring an array which will store my friend's names.
	var friendGroup [4]string
	//Declaring an array which will store my friend's firstname.
	var firstName [4]string
	//Adding names to the array
	friendGroup[0] = "Md Shahbaz Alam"
	friendGroup[1] = "Mintu Burnwal"
	friendGroup[2] = "Mayur Rai"
	friendGroup[3] = "Sahil Soni"

	//Getting the firstnames one by one in a for loop
	for index, name:= range friendGroup {
		//Slicing the names based on whitespace present in the names 
		var slicedName=strings.Fields(name)
		firstName[index] = slicedName[0];
	}

	//Printing first names stored in firstName
	fmt.Println("Our friend group:", firstName);

}