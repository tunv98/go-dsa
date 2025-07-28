package _0250726

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// brute force
func maxprofitBruteforce(prices []int) int {
	var maxP = 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				continue
			}
			temp := prices[j] - prices[i]
			if temp > maxP {
				maxP = temp
			}
		}
	}
	return maxP
}

// skip unnecessary couple
func maxProfit(prices []int) int {
	minIndex := 0
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[minIndex] > prices[i] {
			minIndex = i
			continue
		}
		temp := prices[i] - prices[minIndex]
		if maxP < temp {
			maxP = temp
		}
	}
	return maxP
}

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, maxprofitBruteforce([]int{7, 1, 5, 3, 6, 4}), 5)
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
}
