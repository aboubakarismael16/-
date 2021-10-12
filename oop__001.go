package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	school string
}

func (p Person) eat()  {
	fmt.Println(p.name," it's eating Pizza")
}

func (s Student) eat()  {
	fmt.Println(s.name,"it's eating Kebab")
}
func main() {
	/*
		method是可以继承的，如果匿名字段实现了一个method，
		那么包含这个匿名字段的struct也能调用该method

		存在继承关系时，按照就近原则，进行调用
	 */

	s1 := Student{
		Person {
			name: "ismael",
			age: 12,
		},
		"Nepu",
	}
	fmt.Println(s1)

	var s2 Student
	s2.name = "Issa"
	s2.age = 14
	s2.school = "Tsinghua"

	fmt.Println(s2)
	fmt.Println(s1.name,s1.age,s1.school)
	fmt.Println(s2.name,s2.age,s2.school)
	fmt.Println("-------------------")
	p1 := Person{
		name: "youssouf",
		age: 35,
	}
	p1.eat()
	s2.eat()   // Issa it's eating Kebab
	s1.eat()

}


