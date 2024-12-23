package main

import "fmt"

func main() {
	fmt.Println("Welcome to Sructs")
	// no inheritance in golang; no super or parent
    
	hitesh := User{"Sanjay","sanjay@gmail.com",true,23}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details %+v\n",hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
