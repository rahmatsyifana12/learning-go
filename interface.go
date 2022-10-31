package main

import "fmt"

type HasPerson interface {
	GetName() string
	GetAddress() string
}

func SayHello(hasPerson HasPerson) {
	fmt.Println("Hello", hasPerson.GetName())
}

type Person struct {
	Name string
	Address string
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAddress() string {
	return person.Address
}

func main() {
	rahmat := Person{
		Name: "Rahmat",
		Address: "KPAD Sriwijaya",
	}

	SayHello(rahmat)
}