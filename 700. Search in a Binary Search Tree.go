// You are given the root of a binary search tree (BST) and an integer val.

// Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

 

// Example 1:


// Input: root = [4,2,7,1,3], val = 2
// Output: [2,1,3]
// Example 2:


// Input: root = [4,2,7,1,3], val = 5
// Output: []




Success
Details 
Runtime: 20 ms, faster than 94.64% of Go online submissions for Search in a Binary Search Tree.
Memory Usage: 7.4 MB, less than 76.63% of Go online submissions for Search in a Binary Search Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    
    for root != nil && root.Val != val{
        if root.Val > val{
            root = root.Left
        }else{root = root.Right}
    }
    return root
}



Success
Details 
Runtime: 20 ms, faster than 94.64% of Go online submissions for Search in a Binary Search Tree.
Memory Usage: 7.4 MB, less than 76.63% of Go online submissions for Search in a Binary Search Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val{
        return root
    }
    
    if root.Val > val{
        return  searchBST(root.Left,val)
    }else{
         return searchBST(root.Right,val)
    }
    
    return root
}
