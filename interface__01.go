package main

import (
	"fmt"
	"math"
)

type Shape interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a,b,c float64
}

type Circle struct {
	radius float64
}

func (t Triangle) peri() float64  {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64  {
	p := t.peri() /2
	s := math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return s
}

func (c Circle) peri() float64  {
	return (c.radius * 2) * math.Pi
}

func (c Circle) area() float64  {
	return math.Pow(c.radius,2)*2*math.Pi
}

func TestShape(s Shape)  {
	fmt.Printf("peri : %.3f, area : %.3f\n", s.peri(),s.area())
}

func getShape(s Shape)  {
	if ins,ok := s.(Triangle); ok {
		fmt.Println("the shape is a triangle", ins.a,ins.b,ins.c)
	} else if ins,ok := s.(Circle); ok {
		fmt.Println("the shape is a circle", ins.radius)
	} else if ins,ok := (s).(*Triangle); ok {
		fmt.Printf("%T,%p %p\n",ins, &ins, ins)
		fmt.Println("the pointer shape is a triangle", ins.a,ins.b,ins.c)
	} else {
		fmt.Println("the shape does not exist")
	}
}

func getShape2(s Shape)  {
	switch instance := s.(type) {
	case Triangle:
		fmt.Println("triangle", instance.a, instance.b,instance.c)
	case Circle:
		fmt.Println("circle", instance.radius)
	case *Triangle:
		fmt.Println("pointer shape")
	default:
		fmt.Println("I dont know")

	}
}

func main() {

	t := Triangle{5,3,5}
	fmt.Println(t.peri())
	fmt.Println(t.area())
	fmt.Println(t.a,t.b,t.c)

	fmt.Println("-----------------")
	c := Circle{2}
	fmt.Println(c.peri())
	fmt.Println(c.area())
	fmt.Println(c.radius)

	fmt.Println("---------------")

	var s1 Shape
	s1 = t
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c
	fmt.Println(s2.peri())
	fmt.Println(s2.area())
	fmt.Println("----------------")

	TestShape(t)
	TestShape(s1)
	TestShape(s2)
	TestShape(c)
	fmt.Println("-----------------")
	getShape(t)
	getShape(s1)
	getShape(c)
	getShape(s2)
	fmt.Println("------------------------")
	t2 := &Triangle{}
	t2.a = 3
	t2.b = 4
	t2.c = 2
	getShape(t2)
	getShape2(t2)


}
