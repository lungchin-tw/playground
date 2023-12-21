package generic_test

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func SumOfNumbers[T Number](a T, b T) T {
	return a + b
}

func TestGeneric(t *testing.T) {
	{
		x := SumOfNumbers(5, 6)
		t.Logf("Generic SumOfNumbers: %v, Type: %v", x, reflect.TypeOf(x))
	}

	{
		x := SumOfNumbers(5.1, 6.2)
		t.Logf("Generic SumOfNumbers: %v, Type: %v", x, reflect.TypeOf(x))
	}
}
