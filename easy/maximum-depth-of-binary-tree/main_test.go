package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_emptyTree(t *testing.T) {
	assert.Equal(t, 0, maxDepth(nil))
}

func Test_singleElement(t *testing.T) {
	var start *TreeNode = &TreeNode{1, nil, nil}
	assert.Equal(t, 1, maxDepth(start))
}

func Test_multipleElements(t *testing.T) {
	var one *TreeNode = &TreeNode{2, nil, nil}
	var start *TreeNode = &TreeNode{1, nil, one}
	assert.Equal(t, 2, maxDepth(start))
}
