package main

import "fmt"

func main() {
	res1 := increment()
	fmt.Println(res1)
	r1 := res1()
	fmt.Println(r1)
	r2 := res1()
	fmt.Println(r2)

	res2 := increment()
	fmt.Println(res2)
	r3 := res2()
	fmt.Println(r3)
	fmt.Println(res2())
	fmt.Println(res2())
	fmt.Println(res2())
	fmt.Println(res2())

}

func increment() func()int  {
	i := 0
	func1 := func () int{
		i++
		return i
	}

	return func1
}
