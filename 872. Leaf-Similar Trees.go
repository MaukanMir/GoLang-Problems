// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.



// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

// Two binary trees are considered leaf-similar if their leaf value sequence is the same.

// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

 

// Example 1:


// Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// Output: true
// Example 2:


// Input: root1 = [1,2,3], root2 = [1,3,2]
// Output: false
 

// Constraints:

// The number of nodes in each tree will be in the range [1, 200].
// Both of the given trees will have values in the range [0, 200].

// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Leaf-Similar Trees.
// Memory Usage: 2.6 MB, less than 63.16% of Go online submissions for Leaf-Similar Trees.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    
    one := make([]*TreeNode,0)
    two := make([]*TreeNode,0)
    
    find(root1, &one)
    find(root2,&two)
    
    if len(one) != len(two){return false}
    
    for i:=0; i< len(one); i++{
        if one[i].Val != two[i].Val{
            return false
        }
    }
    
    return true
}

func find(root *TreeNode, got *[]*TreeNode){
    if root == nil{
        return
    }
    
    if root.Left == nil && root.Right == nil{
        *got = append(*got, root)
        return
    }
    
    find(root.Left, got)
    find(root.Right,got)
}
