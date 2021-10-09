package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add) //func(int, int) int
	fmt.Printf("%T\n", oper) // func(int, int, func(int, int) int) int

	res := add(12,13)
	fmt.Println(res)

	fmt.Println(oper) //0x499340

	res2 := oper(12,13,add)
	fmt.Println(res2)

	res3 := oper(5,3, sub)
	fmt.Println(res3 )

	func3 := oper(12,4, func(a int, b int) int {
		return a * b
	})

	fmt.Println(func3)

	func4 := func(a,b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	}

	res4 := oper(12,3,func4)
	fmt.Println(res4)
}

func add(a,b int) int  {
	return a + b
}

func sub(a,b int) int  {
	return a - b
}

func oper(a,b int, fun func(int, int)int) int  {
	fmt.Println(a,b, fun)
	res := fun(a,b)
	return res
}