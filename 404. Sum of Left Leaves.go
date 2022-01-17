Given the root of a binary tree, return the sum of all left leaves.

 

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
Example 2:

Input: root = [1]
Output: 0
 

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-1000 <= Node.val <= 1000

Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Left Leaves.
Memory Usage: 2.8 MB, less than 23.60% of Go online submissions for Sum of Left Leaves.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil{return 0}
    
    return process(root, false)
}

func process(root *TreeNode, is_left bool)int{
    if root.Left == nil && root.Right == nil{
        if(is_left == true){
            return root.Val
        }
        return 0
    }
    
    var total =0 
    if root.Left != nil{
        total += process(root.Left, true)
    }
    
    if(root.Right != nil){
        total += process(root.Right, false)
    }
    
    return total
}
