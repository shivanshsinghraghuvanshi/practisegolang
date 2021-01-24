package treesAndgraphs

type Heap struct {
	Store    []int
	Capacity int
	Length   int
}

func NewHeap(capacity int) *Heap {
	return &Heap{Store: make([]int, capacity),
		Capacity: capacity,
		Length:   0}
}

func (h *Heap) New(capacity int) *Heap {
	return &Heap{
		Store:    make([]int, capacity),
		Capacity: capacity,
		Length:   0,
	}
}

func (h *Heap) ParentIndex(i int) int {
	return (i - 1) / 2
}

func (h *Heap) LeftChild(p int) int {
	return 2*p + 1
}

func (h *Heap) RightChild(p int) int {
	return 2*p + 2
}

func (h *Heap) ParentElement(i int) int {
	return h.Store[(i-1)/2]
}

func (h *Heap) Insert(value int) bool {
	if h.Length < h.Capacity {
		h.Store[h.Length] = value
		h.Length += 1
		i := h.Length - 1
		h.HeapifyUp(i)
		return true
	}
	return false
}

func (h *Heap) Delete(k int) {
	h.Store[k], h.Store[h.Length-1] = h.Store[h.Length-1], h.Store[k]
	h.Store = h.Store[:h.Length-1]
	h.Length -= 1
	if k < h.Length && h.Store[k] > h.Store[(k+1)/2] {
		h.HeapifyUp(k)
	} else {
		h.HeapifyDown(k)
	}
}
func (h *Heap) HeapifyUp(i int) {
	// heapify till root
	for i != 0 && h.Store[(i-1)/2] > h.Store[i] {
		h.Store[(i-1)/2], h.Store[i] = h.Store[i], h.Store[(i-1)/2]
		i = (i - 1) / 2 // <---- moving to next parent until find the root
	}
}

func (h *Heap) HeapifyDown(i int) {
	l := h.LeftChild(i)
	r := h.RightChild(i)
	sm := 0
	if l < h.Length && h.Store[l] < h.Store[r] {
		sm = l
	}
	if r < h.Length && h.Store[r] < h.Store[l] {
		sm = r
	}
	if sm != i {
		h.Store[i], h.Store[sm] = h.Store[sm], h.Store[i]
		h.HeapifyDown(sm)
	}
}
