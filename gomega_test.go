package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

//http://onsi.github.io/gomega/
//https://github.com/onsi/gomega
func TestFarmHasCow(t *testing.T) {
	cases := []struct {
		a, b, c uint
	}{
		{1, 1, 2},
		{1, 2, 30},
		{2, 2, 4},
		{5, 2, 70},
	}

	for _, sample := range cases {
		description := fmt.Sprintf("Sum(%v,%v)", sample.a, sample.b)
		t.Run(description, func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(Sum(sample.a, sample.b)).Should(Equal(sample.c))
		})
	}

	//Example without implicit test description. Failed sample number is present in the output
	for _, sample := range cases {
		t.Run("", func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(Sum(sample.a, sample.b)).Should(Equal(sample.c))
		})
	}
}
