package main

// factorial.
// 0! = 1
// n! = n*(n-1)!

// for example: 2! = ?
// 1. 2! = 2 * (2-1)!
//// 1.1 (2-1)! = 1! = 1 * (1-1)!
// 2. 2! = 2 * (1 * (1 - 1)!)
//// 2.1 (1-1)! = 1
// 3. 2! = 2 * (1 * 1) = 2
// => 2! = 2.

// func main() {

// 	a := f(3)

// 	fmt.Println(a)
// }

func f(n int) int {
	if n == 0 {
		return 1
	}

	return n * f(n-1)
}
