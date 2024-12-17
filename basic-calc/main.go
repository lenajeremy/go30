package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type httpresponse struct {
	Status int         `json:"status"`
	Err    string      `json:"err"`
	Data   interface{} `json:"data"`
}

var p = NewParser()

func main() {

	calculatorHandler := func(w http.ResponseWriter, r *http.Request) {
		var status int
		e := json.NewEncoder(w)
		e.SetIndent("", "   ")

		defer func() {
			log.Printf("%s /calculate ---- status: %d", r.Method, status)
		}()

		originalInput := r.URL.Query().Get("eq")
		if originalInput == "" {
			res := httpresponse{Status: 400, Err: ErrInvalidInput.Error(), Data: nil}
			e.Encode(res)
		}

		input := strings.Trim(originalInput, "\n")          // remove the trailing newline character
		input = strings.Join(strings.Split(input, " "), "") // remove whitespace characters

		if value, err := p.Parse(input); err != nil {
			status = 400
			res := httpresponse{Err: err.Error(), Status: status, Data: nil}
			e.Encode(res)
		} else {
			status = 200
			v := fmt.Sprintf("%s = %.2f", originalInput, sum(value))
			res := httpresponse{Err: "", Status: status, Data: v}
			e.Encode(res)
		}
	}

	http.HandleFunc("/calculate", calculatorHandler)
	fmt.Println("running on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func sum[T Ordered](stack Stack[T]) T {
	var res T
	for _, v := range stack.values {
		res += v
	}
	return res
}

/*
TODO:
------------
HTTP SERVER
------------
1. Set up a server, and make the function available via an endpoint. (DONE✅)
2. Properly handle errors and success responses (DONE✅)
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
