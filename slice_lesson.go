package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[2:4]
	fmt.Printf("%d len = %d, cap = %d ,%p\n", a, len(a), cap(a), &a)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s, len(s), cap(s), s)
	fmt.Println("-----------------------------------------")
	s1 := s
	s1[0] = 100
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s, len(s), cap(s), s)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", a, len(a), cap(a), &a)


	fmt.Println("-----------------------------------------")
	s1 = append(s1, 7,8,9)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s, len(s), cap(s), s)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", a, len(a), cap(a), &a)

	s2 := []int{12,13,14,15,16,17,18,19,20}
	s3 := make([]int, 0)
	for i := 0; i < len(s2); i++ {
		s3 = append(s3, s2[i])
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s2, len(s2), cap(s2), s2)
	fmt.Printf("%d len = %d, cap = %d ,%p\n", s3, len(s3), cap(s3), s3)


}
