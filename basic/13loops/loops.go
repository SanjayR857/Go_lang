package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday","Tuesday","wednesday","Friday"}

	fmt.Println(days)

	for d :=0; d<len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days{
		fmt.Println(days[i])
	}

	for index, day:= range days{
		fmt.Printf("index is %v and value is %v\n",index,day)
	}

	rangeValue := 1 
    
	for rangeValue <10 {
		if rangeValue==5{
			rangeValue++
			continue
		}

		fmt.Println("values is ",rangeValue)
		rangeValue++
		
	}

}