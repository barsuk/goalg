package main

import "fmt"
import "goalg/dijkstra"

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
	fmt.Println(dijkstra.Dj(`b`, adj))
}
