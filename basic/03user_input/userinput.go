package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Userinput values"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("type of this  rating, %T \n", input)

}
