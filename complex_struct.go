package main

import (
	"fmt"
)

type Book struct {
	bookName string
	price float64
}

type student struct {
	name string
	age int
	book Book
}

type student2 struct {
	name string
	age int
	book *Book
}

func main() {
	b1 := Book{
		bookName: "Golang for beginner",
		price: 14.5,
	}

	s1 := student{}
	s1.name = "Ali"
	s1.age = 23
	s1.book = b1
	fmt.Printf("Name : %s, age : %d, Book name : %s, price : %.2f\n", s1.name,s1.age,s1.book.bookName,s1.book.price)
	b1.bookName = "GO in action"
	fmt.Println(b1.bookName,b1.price)
	fmt.Printf("Name : %s, age : %d, Book name : %s, price : %.2f\n", s1.name,s1.age,s1.book.bookName,s1.book.price)



	s2 := student{
		name: "Issa",
		age: 23,
		book: Book{
			bookName: "Python",
			price: 35.6,
		},
	}
	fmt.Printf("Name : %s, age : %d, Book name : %s, price : %.2f\n", s2.name,s2.age,s2.book.bookName,s2.book.price)
	fmt.Println("------------------------------------")

	b2 := Book{
		bookName: "Java script",
		price: 10.10,
	}

	s3 := student2{
		name: "Ismael",
		age: 21,
		book: &b2,
	}
	fmt.Println(b2)
	fmt.Println(s3)
	fmt.Printf("%p\n", &b2)
	fmt.Printf("Name : %s, age : %d, Book name : %s, price : %.2f\n", s3.name,s3.age,s3.book.bookName,s3.book.price)
	b2.bookName = "Docker for you"
	fmt.Println(b2)
	fmt.Printf("Name : %s, age : %d, Book name : %s, price : %.2f\n", s3.name,s3.age,s3.book.bookName,s3.book.price)

}
