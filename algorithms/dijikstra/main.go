package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	fmt.Println("asd")

	graph := Graph{
		Nodes: []string{"A", "B", "C", "D"},
		Edges: map[string][]Edge{
			"A": {
				{
					Node:   "B",
					Weight: 1,
				},
				{
					Node:   "C",
					Weight: 5,
				},
			},
			"B": {
				{
					Node:   "C",
					Weight: 3,
				},
			},
			"C": {
				{
					Node:   "D",
					Weight: 1,
				},
			},
		},

		Lock: sync.RWMutex{},
	}

	p, d := findShortestPath("A", "D", &graph)
	fmt.Println(p)
	fmt.Println(d)
}

type Edge struct {
	Node   string
	Weight int
}

type Graph struct {
	Nodes []string
	Edges map[string][]Edge
	Lock  sync.RWMutex
}

// type PriorityQueue []*Edge

type PriorityQueue struct {
	Items []Edge
	Lock  sync.RWMutex
}

// implement enqueue
func (pq *PriorityQueue) Enqueue(e Edge) {
	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	if len(pq.Items) == 0 {
		pq.Items = append(pq.Items, e)
		return

	}

	insertFlag := false

	// sort the queue by distance value, from lowest -> highest
	for index, item := range pq.Items {
		if item.Weight <= e.Weight {
			continue
		}

		if index == 0 {
			pq.Items = append([]Edge{e}, pq.Items...)
		} else {
			pq.Items = append(pq.Items[:index+1], pq.Items[index:]...)
			pq.Items[index] = e
		}

		insertFlag = true
		break

	}

	if !insertFlag {
		pq.Items = append(pq.Items, e)
	}
}

// implement dequeue
func (pq *PriorityQueue) Dequeue() Edge {
	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	e := pq.Items[0]
	pq.Items = pq.Items[1:]

	return e
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Items) == 0
}

func (pq *PriorityQueue) Size() int {
	return len(pq.Items)
}

func findShortestPath(start, end string, graph *Graph) ([]string, int) {
	var (
		// visited = make(map[string]bool)
		dist = make(map[string]int)
		path = make(map[string]string)
		pq   = PriorityQueue{}
	)

	//
	for _, node := range graph.Nodes {
		dist[node] = math.MaxInt
	}

	// start's weight = 0, due to we will go from this point.
	dist[start] = 0
	path[start] = ""

	// add start item into pq.
	startItem := Edge{
		Node:   start,
		Weight: 0,
	}

	pq.Enqueue(startItem)

	for !pq.IsEmpty() {
		item := pq.Dequeue()

		for _, e := range graph.Edges[item.Node] {
			sumWeight := dist[item.Node] + e.Weight
			if sumWeight < dist[e.Node] {
				dist[e.Node] = sumWeight

				pq.Enqueue(Edge{
					Node:   e.Node,
					Weight: sumWeight,
				})
				path[e.Node] = item.Node
			}
		}
	}

	finalArr := []string{}
	prev, exist := path[end]
	if !exist {
		return nil, 0
	}

	finalArr = append(finalArr, end, prev)
	for prev != start {
		prev, exist = path[prev]
		if !exist {
			return nil, 0
		}

		if prev == "" {
			break
		}

		finalArr = append(finalArr, prev)
	}

	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}

	return finalArr, dist[end]
}
