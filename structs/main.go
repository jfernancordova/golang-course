package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jose := person{
		firstName: "Jose",
		lastName:  "Cordova",
		contactInfo: contactInfo{
			email:   "jfernancordova@gmail.com",
			zipCode: 1426,
		},
	}

	jose.updateName("fernando")
	jose.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
