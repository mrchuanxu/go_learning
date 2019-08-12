func twoSum(nums []int, target int) []int {
    if nums == nil{
        return nil
    }
    var result []int
    imap := map[int]int{}
    for i:=0;i<len(nums);i++{
        if _,ok:=imap[target-nums[i]];!ok{
            imap[nums[i]] =i
        }else{
            result = append(result,imap[target-nums[i]])
            result = append(result,i)
        }
    }
    return result
}