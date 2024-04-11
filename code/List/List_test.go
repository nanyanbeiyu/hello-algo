package List

import (
	"testing"
)

// TestNewMylist 测试新建列表的初始状态
func TestNewMylist(t *testing.T) {
	list := newMylist()
	if list.arrCapacity != 10 {
		t.Errorf("Expected initial capacity to be 10, got %d", list.arrCapacity)
	}
	if list.arrSize != 0 {
		t.Errorf("Expected initial size to be 0, got %d", list.arrSize)
	}
	if cap(list.arr) != 10 {
		t.Errorf("Expected underlying array capacity to be 10, got %d", cap(list.arr))
	}
}

// TestSize 测试 size 方法
func TestSize(t *testing.T) {
	list := newMylist()
	if size := list.size(); size != 0 {
		t.Errorf("Expected size to be 0, got %d", size)
	}
	list.add(1)
	if size := list.size(); size != 1 {
		t.Errorf("Expected size to be 1 after adding an element, got %d", size)
	}
}

// TestCapacity 测试 capacity 方法
func TestCapacity(t *testing.T) {
	list := newMylist()
	if capacity := list.capacity(); capacity != 10 {
		t.Errorf("Expected capacity to be 10, got %d", capacity)
	}
}

// TestGet 测试 get 方法
func TestGet(t *testing.T) {
	list := newMylist()
	list.add(10)
	if value := list.get(0); value != 10 {
		t.Errorf("Expected get(0) to return 10, got %d", value)
	}
}

// TestSet 测试 set 方法
func TestSet(t *testing.T) {
	list := newMylist()
	list.add(10)
	list.set(0, 20)
	if value := list.get(0); value != 20 {
		t.Errorf("Expected set(0, 20) to change the value at index 0 to 20, got %d", value)
	}
}

// TestAdd 测试 add 方法
func TestAdd(t *testing.T) {
	list := newMylist()
	// 添加10个元素，这应该不会触发扩容，因为初始容量是10
	for i := 0; i < 10; i++ {
		list.add(i)
	}
	if list.size() != 10 {
		t.Errorf("Expected size to be 10 after adding ten elements")
	}
	if list.capacity() != 10 {
		t.Errorf("Expected capacity to remain at 10 after adding ten elements")
	}

	// 再添加一个元素，这应该触发扩容，因为当前容量已满
	list.add(10)
	if list.size() != 11 {
		t.Errorf("Expected size to be 11 after adding another element")
	}
	if list.capacity() != 20 { // 扩容后，容量应该是原来的两倍
		t.Errorf("Expected capacity to be doubled after adding an eleventh element")
	}
}

// TestInsert 测试 insert 方法
func TestInsert(t *testing.T) {
	list := newMylist()
	// 先添加一个元素10
	list.add(10)
	// 然后在索引0处插入一个新元素20
	list.insert(0, 20)
	// 验证索引0处的元素是否为20
	if value := list.get(0); value != 20 {
		t.Errorf("Expected insert(0, 20) to insert 20 at index 0, but got %d", value)
	}
	// 验证原来的元素10是否移动到了索引1
	if value := list.get(1); value != 10 {
		t.Errorf("Expected the original element 10 to be shifted to index 1 after insertion, but got %d", value)
	}
	// 验证列表的大小是否正确
	if size := list.size(); size != 2 {
		t.Errorf("Expected size to be 2 after inserting an element, but got %d", size)
	}
}

// TestRemove 测试 remove 方法
func TestRemove(t *testing.T) {
	list := newMylist()
	list.add(10)
	list.add(20)
	removedValue := list.remove(0)
	if removedValue != 10 || list.get(0) != 20 {
		t.Errorf("Expected remove(0) to remove 10 and shift 20 to index 0")
	}
}

// TestExtendCapacity 测试 extendCapacity 方法
func TestExtendCapacity(t *testing.T) {
	list := newMylist()
	for i := 0; i < 10; i++ {
		list.add(i)
	}
	list.extendCapacity()
	if list.capacity() != 20 {
		t.Errorf("Expected capacity to be doubled after extending")
	}
}
