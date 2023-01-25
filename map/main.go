package main

import "fmt"

func main() {
	m := map[int]int{}

	for i := 0; i <= 100; i++ {
		m[i] = i
	}

	// when we use for-lop a map in go, the sequence is not order.
	for k := range m {
		v := m[k]

		fmt.Println(v)
	}
}
