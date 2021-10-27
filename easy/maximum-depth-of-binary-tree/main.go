package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int = 1
	var nextLayer []*TreeNode
	var currentLayer []*TreeNode = []*TreeNode{}

	if root.Left != nil {
		currentLayer = append(currentLayer, root.Left)
	}
	if root.Right != nil {
		currentLayer = append(currentLayer, root.Right)
	}

	for {
		if len(currentLayer) == 0 {
			break
		}

		nextLayer = []*TreeNode{}
		for c := 0; c < len(currentLayer); c++ {
			curVal := currentLayer[c]
			if curVal != nil {
				if curVal.Left != nil {
					nextLayer = append(nextLayer, curVal.Left)
				}
				if curVal.Right != nil {
					nextLayer = append(nextLayer, curVal.Right)
				}

			}
		}

		currentLayer = nextLayer
		depth++
	}

	return depth
}
