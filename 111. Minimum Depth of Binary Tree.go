

Success
Details 
Runtime: 184 ms, faster than 98.98% of Go online submissions for Minimum Depth of Binary Tree.
Memory Usage: 18.6 MB, less than 79.70% of Go online submissions for Minimum Depth of Binary Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil{return 0}
    if root.Right == nil && root.Left == nil{
        return 1
    }
    
    if root.Right == nil{return 1+ minDepth(root.Left)}
    if root.Left == nil{return 1 + minDepth(root.Right)}
    
    return 1 + min(  minDepth(root.Left ), minDepth(root.Right)  )
}



func min(a int, b int)int{
    if a< b{
        return a
    }
    return b
}
