package List

/*
列表（list）是一个抽象的数据结构概念，它表示元素的有序集合，支持元素访问、修改、添加、删除和遍历等操作，无须使用者考虑容量限制的问题。列表可以基于链表或数组实现。
	1.链表天然可以看作一个列表，其支持元素增删查改操作，并且可以灵活动态扩容。
	2.数组也支持元素增删查改，但由于其长度不可变，因此只能看作一个具有长度限制的列表。
*/

/*
列表实现:
主要包含三个重点设计：
	1.初始容量：选取一个合理的数组初始容量。在本示例中，我们选择 10 作为初始容量。
	2.数量记录：声明一个变量 size ，用于记录列表当前元素数量，并随着元素插入和删除实时更新。根据此变量，我们可以定位列表尾部，以及判断是否需要扩容。
	3.扩容机制：若插入元素时列表容量已满，则需要进行扩容。先根据扩容倍数创建一个更大的数组，再将当前数组的所有元素依次移动至新数组。在本示例中，我们规定每次将数组扩容至之前的 2 倍。
*/

// myList
// @Description: 列表类
type myList struct {
	arrCapacity int   // 容量
	arrSize     int   // 数量
	arr         []int // 数组
	extendRatio int   // 扩容倍数
}

func newMylist() *myList {
	return &myList{
		arrCapacity: 10,
		arrSize:     0,
		arr:         make([]int, 10),
		extendRatio: 2,
	}
}

/* size 获取列表元素数量*/
func (l *myList) size() int {
	return l.arrSize
}

/* capacity 获取列表容量*/
func (l *myList) capacity() int {
	return l.arrCapacity
}

/* get 获取指定索引位置的元素*/
func (l *myList) get(index int) int {
	// 索引越界额，则抛出异常
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	return l.arr[index]
}

/* set 设置指定索引位置的元素*/
func (l *myList) set(index, val int) {
	// 索引越界额，则抛出异常
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	l.arr[index] = val
}

/* add 添加元素*/
func (l *myList) add(val int) {
	//当元素满了，则扩容
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = val
	// 元素数量+1
	l.arrSize++
}

/* insert 插入元素*/
func (l *myList) insert(index, val int) {
	// 索引越界额，则抛出异常
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	//当元素满了，则扩容
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	// 将索引以及后面的元素都向后移动一位
	// val 10 index 0
	for j := l.arrSize - 1; j >= index; j-- {
		l.arr[j+1] = l.arr[j]
	}
	l.arr[index] = val
	// 元素数量+1
	l.arrSize++
}

/* remove 删除元素**/
func (l *myList) remove(index int) int {
	// 索引越界额，则抛出异常
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	num := l.arr[index]
	// 将索引以及后面的元素都向前移动一位
	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	// 元素数量-1
	l.arrSize--
	return num
}

/* extendCapacity 扩容*/
func (l *myList) extendCapacity() {
	// 新建一个长度为原数组extendRatio倍的数组，并将原数组复制到新数组中
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	// 更新列表容量
	l.arrCapacity = len(l.arr)
}
func (l *myList) toArray() []int {
	// 仅转换有效长度范围内的列表元素
	return l.arr[:l.arrSize]
}
