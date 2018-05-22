package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

type customer struct {
	id   string
	name string
}

type cat struct {
	name  string
	breed string
}

type naming interface {
	getName() string
}

func (p person) getName() string {
	return fmt.Sprintf("His name is %s %s\n", p.firstName, p.lastName)
}

func (c cat) getName() string {
	return fmt.Sprintf("This is %s. It's a %s\n", c.name, c.breed)
}

func (cs customer) getName() string {
	return fmt.Sprintf("Customer details => id: %s, name: %s\n", cs.id, cs.name)
}

func printName(n naming) {
	fmt.Println(n.getName())
}

func main() {
	p1 := person{"Raditya", "Arya"}
	printName(p1)

	c1 := cat{"Tom", "Sphynx"}
	printName(c1)

	cs1 := customer{"00001", "Jon"}
	printName(cs1)
}
