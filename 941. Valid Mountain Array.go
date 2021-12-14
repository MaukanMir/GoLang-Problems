// Given an array of integers arr, return true if and only if it is a valid mountain array.

// Recall that arr is a mountain array if and only if:

// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

 

// Example 1:

// Input: arr = [2,1]
// Output: false
// Example 2:

// Input: arr = [3,5,5]
// Output: false
// Example 3:

// Input: arr = [0,3,2,1]
// Output: true
 

// Constraints:

// 1 <= arr.length <= 104
// 0 <= arr[i] <= 104

// Success
// Details 
// Runtime: 24 ms, faster than 82.01% of Go online submissions for Valid Mountain Array.
// Memory Usage: 6.7 MB, less than 100.00% of Go online submissions for Valid Mountain Array.



func validMountainArray(arr []int) bool {
    var start =0
    var size = len(arr)
    
    for( start +1 < size && arr[start] < arr[start+1]){
        start++
    }
    
    if start == 0 || start == size-1{return false}
    
    for start +1 < size && arr[start] > arr[start+1]{
        start ++
    }
    
    return start == size-1
}
