package main

import "fmt"

func main() {
	var h1 Human
	h1.Person.name = "xiaoming"
	h1.People.name = "小明"
	fmt.Println(h1)
	//h1.show()
	h1.Person.show()
	h1.People.show()

	h1.Person.show2()
	h1.People.show2()
}

type Person struct {
	name string
}

type People = Person

type Human struct {
	Person
	People
}

func (p Person) show() {
	fmt.Printf("Person:%s\n", p.name)
}

func (p People) show2() {
	fmt.Printf("People:%s\n", p.name)

}
