package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	)


func Test_starCanBeZero(t *testing.T){
	assert.True(t, isMatch("a","ab*") )
	assert.True(t, isMatch("aa","ab*a") )
	assert.True(t, isMatch("abbbb","ab*") )
	assert.True(t, isMatch("awrwrw",".*") )
	assert.True(t, isMatch("e", "a*c*d*e") )
	assert.True(t, isMatch("ae",".*..") )
	assert.False(t, isMatch("","e") )
	assert.False(t, isMatch("e","ed") )
	assert.True(t, isMatch("","") )
	assert.False(t, isMatch("ae","a") )
	assert.False(t, isMatch("ae",".") )
	assert.False(t, isMatch("ae","...") )
	assert.False(t, isMatch("ae","..*..") )
	assert.True(t, isMatch("aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*"))
	assert.True(t, isMatch("cbaacacaaccbaabcb", "c*b*b*.*ac*.*bc*a*"))
}

func Test_charMatches(t *testing.T){
	assert.True(t, charMatches(rune('a'),dot))
	assert.True(t, charMatches(rune('a'),rune('a')))
	assert.False(t, charMatches(rune('a'),rune('b')))
}

func Test_getNextMatchingSequence_singleCharacter(t *testing.T){
	match, isMulti, rest := getNextMatchingSequence("ao")
	assert.Equal(t, rune('a'), match)
	assert.False(t, isMulti)
	assert.Equal(t, "o", rest)
}

func Test_getNextMatchingSequence_multiCharacter(t *testing.T){
	match, isMulti, rest := getNextMatchingSequence(".*o")
	assert.Equal(t, dot, match)
	assert.True(t, isMulti)
	assert.Equal(t, "o", rest)
}


func Test_getNextMatchingSequence_noRest(t *testing.T){
	match, isMulti, rest := getNextMatchingSequence("a")
	assert.Equal(t, rune('a'), match)
	assert.False(t, isMulti)
	assert.Equal(t, "", rest)
}

func Test_matchToInput_singleCharacter_success(t *testing.T){
	success, rest := matchToInput(rune('a'), false, "ab")
	assert.True(t, success)
	assert.Equal(t,1,len(rest))
	assert.Equal(t,"b", rest[0])
}

func Test_matchToInput_singleCharacter_fail(t *testing.T){
	success, rest := matchToInput(rune('x'), false, "ab")
	assert.False(t, success)
	assert.Equal(t,1,len(rest))
	assert.Equal(t,"b", rest[0])
}

func Test_matchToInput_singleCharacter_noInput_fail(t *testing.T){
	success, rest := matchToInput(rune('a'), false, "")
	assert.False(t, success)
	assert.Equal(t,0, len(rest))
}

func Test_matchToInput_multiCharacter(t *testing.T){
	success, rest := matchToInput(rune('a'), true, "abb")
	assert.True(t, success)
	assert.Equal(t, 2, len(rest))
	assert.Equal(t, "abb", rest[0])
	assert.Equal(t, "bb", rest[1])
}

func Test_matchToInput_multiCharacter_dot(t *testing.T){
	success, rest := matchToInput(dot, true, "abb")
	assert.True(t, success)
	assert.Equal(t, 4, len(rest))
	assert.Equal(t, "abb", rest[0])
	assert.Equal(t, "bb", rest[1])
	assert.Equal(t, "b", rest[2])
	assert.Equal(t, "", rest[3])
}
