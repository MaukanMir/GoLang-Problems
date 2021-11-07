// 283. Move Zeroes
// Easy

// 7088

// 197

// Add to List

// Share
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

 

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]
 

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1
 

// Follow up: Could you minimize the total number of operations done?

Success
Details 
Runtime: 20 ms, faster than 74.33% of Go online submissions for Move Zeroes.
Memory Usage: 6.8 MB, less than 69.67% of Go online submissions for Move Zeroes


func moveZeroes(nums []int)  {
    
    var index int
    
    for i:= 0;i<len(nums);i++{
        if nums[i] !=0{
            nums[index] = nums[i]
            index+=1
        }
    }
    
   for i :=index;i<len(nums);i++{
        nums[i] =0
    }
}
