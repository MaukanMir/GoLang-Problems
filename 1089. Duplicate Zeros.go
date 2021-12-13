// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

 

// Example 1:

// Input: arr = [1,0,2,3,0,4,5,0]
// Output: [1,0,0,2,3,0,0,4]
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
// Example 2:

// Input: arr = [1,2,3]
// Output: [1,2,3]
// Explanation: After calling your function, the input array is modified to: [1,2,3]

// Success
// Details 
// Runtime: 4 ms, faster than 94.02% of Go online submissions for Duplicate Zeros.
// Memory Usage: 3.3 MB, less than 100.00% of Go online submissions for Duplicate Zeros.


func duplicateZeros(arr []int)  {
    var size = len(arr)
    var total =0
    for i:=0; i< len(arr); i++{
        if arr[i]==0{total +=1}
    }
    
    for i:= size-1; i>=0; i--{
        if i + total < size{arr[i +total] = arr[i]}
        if arr[i]==0{
            total -=1
            if i +total <size{arr[i+total] =0}
        }
    }
    return
}
