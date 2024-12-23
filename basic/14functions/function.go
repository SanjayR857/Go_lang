package main

import "fmt"

func main() {

	fmt.Println("Welcome to the functions in go lang")
	greeter()
	geeter2()
	result := adder(2,3)
	fmt.Println("result of sum ",result)
	ProResult,_ := proAdder(1,2,3,4,5)
	fmt.Println("Pro Result ",ProResult)
}

func greeter(){
	fmt.Println("Namstey from golang")
}

func geeter2(){
	fmt.Println("Welcome to golang")
}

func adder(x int, y int) int {
	return x+y
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values{
		total += val
	}
	return total, "Hi Pro result"
}