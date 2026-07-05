package main

import "fmt"

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func main() {
	fmt.Println(1, 2, 3, 4, 5, "hello world ")

	var ans = sum(1, 2, 3, 4, 5)
	fmt.Println(ans)
}
