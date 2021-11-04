// 19. Remove Nth Node From End of List
// Medium

// 7397

// 374

// Add to List

// Share
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

 

// Example 1:


// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]
 

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz




Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Remove Nth Node From End of List



func removeNthFromEnd(head *ListNode, n int) *ListNode {
    counter:=1
 
    slow,fast := head,head
    
    
    for counter <=n{
        fast = fast.Next
        counter +=1
    }
    
    if fast == nil{
        return head.Next
    }
    
    for fast.Next != nil{
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    
    return head
}
