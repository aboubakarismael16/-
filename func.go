package main

import "fmt"

func main() {

	res := getSum(100)
	fmt.Println("res = ",res)
	fmt.Println("--------------------")

	getSum2(1,2,4,5,6,7,8,9)
	s1 := []int{10,20,30,40}
	getSum2(s1...)
	fmt.Println("---------------------")
	arr1 := [4]int{1,2,3,4}
	fmt.Println("array in main func : ",arr1)
	func1(arr1)
	fmt.Println("array in main func after modif :", arr1)
	fmt.Println("---------------------")
	s2 := []int{1,2,3,4}
	fmt.Println("slice in main func : ", s2)
	func2(s2)
	fmt.Println("slice s2 in main func after modif :",s2)
	fmt.Println("------------------------------")
	res1, res2 := rectangle(5,3)
	fmt.Println("perimeter :", res1, "area :", res2)
	res3, re4 := rectangle2(5,3)
	fmt.Println("perimeter :", res3, "area :", re4)
	fmt.Println("------------------------")
	func3(-30)
	fmt.Println("------------------------")
	func4()
}

func func3(age int) int {
	if age >= 0 {
		fmt.Println(age)
		return age
	} else {
		fmt.Println("age cannot be under 0 :",-1)
		return -1
	}
}

func func4()  {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			//return
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("func4 over.......")

}

func rectangle(len, wid float64) (float64, float64)  {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter,area
}

func rectangle2(len, wid float64) (peri , area float64)  {
	peri = (len + wid) * 2
	area = len * wid
	return
}

func func2(s []int)  {
	fmt.Println("slice in func2 : ",s)
	s[0] = 100
	fmt.Println("slice in func2 after modif :", s)
}

func func1(arr [4]int)  {
	fmt.Println("array in func1 :", arr)
	arr[0] = 100
	fmt.Println("array in func after modif :",arr)
}

func getSum2(num ...int)  {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Printf("Sum = %d\n", sum)
}

// use return
func getSum(n int)  int {
	sum := 0
	for i := 1 ; i <= n; i++ {
		sum += i
	}
	return sum
}

