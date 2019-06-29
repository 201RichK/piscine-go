package piscine

func BTreeIsBinary(root *TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left != nil {
		return BTreeIsBinary(root.Left) && root.Data > root.Left.Data
	}
	if root.Right != nil {
		return BTreeIsBinary(root.Right) && root.Data < root.Right.Data
	}
	if root.Parent == nil {
		return false
	}
	return false
}

//btreeisbinary_test.go:50:
//  BTreeIsBinary(&solutions.TreeNode{Left:(*solutions.TreeNode)(0xc000080420),
//     Right:(*solutions.TreeNode)(0xc000080480),
//      Parent:(*solutions.TreeNode)(nil), Data:"04"})
//     == true instead of false
