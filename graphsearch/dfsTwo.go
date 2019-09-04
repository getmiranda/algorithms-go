package graphsearch

import (
	"fmt"
)

//Vertex representa un vértice de un grafo
type Vertex struct {
	visited    bool
	value      string
	neighbours []*Vertex
}

//NewVertex es una funcion que crea un nuevo vertice
func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
	}
}

//Connect es una funcion que conecta un vertice con n vertices
func (v *Vertex) Connect(vertex ...*Vertex) { // see comment 4.
	v.neighbours = append(v.neighbours, vertex...)
}

//Graph representa la estructura de un grafo
type Graph struct{}

//Dfs es la funcion que implementa el algoritmo de busqueda por profundidad
func (g *Graph) Dfs(start, meta *Vertex) {
	if start.visited {
		return
	}

	start.visited = true
	fmt.Println(start.value)
	if start.value == meta.value {
		return
	}
	for _, v := range start.neighbours {
		g.Dfs(v, meta)
	}
}
