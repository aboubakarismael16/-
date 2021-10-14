package main

import (
	"fmt"
	"strconv"
)
//1.
type myint int
type mystring string

//2.
type myfun func(a,b int)string

func fun3()myfun  {
	fun := func(a, b int) string {
		return strconv.Itoa(a)+strconv.Itoa(b)
	}
	return fun
}

type myint2 = int

func main() {
	var i1 myint
	i1 = 100
	i2 := 200
	fmt.Println(i1,i2)

	var name mystring
	name = "ismael"
	name2 := "youssouf"
	fmt.Println(name, name2)

	//i2 =i1  //cannot use i1 (type myint) as type int in assignment

	//name2 =name // cannot use name (type mystring) as type string in assignment

	fmt.Printf("%T, %T, %T\n", i1,name,name2) //main.myint, main.mystring, string

	fmt.Println("------------------")
	res := fun3()
	fmt.Println(res(10,20))
	fmt.Println("------------------")
	var i3 myint2
	i3 = 1000
	fmt.Println(i3)
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T,%T,%T\n",i1,i2,i3)  //main.myint,int,int
}
