// 977. Squares of a Sorted Array
// Easy

// 3740

// 128

// Add to List

// Share
// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

 

// Example 1:

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].
// Example 2:

// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]
 

// Constraints:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.

Success
Details 
Runtime: 36 ms, faster than 54.67% of Go online submissions for Squares of a Sorted Array.
Memory Usage: 6.9 MB, less than 73.04% of Go online submissions for Squares of a Sorted Array.


import "sort"
func sortedSquares(nums []int) []int {
    for i:= range nums{
        nums[i] = nums[i]*nums[i]
    }
  sort.Ints(nums)
    return nums
}
