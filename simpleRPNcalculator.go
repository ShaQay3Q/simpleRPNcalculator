package main

import (
	"strconv"
	"strings"
)

func main() {

}

func operate(s *stack, op string) float64 {
	var ret float64
	switch op {
	case "neg":
		ret = -pop(s)
		push(s, ret)
	case "+":
		x := pop(s)
		y := pop(s)
		ret = x + y
		push(s, ret)
	case "-":
		x := pop(s)
		y := pop(s)
		ret = x - y
		push(s, ret)
	case "/":
		x := pop(s)
		y := pop(s)
		ret = x / y
		push(s, ret)
	case "*":
		x := pop(s)
		y := pop(s)
		ret = x * y
		push(s, ret)
	case "summation":
		for !isTheStackEmpty(s) {
			ret = ret + pop(s)
		}
		push(s, ret)
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
	for _, e := range input {
		fl, ok := e.(float64)
		if ok {
			push(&st, fl)
		} else {
			op, _ := e.(string)
			operate(&st, op)
		}
	}

	return pop(&st)
}
