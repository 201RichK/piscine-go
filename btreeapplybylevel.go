package piscine

import (
	"github.com/01-edu/z01"
)

func applyGivenOrder(root *TreeNode, level int, fn interface{}) {
	if root == nil {
		return
	}

	if level == 0 {
		arg := []interface{}{root.Data}
		z01.Call(fn, arg)
	} else if level > 0 {
		applyGivenOrder(root.Left, level-1, fn)
		applyGivenOrder(root.Right, level-1, fn)
	}
}

//Apply the function f level by level
func BTreeApplyByLevel(root *TreeNode, fn interface{}) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		applyGivenOrder(root, i, fn)
	}
}
