package treesAndgraphs

type Heap struct {
	Store    []int
	Capacity int
	Length   int
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
		h.Length += 1
		h.Store = append(h.Store, value)
		i := h.Length - 1
		// heapify till root after insertion
		for i != 0 && h.Store[(i-1)/2] > h.Store[i] {
			h.Store[(i-1)/2], h.Store[i] = h.Store[i], h.Store[(i-1)/2]
			i = (i - 1) / 2 // <---- moving to next parent until find the root
		}
		return true
	}
	return false
}

func (h *Heap) Heapify(index int) {

}
