package piscine

func BTreeIsBinary(root *TreeNode) bool {
	prev := &TreeNode{}
	return isBSTUtil(root, prev)
}

func isBSTUtil(root, prev *TreeNode) bool {
	if root != nil {
		if !isBSTUtil(root.Left, prev) {
			return false
		}

		if prev != nil && root.Data <= prev.Data {
			return false
		}

		prev = root
		return isBSTUtil(root.Right, prev)
	}
	return true
}

//btreeisbinary_test.go:50:
//  BTreeIsBinary(&solutions.TreeNode{Left:(*solutions.TreeNode)(0xc000080420),
//     Right:(*solutions.TreeNode)(0xc000080480),
//      Parent:(*solutions.TreeNode)(nil), Data:"04"})
//     == true instead of false
