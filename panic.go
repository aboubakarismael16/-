package main

import "fmt"

func myfunc(s string)  {
	fmt.Println(s)
}

func funA()  {
	fmt.Println("func A .....")
}

func funB()  {

	defer myfunc("defer fun1......")
	for i :=0; i < 10; i++ {
		fmt.Println(i)

		if i == 5 {
			panic("panic $$$$$$$$$$$")
		}
	}
	defer myfunc("defer fun2........")
}

func main() {

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("it is recovered.......")
		}
	}()
	defer myfunc("defer fun3....")
	funA()
	defer myfunc("defer fun4......")
	funB()

	fmt.Println("main func.......")
}
