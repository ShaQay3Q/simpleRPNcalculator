package main

import (
	"strconv"
	"strings"
)

func main() {

}

func Operation(x float64, y float64, op string) float64 {

	var ret float64
	switch op {
	case "+":
		ret = x + y
	case "-":
		ret = x - y
	case "*":
		ret = x * y
	case "/":
		ret = x / y
	}
	return ret
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

	for i := 0; i < len(strArr); i++ {
		if fl, err := strconv.ParseFloat((strArr[i]), 64); err == nil {
			//fmt.Println(strArr) // 3.1415927410125732
			output = append(output, fl) //mlll hotDog jhgjvgj
		} else {
			output = append(output, strArr[i])
		}
	}
	return output
}

func calculator(s string) float64 {
	input := parse(s)
	st := stack{}
	for i := 0; i < len(input); i++ {
		e := input[i]
		fl, ok := e.(float64)
		if ok {
			push(&st, fl)
		} else {
			op, _ := e.(string)
			if op == "neg" {
				x := pop(&st)
				push(&st, -x)
			} else {
				x := pop(&st)
				y := pop(&st)
				sum := Operation(x, y, op)
				//sum := x + y
				push(&st, sum)
			}
		}
	}

	return pop(&st)
}
