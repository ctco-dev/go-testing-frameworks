package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

//http://onsi.github.io/gomega/
//https://github.com/onsi/gomega
func TestFarmHasCow(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(Sum(1,2)).Should(Equal(uint(30)))
	
	//won't be called
	g.Expect(Sum(1,2)).Should(Equal(uint(29)))
}

