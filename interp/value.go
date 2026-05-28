// GoAWK interpreter value type (not exported).

package interp

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type valueType uint8

const (
	typeNull valueType = iota
	typeStr
	typeNum
	typeNumStr
	typeNumInt
)

// An AWK value (these are passed around by value)
type value struct {
	typ valueType // Type of value
	s   string    // String value (for typeStr and typeNumStr)
	n   float64   // Numeric value (for typeNum)
	l   int64     // Numeric integer value (for typeNumInt)
}

// Create a new null value
func null() value {
	return value{}
}

// Create a new number value
func num(n float64) value {
	return value{typ: typeNum, n: n}
}

// Create a new int64 value
func numInt(l int64) value {
	return value{typ: typeNumInt, l: l}
}

// Create a new string value
func str(s string) value {
	return value{typ: typeStr, s: s}
}

// Create a new value to represent a "numeric string" from an input field
func numStr(s string) value {
	return value{typ: typeNumStr, s: s}
}

// Create a numeric value from a Go bool
func boolean(b bool) value {
	if b {
		return num(1)
	}
	return num(0)
}

// String returns a string representation of v for debugging.
func (v value) String() string {
	switch v.typ {
	case typeStr:
		return fmt.Sprintf("str(%q)", v.s)
	case typeNum:
		return fmt.Sprintf("num(%s)", v.str("%.6g"))
	case typeNumStr:
		return fmt.Sprintf("numStr(%q)", v.s)
	case typeNumInt:
		return fmt.Sprintf("numInt(%d)", v.l)
	default:
		return "null()"
	}
}

// Return true if value is a "true string" (a string or a "numeric string"
// from an input field that can't be converted to a number). If false,
// also return the (possibly converted) number.
/*func (v value) isTrueStrNew() (number, bool) {
	switch v.typ {
	case typeStr:
		return numberInt(0), true
	case typeNumStr:
		f, err := parseFloat(v.s)
		if err != nil {
			return 0, true
		}
		return f, false
	case typeNumInt:
		panic("isTrueStr not implemented for typeNumInt") // TODO
	default: // typeNum, typeNull
		return v.n, false
	}
}*/

func (v value) isTrueStr() (float64, bool) { // TODO switch to above isTrueStrNew
	switch v.typ {
	case typeStr:
		return 0, true
	case typeNumStr:
		f, err := parseFloat(v.s)
		if err != nil {
			return 0, true
		}
		return f, false
	case typeNumInt:
		panic("isTrueStr not implemented for typeNumI64") // TODO
	default: // typeNum, typeNull
		return v.n, false
	}
}

// Return Go bool value of AWK value. For numbers or numeric strings,
// zero is false and everything else is true. For strings, empty
// string is false and everything else is true.
func (v value) boolean() bool {
	switch v.typ {
	case typeStr:
		return v.s != ""
	case typeNumStr:
		f, err := parseFloat(v.s)
		if err != nil {
			return v.s != ""
		}
		return f != 0
	case typeNumInt:
		return v.l != 0
	default: // typeNum, typeNull
		return v.n != 0
	}
}

// Like strconv.ParseFloat, but allow hex floating point without exponent, and
// allow "+nan" and "-nan" (though they both return math.NaN()). Also disallow
// underscore digit separators.
func parseFloat(s string) (float64, error) {
	s = strings.TrimSpace(s)
	if len(s) > 1 && (s[0] == '+' || s[0] == '-') {
		if len(s) == 4 && hasNaNPrefix(s[1:]) {
			// ParseFloat doesn't handle "nan" with sign prefix, so handle it here.
			return math.NaN(), nil
		}
		if len(s) > 3 && hasHexPrefix(s[1:]) && strings.IndexAny(s, "pP") < 0 {
			s += "p0"
		}
	} else if len(s) > 2 && hasHexPrefix(s) && strings.IndexAny(s, "pP") < 0 {
		s += "p0"
	}
	n, err := strconv.ParseFloat(s, 64)
	if err == nil && strings.IndexByte(s, '_') >= 0 {
		// Underscore separators aren't supported by AWK.
		return 0, strconv.ErrSyntax
	} // TODO allow
	return n, err
}

// Return value's string value, or convert to a string using given
// format if a number value. Integers are a special case and don't
// use floatFormat.
func (v value) str(floatFormat string) string {
	if v.typ == typeNum {
		switch {
		case math.IsNaN(v.n):
			return "nan"
		case math.IsInf(v.n, 0):
			if v.n < 0 {
				return "-inf"
			} else {
				return "inf"
			}
		case v.n == float64(int64(v.n)):
			return strconv.FormatInt(int64(v.n), 10)
		default:
			if floatFormat == "%.6g" {
				return strconv.FormatFloat(v.n, 'g', 6, 64)
			}
			return fmt.Sprintf(floatFormat, v.n)
		}
	} else if v.typ == typeNumInt {
		return strconv.FormatInt(v.l, 10)
	}
	// For typeStr and typeNumStr we already have the string, for
	// typeNull v.s == "".
	return v.s
}

// Return value's number value, converting from string if necessary
func (v value) numNew() number {
	switch v.typ {
	case typeStr, typeNumStr:
		// Ensure string starts with a float and convert it
		return parseNumberPrefix(v.s)
	case typeNumInt:
		return numberInt(v.l)
	default: // typeNum, typeNull
		return numberFloat(v.n)
	}
}

func (v value) num() float64 { // TODO switch to numNew above
	switch v.typ {
	case typeStr, typeNumStr:
		// Ensure string starts with a float and convert it
		return parseNumberPrefix(v.s).toFloat()
	default: // typeNum, typeNull
		return v.n
	}
}

// Return value's int64 number value, converting from string if necessary
func (v value) toInt() int64 {
	return v.numNew().toInt()
}

// Return value's float64 number value, converting from string if necessary
func (v value) toFloat() float64 {
	return v.numNew().toFloat()
}

type numberType uint8

const (
	typeFloat numberType = iota
	typeInt
)

// TODO we can consider packing float & int in a single field (a-la C union)
type number struct {
	typ numberType
	f   float64
	l   int64
}

func numberFloat(f float64) number {
	return number{typ: typeFloat, f: f}
}
func numberInt(l int64) number {
	return number{typ: typeInt, l: l}
}

// number toFloat
func (n number) toFloat() float64 {
	if n.typ == typeInt {
		return float64(n.l)
	}
	return n.f
}

// number toInt
func (n number) toInt() int64 {
	if n.typ == typeInt {
		return n.l
	}
	return int64(n.f)
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Like strconv.ParseFloat, but parses at the start of string and
// allows things like "1.5foo"
func parseNumberPrefix(s string) number {
	// Skip whitespace at start
	i := 0
	for i < len(s) && asciiSpace[s[i]] != 0 {
		i++
	}
	start := i

	// Parse optional sign and check for NaN and Inf.
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	if i+3 <= len(s) {
		if hasNaNPrefix(s[i:]) {
			return numberFloat(math.NaN())
		}
		if hasInfPrefix(s[i:]) {
			if s[start] == '-' {
				return numberFloat(math.Inf(-1))
			}
			return numberFloat(math.Inf(1))
		}
	}

	// Parse mantissa: initial digit(s), optional '.', then more digits
	if i+2 < len(s) && hasHexPrefix(s[i:]) {
		return parseHexNumberPrefix(s, start, i+2)
	}
	gotDigit := false
	gotDot := false
	gotExp := false
	for i < len(s) && isDigit(s[i]) {
		gotDigit = true
		i++
	}
	if i < len(s) && s[i] == '.' {
		gotDot = true
		i++
	}
	for i < len(s) && isDigit(s[i]) {
		gotDigit = true
		i++
	}
	if !gotDigit {
		return numberInt(0)
	}

	// Parse exponent ("1e" and similar are allowed, but ParseFloat
	// rejects them)
	end := i
	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		gotExp = true
		i++
		if i < len(s) && (s[i] == '+' || s[i] == '-') {
			i++
		}
		for i < len(s) && isDigit(s[i]) {
			i++
			end = i
		}
	}

	floatStr := s[start:end]
	if !gotDot && !gotExp {
		l, _ := strconv.ParseInt(floatStr, 10, 64)
		// todo if number is too large to fit in int64, shall we parse as float?
		return numberInt(l)
	}
	f, _ := strconv.ParseFloat(floatStr, 64)
	return numberFloat(f) // Returns infinity in case of "value out of range" error
}

func hasHexPrefix(s string) bool {
	return s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func hasNaNPrefix(s string) bool {
	return (s[0] == 'n' || s[0] == 'N') && (s[1] == 'a' || s[1] == 'A') && (s[2] == 'n' || s[2] == 'N')
}

func hasInfPrefix(s string) bool {
	return (s[0] == 'i' || s[0] == 'I') && (s[1] == 'n' || s[1] == 'N') && (s[2] == 'f' || s[2] == 'F')
}

// Helper used by parseNumberPrefix to handle hexadecimal floating point.
func parseHexNumberPrefix(s string, start, i int) number {
	gotDigit := false
	gotDot := false
	gotExp := false // has p or P
	for i < len(s) && isHexDigit(s[i]) {
		gotDigit = true
		i++
	}
	if i < len(s) && s[i] == '.' {
		gotDot = true
		i++
	}
	for i < len(s) && isHexDigit(s[i]) {
		gotDigit = true
		i++
	}
	if !gotDigit {
		return numberInt(0)
	}

	gotExponent := false
	end := i
	if i < len(s) && (s[i] == 'p' || s[i] == 'P') {
		gotExp = true
		i++
		if i < len(s) && (s[i] == '+' || s[i] == '-') {
			i++
		}
		for i < len(s) && isDigit(s[i]) {
			gotExponent = true
			i++
			end = i
		}
	}

	floatStr := s[start:end]
	if !gotDot && !gotExp {
		l, _ := strconv.ParseInt(floatStr, 0, 64)
		return numberInt(l)
	}
	if !gotExponent {
		floatStr += "p0" // AWK allows "0x12", ParseFloat requires "0x12p0"
	}
	f, _ := strconv.ParseFloat(floatStr, 64)
	return numberFloat(f) // Returns infinity in case of "value out of range" error
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isHexDigit(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}
