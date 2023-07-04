package main

import (
	"fmt"
	breadthfirstsearch "section6/breadthFirstSearch"
)

func main() {
	var graph = map[string][]string{
		"you": {"alice", "bob", "claire"},
	}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	// fmt.Println(graph)
	found := breadthfirstsearch.BreadthFirstSRCH(graph, "you")
	fmt.Println(found)
}
