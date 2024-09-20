package main

import (
	"testing"
)

func TestIsBalance(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},         // empty string
		{"{}", true},       // single pair
		{"()", true},       // single pair
		{"[]", true},       // single pair
		{"{}()", true},     // mixed pairs
		{"{[()]}", true},   // nested pairs
		{"{[}]", false},    // mismatched
		{"{(})", false},    // mismatched
		{"{[()]}{}", true}, // multiple balanced pairs
		{"{[()]}}", false}, // extra closing bracket
		{"{[}", false},     // unclosed brackets
		{"{[()]}[", false}, // unclosed brackets
		{"{{{{}}}}", true}, // multiple pairs
		{"[({})]", true},   // nested pairs
		{"))", false},      // unbalanced closing
		{"((()))", true},   // deeply nested pairs
		{"((())", false},   // unbalanced opening
	}

	for _, test := range tests {
		result := isBalance(test.input)
		if result != test.expected {
			t.Errorf("isBalance(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
