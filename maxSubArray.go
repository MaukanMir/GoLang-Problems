// Directions:

// 53. Maximum Subarray
// Easy

// 15551

// 728

// Add to List

// Share
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// A subarray is a contiguous part of an array.

 

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23





// Metrics 

// Success
// Details 
// Runtime: 96 ms, faster than 62.92% of Go online submissions for Maximum Subarray.
// Memory Usage: 9.5 MB, less than 22.86% of Go online submissions for Maximum Subarray.




// Sovled


import "math"
func maxSubArray(nums []int) int {
    current, max_amount := float64(nums[0]),float64(nums[0])
    
    for i:= 1;i<len(nums);i++{
        
        current = math.Max(float64(nums[i]), float64(current) + float64(nums[i]))
        max_amount = math.Max(current,max_amount)
    }
    
    return int(max_amount)
}
