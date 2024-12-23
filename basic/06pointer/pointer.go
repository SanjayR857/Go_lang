package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int 

    fmt.Println("value of pointer is ",ptr)

	myNumber := 23
	// reference means &
	var ptr1 = &myNumber
    
	fmt.Println("value of acutal pointer is ", ptr1)
	fmt.Println("value of acutal pointer  is ", *ptr1)

	*ptr1 = *ptr1 + 2 
	fmt.Println("New value is: ", myNumber)

}