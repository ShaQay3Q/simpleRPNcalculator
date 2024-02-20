package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// main is the enterypoint of this program
func main() {
	// in order to run this program it should be given the file path/filename
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	_ = calculateFromFile(file, os.Stdin, os.Stdout)
}

var operations = map[string]func(*stack, io.Reader, io.Writer) error{
	"neg":       negOperator,
	"+":         sumOperator,
	"-":         minusOperator,
	"/":         divisionOperator,
	"*":         multipicationOperator,
	"drop":      dropOperator,
	"dup":       duplicateOperator,
	"summation": summationOperator,
	"pwr":       powerOperator,
	"printIt":   writeOperator,
	"read":      readOperator,
}

// operate contains all the operators that the calculator able to call
// operate choses an operation and execcutes/calls it
// its is refactored
func operate(s *stack, op string, input io.Reader, output io.Writer) error {

	operation, ok := operations[op]

	if ok {
		err := operation(s, input, output)
		return err
	}

	return nil
}

func readOperator(s *stack, input io.Reader, output io.Writer) error {
	var fl float64
	fmt.Fprintf(output, "enter a number> ")
	fmt.Fscanf(input, "%v", &fl)
	push(s, fl)
	return nil
}

func writeOperator(s *stack, _ io.Reader, output io.Writer) error {
	topValue := pop(s)
	fmt.Fprintf(output, ": %v\n", topValue)
	push(s, topValue)
	return nil
}

func powerOperator(s *stack, input io.Reader, output io.Writer) error {
	exponent := pop(s)
	base := pop(s)
	res := 1.
	absExp := int(math.Abs(exponent))
	for i := 0; i < absExp; i++ {
		res *= base
	}
	if exponent < 0 {
		res = 1. / res
	}
	push(s, res)

	return nil
}

func summationOperator(s *stack, input io.Reader, output io.Writer) error {
	var x float64
	for !isTheStackEmpty(s) {
		x += pop(s)
	}
	push(s, x)

	return nil
}

func duplicateOperator(s *stack, input io.Reader, output io.Writer) error {
	x := pop(s)
	push(s, x)
	push(s, x)

	return nil
}

func dropOperator(s *stack, input io.Reader, output io.Writer) error {
	pop(s)

	return nil
}

func multipicationOperator(s *stack, input io.Reader, output io.Writer) error {
	push(s, pop(s)*pop(s))

	return nil
}

func divisionOperator(s *stack, input io.Reader, output io.Writer) error {
	x := pop(s)
	y := pop(s)

	if y == 0 {
		push(s, y)
		push(s, x)
		return ErrDivisionByZero
	}

	push(s, x/y)

	return nil
}

func minusOperator(s *stack, input io.Reader, output io.Writer) error {
	x := pop(s)
	y := pop(s)
	push(s, x-y)

	return nil
}

// negOperator negates the latest number in the stack
func negOperator(s *stack, input io.Reader, output io.Writer) error {
	push(s, -pop(s))

	return nil
}

// sumOperator sum two numebrs together
func sumOperator(s *stack, input io.Reader, output io.Writer) error {
	push(s, pop(s)+pop(s))

	return nil
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

// calculate does the actual calculation job, calculate
func calculate(s string, input io.Reader, output io.Writer) (float64, error) {

	parsedInput := parse(s)
	st := stack{}
	for _, e := range parsedInput {
		fl, ok := e.(float64)
		if ok {
			push(&st, fl)
		} else {
			op, _ := e.(string)
			err := operate(&st, op, input, output)
			if err != nil {
				return 0, err
			}
		}
	}

	return pop(&st), nil
}

func calculateFromFile(reader io.Reader, input io.Reader, output io.Writer) error {
	fileContent, _ := io.ReadAll(reader)
	// prints anything that passed as argument to eat.
	// it is good for checking for errors!
	// log.Print(fileContent)
	s := string(fileContent)
	_, err := calculate(s, input, output)
	return err
}

var ErrDivisionByZero = errors.New("cannot divide by zero")
