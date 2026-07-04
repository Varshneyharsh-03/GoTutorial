package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//for {
	//	fmt.Println(i)
	//}

	for i := 1; i <= 6; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	for i := range 10 {
		fmt.Println(i)
	}

}
