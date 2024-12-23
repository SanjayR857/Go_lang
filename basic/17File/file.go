package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")
	fmt.Println("Creating a file")
    content := "This needs to go in a file -LearnCode Online.in"
    os.Create("./my")
}
