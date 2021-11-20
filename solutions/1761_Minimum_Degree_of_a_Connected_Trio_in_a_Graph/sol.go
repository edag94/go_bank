package Minimum_Degree_of_a_Connected_Trio_in_a_Graph

import "math"

// fail. This is much more complicated than thought
func minTrioDegree(n int, edges [][]int) int {
	lowestDegree := math.MaxInt

	// build adjacency matrix and visited

	adjacencyMatrix := make([][]bool, n)

	for i := range adjacencyMatrix {
		for j := 0; j < n; j++ {
			adjacencyMatrix[i] = make([]bool, n)
			for k := 0; k < n; k++ {
				adjacencyMatrix[i][k] = false
			}
		}
	}

	// add all edges to matrix
	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		adjacencyMatrix[x-1][y-1] = true // make it 0 indexed
		adjacencyMatrix[y-1][x-1] = true
	}

	// walk the matrix. If a 1, you have 2 edges, so check 3rd for a trio. Check all other nodes to see if 1s. Walk only half of matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if adjacencyMatrix[i][j] {
				// we have 2 nodes, now check all for 1s
				degree := 0
				for k := 0; k < n; k++ {
					if adjacencyMatrix[i][k] && adjacencyMatrix[j][k] {

					} else {
						degree += 1
					}
				}
				if degree != 0 {
					lowestDegree = MinOf(lowestDegree, degree)
				}
			}
		}
	}

	return lowestDegree
}

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
