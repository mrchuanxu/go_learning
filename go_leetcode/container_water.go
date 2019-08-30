const (
    INT_MAX = int(^uint(0) >> 1)
    INT_MIN = ^INT_MAX
)
// brute force
func maxArea(height []int) int {
    // 数组下标是x
    // 数组个数是y
    // 因此 数组下标(x1-x2)*min(y1,y2)
    maxVal := INT_MIN
    hLen := len(height)
    if hLen <=1{
        return height[0]
    }
    for i:=0;i<hLen;i++{ // O(n^2)
        for j:=i+1;j<hLen;j++{
            Val := (j-i)*minVal(height[i],height[j])
            if maxVal < Val{
                maxVal = Val
            }
        }
    }
    return maxVal
    
}

func minVal(y1,y2 int)int{
    if y1>y2{
        return y2
    }
    return y1
}

// 优化 我觉得switch需要恶补
const (
    INT_MAX = int(^uint(0) >> 1)
    INT_MIN = ^INT_MAX
)
func maxArea(height []int) int {
    // 数组下标是x
    // 数组个数是y
    // 因此 数组下标(x1-x2)*min(y1,y2) 维护两个指针
    mxVal,l,r := INT_MIN,0,len(height)-1
    for l<r{
        mxVal = mVal(mxVal,mVal(height[l],height[r],0)*(r-l),1)
        if height[l]<height[r]{
            l++
        }else{
            r--
        }
    }
    return mxVal
}

func mVal(y1,y2,status int)int{
    if status == 1{
        return Max(y1,y2)
    }
    return Min(y1,y2)
}

func Min(y1,y2 int)int{
    if y1>y2{
        return y2
    }else{
        return y1
    }
}

func Max(y1,y2 int)int{
    if y1>y2{
        return y1
    }else{
        return y2
    }
}