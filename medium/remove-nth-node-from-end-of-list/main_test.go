package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestSingleElementReturnsEmpty(t *testing.T){
	singleNode := &ListNode{1, nil}

	assert.Nil(t, removeNthFromEnd(singleNode,1))
}

func TestEmptyNodeListReturnsNil(t *testing.T){
	assert.Nil(t, removeNthFromEnd(nil,1))
}


func TestRemoveElement(t *testing.T){
	five := &ListNode{5, nil}
	four := &ListNode{4, five}
	three := &ListNode{3, four}
	two := &ListNode{2, three}
	start := &ListNode{1, two}
	
	removeNthFromEnd(start,2)

	assert.True(t, three.Next == five)
}


func TestRemoveLastElement(t *testing.T){
	two := &ListNode{2, nil}
	start := &ListNode{1, two}
	
	removeNthFromEnd(start,1)

	assert.Nil(t, start.Next)
}

func TestRemoveFirstElement(t *testing.T){
	two := &ListNode{2, nil}
	start := &ListNode{1, two}
	
	actual := removeNthFromEnd(start,2)

	assert.Equal(t, two, actual)
}
