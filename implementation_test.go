package go21

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPositive(c *C) {
	res, err := PostfixToInfix("2 2 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}
func (s *TestSuite) TestPositive2(c *C) {
	res, err := PostfixToInfix("2 2 + 5 * 25 ^ 12 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "((2 + 2) * 5) ^ 25 + 12")
}
func (s *TestSuite) TestPositive3(c *C) {
	res, err := PostfixToInfix("2 5 ^ 8 * 123 -")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 ^ 5 * 8 - 123")
}
func (s *TestSuite) TestPositive4(c *C) {
	res, err := PostfixToInfix("123 2 * 12 / 6 ^")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "(123 * 2 / 12) ^ 6")
}
func (s *TestSuite) TestPositive5(c *C) {
	res, err := PostfixToInfix("8 4 + 12 / 55 * 12 / 75 ^ 54 * 7 / 6 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "((8 + 4) / 12 * 55 / 12) ^ 75 * 54 / 7 + 6")
}

func (s *TestSuite) TestNegative6(c *C) {
	_, err := PostfixToInfix("2 2 5 + 4 7 *")
	c.Assert(err, ErrorMatches, "Input error: the input expression is not correct as a postfix expression")
}
func (s *TestSuite) TestNegative7(c *C) {
	_, err := PostfixToInfix("2 2* + 4 ^")
	c.Assert(err, ErrorMatches, "Input error: every item in input expression must be separated by space character")
}
func (s *TestSuite) TestNegative8(c *C) {
	_, err := PostfixToInfix("2 2 * 5 $")
	c.Assert(err, ErrorMatches, "Input error: the only allowed characters in input exprression are digits and math operators")
}
func (s *TestSuite) TestNegative9(c *C) {
	_, err := PostfixToInfix("")
	c.Assert(err, ErrorMatches, "Input error: the input expression shall not be empty")
}

func (s *TestSuite) TestPositive10(c *C) {
	res, err := PostfixToInfix("7 1 2 * - 5 3 8 + / + 4 ^")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "(7 - 1 * 2 + 5 / (3 + 8)) ^ 4")
}
func (s *TestSuite) TestPositive11(c *C) {
	res, err := PostfixToInfix("5 4 3 - +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "5 + 4 - 3")
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("2 2 + 81 * 12 / 5 + 12 - 1 ^ 12 * 3 +")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// Output:
	// ((2 + 2) * 81 / 12 + 5 - 12) ^ 1 * 12 + 3
}
