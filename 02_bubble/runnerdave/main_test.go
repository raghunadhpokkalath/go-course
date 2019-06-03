package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`[1 2 3 5]`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestSort(t *testing.T) {
	tables := []struct {
		x []int
		n []int
	}{
		{[]int{1, 3, 2, 5}, []int{1, 2, 3, 5}},
		{[]int{0, 3, 2, 5}, []int{0, 2, 3, 5}},
		{[]int{1, 3, 8, 5}, []int{1, 3, 5, 8}},
		{[]int{-1, -3, -2, -5}, []int{-5, -3, -2, -1}},
		{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{[]int{-1, 1, 0}, []int{-1, 0, 1}},
	}

	for _, table := range tables {
		var unsorted [4]int
		copy(unsorted[:], table.x)
		sorted := bubble(table.x)
		if !equal(sorted, table.n) {
			t.Errorf("Sorting of (%v) was incorrect, got: %v, want: %v.", unsorted, sorted, table.n)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
