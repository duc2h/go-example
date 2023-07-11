package main

import "fmt"

func main() {
	runA()
}

func runA() {
	// a := make([]int, 1, 2)

	// fmt.Printf("1: len: %d, cap: %d, address_0: %d, address_1 %d \n", len(a), cap(a), &a[0])
	// a = append(a, 1)

	// fmt.Printf("2: len: %d, cap: %d, address_0: %d, address_1 %d \n", len(a), cap(a), &a[0])

	// a = append(a, 3) // require new allocation
	// fmt.Printf("A: len: %d, cap: %d, address_0: %d, address_1 %d \n", len(a), cap(a), &a[0])
	// fmt.Println(a)

	// b := append(a, 4) // len = cap => still use old allocation
	// fmt.Printf("B: len: %d, cap: %d, address_0: %d, address_1 %d, address_3: %d \n", len(b), cap(b), &b[0], &b[3])
	// fmt.Println(b)

	// e := append(a, 5) // len = cap => still use old allocation
	// fmt.Printf("E: len: %d, cap: %d, address_0: %d, address_1 %d, address_3: %d \n", len(e), cap(e), &e[0], &e[3])
	// fmt.Println(e)

	// // b has been changed value by e. a has 4 locations (11,22,33,44).
	// // a used 3 locations, b add value 4 into location 44. then e change the value to 5 in location 44.
	// // => value at index 4 of b changed due to e. because they are reference to the same location, when append from a.
	// fmt.Println(b)
	// fmt.Println(a) // a is original array
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	test()

}

func test() {
	a := make([]int, 0, 2)

	// fmt.Printf("1: len: %d, cap: %d, address_0: %d, address_1 %d \n", len(a), cap(a), &a[0])
	a = append(a, 1)
	fmt.Printf("1: len_a: %d, cap_a: %d, address_0: %d \n", len(a), cap(a), &a[0]) // 1, 2, 0x111
	fmt.Println(a)

	b := append(a, 5)
	fmt.Printf("2: len_b: %d, cap_b: %d, address_0: %d, address_1 %d \n", len(b), cap(b), &b[0], &b[1])
	fmt.Println(b)
	fmt.Printf("3: len_a: %d, cap_a: %d, address_0: %d \n", len(a), cap(a), &a[0])
	fmt.Println(a)

	c := append(a, 6)
	fmt.Printf("4: len_c: %d, cap_c: %d, address_0: %d, address_1 %d \n", len(c), cap(c), &c[0], &c[1])
	fmt.Println(c)
	fmt.Printf("5: len_a: %d, cap_a: %d, address_0: %d \n", len(a), cap(a), &a[0])
	fmt.Println(a)
	fmt.Printf("6: len_b: %d, cap_b: %d, address_0: %d, address_1 %d \n", len(b), cap(b), &b[0], &b[1])
	fmt.Println(b)

	a = append(a, 1)
	fmt.Printf("7: len_a: %d, cap_c: %d, address_0: %d, address_1 %d \n", len(a), cap(a), &a[0], &a[1])
	fmt.Println(a)
	fmt.Printf("8: len_c: %d, cap_c: %d, address_0: %d, address_1 %d \n", len(c), cap(c), &c[0], &c[1])
	fmt.Println(c)
	fmt.Printf("9: len_b: %d, cap_b: %d, address_0: %d, address_1 %d \n", len(b), cap(b), &b[0], &b[1])
	fmt.Println(b)
}

type Person struct {
}

func emptySlice() {
	var slice []*Person          // nil slice
	slice1 := []*Person{}        // empty slice
	slice2 := make([]*Person, 0) // empty slice

	fmt.Println(slice == nil)
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
}
