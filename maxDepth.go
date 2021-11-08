// 104. Maximum Depth of Binary Tree
// Easy

// 5140

// 109

// Add to List

// Share
// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

 

// Example 1:


// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Example 2:

// Input: root = [1,null,2]
// Output: 2
// Example 3:

// Input: root = []
// Output: 0
// Example 4:

// Input: root = [0]
// Output: 1

Success
Details 
Runtime: 4 ms, faster than 87.80% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.3 MB, less than 57.21% of Go online submissions for Maximum Depth of Binary Tree.




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{return 0}
    
    left:= maxDepth(root.Left)
    right:= maxDepth(root.Right)
    
    if left > right{
        return  left +1
    }
    return right +1
}
