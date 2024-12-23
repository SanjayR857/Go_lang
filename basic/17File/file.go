package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")
    content := "This needs to go in a file -LearnCode Online.in"
    os.Create("./my")
}
