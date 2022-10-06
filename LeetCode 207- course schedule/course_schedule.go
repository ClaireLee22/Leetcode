package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	fmt.Printf("adjacency list: %+v\n\n", graph)

	visiting := map[int]bool{}
	visited := map[int]bool{}

	for node := 0; node < numCourses; node += 1 {
		fmt.Println("------------------------------------------")
		fmt.Println("dfs node ", node)
		if detectCycle(graph, node, visiting, visited) {
			return false
		}
	}

	return true
}

func detectCycle(
	graph map[int][]int,
	node int,
	visiting, visited map[int]bool) bool {
	if _, found := visited[node]; found {
		fmt.Printf("node %d is already visited(black) -> skip\n", node)
		return false
	}

	if _, found := visiting[node]; found {
		fmt.Printf("node %d is in visiting(gray) -> a cycle is detected\n\n", node)
		return true
	}

	visiting[node] = true
	fmt.Printf("nodes in visiting(gray): %+v\n", visiting)
	fmt.Printf("nodes in visited(black): %+v\n\n", visited)
	for _, descendant := range graph[node] {
		fmt.Printf("current node: node %d\n", node)
		fmt.Printf("visit descendant: node %d\n", descendant)
		if detectCycle(graph, descendant, visiting, visited) {
			return true
		}
	}

	delete(visiting, node)
	visited[node] = true
	fmt.Printf("finish explore node %d\n", node)
	fmt.Printf("nodes in visiting(gray): %+v\n", visiting)
	fmt.Printf("nodes in visited(black): %+v\n\n", visited)
	return false
}

func buildGraph(numNodes int, prerequisites [][]int) map[int][]int {
	graph := map[int][]int{}
	for i := 0; i < numNodes; i++ {
		graph[i] = []int{}
	}

	for _, prerequisite := range prerequisites {
		a, b := prerequisite[0], prerequisite[1]
		graph[b] = append(graph[b], a)
	}

	return graph
}

func main() {
	numCourses := 6
	prerequisites := [][]int{
		{1, 0},
		{3, 0},
		{5, 0},
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 4},
		{3, 5},
	}
	fmt.Println("can finish courses?", canFinish(numCourses, prerequisites))
}

/*
output:
adjacency list: map[0:[1 3 5] 1:[2 3] 2:[] 3:[4] 4:[5] 5:[3]]

------------------------------------------
dfs node  0
nodes in visiting(gray): map[0:true]
nodes in visited(black): map[]

current node: node 0
visit descendant: node 1
nodes in visiting(gray): map[0:true 1:true]
nodes in visited(black): map[]

current node: node 1
visit descendant: node 2
nodes in visiting(gray): map[0:true 1:true 2:true]
nodes in visited(black): map[]

finish explore node 2
nodes in visiting(gray): map[0:true 1:true]
nodes in visited(black): map[2:true]

current node: node 1
visit descendant: node 3
nodes in visiting(gray): map[0:true 1:true 3:true]
nodes in visited(black): map[2:true]

current node: node 3
visit descendant: node 4
nodes in visiting(gray): map[0:true 1:true 3:true 4:true]
nodes in visited(black): map[2:true]

current node: node 4
visit descendant: node 5
nodes in visiting(gray): map[0:true 1:true 3:true 4:true 5:true]
nodes in visited(black): map[2:true]

current node: node 5
visit descendant: node 3
node 3 is in visiting(gray) -> a cycle is detected
*/
