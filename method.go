package main

import "fmt"

type Worker struct {
	name string
	age int
	sex string
}

type Cat struct {
	color string
	age int
}

func (w Worker) work()  {
	fmt.Println(w.name, "is working...")
}

func (p *Worker) rest()  {
	fmt.Println(p.name,"is resting...")  // can also  write *p.name
}

func (p Worker) printInfo()  {
	fmt.Printf("worker name : %s, age : %d, sex : %s\n", p.name,p.age,p.sex)
}

func (c *Cat) printInfo() {
	fmt.Printf("Cat color : %s , age ; %d\n", c.color,c.age)
}

func main() {

	/*
	1.  定义方法的语法
		func (t Type) methodName(parameter list)(return list) {

		}
		func funcName(parameter list)(return list){

		}

		虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
		method里面可以访问接收者的字段
		调用method通过.访问，就像struct里面访问字段一样
	*/

	w1 := Worker{
		name: "ismael",
		age: 23,
		sex: "male",
	}

	w1.work()
	var w2 Worker
	w2.name = "mei er"
	w2.age = 24
	w2.sex = "male"

	w2.work()
	w1.rest()
	w2.rest()
	fmt.Println("--------------------")
	w1.printInfo()

	c := Cat{
		color: "white",
		age: 1,
	}
	c.printInfo()
}
