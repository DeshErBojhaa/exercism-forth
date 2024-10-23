package forth

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyStack  = errors.New("empty stack")
	ErrIllegalOp   = errors.New("illegal operation")
	ErrUndefinedOp = errors.New("undefined operation")
	ErrInvalidEq   = errors.New("invalid equation")
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
	if errors.Is(err, ErrInvalidEq) {
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
	for i, input := range inputs {
		if d.variables[input] != "" {
			inputs[i] = d.variables[input]
		}
	}
	return nil, nil
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
	}
	return nil, nil
}
