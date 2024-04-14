package stack

/*
栈（stack）是一种遵循先入后出逻辑的线性数据结构。
栈的常规操作：
1. 压栈（push）：将元素压入栈顶。
2. 弹栈（pop）：从栈顶弹出元素。
3. 查询栈顶元素（peek）：返回栈顶元素，但不弹出。p
*/

/*在Go中，推荐将slice作为栈的实现方式。*/
// 初始化栈
var stack []int

func push(val int) {
	stack = append(stack, val)
}

func pop() int {
	if len(stack) == 0 {
		return -1
	}
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}

func peek() int {
	if len(stack) == 0 {
		return -1
	}
	return stack[len(stack)-1]
}
