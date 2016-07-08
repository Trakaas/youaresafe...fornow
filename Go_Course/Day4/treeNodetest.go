package main

type intTreeNode struct {
	val      int
	children []intTreeNode
}

func (node *intTreeNode) addChildren(a []int) {
	node.children = append(node.children, a...)
}

func (node *intTreeNode) findVal(val int) {
	if node.val == val {
		return node
	}
	for i := 0; i < len(node.children); i++ {
		result := node.children[i].findVal(val)
		if result != nil {
			return result
		}
	}
	return nil
}

func main() {
	tree := &intTreeNode{val: 3}
	tree.addChildren([]int{7, 8, 9})
	(tree.children[0]).addChildren([]int{4, 5})

}
