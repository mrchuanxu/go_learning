var result []int
var res [][]string

func solveNQueens(n int) [][]string {
    // 这个其实就是8皇后问题类似 回溯法解决就好
    if n==1{
        str :=[]string{"Q"}
        res = append(res,str)
        return res
    }
    if n<=3{
        return res
    }
    result = make([]int,n)
    calQueens(0,n)
    return res
}

func calQueens(row int,m int){
    if row==m{
        strRes := printQueens(result)
        // 判断前面的和后面的
        if len(res)>=1{
            if res[len(res)-1][0]==strRes[0]{
                return
            }
        }
        res = append(res,printQueens(result))
        return
    }
    for col:=0;col<m;col++{
        if isOk(row,col){
            result[row] = col
            calQueens(row+1,m)
        }
    }
}
func isOk(row,col int)bool{
    leftup := col-1
    rightup := col+1
    for i:=row-1;i>=0;i--{
        // 考察当前行
        if result[i]==col{
            return false
        }
        // 考察对角线
        if leftup>=0{
            if result[i]==leftup{
                return false
            }
        }
        if rightup<row{
            if result[i]==rightup{
                return false
            }
        }
        rightup++
        leftup--
    }
    return true
}

func printQueens(resQueen []int)[]string{
    var reStr []string
    var str string = ""
    for row:=0;row<len(resQueen);row++{
        str=""
        for col:=0;col<len(resQueen);col++{
            if resQueen[row]==col{
                str+="Q"
            }else{
                str+="."
            }
        }
        reStr = append(reStr,str)
    }
    return reStr
}