package main

//import "fmt"
import "math"

/*func main() {
    fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
    fmt.Println(maxProfit([]int{7,6,4,3,1}))
    fmt.Println(maxProfit([]int{2,1,4,5,2,9,7}))
}*/

func maxProfit(prices []int) int {
    minPrice := math.MaxInt64
    maxProfit := 0

    for i := range prices {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if prices[i] - minPrice > maxProfit {
            maxProfit = prices[i] - minPrice
        }
    }
    return maxProfit
}
