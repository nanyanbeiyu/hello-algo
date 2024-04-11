package linked_list

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// 测试用例1
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	expected := &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}
	result := addTwoNumbers(l1, l2)
	checkList(t, result, expected)
	// 测试用例2
	l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	l2 = &ListNode{Val: 1}
	expected = &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}
	result = addTwoNumbers(l1, l2)
	checkList(t, result, expected)
	// 测试用例3
	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	expected = &ListNode{Val: 0}
	result = addTwoNumbers(l1, l2)
	checkList(t, result, expected)
}

func checkList(t *testing.T, result, expected *ListNode) {
	for result != nil && expected != nil {
		if result.Val != expected.Val {
			t.Errorf("期望 %v，但得到 %v", expected, result)
			return
		}
		result = result.Next
		expected = expected.Next
	}
	if result != nil || expected != nil {
		t.Errorf("期望 %v，但得到 %v", expected, result)
	}

}
