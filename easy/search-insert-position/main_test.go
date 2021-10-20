package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
	)

func Test_existingElement(t *testing.T){
	assert.Equal(t,2, searchInsert([]int{1,3,5,6}, 5))

}

func Test_nonExistingElement(t *testing.T){
	assert.Equal(t,2, searchInsert([]int{1,3,5,6}, 4))
	assert.Equal(t,1, searchInsert([]int{1,3,5,6}, 2))

}

func Test_smallerThanFirst(t *testing.T){
	assert.Equal(t,0, searchInsert([]int{1,3,5,6}, 0))
}

func Test_biggerThanLast(t *testing.T){
	assert.Equal(t,4, searchInsert([]int{1,3,5,6}, 7))
}
