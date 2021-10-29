
// 1. Two Sum
// Easy

// 25831

// 838

// Add to List

// Share
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

 

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]
 

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.





// Success
// Details 
// Runtime: 4 ms, faster than 96.60% of Go online submissions for Two Sum.
// Memory Usage: 4.3 MB, less than 42.28% of Go online submissions for Two Sum.
// Next challenges:


func twoSum(nums []int, target int) []int {
    visited := make(map[int]int)
    
    for k, v:= range nums{
        if idx,ok := visited[target-v];ok{
            return []int{idx,k}
        }
        visited[v] = k
    }
    
    return []int {}
}
