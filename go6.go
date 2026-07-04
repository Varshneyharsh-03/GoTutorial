package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("you are a adult")
	} else {
		fmt.Println("person is not adult")
	}

	if age >= 18 {
		fmt.Println("you are a adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "student"
	var hasPermission = false

	if hasPermission || role == "admin" {
		fmt.Println("you are an admin")
	}

	if myage := 18; age >= 18 {
		fmt.Println(myage)
	} else if age >= 12 {
		fmt.Println("person is teenage ", myage)
	}

}
