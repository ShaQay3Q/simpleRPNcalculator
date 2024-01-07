package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open(os.Args[1])
	defer file.Close()

	calculateFromFile(file, os.Stdin, os.Stdout)
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
		// printIt supposed to print to a standard output
	case "printIt":
		topValue := pop(s)
		fmt.Fprintf(output, ": %v\n", topValue)
		push(s, topValue)
		// read supposed to read a number form a standard input
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

// calculate does the actual calculation job
func calculate(s string, input io.Reader, output io.Writer) float64 {
	parsedInput := parse(s)
	st := stack{}
	for _, e := range parsedInput {
		fl, ok := e.(float64)
		if ok {
			push(&st, fl)
		} else {
			op, _ := e.(string)
			operate(&st, op, input, output)
		}
	}

	return pop(&st)
}

func calculateFromFile(reader io.Reader, input io.Reader, output io.Writer) {
	fileContent, _ := ioutil.ReadAll(reader)

	s := string(fileContent)
	calculate(s, input, output)
}
