package main

import "fmt"

func main()  {
	sum := getSumRec(5)
	fmt.Println(sum)
	fmt.Println(getFib(12))
}

func getFib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFib(n-1) + getFib(n-2)
}

func getSumRec(n int) int {
	//fmt.Println("*********")
	if n == 1 {
		return 1
	}
	return getSumRec(n-1) + n
}
