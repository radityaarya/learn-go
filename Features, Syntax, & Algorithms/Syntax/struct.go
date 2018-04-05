package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   // to embed a struct to another struct
}

type contact struct {
	phone   int
	zipCode int
}

func main() {
	// declare person #1
	radit := person{"raditya", "arya", contact{8123456, 12345}}

	// declare person #2
	john := person{
		lastName:  "doe",
		firstName: "john",
		contact: contact{
			phone:   8756742,
			zipCode: 12345,
		},
	}

	// declare person #3
	var jane person
	jane.firstName = "jane"
	jane.lastName = "doe"
	jane.contact.phone = 8734542
	jane.contact.zipCode = 54345

	// update struct value
	jane.lastName = "roe"
	radit.lastName = "pradipta"

	fmt.Println(radit, john, jane)
}
