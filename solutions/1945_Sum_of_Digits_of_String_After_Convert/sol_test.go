package Sum_of_Digits_of_String_After_Convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s, k := "leetcode", 2
	ans := getLucky(s, k)
	assert.Equal(t, 6, ans)
}
