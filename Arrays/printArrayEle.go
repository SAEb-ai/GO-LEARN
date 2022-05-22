//Go program to print the elements present in an array
package main

//Importing the Format Package
import "fmt"

//Entry point of Go File
func main() {
	
	//Declaring an array which will store my friend's names.
	var friendGroup [50]string
	//Adding names to the array
	friendGroup[0] = "Shahbaz"
	friendGroup[1] = "Mintu"
	friendGroup[2] = "Mayur"
	friendGroup[3] = "Sahil"
	//Printing names stored in friendGroup
	fmt.Println("Our friend group:", friendGroup);

}