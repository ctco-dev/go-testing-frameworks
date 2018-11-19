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

	//Example without implicit test description. It's hard to tell which sample failed.
	for _, sample := range cases {
		assert.Equal(sample.c, Sum(sample.a, sample.b))
	}

	//Zero benefits in running subtests. Description is not displayed in the output.
	for _, sample := range cases {
		description := fmt.Sprintf("Sum(%v,%v)", sample.a, sample.b)
		t.Run(description, func(t *testing.T) {
			assert.Equal(sample.c, Sum(sample.a, sample.b))
		})
	}
}
