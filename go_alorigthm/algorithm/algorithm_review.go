package main

// // 双向链表
// type doubleNode struct {
// 	val  valstruct
// 	next *doubleNode
// 	pre  *doubleNode
// }

// type singleNode struct {
// 	val  valstruct
// 	next *singleNode
// }

// type valstruct struct {
// 	nums1, nums2 int
// }

// func init() {
// 	// 初始化链表
// 	// 尾插法
// 	tailInsert()
// }

// func tailInsert() (head *doubleNode) {
// 	tmp := &doubleNode{
// 		val:  valstruct{0, 0},
// 		next: nil,
// 	}
// 	head = tmp
// 	for i := 0; i < 10; i++ {
// 		valNode := new(doubleNode)
// 		valNode.val = valstruct{i, i}
// 		tmp.next = valNode
// 		valNode.pre = tmp
// 		tmp = valNode
// 	}
// 	return
// }

// // 回文字符 快慢指针 寻找环的中点  对折

// // 实现单链表反转

// func reveser(head *singleNode) (reverse *singleNode) {
// 	sao := new(singleNode)
// 	tmp := new(singleNode)
// 	for head != nil {
// 		tmp = head.next
// 		head.next = sao
// 		sao = head
// 		head = tmp
// 	}
// 	reverse = sao
// 	return
// }

// // NewStack 一个栈
// type NewStack struct {
// 	// 实现一个只有10的数组栈
// 	arr []int
// }

// // New 初始化栈
// func New() (ns *NewStack) {
// 	return &NewStack{
// 		arr: make([]int, 0),
// 	}
// }

// // Push 推出
// func (ns *NewStack) Push(val int) bool {
// 	ns.arr = append(ns.arr, val)
// 	return true
// }

// // Pop 入栈
// func (ns *NewStack) Pop() bool {
// 	if ns.IsEmpty() {
// 		return false
// 	}
// 	ns.arr = ns.arr[0 : len(ns.arr)-2]
// 	return true
// }

// // IsEmpty 判空
// func (ns *NewStack) IsEmpty() bool {
// 	if len(ns.arr) == 0 {
// 		return true
// 	}
// 	return false
// }

// // Front 首位
// func (ns *NewStack) Front() (val int) {
// 	val = ns.arr[0]
// 	return
// }

// func bianry(arr []int, val int) int {
// 	if len(arr) == 0 {
// 		return -1
// 	}
// 	l, h := 0, len(arr)-1
// 	for l <= h {
// 		mid := l + (h-l)>>1
// 		if arr[mid] > val {
// 			h = mid - 1
// 		}
// 		if arr[mid] < val {
// 			l = mid + 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return -1
// }

// // 堆化 i*2 左 i*2+1右 父i/2

// var heapArr = []int{0, 33, 17, 21, 16, 13, 15, 9, 5, 6, 7, 8, 1, 2}

// func heapInsert(val int) {
// 	// 堆化
// 	heapArr = append(heapArr, val)
// 	// 从下往上堆化
// 	i := len(heapArr) - 1
// 	for i/2 > 0 && heapArr[i] > heapArr[i/2] {
// 		// swap
// 		heapArr[i], heapArr[i/2] = heapArr[i/2], heapArr[i]
// 		i = i / 2 // 继续往上比较
// 	}
// }

// // 把最后一个元素赋值到top，然后从上往下堆化
// func heapDelTop() {
// 	heapArr[1] = heapArr[len(heapArr)-1]
// 	i := 1
// 	hLen := len(heapArr) - 1
// 	for {
// 		maxPos := i
// 		if i*2 < hLen && heapArr[i] < heapArr[i*2] {
// 			maxPos = i * 2 // 左节点
// 		}
// 		// 右节点
// 		if i*2+1 < hLen && heapArr[maxPos] < heapArr[i*2+1] {
// 			maxPos = i*2 + 1
// 		}
// 		if maxPos == i {
// 			break
// 		}
// 		heapArr[i], heapArr[maxPos] = heapArr[maxPos], heapArr[i]
// 		i = maxPos
// 	}

// }

// // graph 无向图
// type singleGraph struct {
// 	v    int               //  顶点个数
// 	IMap map[rune]([]rune) // 哈希表后面一个切片 也可以是链表
// }
