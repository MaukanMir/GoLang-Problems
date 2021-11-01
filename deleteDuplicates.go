
// 83. Remove Duplicates from Sorted List
// Easy

// 3512

// 163

// Add to List

// Share
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

 

// Example 1:


// Input: head = [1,1,2]
// Output: [1,2]
// Example 2:


// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
 

// Constraints:

// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.





// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    current_node := head
    
    for current_node != nil{
        for current_node.Next != nil && current_node.Next.Val == current_node.Val{
            current_node.Next = current_node.Next.Next
        }
        
        current_node = current_node.Next
    }
    return head
}
