package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")
	fmt.Println("Creating a file")
    content := "This needs to go in a file -LearnCode Online.in"
    
	os.Create("./mylcogofile.txt")
	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)
	length,err :=io.WriteString(file,content)

	checkNilErr(err)
		fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string){
	file, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Content of the file is: ", string(file))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}