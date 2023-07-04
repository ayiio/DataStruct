package main

import (
	"fmt"
	"math"
	"section7/dijkstra"
)

func main() {
	//记录关联关系
	var graph = map[string]map[string]int{
		"S": {"A": 6, "B": 2},
	}
	graph["A"] = map[string]int{"E": 1}
	graph["B"] = map[string]int{"A": 3, "E": 5}
	graph["E"] = map[string]int{}

	//记录节点开销
	var costs = map[string]int{
		"A": 6, "B": 2, "E": math.MaxInt,
	}

	//记录节点的父节点
	var parents = map[string]string{
		"A": "S", "B": "S", "E": "",
	}

	//before
	fmt.Println("before, ", costs, parents)

	dijkstra.DijkstraDemo(graph, costs, parents)

	//after
	fmt.Println("after, ", costs, parents)
}
