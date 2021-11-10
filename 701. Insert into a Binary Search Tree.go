// You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

// Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

 

// Example 1:


// Input: root = [4,2,7,1,3], val = 5
// Output: [4,2,7,1,3,5]
// Explanation: Another accepted tree is:

// Example 2:

// Input: root = [40,20,60,10,30,50,70], val = 25
// Output: [40,20,60,10,30,50,70,null,null,25]
// Example 3:

// Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
// Output: [4,2,7,1,3,5]




Success
Details 
Runtime: 20 ms, faster than 100.00% of Go online submissions for Insert into a Binary Search Tree.
Memory Usage: 7.5 MB, less than 39.27% of Go online submissions for Insert into a Binary Search Tree.




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil{return &TreeNode{Val: val}}
    
    if val > root.Val{
        root.Right = insertIntoBST(root.Right,val)
    }else{
        root.Left = insertIntoBST(root.Left,val)
    }
    
    return root
    
}


Success
Details 
Runtime: 28 ms, faster than 90.05% of Go online submissions for Insert into a Binary Search Tree.
Memory Usage: 7.5 MB, less than 39.27% of Go online submissions for Insert into a Binary Search Tree.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    
    node := root
    
    for node != nil{
        if node.Val < val{
            
            if node.Right == nil{
                node.Right = &TreeNode{Val: val}
                return root
            }else{
                node = node.Right
            }
        }else{
            if node.Left == nil{
                node.Left = &TreeNode{Val: val}
                return root
            }else{
                node = node.Left
            }
        }
    }
    
    return &TreeNode{Val: val}
    
}



