
// Given the root of a binary tree, return the sum of values of its deepest leaves.
 

// Example 1:


// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// Output: 15
// Example 2:

// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 19
 

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// 1 <= Node.val <= 100

Success
Details 
Runtime: 52 ms, faster than 97.33% of Go online submissions for Deepest Leaves Sum.
Memory Usage: 8 MB, less than 98.67% of Go online submissions for Deepest Leaves Sum.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    if root == nil{return 0}
    
    max:=-1
    sum :=0
    
    var dfs func(*TreeNode,int)
    dfs = func(root *TreeNode, level int){
        if root == nil{
            return 
        }
        
        if level > max{
            if max < level{
                max = level
                sum = root.Val
            }
        }else if level == max{
            sum += root.Val
        }
        
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    
    dfs(root,0)
    return sum
}
