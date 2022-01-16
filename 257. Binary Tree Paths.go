// Given the root of a binary tree, return all root-to-leaf paths in any order.

// A leaf is a node with no children.

 

// Example 1:


// Input: root = [1,2,3,null,5]
// Output: ["1->2->5","1->3"]
// Example 2:

// Input: root = [1]
// Output: ["1"]

// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
// Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Binary Tree Paths.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    
    seen := []string{}
    dfs(root, strconv.Itoa(root.Val), &seen)
    return seen
}

func dfs(root *TreeNode, ans string, seen *[]string){
    if root.Left == nil && root.Right == nil{
        *seen = append(*seen, ans)
        return
    }
    
    if root.Left != nil{
        dfs(root.Left, ans + "->" + strconv.Itoa(root.Left.Val), seen)
    }
    
    if root.Right != nil{
        dfs(root.Right, ans + "->" + strconv.Itoa(root.Right.Val), seen)
    }
    
    return 
    
    
}
