package main

import (
	_ "image/png"
	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {

	Operation(3.4, 8.9, "*")

	require.Equal(t, Operation(3.4, 8.9, "*"), 30.26)
}

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
}

// https://go.dev/tour/moretypes/11
// https://go.dev/tour/moretypes/14
// https://go.dev/tour/moretypes/15

//func TestAddToTheStack(t *testing.T){
//}
