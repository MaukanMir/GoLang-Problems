// Given the root of a binary tree, collect a tree's nodes as if you were doing this:

// Collect all the leaf nodes.
// Remove all the leaf nodes.
// Repeat until the tree is empty.
 

// Example 1:


// Input: root = [1,2,3,4,5]
// Output: [[4,5,3],[2],[1]]
// Explanation:
// [[3,5,4],[2],[1]] and [[3,4,5],[2],[1]] are also considered correct answers since per each level it does not matter the order on which elements are returned.
// Example 2:

// Input: root = [1]
// Output: [[1]]
 

// Constraints:

// The number of nodes in the tree is in the range [1, 100].
// -100 <= Node.val <= 100

// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Leaves of Binary Tree.
// Memory Usage: 2.1 MB, less than 49.54% of Go online submissions for Find Leaves of Binary Tree



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    
    values := make([][]int,0)
    dfs(root, &values)
    return values
}

func dfs(root *TreeNode, values *[][]int)int{
    if root == nil{
    return -1
    }
    
    left := dfs(root.Left, values)
    right := dfs(root.Right, values)
    
    maxValue := max(left,right)+1
    if len(*values) < maxValue+1{
        *values = append(*values, make([]int,0))
    }
    (*values)[maxValue] = append((*values)[maxValue], root.Val)
    return maxValue
    
}

func max(a,b int)int{
    if a> b{
        return a
    }
    return b
}
