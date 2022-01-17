// Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

 

// Example 1:


// Input: root = [4,2,6,1,3]
// Output: 1
// Example 2:


// Input: root = [1,0,48,null,null,12,49]
// Output: 1

// Success
// Details 
// Runtime: 12 ms, faster than 86.05% of Go online submissions for Minimum Absolute Difference in BST.
// Memory Usage: 6.6 MB, less than 95.35% of Go online submissions for Minimum Absolute Difference in BST.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    
    var result []int
    
    inOrder(root, &result)
    
    sort.Ints(result)
    var minValue =10000
    var current = 10000
    for i:=1; i< len(result); i++{
        current = result[i] - result[i-1]
       minValue = min(minValue,current )
    }
    
   
   
    return minValue
    
}


func min(a int, b int)int{
    if a < b{
        return a
    }
    return b
}

func inOrder(node *TreeNode, result *[]int){
    if(node == nil){return}
        if node != nil{
            
             inOrder(node.Left, result)
            *result = append(*result, node.Val)
            inOrder(node.Right, result)
        }
    
    }
