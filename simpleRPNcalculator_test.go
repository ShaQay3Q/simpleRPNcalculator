package main

import (
	"bytes"
	_ "image/png"
	"os"
	"strings"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {

	require.Empty(t, parse(""))
	require.Equal(t, []any{1.}, parse("1"))
	require.Equal(t, []any{17.5, 2.9, "+"}, parse("17.5 2.9 +"))
	// any -> is a type
}

func TestStack(t *testing.T) {
	s := stack{}

	require.Empty(t, s)

	require.True(t, isTheStackEmpty(&s))

	push(&s, 42.3)
	require.NotEmpty(t, s)

	require.Equal(t, 42.3, pop(&s))

	require.Empty(t, s)
	require.True(t, isTheStackEmpty(&s))

	push(&s, 1.5)
	push(&s, 2.8)
	push(&s, 17.1)

	require.Equal(t, 17.1, pop(&s))
	require.Equal(t, 2.8, pop(&s))
	require.Equal(t, 1.5, pop(&s))

	require.Empty(t, s)
	require.True(t, isTheStackEmpty(&s))

}

func TestCalculate(t *testing.T) {

	require.Equal(t, 3., calculate("3", nil))
	require.Equal(t, 5., calculate("5", nil))
	require.Equal(t, 5.7, calculate("5.7", nil))

	require.Equal(t, 7., calculate("3 4 +", nil))
	require.Equal(t, 12., calculate("3 4 5 + +", nil))
	require.Equal(t, 31., calculate("3 4 + 5 + 6 10 + + 3 +", nil))
	require.Equal(t, -1., calculate("4 3 -", nil))
	require.Equal(t, 27., calculate("3 4 5 + *", nil))

	require.Equal(t, -10., calculate("3 4 * neg 2 +", nil))

	require.Equal(t, -48., calculate("2 3 + 4 33 6 summation neg", nil))
	require.Equal(t, 0., calculate("summation", nil))
	require.Equal(t, 9., calculate("3 dup *", nil))
	require.Equal(t, 5., calculate("2 3 4 drop +", nil))
	require.Equal(t, 66., calculate("2 4 3 pwr +", nil))
	require.Equal(t, -1., calculate("3 0 pwr neg", nil))
	require.Equal(t, (1. / 64), calculate("4 -3 pwr", nil))

}

func TestOperate(t *testing.T) {

	s1 := stack{3.6}
	s2 := stack{3.4, 8.9}
	s3 := stack{3.4, 8.9}
	s4 := stack{2.4, 4.8}
	s5 := stack{2., 5.}
	s6 := stack{1, 2, 3, 4, 5, 6}
	s7 := stack{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s8 := stack{2., 5.}

	operate(&s1, "neg", nil, nil)
	require.Equal(t, stack{-3.6}, s1)

	operate(&s2, "+", nil, nil)
	require.Equal(t, stack{12.3}, s2)

	operate(&s3, "-", nil, nil)
	require.Equal(t, stack{5.5}, s3)

	operate(&s4, "/", nil, nil)
	require.Equal(t, stack{2}, s4)

	operate(&s5, "*", nil, nil)
	require.Equal(t, stack{10}, s5)

	operate(&s6, "summation", nil, nil)
	require.Equal(t, stack{21}, s6)

	operate(&s7, "drop", nil, nil)
	require.Equal(t, stack{1, 2, 3, 4, 5, 6, 7, 8, 9}, s7)

	operate(&s8, "dup", nil, nil)
	require.Equal(t, stack{2, 5, 5}, s8)
}

func TestPrintIt(t *testing.T) {

	s := stack{5.}

	var output bytes.Buffer
	operate(&s, "printIt", nil, &output)
	operate(&s, "printIt", nil, &output)
	require.Equal(t, ": 5\n: 5\n", output.String())
}

func TestRead(t *testing.T) {
	input := strings.NewReader("1.5")
	var output bytes.Buffer
	var s stack
	operate(&s, "read", input, &output)
	operate(&s, "neg", nil, &output)
	operate(&s, "printIt", nil, &output)
	require.Equal(t, "enter a number> : -1.5\n", output.String())
}

func TestReadFromFile001(t *testing.T) {

	file, _ := os.Open("./test-files/001.txt")
	defer file.Close()

	var output bytes.Buffer

	calculateFromFile(file, &output)

	require.Equal(t, ": 10\n", output.String())

}

func TestReadFromFile002(t *testing.T) {

	file, _ := os.Open("./test-files/002.txt")
	defer file.Close()

	var output bytes.Buffer

	calculateFromFile(file, &output)

	require.Equal(t, ": -16\n", output.String())

}

func TestReadFromFile003(t *testing.T) {

	file, _ := os.Open("./test-files/003.txt")
	defer file.Close()

	var output bytes.Buffer

	calculateFromFile(file, &output)

	require.Equal(t, ": 3\n: 3\n", output.String())

}

// https://go.dev/tour/moretypes/11
// https://go.dev/tour/moretypes/14
// https://go.dev/tour/moretypes/15

//func TestAddToTheStack(t *testing.T){
//}
