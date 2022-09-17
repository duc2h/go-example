package main

import "fmt"

func main() {
	runA()
}

func runA() {
	a := make([]int, 1, 2)

	fmt.Printf("1: len: %d, cap: %d, address: %d \n", len(a), cap(a), &a[0])
	a = append(a, 1)

	fmt.Printf("2: len: %d, cap: %d, address: %d \n", len(a), cap(a), &a[0])

	a = append(a, 3) // require new allocation
	fmt.Printf("A: len: %d, cap: %d, address: %d \n", len(a), cap(a), &a[0])
	fmt.Println(a)

	b := append(a, 4) // len = cap => still use old allocation
	fmt.Printf("B: len: %d, cap: %d, address_0: %d, address_3: %d \n", len(b), cap(b), &b[0], &b[3])
	fmt.Println(b)

	e := append(a, 5) // len = cap => still use old allocation
	fmt.Printf("E: len: %d, cap: %d, address_0: %d, address_3: %d \n", len(e), cap(e), &e[0], &e[3])
	fmt.Println(e)
	// b has changed value due to e. a has 4 locations (11,22,33,44).
	// it used 3, b add value 4 into location 44. then e change the value to 5 in location 44.
	// => value at index 4 of b changed follow e. because they are reference to the same location, when append from a.
	fmt.Println(b)
	fmt.Println(a) // a is original array

}
