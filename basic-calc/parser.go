package main

import (
	"fmt"
	"math"
	"strconv"
)

type Parser struct{}

func (p *Parser) Parse(expression string) (Stack[float64], error) {
	var stack Stack[float64]
	var digits = map[rune]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'0': true,
	}
	var signs = map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
		'.': true,
	}

	for _, ch := range expression {
		if !(In(digits, ch) || In(signs, ch)) {
			return stack, fmt.Errorf("invalid character: '%c'", ch)
		}
	}

	// evaluate the input
	var curr float64
	var hasDecimal bool
	var decimalPlaces int
	var prevSign rune = '+'

	for i, ch := range expression {
		if In(digits, ch) {
			chInt, _ := strconv.Atoi(string(ch))
			if hasDecimal {
				decimalPlaces++
				curr = curr + (float64(chInt) / math.Pow(10, float64(decimalPlaces)))
			} else {
				curr = curr*10 + float64(chInt)
			}
		} else if In(signs, ch) {
			if ch == '.' {
				if hasDecimal {
					return stack, fmt.Errorf("%s: encountered multiple decimal in a number", ErrInvalidInput.Error())
				}
				hasDecimal = true
			} else {
				hasDecimal = false
				decimalPlaces = 0

				if (ch == '*' || ch == '/') && curr == 0 {
					return stack, fmt.Errorf("%s: multiplication/division should not operate on left value of zero", ErrInvalidInput.Error())
				}

				if prevSign == '+' {
					stack.Push(curr)
				} else if prevSign == '-' {
					stack.Push(-curr)
				} else if prevSign == '*' {
					prev := stack.Pop()
					stack.Push(curr * (*prev))
				} else if prevSign == '/' {
					prev := stack.Pop()
					stack.Push(curr / *prev)
				}
				curr = 0
				prevSign = ch
			}
		} else {
			continue
		}

		if i == len(expression)-1 && curr > 0 {
			stack.Push(curr)
		}
	}

	return stack, nil
}

func NewParser() *Parser {
	return new(Parser)
}
