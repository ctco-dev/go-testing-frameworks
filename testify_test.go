package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://github.com/stretchr/testify
func TestSomething(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		a, b, c uint
	}{
		{1, 1, 2},
		{1, 2, 30},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, sample := range cases {
		assert.Equal(Sum(sample.a, sample.b), sample.c, "they should be equal")
	}
}
