package main

import (
	"math"
	"math/rand"
	"playground/algorithm/quicksort/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	list := rand.Perm(15)
	t.Log("Before:", list)
	t.Log("After:", core.ToJSONString(QuickSort(list)))

	preValue := math.MinInt
	for _, v := range list {
		assert.GreaterOrEqual(t, v, preValue)
		preValue = v
	}
}

func TestQuickSortEmpty(t *testing.T) {
	t.Log("After:", core.ToJSONString(QuickSort([]int{})))
}

func TestQuickSortOneValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Log("After:", core.ToJSONString(QuickSort([]int{rand.Int()})))
}

func TestQuickSortTwoValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	list := []int{rand.Int(), rand.Int()}
	t.Log("Before:", list)
	t.Log("After:", core.ToJSONString(QuickSort(list)))
}
