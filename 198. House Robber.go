// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

 

// Example 1:

// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 2:

// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.
 

// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
// Memory Usage: 2 MB, less than 67.52% of Go online submissions for House Robber.



func rob(nums []int) int {
    var rob1 =0
    var rob2 =0
    var current =0
    
    for i:=0 ; i< len(nums); i++{
        current = max(rob1 + nums[i], rob2)
        rob1 = rob2
        rob2 = current
    }
    return rob2
}

func max(a,b int) int{
    if(a > b){return a}
    return b
}
