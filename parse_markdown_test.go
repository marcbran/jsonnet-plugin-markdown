package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMarkdownErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "no arguments",
			input:    []any{},
			expected: "markdown must be provided",
		},
		{
			name:     "too many arguments",
			input:    []any{"markdown", "extra"},
			expected: "markdown must be provided",
		},
		{
			name:     "non-string argument",
			input:    []any{123},
			expected: "markdown must be a string",
		},
		{
			name:     "nil argument",
			input:    []any{nil},
			expected: "markdown must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMarkdown().Func(tt.input)

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}
