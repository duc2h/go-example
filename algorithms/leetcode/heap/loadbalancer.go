package main

import (
	"fmt"
	"sync"
)

// https://www.geeksforgeeks.org/calculate-server-loads-using-round-robin-scheduling/

func main() {

	arrival := []int{2, 4, 1, 8, 9}
	brusTime := []int{7, 9, 2, 4, 5}
	r := loadBalancer(3, arrival, brusTime)
	fmt.Println(r)
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue struct {
	Items []Item
	lock  sync.RWMutex
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Items) == 0
}

func (pq *PriorityQueue) Push(e Item) {
	// if the pq is empty
	insertFlag := false

	if len(pq.Items) == 0 {
		pq.Items = append(pq.Items, e)
		return
	}

	// if the pq is not empty
	for index, item := range pq.Items {
		if item.priority > e.priority {
			if index == 0 {
				pq.Items = append([]Item{e}, pq.Items...)
			} else {
				pq.Items = append(pq.Items[:index+1], pq.Items[index:]...)
				pq.Items[index] = item
			}

			insertFlag = true
			break
		}
	}

	if !insertFlag {
		pq.Items = append(pq.Items, e)
	}
}

func (pq *PriorityQueue) Pop() Item {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	item := pq.Items[0]

	pq.Items = pq.Items[1:]
	return item
}

func loadBalancer(n int, arrival []int, brustTime []int) []int {
	var (
		pq = &PriorityQueue{}
		// mapServer = map[int]int{}
		result = []int{}
		// flag      = false
		pqServer = &PriorityQueue{}
	)
	for i, a := range arrival {
		pq.Push(Item{
			value:    i,
			priority: a,
		})
	}

	for i := 1; i <= n; i++ {
		pqServer.Push(Item{
			value:    i,
			priority: 0,
		})
	}

	fmt.Println(pqServer)

	for !pq.IsEmpty() {
		item := pq.Pop()
		arrive := arrival[item.value]
		brust := brustTime[item.value]

		totalTime := arrive + brust

		if pqServer.IsEmpty() {
			return nil
		}

		server := pqServer.Pop()

		if server.priority > arrive {
			pqServer.Push(server)
			result = append(result, -1)
		} else {
			pqServer.Push(Item{value: server.value, priority: totalTime})
			result = append(result, server.value)
		}

		// for i := 1; i <= n; i++ {
		// 	if serverTime, exist := mapServer[i]; exist {
		// 		if serverTime <= arrive {
		// 			mapServer[i] = totalTime
		// 			result = append(result, i)
		// 			flag = true
		// 			break
		// 		}
		// 	}
		// }

		// if flag {
		// 	flag = false
		// } else {
		// 	result = append(result, -1)
		// }
	}

	return result
}
