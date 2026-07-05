package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(num, i)
	}

	fmt.Println(sum)

	map2 := map[string]int{"a": 1, "b": 2}
	for k, v := range map2 {
		fmt.Println(k, v)
	}

	for i, c := range "zolang" {
		fmt.Println(i, string(c))
	}
}
