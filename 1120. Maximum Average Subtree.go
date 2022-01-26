Given the root of a binary tree, return the maximum average value of a subtree of that tree. Answers within 10-5 of the actual answer will be accepted.

A subtree of a tree is any node of that tree plus all its descendants.

The average value of a tree is the sum of its values, divided by the number of nodes.

 

Example 1:


Input: root = [5,6,1]
Output: 6.00000
Explanation: 
For the node with value = 5 we have an average of (5 + 6 + 1) / 3 = 4.
For the node with value = 6 we have an average of 6 / 1 = 6.
For the node with value = 1 we have an average of 1 / 1 = 1.
So the answer is 6 which is the maximum.
Example 2:

Input: root = [0,null,1]
Output: 1.00000
 

Constraints:

The number of nodes in the tree is in the range [1, 104].
0 <= Node.val <= 105

Success
Details 
Runtime: 4 ms, faster than 100.00% of Go online submissions for Maximum Average Subtree.
Memory Usage: 6.2 MB, less than 30.00% of Go online submissions for Maximum Average Subtree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func maximumAverageSubtree(root *TreeNode) float64 {
     var maxVal float64
   
    
    
    dfs(root,&maxVal)
    
    return maxVal
}

func dfs(node *TreeNode, maxVal *float64)(int,int){
    if node == nil{
        return 0,0
    }
    
    leftSum, leftCount := dfs(node.Left, maxVal)
    rightSum, rightCount := dfs(node.Right, maxVal)
    total, count := node.Val+ leftSum + rightSum, leftCount + rightCount +1
    
    avg := float64(total) / float64(count)
    
    if avg > *maxVal{
        *maxVal = avg
    }
    return total, count
}
