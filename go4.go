package main

import "fmt"

const age = 45

var firstname = "harsh"

func main() {
	const name = "golang"
	fmt.Println(firstname)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(host, " ", port)
	fmt.Println(port)
}
