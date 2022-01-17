// Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.

// If the tree has more than one mode, return them in any order.

// Assume a BST is defined as follows:

// The left subtree of a node contains only nodes with keys less than or equal to the node's key.
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
// Both the left and right subtrees must also be binary search trees.
 

// Example 1:


// Input: root = [1,null,2,2]
// Output: [2]
// Example 2:

// Input: root = [0]
// Output: [0]

// Success
// Details 
// Runtime: 28 ms, faster than 7.04% of Go online submissions for Find Mode in Binary Search Tree.
// Memory Usage: 6.3 MB, less than 90.14% of Go online submissions for Find Mode in Binary Search Tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    if root == nil { return []int{} }
  

    var final []int
    var maxValue = -10000
    visited := make(map[int]int)
    got := inOrder(root, visited)
    for _, v:= range got{
        if v > maxValue{ maxValue =v}
    }
    for k,v:= range visited{
        if v== maxValue{
            final = append(final, k)
        }
    }
    
    if len(final) ==0{
        return []int{0}
    }
    
   
    return final
}


func inOrder(node *TreeNode, visited map[int]int) map[int]int {
    if node == nil{
        return visited
    }
    visited[node.Val]++
    inOrder(node.Left,visited)
    inOrder(node.Right,visited)
    return visited
}
