package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a before changed :", a)
	fun1(a)
	fmt.Println("a after changed : ", a)

	fun2(&a)
	fmt.Println("a after changed : ", a)
}

func fun1(num int)  {
	fmt.Println("num inside fun1 :",num)
	 num = 100
	 fmt.Println("num inside fun1 after changed :", num)
}

func fun2(ptr *int)  {
	fmt.Println("ptr inside fun2 :", *ptr)
	*ptr = 100
	fmt.Println("ptr inside fun2 after changed :", *ptr)
}