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

/*func initStack(a []any) stack {
	v := reflect.ValueOf(a)
	v = reflect.Indirect(v)

	var sfl stack
	for i := 0; i < len(a); i++{
		if fl, err := float64(a[i]); err == nil {
			//fmt.Println(strArr) // 3.1415927410125732
			sfl = append(sfl, fl) //mlll hotDog jhgjvgj
		}
	}
	return sfl
}*/

func initStack(a []any) []float64 {

	flt := make([]float64, (len(a)+1)/2)
	for i := 0; i < (len(a)+1)/2; i++ {
		flt[i] = a[i].(float64)
	}
	return flt
}

func initOperator(a []any) []string {

	str := make([]string, (len(a)-1)/2)
	for i := (len(a) + 1) / 2; i < len(a); i++ {
		str[i] = a[i].(string)
	}
	return str
}
