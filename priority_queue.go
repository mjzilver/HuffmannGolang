package main

import "fmt"

// implements heap.Interface
type PriorityQueue []*huffmanNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// sort the queue by frequency high in front low in back
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

// swap the elements (needed for heap.Interface)
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// push the node to the queue in the right position
// highest priority is at the front of the queue = highest frequency
// because we pop the last element from the queue which has the lowest frequency
func (pq *PriorityQueue) Push(x *huffmanNode) {
	node := x
	if len(*pq) == 0 {
		*pq = append(*pq, node)
		return
	}

	// loop through the queue
	for i, n := range *pq {
		if node.freq >= n.freq {
			// insert the node at the current position
			// and append the rest of the queue to the node
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
