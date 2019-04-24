package test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "1", FizzBuzz(1))
	assert.Equal(t, "2", FizzBuzz(2))
	assert.Equal(t, "Fizz", FizzBuzz(3))
	assert.Equal(t, "4", FizzBuzz(4))
	assert.Equal(t, "Buzz", FizzBuzz(5))
	assert.Equal(t, "Fizz", FizzBuzz(6))
	assert.Equal(t, "7", FizzBuzz(7))
	assert.Equal(t, "8", FizzBuzz(8))
	assert.Equal(t, "Fizz", FizzBuzz(9))
	assert.Equal(t, "Buzz", FizzBuzz(10))
	assert.Equal(t, "11", FizzBuzz(11))
	assert.Equal(t, "Fizz", FizzBuzz(12))
	assert.Equal(t, "13", FizzBuzz(13))
	assert.Equal(t, "14", FizzBuzz(14))
	assert.Equal(t, "FizzBuzz", FizzBuzz(15))

}


func FizzBuzz(n int) string {
	result := ""
	if n % 3 == 0 {
		result += "Fizz"
	}
	if n % 5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}

