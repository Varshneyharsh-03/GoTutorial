package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map
	//m := make(map[int]string)
	//m[1] = "1"
	//m[2] = "2"
	//m[3] = "3"
	//fmt.Println(m)
	//fmt.Println(m[3])
	//
	//fmt.Println(m[9])

	//m := make(map[string]int)
	//
	//m["age"] = 30
	//fmt.Println(m["age"], m["phone"])
	//fmt.Println(len(m))
	//
	//delete(m, "age")
	//delete(m, "phone")
	//fmt.Println(m)
	//
	//clear(m)

	m := map[string]int{"age": 3, "phone": 45, "name": 4}
	fmt.Println(m)
	k, ok := m["age"]
	fmt.Println(k)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	m1 := map[string]int{"age": 3, "phone": 45, "name": 4}
	m2 := map[string]int{"age": 3, "phone": 45, "name": 4}

	fmt.Println(maps.Equal(m1, m2))
}
