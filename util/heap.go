package util

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

// IntHeap method - checks if element of i index is less than j index
func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

// IntHeap method -swaps the element of i to j index
func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

//IntHeap method -pushes the item
func (ih *IntHeap) Push(heapintf interface{}) {

	*ih = append(*ih, heapintf.(int))
}

//IntHeap method -pops the item from the heap
func (ih *IntHeap) Pop() interface{} {
	var n int
	var x1 int
	previous := *ih
	n = len(previous)
	x1 = previous[n-1]
	*ih = previous[0 : n-1]
	return x1
}
