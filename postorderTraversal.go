// 145. Binary Tree Postorder Traversal
// Easy

// 3311

// 128

// Add to List

// Share
// Given the root of a binary tree, return the postorder traversal of its nodes' values.

 

// Example 1:


// Input: root = [1,null,2,3]
// Output: [3,2,1]
// Example 2:

// Input: root = []
// Output: []
// Example 3:

// Input: root = [1]
// Output: [1]
// Example 4:


// Input: root = [1,2]
// Output: [2,1]
// Example 5:


// Input: root = [1,null,2]
// Output: [2,1]
 

// Constraints:

// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 99.74% of Go online submissions for Binary Tree Postorder Traversal.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    arr := make([]int,0)
     var search func(*TreeNode)
    
     search = func(node *TreeNode){
        if node == nil{return}
        search(node.Left)
        search(node.Right)
        arr = append(arr, node.Val)
    }
    
    search(root)
    
    return arr
}
