package main

import "fmt"

type Stack []string

func (s *Stack) Push(str string)  {
	*s = append(*s, str)
}

func (s *Stack) isEmpty() bool  {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool)  {
	if s.isEmpty()  {
		return "", false
	} else {
		index := len(*s) - 1
		elemen := (*s)[index]
		*s = (*s)[:index]
		return elemen, true
	}
}

func main()  {
	var stack Stack

	stack.Push("ismael")
	stack.Push("nepu")
	stack.Push("mei er")

	fmt.Println(stack)
	fmt.Println("------------")

	for !stack.isEmpty() {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}

}



