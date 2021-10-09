package main

import "fmt"

func main() {
	/*
		for loop
	*/

	//for i := 100 ; i < 1000; i++ {
	//	x := i / 100
	//	y := i /10 % 10
	//	z := i % 10
	//	if math.Pow(float64(x), 3) + math.Pow(float64(y), 3) + math.Pow(float64(z), 3) == float64(i) {
	//		fmt.Println(i)
	//	}
	//}

	for i := 2; i <100; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}

		if flag {
			fmt.Println(i)
		}
	}

	fmt.Println("----------------")
	var a = 10
	Loop :
		for a < 20 {
			if a == 15 {
				a += 1
				goto Loop
			}
			fmt.Println(a)
			a++
		}


}
