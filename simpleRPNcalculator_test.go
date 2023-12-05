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

func TestIsTheStackEmpty(t *testing.T)
	
	var s stack

	require.True(t, true, isTheStackEmpty(s))
	require.False(t, false, isTheStackEmpty())

}

//func TestAddToTheStack(t *testing.T){
//}
