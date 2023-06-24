package main

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return BTreeSearchItem(root.Left, elem)
}
