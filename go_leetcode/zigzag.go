func convert(s string, numRows int) string {
    // 不要被表象迷惑，其实还是一个数学问题 不要用什么栈和队列
    sLen := len(s)
    if numRows<=1{
        return s
	}
	/** Characters in row 00 are located at indexes k \; (2 \cdot \text{numRows} - 2)k(2⋅numRows−2)
    Characters in row \text{numRows}-1numRows−1 are located at indexes k \; (2 \cdot \text{numRows} - 2) + \text{numRows} - 1k(2⋅numRows−2)+numRows−1
    Characters in inner row ii are located at indexes k \; (2 \cdot \text{numRows}-2)+ik(2⋅numRows−2)+i and (k+1)(2 \cdot \text{numRows}-2)- i(k+1)(2⋅numRows−2)−i.
    **/
    ret := ""
    midCount := (numRows-2)+numRows // 稳定是偶数
    for i:=0;i<numRows;i++{
        for j:=0;j+i<sLen;j+=midCount{
            ret += string(s[j+i]);
            if i!=0&&i!=numRows-1&&(j+midCount-i)<sLen{
                ret+=string(s[j+midCount-i])
            }
        }
    }
    return ret
}