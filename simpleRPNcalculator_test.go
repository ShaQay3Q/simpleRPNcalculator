package main

import (
	_ "image/png"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {

	Operation(3.4, 8.9, "*")

	require.Equal(t, Operation(3.4, 8.9, "*"), 30.26)
}

func TestParse(t *testing.T) {

	require.Empty(t, parse(""))
	require.Equal(t, []any{1}, parse("1"))
	// any -> is a type
}
