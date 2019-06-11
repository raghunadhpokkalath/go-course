package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	var b bytes.Buffer
	out = &b
	expected := "1\n1\n2\n3\n5\n8\n13\n"
	main()
	actual := b.String()
	fmt.Print(actual)
	if expected != actual {
		t.Errorf("Results are not matching exptected %v,actual %v", expected, actual)

	}
}

func TestFib(t *testing.T) {
	var b bytes.Buffer
	out = &b
	testRecords := []struct {
		n        int
		expected string
	}{
		{6, "1\n1\n2\n3\n5\n8\n"},
		{-7, "1\n-1\n2\n-3\n5\n-8\n13\n"},
		{0, ""},
		{1, "1\n"},
	}

	for _, testRec := range testRecords {

		fib(testRec.n)
		actual := b.String()
		if actual != testRec.expected {

			t.Errorf("Results are not matching exptected %v,actual %v", testRec.expected, actual)

		}
		b.Reset()
	}

}
