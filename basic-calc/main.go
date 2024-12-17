package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

func main() {
	fmt.Print("Enter calculator input: ")
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.Trim(input, "\n")                   // remove the trailing newline character
	input = strings.Join(strings.Split(input, " "), "") // remove whitespace characters

	p := NewParser()

	if value, err := p.Parse(input); err != nil {
		panic(err)
	} else {
		fmt.Printf("The Value is: %f", sum(value))
	}
}

func sum[T Ordered](stack Stack[T]) T {
	var res T
	for _, v := range stack.values {
		res += v
	}
	return res
}

func In[T comparable](m map[T]bool, v T) bool {
	_, ok := m[v]
	return ok
}

/*
TODO:
------------
HTTP SERVER
------------
1. Set up a server, and make the function available via an endpoint.
2. Properly handle errors and success responses
3. Add DB for storing previous computations from a particular user
4. Add an endpoint for retrieving these computation

---------------
MAIN CALCULATOR
---------------
1. Adding support for parenthesis. So that parenthesis are considered valid characters.
2. Adding support for slightly complex operators like "**" (power) and "%" (modulo).
3. Add support for mathematical computations like sin(x), cos(x), tan(x)
   as well as custom operators like pow(x, y), sqrt(x), cbrt(x) and maybe nroot(x, n).
4. Add support definition of custom equations.
*/
