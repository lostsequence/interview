package main

import (
	"fmt"
)

// реализовать парралельную запись и чтение
func main() {
	cnt := 10000
	m := make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		go func() {
			m[i] = i + i
		}()
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
