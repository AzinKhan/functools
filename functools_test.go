package functools

import (
	"reflect"
	"sort"
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

func TestMapChan(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}
	input := []int{1, 2, 3, 4, 5}

	resultCh := MapChan(double, input)

	var results []int

	for {
		result, ok := <-resultCh
		if !ok {
			break
		}
		results = append(results, result)
	}

	// Result order is not guaranteed so sort here.
	sort.Slice(results, func(i, j int) bool { return results[i] < results[j] })

	if !reflect.DeepEqual(results, []int{2, 4, 6, 8, 10}) {
		t.Fatal(results)
	}
}

func TestMapLazy(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}
	input := []int{1, 2, 3, 4, 5}
	inputCh := BufferChannel(input)

	resultCh := MapLazy(double, inputCh)

	var results []int
	for result := range resultCh {
		results = append(results, result)
	}

	if !reflect.DeepEqual(results, []int{2, 4, 6, 8, 10}) {
		t.Fatal(results)
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

	testcases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Sums correctly",
			input:    []int{6, 2, 3, 4, 5},
			expected: 20,
		},
		{
			name:     "Single element",
			input:    []int{4},
			expected: 4,
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Reduce(sum, tc.input)
			if result != tc.expected {
				t.Fatal(result)
			}
		})
	}

}
