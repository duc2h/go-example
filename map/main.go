package main

import "fmt"

func main() {
	m := map[int]int{}

	for i := 0; i <= 100; i++ {
		m[i] = i
	}

	// when we use for-lop a map in go, the sequence is not order.
	// the reason due to: map implement hashmap, they store data in a bucket
	// when we insert value to map, map will run hash function to get index => store data into bucket(index)
	//// bucket(index) = hash(key)
	for k := range m {
		v := m[k]

		fmt.Println(v)
	}
}
