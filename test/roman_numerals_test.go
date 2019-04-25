package test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func TestConvert(t *testing.T) {
	tests := []struct {
		expected string
		input int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"XI", 11},
		{"XII", 12},
		{"XIII", 13},
		{"XIV", 14},
		{"XV", 15},
		{"XVI", 16},
		{"XVII", 17},
		{"XVIII", 18},
		{"XIX", 19},
		{"XXI", 21},
		{"XXXI", 31},
		{"XLI", 41},
		{"LI", 51},
		{"LXI", 61},
		{"LXXI", 71},
		{"LXXXI", 81},
		{"XCI", 91},
		{"CI", 101},
		{"DI", 501},
		{"CMI", 901},
		{"MMI", 2001},
		{"MMMCMXCIX", 3999},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Convert(test.input))
	}
}


func Convert(n int) string {

	ciphers := []cipher {
		{"M", "?", "?"},
		{"C", "D", "M"},
		{"X", "L", "C"},
		{"I", "V", "X"}}

	digits := GetDigits(len(ciphers), n)
	roman := ""
	for i, cipher := range ciphers {
		roman += cipher.convert(digits[i])
	}
	return roman
}


func GetDigits(nDigits int, n int) []int {
	result := make([]int, nDigits)
	for i := nDigits - 1; i > 0; i-- {
		n, result[i] = GetDigit(n)
	}
	result[0] = n
	return result
}


func GetDigit(n int) (int, int) {
	return n / 10, n - (n / 10) * 10
}

type cipher struct {
	first, middle, last string
}


func (t *cipher) convert(n int) string {
	result := ""
	if n == 4 {
		n -= 4
		result = t.first + t.middle
	} else if 5 <= n && n <= 8 {
		n -= 5
		result = t.middle
	} else if n == 9 {
		n -= 9
		result = t.first + t.last
	}
	return result + strings.Repeat(t.first, n)
}