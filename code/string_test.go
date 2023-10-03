package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no need escape",
			input: "no need escape",
			want:  "no need escape",
		},
		{
			name: "need escape \\n",
			input: `hello
world
`,
			want: "hello\\nworld\\n",
		},
		{
			name: "need escape \\t",
			input: "hello	world",
			want: "hello\\tworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("encode", func(t *testing.T) {
				got := StringEncode(tt.input)
				require.Equalf(t, tt.want, got, "StringEncode(%v) real res:%v", tt.input, got)
			})
			t.Run("decode", func(t *testing.T) {
				got := StringDecode(tt.want)
				require.Equalf(t, tt.input, got, "StringDecode(%v) real res:%v", tt.want, got)
			})
		})
	}
}
