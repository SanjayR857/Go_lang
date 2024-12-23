package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch and case in golang ")

	rand.Seed(time.Now().UnixNano())
    diceNUmber := rand.Intn(6)+1
	fmt.Println("Value of dice is ", diceNUmber)

	// fallthrough it will execute next one also 
	switch diceNUmber{
	case 1:
		fmt.Println("dice value is 2 and you can open ")
		fallthrough
	case 2:
		fmt.Println("dice value 2 spot")
		fallthrough
	case 3:
		fmt.Println("dice value 3 spot")
		fallthrough
	default:
		fmt.Println("What's that")
	}

}
