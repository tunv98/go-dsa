package _0250712

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// inplace

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func Test_reverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	assert.Equal(t, s, []byte{'o', 'l', 'l', 'e', 'h'})
}
