package main

import (
	"fmt"
)

type cat struct {
	name   string
	weight int
	age    int
}

func main() {
	cat1 := cat{
		name:   "ghost",
		weight: 14,
		age:    2,
	}

	cat2 := cat{
		name:   "nymeria",
		weight: 15,
		age:    3,
	}

	// without pointer. nothing gonna change.
	cat1.updateName()
	fmt.Println("without pointer => ", cat1)

	// use pointer #1 way
	catPointer := &cat1
	catPointer.updateAge()
	fmt.Println("using pointer =>", *catPointer)

	// use pointer #2 way
	cat2.updateAge()
	fmt.Println("using pointer =>", cat2, "\n")

	// ================== now using slice ================== //

	mySlice := []string{"hey", "how", "are", "you"}
	changeSlice(mySlice)
	fmt.Println(mySlice) // slice doesnt need pointer

}

// this function cannot update cat's name, because the function not pass pointer as argument
func (c cat) updateName() {
	c.name = "summer"
}

// and this is it
func (c *cat) updateAge() {
	(*c).age = (*c).age + 2
}

func changeSlice(s []string) {
	s[0] = "bro"
}
