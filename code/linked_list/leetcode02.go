/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
package linked_list

// addTwoNumbers
//
//	@Description: 为两个链表中的数字节点逐个相加，并返回一个新的链表表示相加后的结果
//	@param l1 链表的头节点指针
//	@param l2 链表的头节点指针
//	@return head 表示相加结果的链表的头节点指针。
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode // 尾指针，用于便捷地添加新节点到链表末尾
	carry := 0         // 进位标志

	// 循环处理直到两个链表都被遍历完
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0 // 存储当前节点的值
		// 为避免空指针异常，分别提取l1和l2链表当前节点的值
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 计算当前节点的和及进位
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		// 如果是链表的第一个节点，则初始化头节点和尾指针
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 否则，通过尾指针添加新节点，并更新尾指针位置
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 处理最后的进位
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return
}
