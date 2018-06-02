package main

import (
	"fmt"
)

// Phone struct
type Phone struct {
	countryCode int
	phoneNo     int
}

// Person type
type Person struct {
	firstName string
	lastName  string
	age       int
	phone     Phone
}

func main() {
	var alex Person
	alex.firstName = "Alex"
	alex.lastName = "Hunter"
	alex.age = 28
	alex.phone.countryCode = 91
	alex.phone.phoneNo = 9740288160
	fmt.Println(alex)

	elon := Person{
		firstName: "Elon",
		lastName:  "Musk",
		age:       45,
		phone: Phone{
			countryCode: 91,
			phoneNo:     9740288161,
		},
	}
	fmt.Println(elon)
}
