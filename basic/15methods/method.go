package main

import "fmt"

func main() {
	fmt.Println("Welcome to Sructs")
	// no inheritance in golang; no super or parent
    
	hitesh := User{"Sanjay","sanjay@gmail.com",true,23}
	fmt.Println(hitesh)
	hitesh.NewMail()
	fmt.Printf("Hitesh details %+v\n",hitesh)
	hitesh.GetStatus()
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("IS user active ",u.Status)
}

func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("email of this user is ", u.Email)
}
