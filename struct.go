package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	school string
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func (p person) sayHi() string {
	return "say hi"
}

func (s student) sayHi() string {
	return "student say hi"
}

func (s student) sing(song string) string {
	return "I am sing:" + song
}

type Menu interface {
	sayHi()
	sing(song string)
}

func main() {
	var P person

	P.name = "lisi"
	P.age = 20

	P1 := person{"zhang", 30}

	fmt.Println("p is name:", P.name)
	fmt.Println("p1 is name:", P1.name)

	p, i := older(P, P1)
	fmt.Println("p,i", p.name, i)

	S := student{person{name: "lyb", age: 20}, "小学"}

	fmt.Println(S.school)
	fmt.Println(S.person.name)
	fmt.Println(S.sing("夜来香"))

	fmt.Println(S.sayHi())

	fmt.Println(P.sayHi())

}
