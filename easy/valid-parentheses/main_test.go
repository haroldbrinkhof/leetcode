package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
	)

func Test_emptyString(t *testing.T){
	assert.True(t, isValid(""))
}

func Test_sideBySide(t *testing.T){
	assert.True(t,isValid("[]{}()") )
}

func Test_encapsulating(t *testing.T){
	assert.True(t, isValid("[({})]") )
}

func Test_incorrectOrder(t *testing.T){
	assert.False(t, isValid("[{(})]") )
}

func Test_incorrectCount(t *testing.T){
	assert.False(t, isValid("]") )
	assert.False(t, isValid("[{()]") )
	assert.False(t, isValid("[{(})") )
}

