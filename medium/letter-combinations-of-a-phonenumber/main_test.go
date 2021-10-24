package main
import "testing"
import "github.com/stretchr/testify/assert"

func TestDoubleDigit(t *testing.T){
	input := "23"
	expected := []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}
	assert.Equal(t, expected, letterCombinations(input))
}

func TestSingleDigit(t *testing.T){
	input := "2"
	expected := []string{"a","b","c"}
	assert.Equal(t, expected, letterCombinations(input))
}

func TestOneIsIgnored(t *testing.T){
	input := "12"
	expected := []string{"a","b","c"}
	assert.Equal(t, expected, letterCombinations(input))
}
