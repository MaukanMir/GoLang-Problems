
// 876. Middle of the Linked List
// Easy

// 3603

// 95

// Add to List

// Share
// Given the head of a singly linked list, return the middle node of the linked list.

// If there are two middle nodes, return the second middle node.

 

// Example 1:


// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.
// Example 2:


// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
 

Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
Memory Usage: 2 MB, less than 43.59% of Go online submissions for Middle of the Linked List.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow,fast := head,head
    
    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
