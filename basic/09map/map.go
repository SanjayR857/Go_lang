package main

import "fmt"

func main() {
	fmt.Println("welcome to map coding")

	languages := make(map[string]string)
	languages["1"] = "English"
	languages["2"] = "Kannada"
	languages["3"] = "Hindi"

	fmt.Println("lanuages : ",languages)

	// delecte 
	delete(languages,"1")
	fmt.Println("lanuages : ",languages)

	// loops are intersting in golang 

	for key,value := range languages{
		println(key,value)
	}
	
}