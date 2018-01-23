package dijkstra

import (
	"fmt"
	pq "goalg/pqueue"
)

type VertexData struct {
	Parent   string
	Distance int
}

// ищет кратчайшие пути в графе алгоритмом Дейкстры
// s -- start vertex
// g -- graph to be explored
func Dj(s string, g map[string]map[string]int) map[string]VertexData {

	// множество исследованных узлов
	S := make(map[string]VertexData, len(g))
	S[s] = VertexData{}

	// "узел-дистанция от начального узла"
	// дистанция от начального узла до начального узла -- 0
	dV := make(map[string]int, len(g))
	dV[s] = 0

	// множество V - S
	diffvs := keys(g)
	delete(diffvs, s)

	// costs
	// key -- vertex
	// value -- distance from start vertex
	costs := make(map[string]int, len(diffvs))

	// priority queue
	q := new(pq.PriorityQueue)
	q.StartHeap(len(diffvs))

	// parents
	parent := make(map[string]string, len(diffvs))

	// while S != V
	for current := s; len(diffvs) > 0; {
		for v, le := range g[current] {
			if _, ok := S[v]; ok {
				continue
			}

			d1v := dV[current] + le
			cv, ok := costs[v]
			switch true {
			case !ok:
				q.Insert(v, d1v)
				costs[v] = d1v
				parent[v] = current
			case cv > d1v:
				q.ChangeKey(v, d1v)
				costs[v] = d1v
				parent[v] = current
			}
		}

		// найдём ближайший к началу из узлов в костах
		n, _, err := q.ExtractMin()
		if err != nil {
			panic(fmt.Sprintf("Extracting failed: %s\n", err.Error()))
		}

		dV[n] = costs[n]
		S[n] = VertexData{parent[n], costs[n]}
		delete(diffvs, n)
		delete(costs, n)
		current = n
	}

	return S
}

func nearest(m map[string]int) string {
	var min int
	var nearest string
	c := 0
	for k, v := range m {
		if c == 0 {
			min = v
			nearest = k
			c++
			continue
		}
		if v < min {
			min = v
			nearest = k
		}
	}
	return nearest
}

func keys(m map[string]map[string]int) map[string]bool {
	keys := make(map[string]bool, len(m))
	for k, _ := range m {
		keys[k] = true
	}
	return keys
}
