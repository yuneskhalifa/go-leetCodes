package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return BTreeTransplant(root, node, nil)
	} else if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	} else {
		successor := BTreeMin(node.Right)
		root = BTreeDeleteNode(root, successor)
		node.Data = successor.Data
		return BTreeTransplant(root, successor, nil)
	}
}
