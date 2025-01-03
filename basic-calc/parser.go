package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return new(Parser)
}

func (p *Parser) purifyInput(original string) string {
	input := strings.Trim(original, "\n")               // remove the trailing newline character
	input = strings.Join(strings.Split(input, " "), "") // remove whitespace characters
	return input
}

func (p *Parser) Eval(expression string) (float64, error) {
	expression = p.purifyInput(expression)
	s, err := parseStartEnd(expression, 0, len(expression)-1)
	if err != nil {
		return 0, err
	}
	return sum(s), nil
}

func parseStartEnd(expression string, start, end int) (Stack[float64], error) {
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
		'%': true,
		'.': true,
	}

	var paren = map[rune]bool{
		'(': true,
		')': true,
	}

	var functions = map[string]bool{
		"sin":   true,
		"cos":   true,
		"tan":   true,
		"pow":   true,
		"log":   true,
		"log10": true,
		"ln":    true,
		"sqrt":  true,
		"cbrt":  true,
	}

	// insert all the keywords into the trie for easy search
	var trie = NewTrie()
	for k := range functions {
		trie.Insert(k)
	}

	for _, ch := range expression {
		if !(In(digits, ch) || In(signs, ch) || In(paren, ch) || (ch >= 97 && ch < 123)) {
			return stack, fmt.Errorf("invalid character: '%c'", ch)
		}
	}

	// evaluate the input
	var curr float64
	var hasDecimal bool
	var decimalPlaces int
	var prevSign = "+"
	var kw string
	var kwValid bool

	for i := start; i <= end; i++ {
		ch := rune(expression[i])
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

				if (ch == '*' || ch == '/') && stack.Peek() == nil && curr == 0 {
					return stack, fmt.Errorf("%s: invalid left value for division/multiplication operations", ErrInvalidInput.Error())
				}

				if curr != 0 {
					updateStack(&stack, curr, prevSign)
					curr = 0
				}

				if ch == '*' && i != end && rune(expression[i+1]) == '*' {
					fmt.Println(i, expression, curr, stack)
					// we have a ** operator
					prevSign = "**"
					i += 1
				} else {
					prevSign = string(ch)
				}
			}
		} else if ch == '(' {
			open := 0
			j := i
		closingParenFinder:
			for j <= end {
				if expression[j] == '(' {
					open += 1
				} else if expression[j] == ')' {
					open -= 1
				}

				if open == 0 {
					res, err := parseStartEnd(expression, i+1, j-1)
					if err != nil {
						return stack, err
					}

					s := sum(res)
					if kwValid {
						prevSign = kw
						kw = ""
					}
					updateStack(&stack, s, prevSign)
					i = j
					break closingParenFinder
				} else {
					j += 1
				}
			}

		} else if ch == ')' {
			return stack, fmt.Errorf("%s: closing parenthesis found at index %d without an opening parenthesis", ErrInvalidInput.Error(), i)
		} else {
			kw += string(ch)
			kwValid = In(functions, kw) && expression[i+1] == '('
			println(trie.Has(kw), string(ch))
		}

		if i == end && curr > 0 {
			updateStack(&stack, curr, prevSign)
		}
	}

	return stack, nil
}

func updateStack(stack *Stack[float64], curr float64, sign string) {
	fmt.Println(sign, curr)
	if sign == "+" {
		stack.Push(curr)
	} else if sign == "-" {
		stack.Push(-curr)
	} else if sign == "*" {
		last := stack.Pop()
		if last != nil {
			stack.Push(*last * curr)
		}
	} else if sign == "/" {
		last := stack.Pop()
		if last != nil {
			stack.Push(*last / curr)
		}
	} else if sign == "%" {
		last := stack.Pop()
		if last != nil {
			stack.Push(math.Mod(*last, curr))
		}
	} else if sign == "**" {
		last := stack.Pop()
		if last != nil {
			stack.Push(math.Pow(*last, curr))
		}
	} else if sign == "sin" {
		stack.Push(math.Sin(curr))
	} else if sign == "cos" {
		stack.Push(math.Cos(curr))
	} else if sign == "tan" {
		stack.Push(math.Tan(curr))
	} else if sign == "ln" || sign == "log" {
		stack.Push(math.Log(curr))
	} else if sign == "log10" {
		stack.Push(math.Log10(curr))
	} else if sign == "log2" {
		stack.Push(math.Log2(curr))
	} else if sign == "sqrt" {
		stack.Push(math.Sqrt(curr))
	} else if sign == "cbrt" {
		stack.Push(math.Cbrt(curr))
	}
}

func In[T comparable](m map[T]bool, v T) bool {
	_, ok := m[v]
	return ok
}

func sum[T Ordered](stack Stack[T]) T {
	var res T
	for _, v := range stack.values {
		res += v
	}
	return res
}
