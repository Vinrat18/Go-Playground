package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	lastName  string
	firstName string
	contact   contactInfo
}

type personMap = map[string]person

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func printMap(c personMap) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

func main() {
	// alex := person{"Kumar", "Vinaya"}
	// alex = person{lastName: "Kumar", firstName: "Vinaya"}

	var p person

	p.firstName = "Vinay"
	p.lastName = "Kumar"
	p.contact = contactInfo{
		email: "v@jvk.com",
		zip:   563114,
	}

	p.updateName("Vinaya")

	persons := make(personMap)
	persons[p.firstName] = p

	printMap(persons)
	// delete(colors, 10)
}
