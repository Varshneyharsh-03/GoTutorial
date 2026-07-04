package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")

	default:
		fmt.Println("default")
	}

	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			fmt.Println("int is", i)
		case string:
			fmt.Println("string is", i)
		case bool:
			fmt.Println("bool is", i)
		}
	}

	whoAmI("goland")
}
