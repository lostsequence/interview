package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(len(nums))

	strs := make([]string, 2, 2)
	strs[0] = "hello"
	strs[1] = "aloha"

	fmt.Println(strs[0])
}
