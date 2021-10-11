package main

import (
	"fmt"
)

type Student struct {
	name string
	age int
}

type Worker struct {
	string //  `string`  can not repeat some type
	int

}

func main() {
	s1 := Student{
		name: "Ali",
		age: 23,
	}
	fmt.Printf("name : %s, age : %d\n",s1.name,s1.age)

	s2 := struct {
			name string
			age int
		}{
		name: "ismael",
		age: 12,
		}
	fmt.Printf("name : %s, age : %d\n",s2.name,s2.age)
	fmt.Println("---------------------")
	w1 := Worker{
		"wang er gou",
		30,
	}
	fmt.Println(w1)
	fmt.Println(w1.string)
	fmt.Println(w1.int)
}
