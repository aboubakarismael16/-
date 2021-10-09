package main

import "fmt"

func main() {

	/*
	你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
	特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

		- 如果有很多调用defer，那么defer是采用`后进先出`模式
		- 在离开所在的方法时，执行（报错的时候也会执行）
	 */

	a := 2
	fmt.Println(a)
	defer funcDefer(a)
	a++
	fmt.Println("the final value of a : ",a)
	fmt.Println("-------------------------")
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}

func funcDefer(a int)  {
	fmt.Println("a in defer :",a)
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Starting find Largest number")
	fmt.Println("Largest number in", nums, "is", max)
}
