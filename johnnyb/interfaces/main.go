package main

import "fmt"

type person struct {
	name    string
	address string
	city    string
	state   string
	zip     string
}

type people []person

func main() {
	peeps := people{
		{name: "John", address: "123 Main St", city: "Jamestown", state: "NY", zip: "14701"},
		{name: "Jane", address: "234 Fleet St", city: "Columbia", state: "MD", zip: "21150"},
		{name: "Terry", address: "345 Charles Ave", city: "Gergetown", state: "DC", zip: "20007"},
	}

	fmt.Println(peeps.String())
}

func (p people) String() string {
	var str string
	for _, person := range p {
		str += fmt.Sprintf("%s | %s, %s, %s %s\n", person.name, person.address, person.city, person.state, person.zip)
	}
	return str
}
