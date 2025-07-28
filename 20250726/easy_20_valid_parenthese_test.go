package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// use stack
func isValid(s string) bool {
	brackets := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	stacks := make([]rune, 0, len(s))
	for _, c := range s {
		if closed, ok := brackets[c]; ok {
			stacks = append(stacks, closed)
			continue
		}
		if len(stacks) == 0 || stacks[len(stacks)-1] != c {
			return false
		}
		stacks = stacks[:len(stacks)-1]
	}
	return len(stacks) == 0
}

func Test_isValid(t *testing.T) {
	assert.Equal(t, true, isValid("([])"))
	assert.Equal(t, false, isValid("(])"))
	assert.Equal(t, false, isValid("]"))
}
