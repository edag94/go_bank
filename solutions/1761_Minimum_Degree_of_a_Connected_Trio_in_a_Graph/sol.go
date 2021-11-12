package Minimum_Degree_of_a_Connected_Trio_in_a_Graph

func minTrioDegree(n int, edges [][]int) int {
	// build adjacency matrix and visited

	adjacencyMatrix := make([][]int, n)
	for i := range adjacencyMatrix {
		for j := 0; j < n; j++ {
			adjacencyMatrix[i] = make([]int, n)
		}
	}

	// add all edges to matrix
	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		adjacencyMatrix[x][y] = 1
	}

	// walk the matrix. If
}
