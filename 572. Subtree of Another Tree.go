Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

 

Example 1:


Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
Example 2:


Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
 

Constraints:

The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-104 <= root.val <= 104
-104 <= subRoot.val <= 104


Success
Details 
Runtime: 15 ms, faster than 79.17% of Go online submissions for Subtree of Another Tree.
Memory Usage: 7 MB, less than 32.95% of Go online submissions for Subtree of Another Tree.
Next challenges:

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil{
        return false
    }
    
    if root.Val == subRoot.Val{
        if areEquals(root, subRoot) == true{
            return true
        }
    }
    
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right,subRoot)
}

func areEquals(one,two *TreeNode)bool{
    if one == nil && two == nil{
        return true
    }
    
    if one == nil || two == nil || one.Val != two.Val{
        return false
    }
    
    return areEquals(one.Left, two.Left) && areEquals(one.Right,two.Right)
}
