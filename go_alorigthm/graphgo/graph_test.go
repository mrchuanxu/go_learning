package graphgo

import (
	// "testing"
	"fmt"
)

// graph
type iMap struct{
	imap map[rune]([]rune)
	v int // 顶点个数
}
// 无🤮
func (bMap *iMap) addEdge(x,y rune){
	bMap.imap[x] = append(bMap.imap[x],y)
	bMap.imap[y] = append(bMap.imap[y],x)
}
// 

// build graph
func (bMap *iMap) findPathBfs(s,t rune){
	if s==t{
		return
	}
	visited := map[rune]bool{}
	queue := New(bMap.v) // 队列存储访问的数组 map
	prev := map[rune]rune{}
	// 把第一个元素入队列
	queue.enqueue(s)
	visited[s] = true
	// 开始访问
	for queue.Len()!=0{
		w,_:= queue.dequeue().(rune)
		for i:=0;i<len(bMap.imap[w]);i++{
			ele := bMap.imap[w][i]
			if !visited[ele]{
				prev[ele] = w
				if ele == t{
					toPrint(prev,s,t)
					return
				}
				queue.enqueue(ele)
				visited[ele] = true
			}
		}
	}
}

func toPrint(prev map[rune]rune,s rune,t rune){
	if _,ok:=prev[t];ok&&t!=s{
		toPrint(prev,s,prev[t])
	}
	fmt.Println(t)
}
// 求s到t的最短路径

var found bool = false

func (bMap *iMap)findPathDfs(s rune,t rune){
	found = false
	if s == t{
		return
	}

	visited := map[rune]bool{}
	prev := map[rune]rune{}
	bMap.recurf(prev,visited,s,t)
	toPrint(prev,s,t)
}

func (bMap *iMap)recurf(prev map[rune]rune,visited map[rune]bool,s rune,t rune){
	if found == true{
		return
	}
	if s == t{
		found = true
		return
	}
	visited[s] = true
	for i:=0;i<len(bMap.imap[s]);i++{
		q := bMap.imap[s][i]
		if !visited[q]{
			prev[q] = s
			bMap.recurf(prev,visited,q,t)
		}
	}
}