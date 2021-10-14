package main

import (
	"fmt"
	"github.com/andrewkroh/beats/x-pack/elastic-agent/pkg/agent/errors"
)

func main() {
	//f,err := os.Open("test.txt")
	//if err != nil {
	//	fmt.Println(err)   //open test.txt: no such file or directory
	//	return
	//}
	//fmt.Println(f.Name(), "exist....")

	err := checkAge(-13)
	if err != nil {
		fmt.Println(err)  // age can not be < than 0: /home/aboubakar/Documents/归档/error.go[25]: unknown error
		return
	}
}

func checkAge(age int) error {
	if age < 0 {
		return errors.New("age can not be < than 0")
	}
	fmt.Println(age,"it is ok...")
	return nil
}
