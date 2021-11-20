package Minimum_Degree_of_a_Connected_Trio_in_a_Graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	n := 6
	edges := [][]int{[]int{1, 2}, []int{1, 3}, []int{3, 2}, []int{4, 1}, []int{5, 2}, []int{3, 6}}
	sol := minTrioDegree(n, edges)
	assert.Equal(t, 3, sol)
}
