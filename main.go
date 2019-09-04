package main

import (
	"fmt"

	"github.com/getmiranda/algorithms/graphsearch"
)

//Graph: http://graphonline.ru/en/?graph=wxniyeUKFetPwYFU
var nodes = map[int][]int{
	0: []int{3},
	1: []int{3, 5},
	2: []int{4, 5},
	3: []int{0, 1, 5, 7},
	4: []int{2, 6},
	5: []int{1, 2, 3, 6},
	6: []int{4, 5},
	7: []int{3},
}

func main() {

	//BUSQUEDA POR AMPLITUD
	visited := []int{}

	graphsearch.Bfs(0, 2, nodes, func(node int) {
		visited = append(visited, node)
	})

	fmt.Println("BFS-> ", visited)

	//BUSQUEDA POR PROFUNDIDAD
	visited = []int{}

	graphsearch.Dfs(0, 5, nodes, func(node int) {
		visited = append(visited, node)
	})

	fmt.Println("DFS-> ", visited)

	// //BUSQUEDA POR PROFUNDIDAD TWO
	fmt.Println("DFS TWO")
	v0 := graphsearch.NewVertex("0")
	v1 := graphsearch.NewVertex("1")
	v2 := graphsearch.NewVertex("2")
	v3 := graphsearch.NewVertex("3")
	v4 := graphsearch.NewVertex("4")
	v5 := graphsearch.NewVertex("5")
	v6 := graphsearch.NewVertex("6")
	v7 := graphsearch.NewVertex("7")

	g := graphsearch.Graph{}

	v0.Connect(v3)
	v1.Connect(v3, v5)
	v2.Connect(v4, v5)
	v3.Connect(v0, v1, v5, v7)
	v4.Connect(v2, v6)
	v5.Connect(v1, v2, v3, v6)
	v6.Connect(v4, v5)
	v7.Connect(v3)
	g.Dfs(v0, v3)

	//BUSQUEDA POR A*
	fmt.Println("A*")
	mapData := graphsearch.Read_map(graphsearch.MAP0) //Read map data and create the map_data
	graph := graphsearch.NewGraph(mapData)            //Create a new graph
	nodesPath := graphsearch.Astar(graph)             //Get the shortest path

	//Returns a list of nodes from START to STOP avoiding all obstacles if possible

	fmt.Println(nodesPath)
}
