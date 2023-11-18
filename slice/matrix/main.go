package main

import "fmt"

func main() {
	a := [][]string{}

	a = append(a, []string{"1"})
	a = append(a, []string{"2"})

	for i, v := range a {
		fmt.Println(i)
		fmt.Println(v)

		fmt.Println("---------------------")
	}
}
