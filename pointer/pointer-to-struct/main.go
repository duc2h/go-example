package main

import "fmt"

// pointer to struct https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/

/*
In Go, when we print the contents of a struct using the fmt.Println() function, it prints the values along with their respective keys in curly braces. So, the output you see in the console, &{a}, means that it's a pointer to the user struct with the name field set to "a". The & symbol indicates that c is a pointer to the memory address of a, while the curly braces show the content of the struct being pointed to.

If you want the physical address of a, you can use the fmt.Printf() function with the %p format specifier to print the address in hexadecimal format. Here's an example of how to do it:

package main

import "fmt"

func main() {

 type user struct {
  name string
 }

 a := user{
  name: "a",
 }

 c := &a
 fmt.Printf("Physical address of a: %pn", &a)
 fmt.Println("Value of c:", c)
}


This will output:

Physical address of a: 0x1040b420
Value of c: &{a}


You can see that the physical address of a is printed in hexadecimal format with %p and it's different from the value of pointer c.
*/

func main() {

	type user struct {
		name string
	}

	a := &user{
		name: "a",
	}

	c := &a
	fmt.Println(a)
	fmt.Println(*c)
}
