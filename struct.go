package main

import "fmt"


type Person struct {
	name string
	age int
	sex string
	address string
}


func main() {
	var p1 Person
	fmt.Println(p1)

	p1.name = "Ismael"
	p1.age = 20
	p1.sex = "male"
	p1.address = "Daqing"
	fmt.Printf("name : %s, age : %d, sex : %s, address : %s\n", p1.name,p1.age,p1.sex,p1.address)

	p2 := Person{
		name: "Nepu",
		age: 30,
		sex: "unknow",
		address: "Daqing",
	}
	fmt.Printf("name : %s, age : %d, sex : %s, address : %s\n", p2.name,p2.age,p2.sex,p2.address)

	p4 := Person{"aboubakar",21,"male","Beijing"}
	fmt.Println(p4)
	fmt.Println("-------------------------")
	 p5 := p1
	 fmt.Printf("%p, %T\n",&p1,p1)   // 0xc0000b6040, main.Person
	 fmt.Printf("%p, %T\n", &p5,p1)  // 0xc0000b6100, main.Person
	 fmt.Println(p5)
	 p5.name = "youssouf"
	 fmt.Println(p5)
	 fmt.Println(p1)
	 fmt.Println("-----------------------")
	 var pptr *Person
	 pptr = &p1
	 fmt.Printf("%p, %T\n", &p1,p1)   // 0xc0000b6100, main.Person
	 fmt.Printf("%p, %T\n",pptr,pptr)  // 0xc000024080, *main.Person
	 fmt.Println(p1)
	 fmt.Println(*pptr)
	 pptr.name = "youssouf"
	 fmt.Println(pptr)
	 fmt.Println(p1)
	 fmt.Println("------------------")
	 pptr2 := new(Person)
	 fmt.Printf("%T\n",pptr2)  //*main.Person
	 fmt.Println(pptr2)
	 pptr2.name  = "houde"
	 pptr2.age = 13
	 pptr2.sex = "Unknow"
	 pptr2.address = "Nepu"
	 fmt.Println(*pptr2)
}

