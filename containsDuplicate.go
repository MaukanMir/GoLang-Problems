// <-- Directions -->
// 217. Contains Duplicate
// Easy

// 2772

// 890

// Add to List

// Share
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
// <-- End Of Directions -->

func containsDuplicate(nums []int) bool {
    visited := map[int]bool{}
    
    for i:= range nums{
        check:= nums[i]
        if visited[check] != true{
            visited[check] = true
        }else{
            return true
        }
    }
    return false
}
