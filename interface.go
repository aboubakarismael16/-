package main

import "fmt"

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlaskDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "mouse start working")
}

func (m Mouse) end()  {
	fmt.Println(m.name,"mouse finished working...")
}

func (f FlaskDisk) start()  {
	fmt.Println(f.name, "flaskDisk start working")
}

func (f FlaskDisk) end()  {
	fmt.Println(f.name, "flaskDisk finished working....")
}

func TestInfo(usb USB)  {
	usb.start()
	usb.end()
}

func main() {

	/*
	定义接口
	type interface_name interface {
			method_name1 [return_type]
			method_name2 [return_type]
			method_name3 [return_type]
			...
			method_namen [return_type]
	}


	 实现接口方法
		func (struct_name_variable struct_name) method_name1() [return_type] {

		}

	*/


	m := Mouse{
		name: "hero",
	}

	f := FlaskDisk{
		name: "ADATA",
	}

	fmt.Println(m)
	fmt.Println(f)

	//m.start()
	//m.end()
	TestInfo(m)
	//f.start()
	//f.end()
	TestInfo(f)
}
