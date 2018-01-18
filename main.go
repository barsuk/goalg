package main

import "fmt"
import "goalg/dijkstra"
import pq "goalg/pqueue"

var adj = map[string]map[string]int{
	"a": {
		"b": 1,
		"d": 4,
		"f": 8,
		"g": 4,
	},
	"b": {
		"c": 1,
		"d": 3,
		"e": 1,
	},
	"c": {
		"a": 2,
		"g": 7,
	},
	"d": {
		"c": 2,
		"g": 5,
	},
	"e": {
		"b": 9,
		"d": 4,
		"f": 3,
	},
	"f": {
		"a": 1,
		"e": 5,
	},
	"g": {
		"e": 1,
		"f": 2,
	},
}

func main() {
	fmt.Println(adj)
	fmt.Println(dijkstra.Dj(`f`, adj))

	h := new(pq.Heap)
	h.StartHeap(16)

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(5)
	h.Insert(-6)
	h.Insert(-2)
	fmt.Println(*h)

	h.ExtractMin()
	fmt.Println(*h)
	h.ExtractMin()
	fmt.Println(*h)
	h.ExtractMin()
	fmt.Println(*h)
	h.ExtractMin()
	fmt.Println(*h)
	h.ExtractMin()
	fmt.Println(*h)

	min, err := h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%d has been extracted\n", min)
	}

	min, err = h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%d has been extracted\n", min)
	}

	min, err = h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%d has been extracted\n", min)
	}

	min, err = h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%d has been extracted\n", min)
	}

	min, err = h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%d has been extracted\n", min)
	}

}
