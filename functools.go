package functools

import "sync"

// Map applies the given function to every element in the provided slice.
func Map[A any, B any](fn func(A) B, as []A)  []B{
	results := make([]B, len(as))
	for i, a := range as {
		results[i] = fn(a)
	}
	return results
}

// MapAsync applies the given function to each element in the provided slice in
// parallel, with each element mapped in a separate goroutine.
func MapAsync[A any, B any](fn func(A) B, as []A) []B{
	n := len(as)
	results := make([]B, n)

	var wg sync.WaitGroup
	wg.Add(n)

	for i, a := range as {
		i := i
		a := a
		go func() {
			defer wg.Done()
			results[i] = fn(a)
		}()
	}

	wg.Wait()
	return results
}

// Filter returns the provided slice with any elements not satisfying fn
// removed. The resulting slice can be smaller than the input slice.
func Filter[A any](fn func(A) bool, as []A) []A {
	results := make([]A, 0, len(as))

	for _, a := range as {
		if !fn(a) {
			continue
		}
		results = append(results, a)
	}
	return results
}

// Reduce applies fn to the given slice's elements cumulatively.
func Reduce[A any](fn func(A, A) A, as []A) A {
	var result A
	if len(as) == 0 {
		return result
	}
	for _, a := range as {
		result = fn(result, a)
	}
	return result
}