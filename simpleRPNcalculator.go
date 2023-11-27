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

func parse(s string) []any {

	//a := regexp.MustCompile(` `)
	//b := a.Split(s, -1)

	if len(s) == 0 {
		return nil
	}

	input := []any{}

	/*for i := 0; i < len(s.body); i++ {
		a := position{x: s.body[i].x, y: s.body[i].y}
		ret_s.body = append(ret_s.body, a)
	}*/

	strArr := strings.Split(s, " ")

	for i := 0; i < len(strArr); i++ {
		if fl, err := strconv.ParseFloat((strArr[i]), 64); err == nil {
			//fmt.Println(strArr) // 3.1415927410125732
			input = append(input, fl)
		} else {
			input = append(input, strArr[i])
		}

		//v := s[i]
	}
	return input
}

/*func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}*/
