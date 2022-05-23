//Go program to showcase implementation of Structures in Go
package main

//Importing the Format Package
import "fmt"
 
var friendList = make([]Name, 0)

//Defines a structure
type Name struct {
	firstName string
	middleName string
	lastName string
	age uint
}

//Entry point of Go File
func main() {

	//adding names
	addName("Md", "Shahbaz", "Alam", 21)
	addName("Mr", "Mintu", "Burnwal", 22)
	addName("Mr", "Mayur", "Rai", 22)
	addName("Mr", "Sahil", "Soni", 22)
	//displaying my friend's names
	displayFriendList()

}

//Function to add a friend name
func addName(firstName string, middleName string, lastName string, age uint) {

	var name = Name {
		firstName: firstName,
		middleName: middleName,
		lastName: lastName,
		age: age,
	}

	//Adding friend's names
	friendList = append(friendList, name)

}

//Function to display the names of all of my firends
func displayFriendList() {
	for _, name:= range friendList {
		fmt.Println(name.firstName, name.middleName, name.lastName)
	}
}