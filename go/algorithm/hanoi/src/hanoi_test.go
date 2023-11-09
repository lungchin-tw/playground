package main

import (
	"playground/algorithm/hanoi/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHanoi_OneDisk(t *testing.T) {
	assert.Equal(t, 1, RunTowerOfHanoi(t, 1))
}

func TestHanoi_TwoDisk(t *testing.T) {
	assert.Equal(t, 3, RunTowerOfHanoi(t, 2))
}

func TestHanoi_ThreeDisk(t *testing.T) {
	assert.Equal(t, 7, RunTowerOfHanoi(t, 3))
}

func TestHanoi_FourDisk(t *testing.T) {
	assert.Equal(t, 15, RunTowerOfHanoi(t, 4))
}

func RunTowerOfHanoi(t *testing.T, numDisk int) int {
	t.Log(core.CurFuncDesc(), "Num Disk:", numDisk)
	algo, err := NewTowerOfHanoi(numDisk, "L", "R", "C")
	assert.NoError(t, err)
	assert.NotNil(t, algo)
	algo.Run()
	t.Log("Move Counts:", algo.MoveCount())
	return algo.MoveCount()
}
