
// 这个算法是最快的，好像都是围绕中间的算起 一直算到结尾 遍历一次就可以了
func longestPalindrome(s string) string {
    nLen := len(s)
    if nLen <= 1{
        return s
    }
    start,end:=0,0
    for i:=0;i<nLen;i++{
        len1 := maxPalindrome(s,i,i)
        len2 := maxPalindrome(s,i,i+1)
        maxLen := MaxLen(len1,len2)
        if maxLen>end-start{
            start = i-(maxLen-1)/2
            end = i+maxLen/2
        }
    }
    return s[start:end+1]
}

// 先从中间算起吧
func maxPalindrome(s string, left,right int)int{
    nLen := len(s)
    for left >= 0 && right<nLen && s[left]==s[right]{
        left--
        right++
    }
    return right-left-1
}

func MaxLen(len1,len2 int)int{
    if len1>len2{
        return len1
    }
    return len2
}
