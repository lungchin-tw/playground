package main

import "fmt"

func QuickSort(list []int) []int {
	r := NewQuickSortRunner(list)
	r.Sort()
	fmt.Println("MoveCount:", r.MoveCount())
	fmt.Println("SwapCount:", r.SwapCount())
	return list
}

type quickSortRunner struct {
	list      []int
	moveCount int
	swapCount int
}

func (r *quickSortRunner) MoveCount() int {
	return r.moveCount
}

func (r *quickSortRunner) SwapCount() int {
	return r.swapCount
}

func NewQuickSortRunner(list []int) *quickSortRunner {
	return &quickSortRunner{
		list: list,
	}
}

func (r *quickSortRunner) Sort() {
	r.sort(0, len(r.list)-1)
}

func (r *quickSortRunner) sort(startIndex, endIndex int) {
	fmt.Println("list:", r.list)
	fmt.Printf("StarIndex: %v, EndIndex: %v\n", startIndex, endIndex)
	if len(r.list) == 0 {
		return
	} else if (startIndex < 0) || (startIndex >= len(r.list)) {
		return
	} else if (endIndex < 0) || (endIndex >= len(r.list)) {
		return
	} else if endIndex <= startIndex {
		return
	}

	curIndex := startIndex - 1
	swapMarker := curIndex
	pivotValue := r.list[endIndex]
	for curIndex < endIndex {
		curIndex++
		r.moveCount++
		if r.list[curIndex] > pivotValue {
			continue
		}

		swapMarker++
		if swapMarker >= curIndex {
			continue
		}

		r.list[swapMarker], r.list[curIndex] = r.list[curIndex], r.list[swapMarker]
		r.swapCount++

	}

	fmt.Println("SwapMarker:", swapMarker)

	// left Half
	r.sort(startIndex, swapMarker-1)

	// Right Half
	r.sort(swapMarker+1, endIndex)
}
