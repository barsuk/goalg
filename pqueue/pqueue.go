package pqueue

import "fmt"

type PriorityQueue struct {
	Heap     []int
	Position []string       // position of key in the Heap slice
	Index    map[string]int // position of key in the Heap slice
}

func (pq *PriorityQueue) StartHeap(n int) {
	pq.Heap = make([]int, 0, n)
	pq.Position = make([]string, 0, n)
	pq.Index = make(map[string]int, n)
}

func (pq *PriorityQueue) HeapifyUp(i int) error {
	h := &(pq.Heap)
	pos := &(pq.Position)
	index := &(pq.Index)

	if i > 0 {
		var j int = (i - 1) / 2
		if (*h)[i] < (*h)[j] {
			(*h)[i], (*h)[j], (*pos)[i], (*pos)[j] = (*h)[j], (*h)[i], (*pos)[j], (*pos)[i]

			(*index)[(*pos)[i]] = i
			(*index)[(*pos)[j]] = j

			pq.HeapifyUp(j)
		}
	}

	return nil
}

func (pq *PriorityQueue) Insert(k string, v int) {
	h := &(pq.Heap)

	*h = append(*h, v)
	i := len(*h) - 1

	pq.Position = append(pq.Position, k)

	pq.Index[k] = i

	pq.HeapifyUp(i)
}

func (pq *PriorityQueue) FindMin() (string, int, error) {
	h := &(pq.Heap)
	p := &(pq.Position)

	if len(*h) == 0 {
		return ``, 0, fmt.Errorf(`heap is empty yet!`)
	}
	return (*p)[0], (*h)[0], nil
}

func (pq *PriorityQueue) ExtractMin() (string, int, error) {
	key, min, err := pq.FindMin()
	if err != nil {
		return ``, 0, fmt.Errorf(`unable to find minimal: %s`, err.Error())
	}

	pq.Delete(0)

	return key, min, nil
}

func (pq *PriorityQueue) Delete(i int) error {
	h := &(pq.Heap)
	p := &(pq.Position)
	index := &(pq.Index)

	(*h)[i] = (*h)[len(*h)-1]
	(*h)[len(*h)-1] = 0
	*h = (*h)[:len(*h)-1]

	delete(*index, (*p)[i])

	(*p)[i] = (*p)[len(*p)-1]
	(*p)[len(*p)-1] = ``
	*p = (*p)[:len(*p)-1]

	if len(*p) > 0 {
		(*index)[(*p)[i]] = i
	}

	err := pq.HeapifyDown(i)
	if err != nil {
		return fmt.Errorf(`unable to heapify down: %s`, err.Error())
	}

	return nil
}

// Orders heap after deleting an element in position i (index starts from 0!)
func (pq *PriorityQueue) HeapifyDown(i int) error {
	h := &(pq.Heap)
	pos := &(pq.Position)
	index := &(pq.Index)

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
		(*h)[i], (*h)[j], (*pos)[i], (*pos)[j] = (*h)[j], (*h)[i], (*pos)[j], (*pos)[i]

		(*index)[(*pos)[i]] = i
		(*index)[(*pos)[j]] = j

		pq.HeapifyDown(j)
	}

	return nil
}

func (pq *PriorityQueue) ChangeKey(k string, new int) error {
	h := &(pq.Heap)
	index := &(pq.Index)

	position, ok := (*index)[k]
	if !ok {
		return fmt.Errorf(`no such index: %s`, k)
	}

	oldValue := (*h)[position]
	(*h)[position] = new

	switch true {
	case oldValue > new:
		return pq.HeapifyDown(position)
	case oldValue < new:
		return pq.HeapifyUp(position)
	case oldValue == new:
		return nil
	}

	return nil
}
