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
    // 其实走n+1 快慢指针会更快
    if n <= 0{
        return head
    }
    k := 0
    sHead := head
    nHead := new(ListNode)
    nHead.Next = head
    for head!=nil{
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
    head = nHead
    for k = k-n;k>0;k--{
        head = head.Next
    }
    
    head.Next = head.Next.Next
    return sHead
}

// n+1 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0{
        return head
    }
    sHead := new(ListNode)
    sHead.Next = head
    second := sHead
    first := sHead

    for i:=0;i<n+1;i++{
        first = first.Next
    }

    for first!=nil{
        first = first.Next
        second = second.Next
    }
    second.Next = second.Next.Next
    return sHead.Next
}