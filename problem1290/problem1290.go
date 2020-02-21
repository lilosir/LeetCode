/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getDecimalValue(head *ListNode) int {
    var binSlice []int
    
    for head != nil {
        binSlice = append(binSlice, head.Val)
        head = head.Next
    }
    
    result := 0
    for i:=len(binSlice)-1; i>=0; i-- {
        result += binSlice[i] * (1 << (len(binSlice)-1-i))
    }
    
    return result
}