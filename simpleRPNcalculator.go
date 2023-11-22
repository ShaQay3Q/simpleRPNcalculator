package main

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
	
	return ""
}
