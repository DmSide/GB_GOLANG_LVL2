package _defer

import (
	"errors"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestSafeDivision(c *C) {
	var err error
	err = SafeDivision()
	c.Assert(err, NotNil)
	err2 := NewErrorWithTrace("division by zero")
	c.Assert(errors.Is(err, err2), Equals, true)
}

