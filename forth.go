package forth

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyStack      = errors.New("empty stack")
	ErrIllegalOp       = errors.New("illegal operation")
	ErrUndefinedOp     = errors.New("undefined operation")
	ErrInvalidEq       = errors.New("invalid equation")
	ErrDivByZero       = errors.New("divide by zero")
	ErrOneValueInStack = errors.New("only one value on the stack")
)

type data struct {
	variables map[string]string
}

func (d *data) solveInput(inputs []string) error {
	va := inputs[0]
	if _, err := strconv.ParseInt(va, 10, 64); err == nil {
		return ErrIllegalOp
	}
	val, err := d.solveEq(inputs[1:])
	if err != nil {
		d.variables[va] = strings.Join(inputs[1:], " ")
		return nil
	}
	value := make([]string, len(val))
	for i, v := range val {
		value[i] = strconv.Itoa(v)
	}
	d.variables[va] = strings.Join(value, " ")

	return nil
}

func (d *data) solveEq(inputs []string) ([]int, error) {
	tmpInp := make([]string, 0)
	for _, input := range inputs {
		if d.variables[input] != "" {
			vv := strings.Split(d.variables[input], " ")
			tmpInp = append(tmpInp, vv...)
			continue
		}
		if _, err := strconv.Atoi(input); err == nil {
			tmpInp = append(tmpInp, input)
			continue
		}
		return nil, ErrUndefinedOp
	}
	inputs = tmpInp
	stack := make([]int, 0)
	for _, input := range inputs {
		if val, err := strconv.Atoi(input); err == nil {
			stack = append(stack, val)
			continue
		}
		switch input {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				if len(stack) == 1 {
					return nil, ErrOneValueInStack
				}
				return nil, ErrEmptyStack
			}
			switch input {
			case "+":
				stack[len(stack)-2] += stack[len(stack)-1]
			case "-":
				stack[len(stack)-2] -= stack[len(stack)-1]
			case "*":
				stack[len(stack)-2] *= stack[len(stack)-1]
			case "/":
				if stack[len(stack)-1] == 0 {
					return nil, ErrDivByZero
				}
				stack[len(stack)-2] /= stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case "dup":
			if len(stack) < 1 {
				return nil, ErrEmptyStack
			}
			stack = append(stack, stack[len(stack)-1])
		case "drop":
			if len(stack) < 1 {
				return nil, ErrEmptyStack
			}
			stack = stack[:len(stack)-1]
		case "swap":
			if len(stack) < 2 {
				if len(stack) == 1 {
					return nil, ErrOneValueInStack
				}
				return nil, ErrEmptyStack
			}
			stack[len(stack)-2], stack[len(stack)-1] = stack[len(stack)-1], stack[len(stack)-2]
		case "over":
			if len(stack) < 2 {
				if len(stack) == 1 {
					return nil, ErrOneValueInStack
				}
				return nil, ErrEmptyStack
			}
			stack = append(stack, stack[len(stack)-2])
		}
	}
	return stack, nil
}

func Forth(input []string) ([]int, error) {
	for i, line := range input {
		input[i] = strings.ToLower(line)
	}
	d := data{
		variables: map[string]string{
			"+":    "+",
			"-":    "-",
			"*":    "*",
			"/":    "/",
			"dup":  "dup",
			"swap": "swap",
			"over": "over",
			"drop": "drop",
		},
	}
	for _, line := range input {
		split := strings.Split(line, " ")
		if split[0] == ":" {
			if err := d.solveInput(split[1 : len(split)-1]); err != nil {
				return nil, err
			}
			continue
		}
		return d.solveEq(split)
	}
	return nil, nil
}
