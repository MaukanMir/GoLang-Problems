// 169. Majority Element
// Easy

// 6962

// 295

// Add to List

// Share
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

Success
Details 
Runtime: 16 ms, faster than 93.66% of Go online submissions for Majority Element.
Memory Usage: 6.1 MB, less than 66.67% of Go online submissions for Majority Element.




func majorityElement(nums []int) int {
    
    options :=0
    count := 0
    
    for i:= range nums{
        if count ==0{
            options = nums[i]
        }
        
        if (options == nums[i]){
            count +=1
        }else{
            count -=1
        }
    }
    return options
}
