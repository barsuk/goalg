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

	h := new(pq.PriorityQueue)
	h.StartHeap(18)

	h.Insert(`a`, 1)
	h.Insert(`b`, 2)
	h.Insert(`c`, 3)
	h.Insert(`d`, 5)
	h.Insert(`e`, -5)
	h.Insert(`f`, -1)
	h.Insert(`g`, 0)
	h.Insert(`h`, 4)
	//h.Insert(`i`, 15)
	//h.Insert(`j`, -3)
	//h.Insert(`k`, -4)
	//h.Insert(`l`, -6)
	//h.Insert(`m`, -8)
	//h.Insert(`n`, -9)
	//h.Insert(`o`, 12)
	//h.Insert(`p`, 22)
	//h.Insert(`q`, 23)
	//h.Insert(`r`, 24)
	//h.Insert(`s`, 25)
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

	h.ExtractMin()
	fmt.Println(*h)
	h.ExtractMin()
	fmt.Println(*h)

	key, min, err := h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%s, %d have been extracted\n", key, min)
	}

	key, min, err = h.ExtractMin()
	if err != nil {
		fmt.Printf("Extracting failed: %s\n", err.Error())
	} else {
		fmt.Printf("%s, %d have been extracted\n", key, min)
	}

}
