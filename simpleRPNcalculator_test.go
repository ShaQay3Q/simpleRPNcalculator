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

	require.InEpsilon(t, 17.1, pop(&s), 0.001)
	require.InDelta(t, 2.8, pop(&s), 0.001)
	require.Equal(t, 1.5, pop(&s))

	require.Empty(t, s)
	require.True(t, isTheStackEmpty(&s))

}

func TestCalculate(t *testing.T) {

	fl, _ := calculate("3", nil, nil)

	require.Equal(t, 3., fl)

	fl, _ = calculate("5", nil, nil)
	require.Equal(t, 5., fl)

	fl, _ = calculate("5.7", nil, nil)
	require.Equal(t, 5.7, fl)

	fl, _ = calculate("3 4 +", nil, nil)
	require.Equal(t, 7., fl)

	fl, _ = calculate("3 4 5 + +", nil, nil)
	require.Equal(t, 12., fl)

	fl, _ = calculate("3 4 + 5 + 6 10 + + 3 +", nil, nil)
	require.Equal(t, 31., fl)

	fl, _ = calculate("4 3 -", nil, nil)
	require.Equal(t, -1., fl)

	fl, _ = calculate("3 4 5 + *", nil, nil)
	require.Equal(t, 27., fl)

	fl, _ = calculate("3 4 * neg 2 +", nil, nil)
	require.Equal(t, -10., fl)

	fl, _ = calculate("2 3 + 4 33 6 summation neg", nil, nil)
	require.Equal(t, -48., fl)

	fl, _ = calculate("summation", nil, nil)
	require.Equal(t, 0., fl)

	fl, _ = calculate("3 dup *", nil, nil)
	require.Equal(t, 9., fl)

	fl, _ = calculate("2 3 4 drop +", nil, nil)
	require.Equal(t, 5., fl)

	fl, _ = calculate("2 4 3 pwr +", nil, nil)
	require.Equal(t, 66., fl)

	fl, _ = calculate("3 0 pwr neg", nil, nil)
	require.Equal(t, -1., fl)

	fl, _ = calculate("4 -3 pwr", nil, nil)
	require.Equal(t, (1. / 64), fl)

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

	err := operate(&s1, "neg", nil, nil)

	require.Nil(t, err)
	require.Equal(t, stack{-3.6}, s1)

	err = operate(&s2, "+", nil, nil)

	require.Nil(t, err)
	require.Equal(t, stack{12.3}, s2)

	err = operate(&s3, "-", nil, nil)

	require.Nil(t, err)
	require.Equal(t, stack{5.5}, s3)

	err = operate(&s4, "/", nil, nil)
	require.Nil(t, err)
	require.Equal(t, stack{2}, s4)

	err = operate(&s5, "*", nil, nil)
	require.Nil(t, err)
	require.Equal(t, stack{10}, s5)

	err = operate(&s6, "summation", nil, nil)
	require.Nil(t, err)
	require.Equal(t, stack{21}, s6)

	err = operate(&s7, "drop", nil, nil)
	require.Nil(t, err)
	require.Equal(t, stack{1, 2, 3, 4, 5, 6, 7, 8, 9}, s7)

	err = operate(&s8, "dup", nil, nil)
	require.Nil(t, err)
	require.Equal(t, stack{2, 5, 5}, s8)
}

func TestPrintIt(t *testing.T) {

	s := stack{5.}

	var output bytes.Buffer
	err := operate(&s, "printIt", nil, &output)
	require.Nil(t, err)
	err = operate(&s, "printIt", nil, &output)
	require.Nil(t, err)
	require.Equal(t, ": 5\n: 5\n", output.String())
}

func TestRead(t *testing.T) {
	input := strings.NewReader("1.5")
	var output bytes.Buffer
	var s stack
	err := operate(&s, "read", input, &output)
	require.Nil(t, err)
	err = operate(&s, "neg", nil, nil)
	require.Nil(t, err)
	err = operate(&s, "printIt", nil, &output)
	require.Nil(t, err)
	require.Equal(t, "enter a number> : -1.5\n", output.String())
}

func TestReadFromFile001(t *testing.T) {

	file, _ := os.Open("./test-files/001.txt")
	defer file.Close()

	var output bytes.Buffer

	err := calculateFromFile(file, nil, &output)

	require.Nil(t, err)
	require.Equal(t, ": 10\n", output.String())

}

func TestReadFromFile002(t *testing.T) {

	file, _ := os.Open("./test-files/002.txt")
	defer file.Close()

	var output bytes.Buffer

	err := calculateFromFile(file, nil, &output)

	require.Nil(t, err)
	require.Equal(t, ": -16\n", output.String())

}

func TestReadFromFile003(t *testing.T) {

	file, _ := os.Open("./test-files/003.txt")
	defer file.Close()

	var output bytes.Buffer

	err := calculateFromFile(file, nil, &output)

	require.Nil(t, err)
	require.Equal(t, ": 3\n: 3\n", output.String())

}

func TestReadFromFile004(t *testing.T) {
	file, _ := os.Open("./test-files/004.txt")
	defer file.Close()

	var output bytes.Buffer
	//passing the variable instead f typing it in terminal
	input := strings.NewReader("40")

	err := calculateFromFile(file, input, &output)

	require.Nil(t, err)
	require.Equal(t, "enter a number> : -40\n", output.String())
}

func TestReadFromFile005(t *testing.T) {
	file, _ := os.Open("./test-files/005.txt")
	defer file.Close()

	var output bytes.Buffer

	err := calculateFromFile(file, nil, &output)

	require.NotNil(t, err)
	require.Equal(t, ErrDivisionByZero, err)
}

//func FuzzCalculate(f *testing.F) {
//	testcases := []string{
//		"3 4 + 5 + 6 10 + + 3 +",
//		"3 4 5 + +",
//		"3 0 pwr neg",
//		"3 dup *",
//		"0 17 / printIt",
//	}
//	for _, tc := range testcases {
//		f.Add(tc) // Use f.Add to provide a seed corpus
//	}
//	f.Fuzz(func(t *testing.T, a string) {
//		var output bytes.Buffer
//		_, _ = calculate(a, nil, &output)
//	})
//}

// https://go.dev/tour/moretypes/11
// https://go.dev/tour/moretypes/14
// https://go.dev/tour/moretypes/15
