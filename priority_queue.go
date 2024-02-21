package main

import "fmt"

type PriorityQueue []*huffmanNode

// highest priority is at the front of the queue = highest frequency
func (pq *PriorityQueue) Push(x *huffmanNode) {
	node := x
	if len(*pq) == 0 {
		*pq = append(*pq, node)
		return
	}

	for i, n := range *pq {
		if node.freq >= n.freq {
			// insert the node at the current position and append the rest of the queue to the node
			*pq = append((*pq)[:i], append(PriorityQueue{node}, (*pq)[i:]...)...)
			return
		} else if i == len(*pq)-1 {
			// if the node has the lowest frequency append it to the end of the queue
			*pq = append(*pq, node)
			return
		}
	}
}

// get the lowest frequency node and remove it from the queue
func (pq *PriorityQueue) Pop() *huffmanNode {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// print the queue for debugging
func (pq *PriorityQueue) Print() {
	for _, node := range *pq {
		fmt.Printf("%s:%d ", node.String(), node.freq)
	}
	fmt.Println()
}
