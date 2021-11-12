package Minimum_Degree_of_a_Connected_Trio_in_a_Graph

func minTrioDegree(n int, edges [][]int) int {
	// build adjacency matrix and visited

	adjacencyMatrix := make([][]int, n)
	visitedMatrix := make([][]int, n)

	for i := range adjacencyMatrix {
		for j := 0; j < n; j++ {
			adjacencyMatrix[i] = make([]int, n)
			visitedMatrix[i] = make([]int, n)
			for k := 0; k < n; k++ {
				adjacencyMatrix[i][k] = 0
				visitedMatrix[i][k] = 0
			}
		}
	}

	// add all edges to matrix
	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		adjacencyMatrix[x][y] = 1
	}

	// walk the matrix. If a 1, you have 2 edges, so check 3rd for a trio. Check all other nodes to see if 1s.
}
