
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

 

Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false

Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
Memory Usage: 2.2 MB, less than 22.52% of Go online submissions for Same Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if(p == nil && q == nil){return true}
    if p == nil || q == nil{return false}
    if(p.Val != q.Val){return false}
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right,q.Right)
}
