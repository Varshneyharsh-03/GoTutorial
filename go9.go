package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2, 5)
	fmt.Println(cap(nums2))
	nums2[0] = 1
	nums2 = append(nums2, 45)
	nums2 = append(nums2, 35)
	nums2 = append(nums2, 25)
	//nums2 = append(nums2, 15)
	fmt.Println(nums2, cap(nums2))

	nums3 := []int{}
	nums3 = append(nums3, 1, 2, 56)
	//nums3[4] = 34
	fmt.Println(nums3, len(nums3), cap(nums3))

	var num5 = []int{1, 2, 3, 4, 5}
	fmt.Println(num5[0:2])
	fmt.Println(num5[1:])
	fmt.Println(num5[:2])

	var nums7 = []int{1, 2}
	var nums8 = []int{1, 2}

	fmt.Println(slices.Equal(nums7, nums8))

	var nums9 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums9)
}
