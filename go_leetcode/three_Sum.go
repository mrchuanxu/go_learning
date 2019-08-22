func threeSum(nums []int) [][]int {
    res := [][]int{}
    if len(nums)==0{
        return res
    }
    // 相当于取三个元素 排列组合
    // 是否重复？ 重复的解法，存在重复的元素 要是排序了 就不会重复了！
    // 将nums []int排序 快排吧
    
    // 决定用指针来计算
    nLen := len(nums)
    quickSort(nums,0,nLen-1)
    for i:=0;i< nLen;i++{
        j,k := i+1,nLen-1
        if i>0&&(nums[i]==nums[i-1]){
            continue
        }
        for j < k {
            if k<(nLen-1)&&(nums[k]==nums[k+1]){ // 多加几个判断 让运行跳过下面的
                k--
                continue
            }
            if(nums[i]+nums[j]+nums[k]>0){ // 从后往前指
                    k--
            }else if(nums[i]+nums[j]+nums[k]<0){
                    j++
            }else{
                tmpRes := []int{}
                tmpRes = append(tmpRes,nums[i])
                tmpRes = append(tmpRes,nums[j])
                tmpRes = append(tmpRes,nums[k])
                res = append(res,tmpRes)
                k--
                j++
                tmpRes = nil  
                }
            }
    }
    return res
}

func quickSort(nums []int,start int,end int){
    if len(nums) == 0{
        return
    }
    if start >= end{
        return
    }
    p := partition(nums,start,end)
    quickSort(nums,start,p-1)
    quickSort(nums,p+1,end)
}

func partition(nums []int,start int,end int)int{
    pivot := nums[end] //  取最后一个元素作为分区点
    i := start
    for j := i;j<end;j++{
        if nums[j]<pivot{
            nums[i],nums[j] = nums[j],nums[i]
            i++
        }
    }
    nums[end],nums[i] = nums[i],nums[end]
    return i
}