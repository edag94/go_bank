package Minimum_Number_of_Increments_on_Subarrays_to_Form_a_Target_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var testArray []int = []int{3, 1, 5, 4, 2}
	ans := minNumberOperations(testArray)
	assert.Equal(t, 7, ans)
}
