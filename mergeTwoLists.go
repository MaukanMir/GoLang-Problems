// 21. Merge Two Sorted Lists
// Easy

// 8963

// 893

// Add to List

// Share
// Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.

 

// Example 1:


// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: l1 = [], l2 = []
// Output: []
// Example 3:

// Input: l1 = [], l2 = [0]
// Output: [0]
 

// Constraints:

// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both l1 and l2 are sorted in non-decreasing order.


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.6 MB, less than 56.21% of Go online submissions for Merge Two Sorted Lists.


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil{return l2}
    if l2 == nil{return l1}
    
    if l1.Val < l2.Val{
        l1.Next = mergeTwoLists(l1.Next,l2)
        return l1
    }else{
        l2.Next = mergeTwoLists(l2.Next,l1)
        return l2
    }
    
}
