package main

import "fmt"

type Human interface {
	candoit() string
}

type Man struct {
}
type Woman struct {
}

func (m Man) candoit() string {
	return "I'm go to classroom."
}

func (w Woman) candoit() string {
	return "I,m going to job"
}

func main() {
	m := new(Man)
	w := new(Woman)

	canArr := [...]Human{m, w}
	for n, _ := range canArr {
		fmt.Println("I'm a human, and I want to : ", canArr[n].candoit())
	}
}
