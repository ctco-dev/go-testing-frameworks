package main

import (
	"flag"
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	n1 int
	n2 int
}

var _ = Suite(&MySuite{})

var gocheck = flag.Bool("gocheck", false, "Include gocheck tests")

//SetUpSuite will run once when suite starts running
func (s *MySuite) SetUpSuite(c *C) {
	fmt.Println("Setup suite")
	if !*gocheck {
		c.Skip("-gocheck not provided")
	}
}

//SetUpTest will run before each test in a suite
func (s *MySuite) SetUpTest(c *C) {
	fmt.Println("Setup test")
}

//TearDownSuite will run once after suite stops running
func (s *MySuite) TearDownSuite(c *C) {
	fmt.Println("Tear down suite")
}

//SetUpTest will run after each test in a suite
func (s *MySuite) TearDownTest(c *C) {
	fmt.Println("Tear down test")
}

func (s *MySuite) TestSum_GC_Check(c *C) {
	sum := Sum(1, 2)
	c.Check(sum, Equals, uint(2))
	fmt.Println("Printed after a failed check.")
}

func (s *MySuite) TestSum_GC_Assert(c *C) {
	sum := Sum(1, 2)
	c.Assert(sum, Equals, uint(2))
	fmt.Println("This will not be printed after a failed assert.")
}

func (s *MySuite) TestDivide_GC_Error(c *C) {
	_, err := Divide(1, 0)
	c.Check(err, ErrorMatches, `Division.*`)
}

func (s *MySuite) TestDivide_GC_Panic(c *C) {
	c.Check(func() { DivideUnsafe(1, 0) }, PanicMatches, `runtime error.*`)
}

func (s *MySuite) TestDivide_GC_NotNil(c *C) {
	a := struct{}{}
	c.Check(a, Not(IsNil))
}

func (s *MySuite) BenchmarkSum_GC(c *C) {
	for i := 0; i < c.N; i++ {
		_ = Sum(1, 2)
	}
}
