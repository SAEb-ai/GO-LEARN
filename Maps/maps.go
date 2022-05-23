//Go program to showcase implementation of maps in Go
package main

//Importing the Format Package
import "fmt"

//Making a empty list of names of my friends
var friendList = make([]map[string]string, 0)

//Entry point of Go File
func main() {

	//adding names
	addName("Md", "Shahbaz", "Alam")
	addName("Mr", "Mintu", "Burnwal")
	addName("Mr", "Mayur", "Rai")
	addName("Mr", "Sahil", "Soni")
	//displaying my friend's names
	displayFriendList()

}

//Function to add a friend name
func addName(firstName string, middleName string, lastName string) {

	//Create an empty map
	var name = make(map[string]string)
	//Add values to the map
	name["firstname"] = firstName
	name["middlename"] = middleName
	name["lastname"] = lastName
	friendList = append(friendList, name)

}

//Function to display the names of all of my firends
func displayFriendList() {
	for _, name:= range friendList {
		fmt.Println(name["firstname"], name["middlename"], name["lastname"])
	}
}