package main

import "fmt"

// var start with captile means its public 
const LoginToken string = "sdasdsdsad"

func main() {
	var username string = "Sanjay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallvalue uint8 =255
	fmt.Println(smallvalue)
	fmt.Printf("Variable is of type: %T \n", smallvalue)

	var numberValue int = 234
	fmt.Println(numberValue)
	fmt.Printf("Variable is of type: %T \n", numberValue)

	var floatvalue  float32 = 34.343
	fmt.Println(floatvalue)
	fmt.Printf("Varable is of type: %T \n", floatvalue)

	//defaulpt values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Varable is of type: %T \n",  anotherVariable)

	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("Varable is of type: %T \n",  anotherVariable1)

	//another away to create variable no var style 
	string1 :=  "String"
	println(string1) 
}
