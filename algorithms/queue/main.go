package main

import (
	"container/list"
	"fmt"
)

func main() {
	runQueueSlice()
}

// make a queue with slice
func runQueueSlice() {
	q := []string{"1", "2"}
	// q = enqueueSlice(q, "3")
	// q = enqueueSlice(q, "4")

	s, q := dequeueSlice(q)
	fmt.Println(s)
	fmt.Println(q)
	fmt.Println("==============")

	// dequeueSlice(q)
	// fmt.Println(s)
	// fmt.Println(q)
	// fmt.Println("==============")

}

// dequeue
// this way can make memory leak, if the array's type are pointer or "header" type (slice, string)
// more detail: https://stackoverflow.com/questions/55045402/memory-leak-in-golang-slice
func dequeueSlice(q []string) (string, []string) {
	num := q[0]
	q = q[1:]

	return num, q
}

// enqueue
func enqueueSlice(q []string, ele string) []string {
	q = append(q, ele)

	return q
}

// ///
// in order to avoid memory leak, we can use linked list.
func runLL() {
	q := list.New()
	enqueueLL(q, "1")
	enqueueLL(q, "2")
	enqueueLL(q, "3")
	enqueueLL(q, "4")
	enqueueLL(q, "5")

	dequeueLL(q)

	dequeueLL(q)

}

func enqueueLL(q *list.List, ele string) {
	q.PushBack(ele)
}

func dequeueLL(q *list.List) string {
	str := q.Front()
	q.Remove(str)
	return str.Value.(string)
}
