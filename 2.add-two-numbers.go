package leetcode

// https://leetcode-cn.com/problems/add-two-numbers/
// 两数相加

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return dfs(l1, l2, 0)
}

func dfs(l, r *ListNode, n int) *ListNode {
	if l == nil && r == nil && n == 0 {
		return nil
	}

	if r != nil {
		n += r.Val
		r = r.Next
	}

	if l != nil {
		n += l.Val
		l = l.Next
	}

	node := new(ListNode)
	node.Val = n % 10
	node.Next = dfs(l, r, n/10)
	return node
}
