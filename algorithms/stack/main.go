package main

import "fmt"

func main() {
	run()
}

// use slice
func run() {
	s := []string{"1", "2"}
	s = add(s, "3")
	s = add(s, "4")

	fmt.Println(s)

	str, s := remove(s)
	fmt.Println(str)
	fmt.Println(s)
	fmt.Println("=============")

	str, s = remove(s)
	fmt.Println(str)
	fmt.Println(s)
	fmt.Println("=============")

}

// add
func add(s []string, ele string) []string {
	s = append(s, ele)

	return s
}

func remove(s []string) (string, []string) {
	last := len(s) - 1
	str := s[last]
	s = s[:last]

	return str, s
}

// using linked list to avoid memory leaks
