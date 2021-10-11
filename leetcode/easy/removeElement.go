package main

import "fmt"

func main() {
    fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
    fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}

func removeElement(nums []int, val int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            if i+1 < len(nums) {
                nums = append(nums[:i], nums[i+1:]...)
                i--
            } else {
                nums = nums[:len(nums)-1]
            }
        }
	//fmt.Printf("index: %d, nums: %v\n", i, nums)
    }

    return len(nums)
}

