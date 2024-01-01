package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := os.Args[1]
	calculator(input)
}

// operate contains all the operators that the calculator able to call
func operate(s *stack, op string, input io.Reader, output io.Writer) {
	switch op {
	case "neg":
		push(s, -pop(s))
	case "+":
		push(s, pop(s)+pop(s))
	case "-":
		x := pop(s)
		y := pop(s)
		push(s, x-y)
	case "/":
		x := pop(s)
		y := pop(s)
		push(s, x/y)
	case "*":
		push(s, pop(s)*pop(s))
	case "drop":
		pop(s)
	case "dup":
		x := pop(s)
		push(s, x)
		push(s, x)
	case "summation":
		var x float64
		for !isTheStackEmpty(s) {
			x = x + pop(s)
		}
		push(s, x)
	case "pwr":
		exponent := pop(s)
		base := pop(s)
		res := 1.
		absExp := int(math.Abs(exponent))
		for i := 0; i < absExp; i++ {
			res = res * base
		}
		if exponent < 0 {
			res = 1. / res
		}
		push(s, res)
	case "res":
		res := pop(s)
		fmt.Fprintf(output, ": %v\n", res)
		push(s, res)
	case "read":
		var fl float64
		fmt.Fprintf(output, "enter a number> ")
		fmt.Fscanf(input, "%v", &fl)
		push(s, fl)
	}
}

type stack []float64

func isTheStackEmpty(s *stack) bool {
	return len(*s) == 0
}

func push(s *stack, fl float64) {

	*s = append(*s, fl)
}

func pop(s *stack) float64 {

	i := len(*s) - 1
	fl := (*s)[i]
	(*s) = (*s)[:i]
	return fl
}

// parse pasre the input string into array of any
func parse(s string) []any {

	if len(s) == 0 {
		return nil
	}
	// array of zero elements
	output := []any{}

	strArr := strings.Split(s, " ")

	for i := range strArr {
		if fl, err := strconv.ParseFloat((strArr[i]), 64); err == nil {
			output = append(output, fl)
		} else {
			output = append(output, strArr[i])
		}
	}
	return output
}

// calculator does the actual calculation job
func calculator(s string) float64 {
	input := parse(s)
	st := stack{}
	for _, e := range input {
		fl, ok := e.(float64)
		if ok {
			push(&st, fl)
		} else {
			op, _ := e.(string)
			operate(&st, op, os.Stdin, os.Stdout)
		}
	}

	return pop(&st)
}
