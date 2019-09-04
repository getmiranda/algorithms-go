package graphsearch

//Dfs is una funcion que implementa el algoritmo de busqueda por profundidad
func Dfs(node int, meta int, nodes map[int][]int, fn func(int)) {
	dfsRecursive(node, meta, map[int]bool{}, fn, nodes)
}

func dfsRecursive(node int, meta int, v map[int]bool, fn func(int), nodes map[int][]int) {
	v[node] = true
	fn(node)
	if node == meta {
		return
	}
	for _, n := range nodes[node] {

		if _, ok := v[n]; !ok {

			dfsRecursive(n, meta, v, fn, nodes)

		}
	}

}
