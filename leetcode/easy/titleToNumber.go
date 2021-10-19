package main

// 171. Excel Sheet Column Number
/*import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(titleToNumber("ZAA"))
}*/

func titleToNumber(columnTitle string) int {
    runes := []rune(columnTitle)
    reverseSlice(runes)
    columnNumber := 0

    for i, v := range runes {
        columnNumber += int(math.Pow(26, float64(i))) * (int(v) - 64)
    }

    return columnNumber
}

func reverseSlice(runes []rune) {
    for i := 0; i < len(runes) / 2; i++ {
        runes[i], runes[len(runes) - 1 - i] = runes[len(runes) - 1 - i], runes[i]
    }
}
