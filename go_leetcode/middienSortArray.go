func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // 其实要求O(log(m+n))  归并的思路可以尝试用一下 其实我觉得就是归并排序的变形
    var res float64
    // 两个有序的数组排序 取出中位数
    // 边界条件
    if len(nums1) == 0&&len(nums2)!=0{
        if len(nums2)>1{
            if len(nums2)%2 == 0{
                return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2-1])/2
            }else{
                return float64(nums2[len(nums2)/2])
            }
            }else{
                return float64(nums2[0])
            }
    }
    if len(nums2) == 0&&len(nums1)!=0{
        if len(nums1)>1{
        if len(nums1)%2 == 0{
        return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1])/2 
        }else{
            return float64(nums1[len(nums1)/2])
        }
        }else{
            return float64(nums1[0])
        }
    }
    if len(nums2)==0&&len(nums1)==0{
        return 0.0
    }
    nLen := len(nums1)+len(nums2)
    resArr := []int{}
    if nLen%2==0{
        pos1,pos2 := nLen/2-1,nLen/2
        // 计数 取出这两个位置的数据
        i,j,k:=0,0,0
        for i<len(nums1)&&j <len(nums2){
            if nums1[i]<nums2[j]{
                resArr = append(resArr,nums1[i])
                k++
                i++ // 考虑边界条件
            }else{
                resArr = append(resArr,nums2[j])
                k++
                j++
            }
        }
        for i<len(nums1){
            resArr = append(resArr,nums1[i])
            if k == pos2{
                res = float64((resArr[pos1]+resArr[pos2]))/2
                return res
            }
            k++
            i++
        }
        for j<len(nums2){
            resArr = append(resArr,nums2[j])
            if k == pos2{
                res = float64((resArr[pos1]+resArr[pos2]))/2
            return res
            }
            k++
            j++
        }
        res = float64((resArr[pos1]+resArr[pos2]))/2
        return res
    }else{
        pos := nLen/2
         i,j,k:=0,0,0
        for i<len(nums1)&&j <len(nums2){
            if nums1[i]<nums2[j]{
                resArr = append(resArr,nums1[i])
                k++
                i++ // 考虑边界条件
            }else{
                resArr = append(resArr,nums2[j])
                k++
                j++
            }
        }
        for i<len(nums1){
            resArr = append(resArr,nums1[i])
             if k == pos{
                res = float64(resArr[pos])
                return res
            }
            k++
            i++
        }
        for j<len(nums2){
            resArr = append(resArr,nums2[j])
            if k == pos{
                res = float64(resArr[pos])
                return res
            }
            k++
            j++
        }
        res = float64(resArr[pos])
        return res
    }
    return res
}