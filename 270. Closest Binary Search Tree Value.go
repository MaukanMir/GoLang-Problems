// Given the root of a binary search tree and a target value, return the value in the BST that is closest to the target.

 

// Example 1:


// Input: root = [4,2,5,1,3], target = 3.714286
// Output: 4
// Example 2:

// Input: root = [1], target = 4.428571
// Output: 1
 

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// 0 <= Node.val <= 109
// -109 <= target <= 109

// Success
// Details 
// Runtime: 4 ms, faster than 91.92% of Go online submissions for Closest Binary Search Tree Value.
// Memory Usage: 5 MB, less than 100.00% of Go online submissions for Closest Binary Search Tree Value.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    if(root == nil){
        return 0
    }
    
    minValue := (root.Val)
  
    for root != nil{
        
    
        if math.Abs( float64(root.Val) - target)  < math.Abs( float64(minValue) - target) {
            
            minValue = root.Val
        }
        
        if target > float64(root.Val) {
            root = root.Right
        }else{
            root = root.Left
        }
    }
    
    return minValue
}
