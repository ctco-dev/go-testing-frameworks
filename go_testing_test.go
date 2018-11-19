package main

import (
	"testing"
)

//https://blog.alexellis.io/golang-writing-unit-tests/
//https://golang.org/pkg/testing/

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSumMultiple(t *testing.T) {

	cases := []struct {
		a, b, c uint
	}{
		{1, 1, 2},
		{1, 2, 30},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, sample := range cases {
		total := Sum(sample.a, sample.b)
		if total != sample.c {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", sample.a, sample.b, total, sample.c)
		}
	}
}
