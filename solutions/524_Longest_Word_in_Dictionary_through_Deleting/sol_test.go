package Longest_Word_in_Dictionary_through_Deleting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	ans := findLongestWord(s, dictionary)
	assert.Equal(t, "apple", ans)
}
