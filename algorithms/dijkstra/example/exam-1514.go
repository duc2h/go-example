package main

import (
	"fmt"
	"math"
	"sync"
)

// https://leetcode.com/problems/path-with-maximum-probability/

// You are given an undirected weighted graph of n nodes (0-indexed), represented by an edge list where
// edges[i] = [a, b] is an undirected edge connecting the nodes a and b with a probability of success of traversing
// that edge succProb[i].

// Given two nodes start and end, find the path with the maximum probability of success to go from start to end
// and return its success probability.

// If there is no path from start to end, return 0. Your answer will be accepted if it differs from the correct
// answer by at most 1e-5.

///// example
// Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
// Output: 0.25000
// Explanation: There are two paths from start to end, one having a probability of success = 0.2
// and the other has 0.5 * 0.5 = 0.25.

// create graph
type Edge1514 struct {
	Node   int
	Weight float64
}
type Graph1514 struct {
	Edges map[int][]Edge1514
}

// priority queue
type PriorityQueue struct {
	Items []Edge1514
	Lock  sync.RWMutex
}

// implement enqueue
func (pq *PriorityQueue) Enqueue(e Edge1514) {
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
			pq.Items = append([]Edge1514{e}, pq.Items...)
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
func (pq *PriorityQueue) Dequeue() Edge1514 {
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

// create graph
func createGraph(n int, edges [][]int, succProb []float64) Graph1514 {
	g := Graph1514{
		Edges: map[int][]Edge1514{},
	}

	for i, e := range edges {
		if len(succProb) <= i {
			continue
		}
		key1 := e[0]
		key2 := e[1]

		if _, ok := g.Edges[key1]; ok {
			v := g.Edges[key1]
			v = append(v, Edge1514{
				Node:   key2,
				Weight: succProb[i],
			})
			g.Edges[key1] = v
		} else {
			g.Edges[key1] = []Edge1514{
				{
					Node:   key2,
					Weight: succProb[i],
				},
			}
		}

		if _, ok := g.Edges[key2]; ok {
			v := g.Edges[key2]
			v = append(v, Edge1514{
				Node:   key1,
				Weight: succProb[i],
			})
			g.Edges[key2] = v
		} else {
			g.Edges[key2] = []Edge1514{
				{
					Node:   key1,
					Weight: succProb[i],
				},
			}
		}
	}

	return g
}

// findPath
func findPath(n int, graph Graph1514, start, end int) float64 {
	cost := map[int]float64{}
	for i := 0; i < n; i++ {
		cost[i] = -math.MaxFloat64
	}

	cost[start] = 1

	pq := PriorityQueue{
		Items: []Edge1514{
			{
				Node:   start,
				Weight: 0,
			},
		},
	}

	for !pq.IsEmpty() {
		item := pq.Dequeue()

		for _, e := range graph.Edges[item.Node] {
			fmt.Println(e)

			c := e.Weight * cost[item.Node]
			if c > cost[e.Node] {
				cost[e.Node] = c

				// enqueue

				pq.Enqueue(Edge1514{
					Node:   e.Node,
					Weight: e.Weight,
				})
			}
		}
	}

	fmt.Println(cost)

	if v := cost[end]; v > 0 {
		return v
	}

	//
	return 0
}

func main() {

	e := [][]int{
		{1, 4},
		{2, 4},
		{0, 4},
		{0, 3},
		{0, 2},
		{2, 3},
	}

	fmt.Println(e)
	s := []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04}
	fmt.Println(maxProbability(5, e, s, 3, 4))

}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := createGraph(n, edges, succProb)
	fmt.Println(graph)

	return findPath(n, graph, start_node, end_node)
}
