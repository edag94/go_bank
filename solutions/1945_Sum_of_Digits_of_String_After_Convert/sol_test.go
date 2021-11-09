package Sum_of_Digits_of_String_After_Convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// s0, k0 := "leetcode", 2
	// ans0 := getLucky(s0, k0)
	// assert.Equal(t, 6, ans0)

	s1, k1 := "dbvmfhnttvr", 5
	ans1 := getLucky(s1, k1)
	assert.Equal(t, 5, ans1)
}
