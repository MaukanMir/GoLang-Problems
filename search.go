// 704. Binary Search
// Easy

// 2528

// 79

// Add to List

// Share
// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

 

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
 

Success
Details 
Runtime: 28 ms, faster than 98.82% of Go online submissions for Binary Search.
Memory Usage: 6.9 MB, less than 97.07% of Go online submissions for Binary Search.


func search(nums []int, target int) int {
    left,right := (0), (len(nums)-1)
    
    for left <= right{
        pivot := (left + (right -left)/2)
        if nums[pivot] == target{
            return pivot
        }else if nums[pivot] < target{
            left = pivot +1
            
        }else{
            right = pivot -1
        }
    }
    
    return -1
}
