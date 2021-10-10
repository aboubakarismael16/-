package main

import "fmt"

func main() {

	arr1 := fun1()
	fmt.Printf("type : %T, add : %p, Value : %v\n", arr1,&arr1,arr1)

	arr2 := fun2()
	fmt.Printf("type : %T, add : %p, Value : %v\n", arr2,&arr2,arr2)

}

func fun1() [4]int {
	arr := [4]int{1,2,3,4}
	return arr
}

func fun2() *[4]int  {
	arr := [4]int{5,6,7,8}
	return &arr
}