
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

 

Example 1:


Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:


Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
 

Constraints:

The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104

Success
Details 
Runtime: 8 ms, faster than 93.15% of Go online submissions for Kth Smallest Element in a BST.
Memory Usage: 6.4 MB, less than 85.96% of Go online submissions for Kth Smallest Element in a BST.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    
    if root == nil{
        return 0
    }
    var values []int
    
    dfs(root, &values)
    
    return values[k-1]
    
}

func dfs(node *TreeNode, values *[]int){
    if node == nil{
        return
    }
    
    dfs(node.Left,values)
    *values = append(*values, node.Val)
    dfs(node.Right, values)
}
