package go21

import (
	"fmt"
	"strings"
)

const (
	operators = "+-*/^"
	digits    = "1234567890"
	EXPR      = 0
	SIGN      = 1
	TIER1     = 1
	TIER2     = 2
	TIER3     = 3
	PLAIN     = 100
)

type Atom struct {
	Value string
	Type  int
	Sign  int
}

//PostfixToInfix converts math statement
func PostfixToInfix(input string) (string, error) {
	var err error = nil
	if input == "" {
		return "", fmt.Errorf("Input error: the input expression shall not be empty")
	}
	var arr = strings.Split(input, " ")
	if err := testArray(arr); err != nil {
		return "", err
	}

	atoms := []Atom{}
	for _, str := range arr {
		if strings.ContainsAny(str, digits) {
			atoms = append(atoms, Atom{str, EXPR, PLAIN})
		} else {
			if strings.ContainsAny(str, "+-") {
				atoms = append(atoms, Atom{str, SIGN, TIER1})
			} else if strings.ContainsAny(str, "*/") {
				atoms = append(atoms, Atom{str, SIGN, TIER2})
			} else if strings.ContainsAny(str, "^") {
				atoms = append(atoms, Atom{str, SIGN, TIER3})
			}
		}
	}
	i, err := next(atoms)
	for i != 0 && err == nil {
		if atoms[i-2].Sign < atoms[i].Sign {
			atoms[i-2].Value = "(" + atoms[i-2].Value + ")"
		}
		if atoms[i-1].Sign < atoms[i].Sign {
			atoms[i-1].Value = "(" + atoms[i-1].Value + ")"
		}
		atoms[i-2].Value = atoms[i-2].Value + " " + atoms[i].Value + " " + atoms[i-1].Value
		atoms[i-2].Sign = atoms[i].Sign
		atoms = remove(atoms, i-1)
		atoms = remove(atoms, i-1)
		i, err = next(atoms)
	}
	return atoms[0].Value, err
}

func next(atoms []Atom) (int, error) {
	var err error = nil
	if len(atoms) == 1 {
		return 0, nil
	}
	i := 2
	for err == nil && !(atoms[i-2].Type == EXPR && atoms[i-1].Type == EXPR && atoms[i].Type == SIGN) {
		i++
		if i == len(atoms) {
			err = fmt.Errorf("Input error: the input expression is not correct as a postfix expression")
		}
	}
	return i, err
}

func remove(atoms []Atom, i int) []Atom {
	return append(atoms[:i], atoms[i+1:]...)
}

func testArray(arr []string) error {
	if len(arr)%2 == 0 {
		return fmt.Errorf("Input error: the input expression is not correct as a postfix expression")
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if !strings.ContainsAny(string(arr[i][j]), operators+digits) {
				return fmt.Errorf("Input error: the only allowed characters in input exprression are digits and math operators")
			}
		}
		if strings.ContainsAny(arr[i], operators) &&
			strings.ContainsAny(arr[i], digits) {
			return fmt.Errorf("Input error: every item in input expression must be separated by space character")
		}

	}
	return nil
}
