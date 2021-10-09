package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := "hello"

	fmt.Println(len(s1))

	a := "h"
	b := 104
	fmt.Printf("%c,%s,%c\n",s1[0],a,b)

	for _,v := range s1 {
		//fmt.Println(v)
		fmt.Printf("%c\n", v)
	}

	fmt.Println("----------------------")
	slice := []byte{65,66,67,68}
	s2 := string(slice)

	fmt.Println(s2)

	s3 := "abcde"
	slice1 := []byte(s3)
	fmt.Println(slice1)
	fmt.Println("-------------------")

	s4 := "hello nepu"
	fmt.Println(strings.Contains(s4,"advd"))

	fmt.Println("----------------------")

	s5 := "false"
	b1, err := strconv.ParseBool(s5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1,b1)

	ss5 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n",ss5,ss5)

	s6 := "100"
	i2, err := strconv.ParseInt(s6,2,64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2,i2)

	ss6 := strconv.FormatInt(i2,20)
	fmt.Printf("%T,%s\n", ss6,ss6)

	i3, err := strconv.Atoi("-42")
	fmt.Printf("%T,%d\n", i3,i3)

	s7 := strconv.Itoa(-42)
	fmt.Printf("%T, %s\n", s7,s7)
}
