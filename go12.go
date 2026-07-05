package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

func processIt(fn func(a int) int) {
	fn(2)
}

func processIt2() func(a int) int {
	return func(a int) int { return 0 }
}
func main() {
	sum := add(1, 2)
	fmt.Println(sum)
	fmt.Println(getLanguages())

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)
	fn := func(a int) int { return a + 1 }
	processIt(fn)

	newfn := processIt2()
	fmt.Println(newfn(6))

}
