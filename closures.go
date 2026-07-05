package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}
func main() {
	var increment = counter()
	var count1 = increment()
	var count2 = increment()
	fmt.Println(count1)
	fmt.Println(count2)

}
