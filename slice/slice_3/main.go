package main

import "fmt"

func main() {
	n := make([]int, 6, 6)
	for i := 0; i < len(n); i++ {
		n[i] = i
	}

	n1 := n[1:3]
	fmt.Println(n1)
	fmt.Println(len(n1), cap(n1))

	n1[0] = 777
	// fmt.Println(n)
	// fmt.Println(n1)

	n1 = n1[:cap(n1)]
	// fmt.Println(n1)
	// fmt.Println(len(n1), cap(n1))

	n1 = append(n1, 999)
	// fmt.Println(n1)
	// fmt.Println(len(n1), cap(n1))

	n1[0] = 888
	// fmt.Println(n)
	// fmt.Println(n1)
}
