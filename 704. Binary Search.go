Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Success
Details 
Runtime: 28 ms, faster than 96.95% of Go online submissions for Binary Search.
Memory Usage: 6.7 MB, less than 81.92% of Go online submissions for Binary Search.


func search(nums []int, target int) int {
    
    left, right := 0, len(nums)-1
    
    for left<= right{
        mid := left + (right - left)/2
        if nums[mid] == target{return mid
        }else if nums[left] < target{left++
        }else if nums[right] > target{right--}
    }
    return -1
}
