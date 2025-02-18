package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	/* false if left is > than node */
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	/* false if right is < than node */
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	/* false if, recursively, the left or right is not a BST */
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}
	/* passing all that, it's a BST */
	return true

}

//btreeisbinary_test.go:50:
//  BTreeIsBinary(&solutions.TreeNode{Left:(*solutions.TreeNode)(0xc000080420),
//     Right:(*solutions.TreeNode)(0xc000080480),
//      Parent:(*solutions.TreeNode)(nil), Data:"04"})
//     == true instead of false
