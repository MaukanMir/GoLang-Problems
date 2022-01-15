

Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for N-ary Tree Preorder Traversal.
Memory Usage: 4.1 MB, less than 53.66% of Go online submissions for N-ary Tree Preorder Traversal.


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    
   var res []int
    var dfs func(*Node)
    dfs = func(root *Node){
        if root == nil{
            return
        }
        
        res = append(res,root.Val)
        for _ ,i:= range root.Children{
            dfs(i)
        }
    }
    
    dfs(root)
    
    return res
    
