package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left && root.Right {
		return (BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right))
	}
	return false
}

//btreeisbinary_test.go:50:
//  BTreeIsBinary(&solutions.TreeNode{Left:(*solutions.TreeNode)(0xc000080420),
//     Right:(*solutions.TreeNode)(0xc000080480),
//      Parent:(*solutions.TreeNode)(nil), Data:"04"})
//     == true instead of false
