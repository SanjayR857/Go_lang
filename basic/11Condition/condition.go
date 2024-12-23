package main

import "fmt"

func main() {

	fmt.Println("Welcome to IF condition")
    loginCount :=10
    var result string
	if loginCount<10{
		result = "Watch out"
	} else if loginCount>10 {
		result =" regular flow"
	} else{
		result = "number is 10"
	}
	fmt.Println(result)

	if 9%2!=0 {
		fmt.Println("its not even number")

	} else {
		fmt.Println("its even number")
	}

	if num :=3; num<10 {
		fmt.Println("Number is less than 10")

	} else {
		fmt.Println("Number is Not less than 10")
	}

}