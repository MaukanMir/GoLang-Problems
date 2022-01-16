// Given a n-ary tree, find its maximum depth.

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

 

// Example 1:



// Input: root = [1,null,3,2,4,null,5,6]
// Output: 3
// Example 2:



// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: 5
 

// Constraints:

// The total number of nodes is in the range [0, 104].
// The depth of the n-ary tree is less than or equal to 1000.

// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Depth of N-ary Tree.
// Memory Usage: 3.3 MB, less than 60.71% of Go online submissions for Maximum Depth of N-ary Tree.


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil{return 0}
    
    ans:= 0
    
    for _, child := range root.Children{
        ans = max(ans, maxDepth(child))
    }
    return ans +1
}

func max(a int, b int)int{
    if a>b{
        return a
    }
    return b
}
