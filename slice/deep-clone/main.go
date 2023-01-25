package main

import "fmt"

func main() {

	s := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8}
	removeSlice(s, 0)
	fmt.Printf("After remove slice: %v \n", s)
	fmt.Printf("len: %d, cap: %d \n", len(s), cap(s))
	fmt.Printf("%p \n", s)

	// s1 := removeSlice1(s, 0)
	// fmt.Printf("After remove slice s: %v \n", s)
	// fmt.Printf("After remove slice s1: %v \n", s1)
	// fmt.Printf("len: %d, cap: %d \n", len(s), cap(s))
	// fmt.Printf("len: %d, cap: %d \n", len(s1), cap(s1))
	// fmt.Printf("address_s0-0: %d, address_s1-0: %d \n", &s[0], &s1[0])
}

// error
func removeSlice(s []byte, i int) {
	s = append(s[:i], s[i+1:]...)
	fmt.Printf("Remove slice: %v \n", s)
	fmt.Printf("len: %d, cap: %d \n", len(s), cap(s))
	fmt.Printf("%p \n", s)

}

func removeSlice1(s []byte, i int) []byte {
	s = append(s[:i], s[i+1:]...)
	fmt.Println(&s)
	fmt.Printf("Remove slice: %v \n", s)
	return s
}
