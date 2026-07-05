package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var names [3]string
	names[0] = "golang"
	fmt.Println(names)

	nums1 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums1)

}
