Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

 

Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
 

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length

Success
Details 
Runtime: 48 ms, faster than 100.00% of Go online submissions for Max Consecutive Ones III.
Memory Usage: 7 MB, less than 81.37% of Go online submissions for Max Consecutive Ones III.




func longestOnes(nums []int, k int) int {
    
    windowStart, windowEnd := 0,0
    
    for windowEnd < len(nums){
        if nums[windowEnd] ==0{
            k-=1
        }
        
        if k<0{
            
            if nums[windowStart] ==0{
                k+=1
            }
            windowStart +=1
                
        }
        windowEnd +=1
    }
    
    return windowEnd - windowStart
}
