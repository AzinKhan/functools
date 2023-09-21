package functools

import "sync"

// Map applies the given function to every element in the provided slice.
func Map[A any, B any](fn func(A) B, as []A) []B {
	results := make([]B, len(as))
	for i, a := range as {
		results[i] = fn(a)
	}
	return results
}

// MapAsync applies the given function to each element in the provided slice in
// parallel, with each element mapped in a separate goroutine. The ordering of
// the input slice is maintained in the output.
func MapAsync[A any, B any](fn func(A) B, as []A) []B {
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

// MapChan provides a similar functionality to MapAsync except the results are
// returned via a channel, instead of being collected as a slice. As a result,
// this function is non-blocking. The returned channel is read-only and is
// closed once all the elements have been mapped over. The ordering of the
// input elements is not maintainted in the output; the results are written to
// the channel as they arrive.
func MapChan[A any, B any](fn func(A) B, as []A) <-chan B {
	n := len(as)
	results := make(chan B, n)

	go func() {
		var wg sync.WaitGroup
		wg.Add(n)
		for _, a := range as {
			a := a
			go func() {
				defer wg.Done()
				results <- fn(a)
			}()
		}
		wg.Wait()
		close(results)
	}()

	return results
}

// MapLazy maps the given function over the elements received on the channel
// `as`. The results are written to the returned channel. The caller is assumed
// to control the closing of the input channel. Once that is closed, the
// returned results channel is also closed. The returned results channel is not
// buffered.
func MapLazy[A any, B any](fn func(A) B, inCh <-chan A) <-chan B {
	results := make(chan B)
	go func() {
		for {
			a, ok := <-inCh
			if !ok {
				close(results)
				return
			}
			results <- fn(a)
		}
	}()

	return results
}

// Filter returns the provided slice with any elements not satisfying fn
// removed. The resulting slice can be smaller than the input slice. A new
// slice is created for the purposes of this function, the original slice is
// not modified.
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

// Reduce applies fn to the given slice's elements cumulatively. If an empty
// list is passed, then the zero-value of the type A is returned.
func Reduce[A any](fn func(A, A) A, as []A) A {
	var result A
	if len(as) == 0 {
		return result
	}
	result = as[0]
	for _, a := range as[1:] {
		result = fn(result, a)
	}
	return result
}

// BufferChannel takes a slice and returns a buffered, closed, channel of the
// elements therein.
func BufferChannel[A any](as []A) <-chan A {
	ch := make(chan A, len(as))
	for _, a := range as {
		ch <- a
	}
	close(ch)
	return ch
}

// UnbufferChannel reads elements that come through the given channel and puts
// them into a slice. It blocks until the channel is closed.
func UnbufferChannel[A any](ch <-chan A) []A {
	var result []A
	for a := range ch {
		result = append(result, a)
	}
	return result
}

// FindFirst returns a pointer to the first element that matches the condition
// If no elements match the condition then a sentinel value of nil is returned.
func FindFirst[A any](condition func(A) bool, as []A) *A {
	for _, a := range as {
		if condition(a) {
			return &a
		}
	}
	return nil
}
