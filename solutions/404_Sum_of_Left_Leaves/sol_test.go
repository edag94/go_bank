package Sum_of_Left_Leaves

import (
	"testing"

	utils "github.com/edag94/go_bank/utils"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	treeArray := []int{1, 2, 3, 4, 5}
	input := utils.IntArrayToTreeNode(treeArray)
	sum := sumOfLeftLeaves(input)
	assert.Equal(t, sum, 4)
}
