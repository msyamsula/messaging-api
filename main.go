package main

import (
	"fmt"

	"github.com/msyamsula/messaging-api/graph/service"
	graphSolverO "github.com/msyamsula/messaging-api/graph/service/object"
)

func main() {
	g := service.Graph{
		G: map[string][]string{
			"1": {"2"},
			"2": {"3", "1", "5"},
			"3": {"4", "2"},
			"4": {"5", "3"},
			"5": {"2", "4"},
		},
		E: map[service.Edge]int64{},
		N: map[string]bool{
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
		},
	}
	gs := graphSolverO.New(g)

	// fmt.Println(gs.DFS("1", "5"))
	path, _ := gs.DFS("4", "3")
	fmt.Println(path)
}
