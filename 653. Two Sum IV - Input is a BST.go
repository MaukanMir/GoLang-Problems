// Given the root of a Binary Search Tree and a target number k, return true if there exist two elements in the BST such that their sum is equal to the given target.

 

// Example 1:


// Input: root = [5,3,6,2,4,null,7], k = 9
// Output: true
// Example 2:


// Input: root = [5,3,6,2,4,null,7], k = 28
// Output: false
// Example 3:

// Input: root = [2,1,3], k = 4
// Output: true
// Example 4:

// Input: root = [2,1,3], k = 1
// Output: false
// Example 5:

// Input: root = [2,1,3], k = 3
// Output: true
 

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// -104 <= Node.val <= 104
// root is guaranteed to be a valid binary search tree.
// -105 <= k <= 105

Success
Details 
Runtime: 20 ms, faster than 95.97% of Go online submissions for Two Sum IV - Input is a BST.
Memory Usage: 7.9 MB, less than 32.21% of Go online submissions for Two Sum IV - Input is a BST.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    var list = make(map[int]bool)
    return dfs(root,k,list)

}

func dfs(root *TreeNode, k int, list map[int]bool)bool{
        if root == nil{return false}
        
        if _,check := list[root.Val];check{return true}
        
        list[k- root.Val]= true
        
        return dfs(root.Left,k, list) || dfs(root.Right,k, list)
    }


