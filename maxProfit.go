// 121. Best Time to Buy and Sell Stock
// Easy

// 11696

// 442

// Add to List

// Share
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.




// Success
// Details 
// Runtime: 120 ms, faster than 86.80% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.3 MB, less than 97.77% of Go online submissions for Best Time to Buy and Sell Stock.



func maxProfit(prices []int) int {
    min_price := 100000000000
    var max_price int
    
    for i:=0;i<len(prices);i++{
        if prices[i] < min_price{
            min_price = prices[i]
        }else if max_price < prices[i] - min_price{
            max_price = prices[i] - min_price
        }
    }
    
    return max_price
}

// Success
// Details 
// Runtime: 112 ms, faster than 99.56% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.2 MB, less than 96.10% of Go online submissions for Best Time to Buy and Sell Stock.

func maxProfit(prices []int) int {
    
    var l,r =(0) ,(1)
    var maxProfit =float64(0)
    
    for r < (len(prices)) {
        if  prices[l] < prices[r]{
            var profit = float64(prices[r] - prices[l])
            maxProfit = math.Max(profit,maxProfit)
        }else{
            l = r
        }
        r+=1
    }
    
    return int(maxProfit)
}
