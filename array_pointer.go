package main

import "fmt"

func main() {
	a1,b,c,d := 1,2,3,4
	arr1 := [4]int{a1,b,c,d}
	arr2 := []*int{&a1,&b,&c,&d}
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1[0] = 100
	fmt.Println(arr1)
	fmt.Println(a1)

	*arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a1)

	b = 1000
	fmt.Println(b)
	fmt.Println(arr1)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(*arr2[i])
	}

}
