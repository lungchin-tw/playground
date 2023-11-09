package main

import (
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	innerErr := fmt.Errorf("Inner Error")
	outterErr := fmt.Errorf("Outter Error: %w", innerErr)
	t.Log(outterErr)
}
