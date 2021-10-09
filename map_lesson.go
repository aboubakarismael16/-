package main

import (
	"fmt"
	"sort"
)

func main() {

	/*
		1. make map
		2. map nil
		4. take a value from map
		5.delete value from map
			if the value exist then it can be delete but
			if the value does not exist the operation can not affect map

		6.len map
			len()
	*/

	// 1. make map
	var map1 map[int]string /*如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对*/
	map2 := make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 72}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println("-------------------------")
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	//2. map nil
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "nepu"

	fmt.Println(map1[3])

	//3.take a value from map
	v, ok := map1[40]
	if ok {
		fmt.Printf("the value id : %s\n", v)
	} else {
		fmt.Printf("the value is not exist : %s\n", v)
	}
	//4. modify a value in map
	map1[2] = "Golang"
	fmt.Println(map1)

	//5. delete a value from map
	delete(map1, 3)
	fmt.Println(map1)

	//6.map len
	fmt.Println(len(map1))
	fmt.Println("-------------------------")

	map4 := make(map[int]string)
	map4[1] = "china"
	map4[2] = "chad"
	map4[3] = "france"
	map4[4] = "usa"
	map4[5] = "japon"

	// not in order
	for k, v := range map4 {
		fmt.Printf("key %d ----> value %s\n", k, v)
	}

	fmt.Println("-------------------------")
	// make in order
	keys := make([]int, 0, len(map4))

	for k, _ := range map4 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, map4[k])
	}
	fmt.Println("---------------------")

	map5 := make(map[string]string)
	map5["name"] = "ismael"
	map5["age"] = "28"
	map5["sex"] = "male"
	map5["address"] = "NDJXXXRueXXXporte"

	map6 := make(map[string]string)
	map6["name"] = "Fatima"
	map6["age"] = "25"
	map6["sex"] = "female"
	map6["address"] = "ABCHXXXRueXXXporte"

	map7 := make(map[string]string)
	map7["name"] = "youssouf"
	map7["age"] = "30"
	map7["sex"] = "male"
	map7["address"] = "NDJXXXRueXXXporte"

	s_map := make([]map[string]string, 0, 3)
	s_map = append(s_map, map5)
	s_map = append(s_map, map6)
	s_map = append(s_map, map7)

	for k, val := range s_map {
		fmt.Printf("person %d\n", k+1)
		fmt.Printf("\t name : %s\n", val["name"])
		fmt.Printf("\t age : %s\n", val["age"])
		fmt.Printf("\t sex : %s\n", val["sex"])
		fmt.Printf("\t address : %s\n", val["address"])
	}
}
