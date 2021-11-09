// 226. Invert Binary Tree
// Easy

// 6838

// 96

// Add to List

// Share
// Given the root of a binary tree, invert the tree, and return its root.

 

// Example 1:


// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Example 2:


// Input: root = [2,1,3]
// Output: [2,3,1]
// Example 3:

// Input: root = []
// Output: []
 

// Constraints:

// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100



Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
Memory Usage: 2.1 MB, less than 46.77% of Go online submissions for Invert Binary Tree.




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root != nil{
        root.Left,root.Right = root.Right,root.Left
        invertTree(root.Left)
        invertTree(root.Right)
    }
    return root
}




Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
Memory Usage: 2.2 MB, less than 46.77% of Go online submissions for Invert Binary Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root != nil{
        queue := []*TreeNode{root}
        for len(queue) >0{
            current:= queue[0]
            queue = queue[1:]
            
            if current == nil{
                continue
            }
            current.Left, current.Right = current.Right,current.Left
            
            queue = append (queue,current.Left,current.Right)
        }
    }
    return root
}
