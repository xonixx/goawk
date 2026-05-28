package interp

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestParseFloat(t *testing.T) {
	tests := []struct {
		in      string
		want    float64
		wantErr bool
		name    string
	}{
		{"12.5", 12.5, false, "decimal"},
		{"\t\n  -42.25 \r", -42.25, false, "trim whitespace"},
		{"+NaN", math.NaN(), false, "positive nan"},
		{"-nan", math.NaN(), false, "negative nan"},
		{"+inf", math.Inf(1), false, "positive infinity"},
		{"0x10", 16, false, "hex integer without exponent"},
		{"-0x1.8", -1.5, false, "signed hex fraction without exponent"},
		{"0x1.8P2", 6, false, "uppercase hex exponent"},
		{"1_000", 0, true, "underscore rejected"},
		{"12x", 0, true, "trailing text rejected"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFloat(tt.in)
			if tt.wantErr {
				if !errors.Is(err, strconv.ErrSyntax) {
					t.Fatalf("parseFloat(%q) error = %v, want syntax error", tt.in, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseFloat(%q) unexpected error: %v", tt.in, err)
			}
			if !equalFloats(got, tt.want) {
				t.Fatalf("parseFloat(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		in      string
		want    number
		wantErr bool
		name    string
	}{
		{"42", numberInt(42), false, "decimal integer"},
		{"\t\n  -42 \r", numberInt(-42), false, "trim whitespace"},
		{"+17", numberInt(17), false, "positive integer"},
		{"12.5", numberFloat(12.5), false, "decimal fraction"},
		{"6.02e2", numberFloat(602), false, "exponent"},
		{"+NaN", numberFloat(math.NaN()), false, "positive nan"},
		{"-nan", numberFloat(math.NaN()), false, "negative nan"},
		{"0x10", numberFloat(16), false, "hex integer without exponent"},
		{"-0x1.8", numberFloat(-1.5), false, "signed hex fraction without exponent"},
		{"0x1.8P2", numberFloat(6), false, "uppercase hex exponent"},
		{"1_000", numberInt(0), true, "integer underscore rejected"},
		{"1_000.5", numberInt(0), true, "float underscore rejected"},
		{"12x", numberInt(0), true, "trailing text rejected"},
		{"", numberInt(0), true, "empty rejected"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNumber(tt.in)
			if tt.wantErr {
				if !errors.Is(err, strconv.ErrSyntax) {
					t.Fatalf("parseNumber(%q) error = %v, want syntax error", tt.in, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseNumber(%q) unexpected error: %v", tt.in, err)
			}
			if !equalNumbers(got, tt.want) {
				t.Fatalf("parseNumber(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

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

func equalFloats(a, b float64) bool {
	return a == b || math.IsNaN(a) && math.IsNaN(b)
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
