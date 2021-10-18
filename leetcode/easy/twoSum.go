package main

// numbers should be sorted in non-decreasing order
func twoSum(numbers []int, target int) []int {
    result := make([]int, 0, 2)

    i := 0
    j := len(numbers) - 1

    for i < j {
        if numbers[i] + numbers[j] == target {
            result = append(result, i+1, j+1)
            return result
        } else if numbers[i] + numbers[j] > target {
            j--
        } else {
            i++
        }
    }

    return result
}
