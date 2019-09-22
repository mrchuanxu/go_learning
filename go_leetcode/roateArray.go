func rotate(nums []int, k int)  {
    // 其实也满多方法的 第一个方法就是 冒泡直接交换
    if len(nums)<=1||k==0{
        return
    }
    // 冒泡
    mCount := len(nums)-k
    for i:=len(nums)-1;i>=mCount;i--{
        for j:=len(nums)-1;j>0;j--{
            nums[j], nums[j-1]= nums[j-1],nums[j]
        }
    }
   
}