package main

import "fmt"

func main() {
	var m map[int]string

	fmt.Println(m[1] == "")
	_, ok := m[1]
	fmt.Println(ok)

	m[1] = "bar"
	fmt.Println(m[1])
}
