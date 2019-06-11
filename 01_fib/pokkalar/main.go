package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	sign, prev, curr := 1, 0, 1
	if n < 0 {
		sign = -1
	}
	for i := 0; i < n*sign; i++ {
		fmt.Fprintln(out, curr)
		sum := prev + curr*sign
		prev = curr
		curr = sum
	}
}

func main() {
	fib(7)
}
