package main

import "fmt"

// type Dog struct {
// 	Animal
// }
// type Animal struct {
// 	Age int
// }

// func (a *Animal) Move() {
// 	fmt.Println("Animal moved")
// }
// func (a *Animal) SayAge() {
// 	fmt.Printf("Animal age: %d\n", a.Age)
// }
// func main() {
// 	d := Dog{}
// 	d.Age = 3
// 	d.Move()
// 	d.SayAge()
// }

// type Person struct{
//   firstName string
//   lastName string
// }

// func (p Person) name() string {
//   return p.firstName + " " + p.lastName
// }
// func main(){
//   p := Person{"Hendy", "Good Person"}
//   fmt.Println(p.name())
// }

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

func main() {
	author1 := author{
		"Hendy",
		"Nurfrianto",
		"Golang Programing",
	}
	send := post{
		"Golang is Happy",
		"Inheritance In Golang",
		author1,
	}
	send.details()
}
