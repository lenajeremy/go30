package main

import (
	"errors"
	"testing"
)

type subtest struct {
	input      string
	wants      float64
	shouldFail bool
	expErr     error
}

func (st subtest) run(t *testing.T) {
	res, err := p.Eval(st.input)
	if err != nil {
		if !st.shouldFail {
			t.Errorf("%s failed! expected value: %f, got error: %s", st.input, st.wants, err)
		} else if !errors.Is(err, st.expErr) {
			t.Errorf("%s failed with unexpected error! expected error: %s; got error: %s", st.input, st.expErr, err)
		}
	}
	if st.shouldFail {
		t.Errorf("%s was supposed to fail with error: %s but returned with res: %f", st.input, st.expErr, res)
	} else if res != st.wants {
		t.Errorf("%s failed: expected %f, got %f", st.input, st.wants, res)
	}
}

func TestParser_PurifyInput(t *testing.T) {
	subtests := []struct {
		input   string
		expects string
	}{
		{"2 + 2", "2+2"},
		{"50 * 29 + 10 0.0 - 50", "50*29+100.0-50"},
		{"500 - 500 - 0 - 0.001 + 20 * 5 / 20 - 5 * (1/2)", "500-500-0-0.001+20*5/20-5*(1/2)"},
	}

	for _, st := range subtests {
		res := p.purifyInput(st.input)
		if res != st.expects {
			t.Errorf("purifying input %s failed, expected %s, got %s", st.input, st.expects, res)
		}
	}
}

func TestBasicExpressions(t *testing.T) {
	subTests := []subtest{
		{"2 + 2", 4, false, nil},
		{"4 * 4", 16, false, nil},
		{"7 * 15 - 3", 102, false, nil},
	}

	for _, st := range subTests {
		t.Run(st.input, st.run)
	}
}

func TestOperationsWithParenthesis(t *testing.T) {
	subTests := []subtest{
		{"(2 + 2)", 4, false, nil},
		{"(2 + 2)*15", 60, false, nil},
		{"(0 + 2)- 15 * (2 - 10) * (0 - 50)", -5998, false, nil},
		{"(2.7 + 5.3)*15 + 9.8", 129.8, false, nil},
	}
	for _, st := range subTests {
		t.Run(st.input, st.run)
	}
}

func TestModuleOperator(t *testing.T) {
	subtests := []struct {
		input   string
		expects float64
	}{
		{"5 % 2", 1},
		{"50*20 % 5", 0},
		{"50 * 20 % 3", 1},
		{"50 % 20 * 3", 30},
	}

	for _, st := range subtests {
		res, _ := p.Eval(st.input)
		if res != st.expects {
			t.Errorf("modulo operation %s failed, expected %f, got %f", st.input, st.expects, res)
		}
	}
}

func TestPowerOperator(t *testing.T) {
	subTests := []subtest{
		{"2 ** 2", 4, false, nil},
		{"2 ** 5", 32, false, nil},
		{"5 ** 5", 3125, false, nil},
	}

	for _, st := range subTests {
		t.Run(st.input, st.run)
	}
}

func TestInvalidOperations(t *testing.T) {
	subTests := []subtest{
		{"5 / 0", 0, true, nil},
		{"6 - m", 0, true, ErrInvalidInput},
	}

	for _, st := range subTests {
		t.Run(st.input, st.run)
	}
}
