package xrunes

import (
	"reflect"
	"testing"
)

func TestUnderscore(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		options  []HyphenMinusOption
		expected []rune
	}{
		{
			name:     "Empty input",
			input:    []rune{},
			options:  nil,
			expected: []rune{},
		},
		{
			name:     "Simple lowercase",
			input:    []rune("hello world"),
			options:  nil,
			expected: []rune("hello_world"),
		},
		{
			name:     "Simple uppercase",
			input:    []rune("Hello World"),
			options:  nil,
			expected: []rune("hello_world"),
		},
		{
			name:     "Preserve case",
			input:    []rune("Hello World"),
			options:  []HyphenMinusOption{func(params *HyphenMinusParams) { params.PreserveCase = true }},
			expected: []rune("Hello_World"),
		},
		{
			name:     "Screaming case",
			input:    []rune("Hello World"),
			options:  []HyphenMinusOption{Screaming},
			expected: []rune("HELLO_WORLD"),
		},
		{
			name:     "Mixed case with numbers",
			input:    []rune("Hello123 World456"),
			options:  nil,
			expected: []rune("hello123_world456"),
		},
		{
			name:     "Multiple spaces and special characters",
			input:    []rune("Hello   World---Test"),
			options:  nil,
			expected: []rune("hello_world_test"),
		},
		{
			name:     "Trailing underscore",
			input:    []rune("Hello World "),
			options:  nil,
			expected: []rune("hello_world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Underscore(tt.input, tt.options...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %q, got %q", string(tt.expected), string(result))
			}
		})
	}
}

func TestDasherize(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		options  []HyphenMinusOption
		expected []rune
	}{
		{
			name:     "Empty input",
			input:    []rune{},
			options:  nil,
			expected: []rune{},
		},
		{
			name:     "Simple lowercase",
			input:    []rune("hello world"),
			options:  nil,
			expected: []rune("hello-world"),
		},
		{
			name:     "Simple uppercase",
			input:    []rune("Hello World"),
			options:  nil,
			expected: []rune("hello-world"),
		},
		{
			name:     "Preserve case",
			input:    []rune("Hello World"),
			options:  []HyphenMinusOption{func(params *HyphenMinusParams) { params.PreserveCase = true }},
			expected: []rune("Hello-World"),
		},
		{
			name:     "Screaming case",
			input:    []rune("Hello World"),
			options:  []HyphenMinusOption{Screaming},
			expected: []rune("HELLO-WORLD"),
		},
		{
			name:     "Mixed case with numbers",
			input:    []rune("Hello123 World456"),
			options:  nil,
			expected: []rune("hello123-world456"),
		},
		{
			name:     "Multiple spaces and special characters",
			input:    []rune("Hello   World---Test"),
			options:  nil,
			expected: []rune("hello-world-test"),
		},
		{
			name:     "Trailing dash",
			input:    []rune("Hello World "),
			options:  nil,
			expected: []rune("hello-world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Dasherize(tt.input, tt.options...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %q, got %q", string(tt.expected), string(result))
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected []rune
	}{
		{
			name:     "Empty input",
			input:    []rune{},
			expected: []rune{},
		},
		{
			name:  "Simple lowercase",
			input: []rune("hello world"),

			expected: []rune("helloWorld"),
		},
		{
			name:  "Simple uppercase",
			input: []rune("Hello World"),

			expected: []rune("helloWorld"),
		},
		{
			name:  "Mixed case with numbers",
			input: []rune("Hello123 World456"),

			expected: []rune("hello123World456"),
		},
		{
			name:     "Multiple spaces and special characters",
			input:    []rune("Hello   World---Test"),
			expected: []rune("helloWorldTest"),
		},
		{
			name:     "Trailing dash",
			input:    []rune("Hello World "),
			expected: []rune("helloWorld"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CamelCase(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %q, got %q", string(tt.expected), string(result))
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected []rune
	}{
		{
			name:     "Empty input",
			input:    []rune{},
			expected: []rune{},
		},
		{
			name:  "Simple lowercase",
			input: []rune("hello world"),

			expected: []rune("HelloWorld"),
		},
		{
			name:  "Simple uppercase",
			input: []rune("Hello World"),

			expected: []rune("HelloWorld"),
		},
		{
			name:  "Mixed case with numbers",
			input: []rune("Hello123 World456"),

			expected: []rune("Hello123World456"),
		},
		{
			name:     "Multiple spaces and special characters",
			input:    []rune("Hello   World---Test"),
			expected: []rune("HelloWorldTest"),
		},
		{
			name:     "Trailing dash",
			input:    []rune("Hello World "),
			expected: []rune("HelloWorld"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PascalCase(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %q, got %q", string(tt.expected), string(result))
			}
		})
	}
}
