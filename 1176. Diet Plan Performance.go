

Success
Details 
Runtime: 20 ms, faster than 100.00% of Go online submissions for Diet Plan Performance.
Memory Usage: 8 MB, less than 75.00% of Go online submissions for Diet Plan Performance.



func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
    
    var right, left = 0,0
    var score, total = 0,0
    
    for left < len(calories) && right < len(calories){
        total += calories[right]
        
        if right - left >= k{
            
        total -= calories[left]
        left +=1
        }
        
        if right - left == k-1{
            if total < lower{
                score -=1
            }
            if total> upper{
                score +=1
            }
        }
        
        right ++
    }
    
    
    return score
    
}
