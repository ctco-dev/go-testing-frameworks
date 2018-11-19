package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{5, 2, 70},
	}

	for _, sample := range cases {
		description := fmt.Sprintf("Sum(%v,%v)", sample.a, sample.b)

		//it seems that uint's are printed in hex format (see https://github.com/stretchr/testify/issues/400)
		assert.Equal(sample.c, Sum(sample.a, sample.b), description)
	}
}
