package main

import (
	_ "image/png"
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

func TestCalculator(t *testing.T) {

	require.Equal(t, 3., calculator("3"))
	require.Equal(t, 5., calculator("5"))
	require.Equal(t, 5.7, calculator("5.7"))

	require.Equal(t, 7., calculator("3 4 +"))
	require.Equal(t, 12., calculator("3 4 5 + +"))
	require.Equal(t, 31., calculator("3 4 + 5 + 6 10 + + 3 +"))
	require.Equal(t, -1., calculator("4 3 -"))
	require.Equal(t, 27., calculator("3 4 5 + *"))

	require.Equal(t, -10., calculator("3 4 * neg 2 +"))

	require.Equal(t, -48., calculator("2 3 + 4 33 6 summation neg"))
	require.Equal(t, 0., calculator("summation"))
	require.Equal(t, 9., calculator("3 dup *"))
	require.Equal(t, 5., calculator("2 3 4 drop +"))
	require.Equal(t, 66., calculator("2 4 3 pwr +"))
	require.Equal(t, -1., calculator("3 0 pwr neg"))
}

func TestOperate(t *testing.T) {

	s1 := stack{3.6}
	//s2 := stack{3.4, 8.9}
	//s3 := stack{3.4, 8.9}
	//s4 := stack{2.4, 4.8}
	//s5 := stack{2., 5.}
	//s6 := stack{1, 2, 3, 4, 5, 6}
	//s7 := stack{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s8 := stack{2., 5.}

	operate(&s1, "neg")
	require.Equal(t, stack{-3.6}, s1)

	/*
		require.Equal(t, -3.6, operate(&s1, "neg"))
		require.Equal(t, 12.3, operate(&s2, "+"))
		require.Equal(t, 5.5, operate(&s3, "-"))
		require.Equal(t, 2., operate(&s4, "/"))
		require.Equal(t, 10., operate(&s5, "*"))
		require.Equal(t, 21., operate(&s6, "summation"))
		// ? is it what you were after?
		//require.EqualValues(t, 1, 2, 3, operate(&s7, "drop"))
		require.Equal(t, 9., operate(&s7, "drop"))
		require.Equal(t, 8, len(s7))
		require.Equal(t, 5., operate(&s8, "dup"))
		require.Equal(t, 3, len(s8))
		require.Equal(t, stack{2, 3}, operate(&stack{2, 3, 4}, "drop"))
	*/
}

// https://go.dev/tour/moretypes/11
// https://go.dev/tour/moretypes/14
// https://go.dev/tour/moretypes/15

//func TestAddToTheStack(t *testing.T){
//}
