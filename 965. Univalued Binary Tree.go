// A binary tree is uni-valued if every node in the tree has the same value.

// Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

 

// Example 1:


// Input: root = [1,1,1,1,1,null,1]
// Output: true
// Example 2:


// Input: root = [2,2,2,5,2]
// Output: false


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Univalued Binary Tree.
Memory Usage: 2.3 MB, less than 66.67% of Go online submissions for Univalued Binary Tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil{return true}
    
    if root.Left != nil{
        if root.Val != root.Left.Val{return false}
    }
    
    if(root.Right != nil){
        if root.Val != root.Right.Val{return false}
    }
    
    return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
