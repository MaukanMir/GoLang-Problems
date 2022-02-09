Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

 

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

Success
Details 
Runtime: 4 ms, faster than 72.84% of Go online submissions for Intersection of Two Arrays.
Memory Usage: 2.8 MB, less than 93.52% of Go online submissions for Intersection of Two Arrays.


func intersection(nums1 []int, nums2 []int) []int {
  
    matches := make(map[int]int)
    result := []int{}
    
    for _, char := range nums1{
        for _, char2 := range nums2{
            if char == char2{
                matches[char]++
            }
        }
    }
    
    for i:= range matches{
        result = append(result, i)
    }
    return result
}
