/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 其实就是将 k-n+1个
    // 要是求n就需要知道整个链有多长
    if n <= 0{
        return head
    }
    k := 1
    sHead := head
    nHead := head
    for head.Next!=nil{
        head = head.Next
        k++
    }
    if n>k{
        return sHead
    }
    if n==k{
        sHead = sHead.Next
        return sHead
    }
    nLen := k-n+1
    for i:=1;i<nLen-1;i++{
        nHead = nHead.Next
    }
    
    nHead.Next = nHead.Next.Next
    return sHead
}