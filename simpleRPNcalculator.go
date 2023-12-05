package main

import (
	"strconv"
	"strings"
)

func main() {

}

func Operation(x float64, y float64, op string) (i float64) {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return
}

type stack struct {
	a    float64
	b    float64
	oprt string
}

func isTheStackEmpty(s *stack) bool {
	if s == nil {
		return true
	}
	return false
}

// parse pasre the input string into array of any
func parse(s string) []any {

	//a := regexp.MustCompile(` `)
	//b := a.Split(s, -1)

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

		//v := s[i]
	}
	return output
}

// l is to find the fist operator's house in the parsed array
//var l = (len(parse().) + 1) / 2
