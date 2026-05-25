package interp

import (
	"math"
	"testing"
)

func TestParseNumberPrefix(t *testing.T) {
	tests := []struct {
		in   string
		want number
		name string
	}{
		{"", numberInt(0), "empty"},
		{"foo", numberInt(0), "non-numeric"},
		{"\t\n\v\f\r 42x", numberInt(42), "ascii whitespace"},
		{"12345xyz", numberInt(12345), "decimal integer prefix"},
		{" -123x", numberInt(-123), "negative integer"},
		{"12.5kg", numberFloat(12.5), "decimal fraction"},
		{".75x", numberFloat(0.75), "leading dot"},
		{"7.foo", numberFloat(7), "trailing dot"},
		{"6.02e2x", numberFloat(602), "exponent"},
		{"1e+x", numberFloat(1), "incomplete exponent"},
		{"+nanx", numberFloat(math.NaN()), "positive nan"},
		{"-NaNx", numberFloat(math.NaN()), "negative nan"},
		{"INFx", numberFloat(math.Inf(1)), "positive infinity"},
		{"-infx", numberFloat(math.Inf(-1)), "negative infinity"},
		{"0x10tail", numberInt(16), "hex integer"},
		{"-0xf.fp-1tail", numberFloat(-7.96875), "hex fraction"},
		{"0x.tail", numberInt(0), "hex without digits"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseNumberPrefix(tt.in)
			if !equalNumbers(got, tt.want) {
				t.Fatalf("parseNumberPrefix(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

func equalNumbers(a, b number) bool {
	if a.typ != b.typ {
		return false
	}
	switch a.typ {
	case typeInt:
		return a.l == b.l
	case typeFloat:
		return a.f == b.f || math.IsNaN(a.f) && math.IsNaN(b.f)
	default:
		return false
	}
}
