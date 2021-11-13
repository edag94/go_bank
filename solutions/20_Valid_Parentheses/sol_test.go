package Valid_Parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := "([)]"
	ans := isValid(s)
	assert.Equal(t, false, ans)
}
