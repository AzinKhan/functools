package functools

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}
	input := []int{1, 2, 3, 4, 5}

	result := Map(double, input)

	if !reflect.DeepEqual(result, []int{2, 4, 6, 8, 10}) {
		t.Fail()
	}
}

func TestMapAsync(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}
	input := []int{1, 2, 3, 4, 5}

	result := MapAsync(double, input)

	if !reflect.DeepEqual(result, []int{2, 4, 6, 8, 10}) {
		t.Fatal(result)
	}
}

func TestFilter(t *testing.T) {
	isEven := func(x int) bool {
		return x%2 == 0
	}
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := Filter(isEven, input)

	if !reflect.DeepEqual(result, []int{2, 4, 6, 8, 10}) {
		t.Fatal(result)
	}
}

func TestReduce(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}
	input := []int{1, 2, 3, 4, 5}

	result := Reduce(sum, input)

	if result != 15 {
		t.Fatal(result)
	}

}
