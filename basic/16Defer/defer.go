package main
//defer will execute at last 
import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("First")
	defer fmt.Println("Hello")
	defer fmt.Println("2")
	myDefer()
}

func myDefer(){
	for i := 0 ; i<5 ; i++ {
		defer  fmt.Println(i)
	}
}