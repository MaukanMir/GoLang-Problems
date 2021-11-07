




// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
// Memory Usage: 2.1 MB, less than 99.64% of Go online submissions for Binary Tree Preorder Traversal.


// 144. Binary Tree Preorder Traversal
// Easy

// 3121

// 106

// Add to List

// Share
// Given the root of a binary tree, return the preorder traversal of its nodes' values.

 

// Example 1:


// Input: root = [1,null,2,3]
// Output: [1,2,3]
// Example 2:

// Input: root = []
// Output: []
// Example 3:

// Input: root = [1]
// Output: [1]
// Example 4:


// Input: root = [1,2]
// Output: [1,2]
// Example 5:


// Input: root = [1,null,2]
// Output: [1,2]




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }

    var ret []int
    p := root
    stack := make([]*TreeNode, 0)
    stack = append(stack, p)

    for len(stack) != 0 {
        visit := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        ret = append(ret, visit.Val)
        
        if visit.Right != nil {
            stack = append(stack, visit.Right)
        }
        if visit.Left != nil {
            stack = append(stack, visit.Left)
        }
    }
    return ret
}
