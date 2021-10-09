package main

import "fmt"

func main() {
	/*
		Unamed func can only call one time

			func2 := func (){
			fmt.Println("it is also unamed func")
		}
		func2() // that one can call all the time
	*/

	func (){
		fmt.Println("it is unamed func")
	}()

	func2 := func (){
		fmt.Println("it is also unamed func")
	}
	func2()

	func (a, b int) {
		fmt.Println(a,b)
	}(1,2)

	res := func (a,b int) int {
		return a+b
	}(10, 20)  // this return the parameter
	fmt.Println(res)


	res2 := func(a,b int) int {
		return a+b
	} // return the address
	fmt.Println(res2)  // 0x4978a0
}
