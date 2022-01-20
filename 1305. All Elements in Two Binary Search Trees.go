// Given two binary search trees root1 and root2, return a list containing all the integers from both trees sorted in ascending order.

 

// Example 1:


// Input: root1 = [2,1,4], root2 = [1,0,3]
// Output: [0,1,1,2,3,4]
// Example 2:


// Input: root1 = [1,null,8], root2 = [8,1]
// Output: [1,1,8,8]

// Success
// Details 
// Runtime: 112 ms, faster than 60.00% of Go online submissions for All Elements in Two Binary Search Trees.
// Memory Usage: 8.4 MB, less than 55.38% of Go online submissions for All Elements in Two Binary Search Trees.



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    var result []int
    
    findNodes(root1, &result)
    findNodes(root2,&result)
    
    sort.Ints(result)
    return result
}

func findNodes(node *TreeNode, result*[]int){
    if node == nil{
        return 
    }
    
    if(node != nil){
        *result = append(*result, node.Val)
    }
    
    if(node.Left != nil){ findNodes(node.Left, result)}
    if(node.Right !=nil){findNodes(node.Right,result)}
}
