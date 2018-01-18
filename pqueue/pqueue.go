package pqueue

import "fmt"

type Heap []int

func (h *Heap) StartHeap(n int) {
	*h = make([]int, 0, n)
}

func (h *Heap) HeapifyUp(i int) {
	if i > 0 {
		var j int = (i - 1) / 2
		if (*h)[i] < (*h)[j] {
			(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
			h.HeapifyUp(j)
		}
	}
}

func (h *Heap) Insert(v int) {
	*h = append(*h, v)
	i := len(*h) - 1
	h.HeapifyUp(i)
}

func (h *Heap) FindMin() (int, error) {
	if len(*h) == 0 {
		return 0, fmt.Errorf(`heap is empty yet!`)
	}
	return (*h)[0], nil
}

func (h *Heap) ExtractMin() (int, error) {
	min, err := h.FindMin()
	if err != nil {
		return 0, fmt.Errorf(`unable to find minimal: %s`, err.Error())
	}

	h.Delete(0)

	return min, nil
}

func (h *Heap) Delete(i int) error {
	(*h)[i] = (*h)[len(*h)-1]
	(*h)[len(*h)-1] = 0
	*h = (*h)[:len(*h)-1]

	err := h.HeapifyDown(i)
	if err != nil {
		return fmt.Errorf(`unable to heapify down: %s`, err.Error())
	}

	return nil
}

// Orders heap after deleting an element in position i (index starts from 0!)
func (h *Heap) HeapifyDown(i int) error {
	var j, left, right int

	n := len(*h)

	switch p := 2*i + 2; true {
	case p > n:
		return nil
	case p < n:
		left = 2*i + 1
		right = 2*i + 2
		if (*h)[left] < (*h)[right] {
			j = left
		} else {
			j = right
		}
	case p == n:
		j = 2*i + 1
	}

	if (*h)[j] < (*h)[i] {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		h.HeapifyDown(j)
	}

	return nil
}
