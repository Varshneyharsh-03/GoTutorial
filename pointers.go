package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

func changeByRef(num *int) {
	*num = 5
	fmt.Println("In changeByRef", num)
}
func main() {
	num := 1
	changeNum(num)
	fmt.Println("memory address ", &num)
	fmt.Println("after change num in main fn : ", num)
	changeByRef(&num)
	fmt.Println("after change num in main fn after ref : ", num)
}
