package main

import "fmt"

func main() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
