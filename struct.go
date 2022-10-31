package main

import "fmt"

type Person struct {
	Name	string
	Address	string
	Age 	int
	Maried	bool
}

func (person* Person) sayHello() {
	fmt.Printf("%p\n", person)
	fmt.Println("Hello, my name is", person.Name)
}

func sayHi(person* Person) {
	fmt.Printf("%p\n", person)
	fmt.Println("Hi, my name is", person.Name)
}

func main() {
	var rahmat Person;
	rahmat.Name = "Rahmat"
	rahmat.Address = "Jln KPAD Sriwijaya"
	rahmat.Age = 22
	fmt.Printf("%p\n", &rahmat)

	// dani := Person{
	// 	Name: "Dani Eka",
	// 	Address: "Bandung",
	// 	Age: 32,
	// }
	// fmt.Println(dani)

	rahmat.sayHello()
	sayHi(&rahmat)
}