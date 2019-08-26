const (
	MaxUint =^uint(0)
    MaxInt = int(MaxUint >> 1)
)
func threeSumClosest(nums []int, target int) int {
    // target - sum 取 minValue 比较即可
    // 需要排序吗？ 需要 为什么？因为排序使得数字相加会相近
    nLen := len(nums)
    if nLen==0{
        return 0
    }
    res := 0
    minVal:=MaxInt
    for i:=0;i<nLen;i++{
        j:=i+1
        k:=nLen-1
        // 只有一个解决方案 也就是说 每个元素都是独立的 没有重复的
        vali := target - nums[i]
        for ;j<k;j++{
            valij := vali-nums[j]
            for ;k<nLen&&k>j;k--{
                val := valij-nums[k]
                if minVal>Abs(val){
                    minVal = Abs(val)
                    res = target - val
                }
            }
            k = nLen-1
        }
    }
    return res
}

func Abs(n int)int{
    if n<0{
        return -n
    }
    return n
}