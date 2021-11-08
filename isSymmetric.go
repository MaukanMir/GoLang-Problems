// 101. Symmetric Tree
// Easy

// 7726

// 192

// Add to List

// Share
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

 

// Example 1:


// Input: root = [1,2,2,3,4,4,3]
// Output: true
// Example 2:


// Input: root = [1,2,2,null,3,null,3]
// Output: false
 

// Constraints:

// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
 


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool { 
    return isMirror( root,root)  }

func isMirror(left, right *TreeNode)bool{
    if left == nil && right == nil{return true}
    if left == nil || right == nil{return false}
    
    return left.Val == right.Val && isMirror(left.Left,right.Right) && isMirror(left.Right,right.Left)
}



Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 54.85% of Go online submissions for Symmetric Tree.
