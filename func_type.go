package main

import "fmt"

func main() {
	var c func(int,int)
	c = funType
	fmt.Println(c) // 0x4976e0
	fmt.Printf("%T\n",c) //func(int, int)
	fmt.Println(funType) //0x4992a0
	fmt.Println("------------------")
	funType(10, 20)
	c(100, 200)
	fmt.Println("------------------")
	res1 := funType2
	res2 := funType2(12,13)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res1(1,2))
	//res2()  //cannot call non-function res2 (type int)

}

func funType(a, b int)  {
	fmt.Println("a and b value :", a, b)
}

func funType2(a,b int) int {
	return a+b
}
