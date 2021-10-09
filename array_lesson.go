package main

import "fmt"

func main()  {
	var arr [4] int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4


	fmt.Println(arr, "cap =",cap(arr),"len =", len(arr))

	var a = [...]int{1:3,3:4}
	fmt.Println(a)

	var b = []int{1,2,3,4}
	fmt.Println(b,len(b), cap(b))

	b[3] = 12
	fmt.Println(b)

	var c = []int{'A','B','C','D'}
	fmt.Printf("%T, %d len = %d,cap = %d \n",c, c, len(c), cap(c))

	fmt.Println("--------------")
	f := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(f); i++ {
		fmt.Printf("%d th element of a is %.2f\n",i , f[i])
	}

	fmt.Println("--------------")

	sum := float64(0)
	for _,v := range f {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println("--------------")
	arr3 := [5]int{8,4,23,15,7}

	for i := 0 ; i < len(arr3); i++ {
		for j := 0; j < len(arr3)-1; j++ {
			if arr3[j] > arr3[j+1] {
				arr3[j] ,arr3[j+1] = arr3[j+1], arr3[j]
			}
		}
	}
	fmt.Println(arr3)
	fmt.Println("-------------")

	arr4 := [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}

	fmt.Printf("er wei shu zu : %p len = %d\n", &arr4, len(arr4))
	fmt.Printf("yi wei shu zu : %p len = %d\n", &arr4[0], len(arr4[0]))

	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Print(arr4[i][j],"\t")
		}

		fmt.Println()
	}


}
