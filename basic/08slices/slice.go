package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var numbers = []string{"1", "2", "3"}
	fmt.Println(numbers)

	numbers = append(numbers, "3", "4")
	fmt.Println(numbers)
	fmt.Println(numbers[1:])
    
	// array
	highsore := make([]int, 4)
	highsore[0] = 234
	highsore[1] = 234
	highsore[2] = 234
	highsore[3] = 234
    
	//to list 
	highsore = append(highsore, 444,555,333,666)
	fmt.Println("highscore ",highsore)
	sort.Ints(highsore)
	fmt.Println(highsore)

	// how to remove a value slices based on index 

	var courese = []string{"reactjs","javascript","swift","python"}
	fmt.Println("courese ",courese)
	courese1 := append(courese[2:])
	var index int = 2
	courese2 := append(courese[:index],courese[index:]...)
	fmt.Println("courese ",courese1)
	fmt.Println("courese ",courese2)
	
}
