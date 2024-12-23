package main

import "fmt"

func main() {
	fmt.Println("Welcome to arry in golang")
    
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Bananna"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato","benas","mushroom"}
	print(vegList)
	print(len(vegList))
}
