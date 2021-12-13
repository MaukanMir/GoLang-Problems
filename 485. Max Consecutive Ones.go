Given a binary array nums, return the maximum number of consecutive 1's in the array.

 

// Example 1:

// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
// Example 2:

// Input: nums = [1,0,1,1,0,1]
// Output: 2
 

// Constraints:

// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.

Success
Details 
Runtime: 36 ms, faster than 86.88% of Go online submissions for Max Consecutive Ones.
Memory Usage: 6.7 MB, less than 97.51% of Go online submissions for Max Consecutive Ones.


import "math"
func findMaxConsecutiveOnes(nums []int) int {
    var curr,max = float64(0),float64(0)
    
    for i:=0; i< len(nums); i++{
        if(nums[i] ==1){
            curr +=1
        }else if(nums[i] != 1){ 
            max =  math.Max(float64(max),float64(curr))
            curr =0
        }
    }
    return int(math.Max(float64(max),float64(curr)))
}
