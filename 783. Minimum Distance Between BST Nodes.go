// Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

 

// Example 1:


// Input: root = [4,2,6,1,3]
// Output: 1
// Example 2:


// Input: root = [1,0,48,null,null,12,49]
// Output: 1

// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Distance Between BST Nodes.
// Memory Usage: 2.5 MB, less than 60.00% of Go online submissions for Minimum Distance Between BST Nodes.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    var result []int
    DFS(root,&result)
    sort.Ints(result)
    total :=100000
    fmt.Println(result)
    for i:= 1; i< len(result); i++{
        var got = result[i] - result[i-1]
        total = min( got, total)
    }
                    
    return total                
}

func DFS(root *TreeNode, result*[]int){
    if root != nil{
        *result = append(*result, root.Val)
        DFS(root.Left, result)
        DFS(root.Right, result)
    }
}

func min(a int, b int)int{
    if a< b{
        return a
    }
    
    return b
}
