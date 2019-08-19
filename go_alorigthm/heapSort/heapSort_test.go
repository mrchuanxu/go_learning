package heapSort
// 堆是一棵完全二叉树
// 堆中的每一个节点的值都必须大雨等于或者小于等于其子树🌲中每个节点的值

// 用数组来存储完全二叉树 节省空间

type iHeap struct{
	isliHeap []int
	count int
	capacity int
}

// 堆化 heapify 从下往上 和从上往下

// 建堆 TODO:
func (iheap *iHeap)heapify(arr []int,capacity int,mid int)(*iHeap){
	// 建堆要将整个数据建立起来
	// 只需要取中间的开始持续堆化就好
	for {
		maxPos := mid
		if mid*2<=capacity&&arr[mid]<arr[mid*2]{ maxPos = mid*2 }
		if (mid*2+1)<=capacity&&arr[maxPos]<arr[mid*2+1]{maxPos = mid*2+1}
		if maxPos == mid{
			break
		}
		arr[mid],arr[maxPos] = arr[maxPos],arr[mid]
		mid = maxPos
	}
	iheap.isliHeap = arr
	iheap.count = len(arr)
	iheap.capacity = len(arr)
	return iheap
}
func (iheap *iHeap)buildHeap(arr []int){
	for i:=len(arr)/2;i>=1;i--{
		iheap = iheap.heapify(arr,len(arr),i)
	}
}

// 堆排序
func (iheap *iHeap)SortHeap()(*iHeap){
	kNum := iheap.count
	i := 1
	for kNum>1{
		iheap.isliHeap[kNum],iheap.isliHeap[i] = iheap.isliHeap[i],iheap.isliHeap[kNum]
		kNum--
		iheap.heapify(iheap.isliHeap,kNum,1) // 持续与1堆化
	}
	return iheap
}
// 插入堆
func (iheap *iHeap)insertData(val int)(*iHeap){
	if iheap.count >= iheap.capacity{
		// errors.New("heap is full")
		return iheap
	}
	iheap.count++
	// 开始持续堆化
	iparent := iheap.count>>1
	ison := iheap.count
	currentHeap := iheap.isliHeap
	for (iparent/2>0)&&currentHeap[ison]>currentHeap[iparent]{
		currentHeap[ison],currentHeap[iparent] = currentHeap[iparent],currentHeap[ison]
		iparent/=2
	}
	iheap.isliHeap = currentHeap
	return iheap
}

// 删除堆顶元素 将堆顶元素与最后的的元素交换 然后持续堆化
func (iheap *iHeap)deleteUpElement()bool{
	if iheap.count == 0{
		return false
	}
	iheap.isliHeap[1] = iheap.isliHeap[iheap.count]
	iheap.isliHeap[iheap.count] = 0
	iheap.count--
	// 开始持续堆化
	i := 1
	flag := false
	for{
		maxPos := i
		if 2*i<=iheap.capacity&&iheap.isliHeap[i]<iheap.isliHeap[i*2]{
			maxPos = i*2
		}
		if 2*i+1<=iheap.capacity&&iheap.isliHeap[maxPos]<iheap.isliHeap[2*i+1]{
			maxPos = i*2+1
		}
		if maxPos == i{
			flag = true
			break
		}
		iheap.isliHeap[i],iheap.isliHeap[maxPos] = iheap.isliHeap[maxPos],iheap.isliHeap[i]
		i = maxPos
	}
	return flag
}





