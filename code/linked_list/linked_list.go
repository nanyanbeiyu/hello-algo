package main

import "fmt"

/*
链表（linked list）：链表是一种物理存储单元上非连续、非顺序的存储结构，它是链表中第一个结点的地址引用，而不是像数组一样引用数组中元素的位置。
	链表由一系列结点组成，每个结点由一个存储数据元素的值和一个指向下一个结点的指针组成。
链表的组成单位是节点（node）对象。节点对象由两部分组成：数据元素和指针（指向下一个节点的地址）。
	1. 链表的首节点称为“头结点”，最后一个节点被称为“尾节点”。
	2. 尾节点指向的是“空”。
*/

// ListNode
// @Description: 链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一个节点
}

// NewListNode
//
//	@Description: 初始化节点
//	@param val
//	@return *ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

/*
链表常见操作：

 1. 初始化链表 NewListNode ：建立链表分为两步
    第一步是初始化各个节点对象。
    第二步是构建节点之间的引用关系。初始化完成后，我们就可以从链表的头节点出发，通过引用指向 next 依次访问所有节点。

 2. 插入节点 insertNode ：在链表头部插入节点，在链表尾部插入节点，在链表中间插入节点。
    假设我们想在相邻的两个节点 n0 和 n1 之间插入一个新节点 P ，则只需改变两个节点引用（指针）即可，时间复杂度为 O(1)。

 3. 删除节点 removeNode ：删除链表中的节点。在链表中删除节点也非常方便，只需改变一个节点的引用（指针）即可。

 4. 访问节点 access ：在链表中访问节点的效率较低，因为我们可以在O(1)时间下访问数组中的任意元素。链表则不然，程序需要从头节点出发，逐个向后遍历，直至找到目标节点

 5. 查找节点 findNode ：遍历链表，查找其中值为 target 的节点，输出该节点在链表中的索引。此过程也属于线性查找。

 6. 遍历链表 traverse ：遍历链表，输出链表中所有节点的值。
*/
/*在链表n1后添加节点p*/
func insertNode(n1, p *ListNode) {
	n := n1.Next // 获取n1的下一个节点n2
	p.Next = n   // 将p的next指向n2
	n1.Next = p  // 将n1的next指向p
}

/*删除链表的节点n1后的首个节点*/
func removeNode(n1 *ListNode) {
	if n1.Next == nil {
		return
	}
	// n1->p->n2
	P := n1.Next // 获取n1的下一个节点p
	n2 := P.Next // 获取p的下一个节点n2
	n1.Next = n2 // 将n1的next指向n2
}

/*访问链表中索引为index的节点*/
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

/* 在链表中查找值为 target 的首个节点 */
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

/*遍历链表*/
func traverse(head *ListNode) {
	for head != nil {
		if head.Next == nil {
			fmt.Printf("%d\n", head.Val)
			break
		}
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func main() {
	// 初始化各个节点
	node1 := NewListNode(1)
	node2 := NewListNode(3)
	node3 := NewListNode(2)
	node4 := NewListNode(4)
	//构建节点之间的引用
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	// 访问节点
	res := access(node1, 1)
	fmt.Println(res.Val)
	// 插入节点
	node5 := NewListNode(5)
	insertNode(node1, node5)
	// 遍历链表
	traverse(node1)
	// 删除节点
	removeNode(node1)
	traverse(node1)
	// 查找节点
	fmt.Println(findNode(node1, 3))
}
