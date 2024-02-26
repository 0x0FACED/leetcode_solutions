package task100

import (
	"leetcode_solutions/internal/queue"
	"leetcode_solutions/internal/treenode"
)

func IsSameTree(p *treenode.TreeNode, q *treenode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	var queue = queue.New()
	queue.PushLast(p)
	queue.PushLast(q)

	for queue.Head() != nil {
		pNode := queue.PopFirst().(*treenode.TreeNode)
		qNode := queue.PopFirst().(*treenode.TreeNode)

		if pNode.Val != qNode.Val {
			return false
		}

		if pNode.Left != nil && qNode.Left != nil {
			queue.PushLast(pNode.Left)
			queue.PushLast(qNode.Left)
		} else if pNode.Left != nil || qNode.Left != nil {
			return false
		}

		if pNode.Right != nil && qNode.Right != nil {
			queue.PushLast(pNode.Right)
			queue.PushLast(qNode.Right)
		} else if pNode.Right != nil || qNode.Right != nil {
			return false
		}
		queue.PrintQueue()
	}

	return true
}
