package main

import "fmt"

func main() {
	var strs []string
	initSlice(strs)
	fmt.Println(strs)
}

func initSlice(strs []string) {
	strs = make([]string, 0)
	strs = append(strs, "hello")
}
