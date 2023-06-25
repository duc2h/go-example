package main

import "fmt"

// fibonacci formula:
// f(1) = 1
// f(2) = 2
// n > 2 => f(n) = f(n-1) + f(n-2)

// for example: n = 4
//1. f(4) = f(3) + f(2)
//2. f(4) = (f(2) + f(1)) + 1
//3. f(4) = 1 + 1 + 1 = 3
// => f(4) = 3

func main() {

	a := fibonacci(4)

	fmt.Println(a)
}

func fibonacci(n int) int {
	if n == 2 || n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)

}
