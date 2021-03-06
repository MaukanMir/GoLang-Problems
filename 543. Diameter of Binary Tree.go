// Given the root of a binary tree, return the length of the diameter of the tree.

// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

// The length of a path between two nodes is represented by the number of edges between them.

 

// Example 1:


// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
// Example 2:

// Input: root = [1,2]
// Output: 1
 

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// -100 <= Node.val <= 100

Success
Details 
Runtime: 3 ms, faster than 95.42% of Go online submissions for Diameter of Binary Tree.
Memory Usage: 4.7 MB, less than 21.48% of Go online submissions for Diameter of Binary Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    
    width:= 0
    
    dfs(root,&width)
    return width
    
}

func dfs( root *TreeNode, width *int)int{
    if root == nil{
        return 0
    }
    
    leftPath := dfs(root.Left, width)
    rightPath := dfs(root.Right, width)
    
    *width = max(*width, leftPath + rightPath)
    
    return max( leftPath, rightPath)+1
}

func max(a,b int)int{
    if a> b{
        return a
    }
    
    return b
}
