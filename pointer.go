package main

import "fmt"

func main() {
	a := 10 /* 声明实际变量 */
	fmt.Println(a)
	fmt.Println("address of integer :", &a) //0xc000020098

	var p1 *int                               /* 声明指针变量 */
	p1 = &a                                   /* 指针变量的存储地址 */
	/* 指针变量的存储地址 */
	fmt.Println("address of pointer :", p1)   //0xc000020098
	fmt.Println("address of p itself :", &p1) //0xc00000e030
	/* 使用指针访问值 */
	fmt.Println(*p1)

	a = 100
	fmt.Println(&a)
	fmt.Println(a)

	*p1 = 200
	fmt.Println(a)
	fmt.Println("-----------------")

	var p2 **int
	fmt.Printf("%T, %T,%T\n", a, p1, p2) //int, *int,**int
	p2 = &p1
	fmt.Println(&a, p1, *p2) // 0xc000020098 0xc000020098 0xc000020098
	fmt.Println(**p2)        // 200

	fmt.Println("-----------------")
	arr := [4]int{1,2,3,4}

	var p3 *[4]int
	p3 = &arr
	fmt.Printf("%T,%T\n",arr,p3)
	fmt.Printf("%d\n",arr)
	fmt.Println(&p3)
	fmt.Println(p3)  // &[1 2 3 4]

	(*p3)[0] = 100
	fmt.Println(arr)

	p3[0] = 200
	fmt.Println(arr)

}
