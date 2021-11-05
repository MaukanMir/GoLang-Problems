// 206. Reverse Linked List
// Easy

// 9077

// 162

// Add to List

// Share
// Given the head of a singly linked list, reverse the list, and return the reversed list.

 

// Example 1:


// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Example 2:


// Input: head = [1,2]
// Output: [2,1]
// Example 3:

// Input: head = []
// Output: []
 

// Constraints:

// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000



Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Reverse Linked List.



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head != nil && head.Next ==nil{return head}
    
     var prev,curr *ListNode  = nil,head

    
    for curr != nil{
        next_node := curr.Next
        curr.Next = prev
        prev = curr
        curr = next_node
    }
   return prev
}
