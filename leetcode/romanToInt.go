package main

import "fmt"

func main() {
    three := "III"
    //four := "IV"
    //nine := "IX"
    //fifty_eight := "LVIII"
    long_number:= "MCMXCIV"

    fmt.Printf("%s -> %d\n", three, romanToInt(three))
    fmt.Printf("%s -> %d\n", long_number, romanToInt(long_number))
    fmt.Printf("%s -> %d\n", "XCIX", romanToInt("XCIX"))
}

func romanToInt(s string) int {
    result := 0

    r := []rune(s)

    rn := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i, v := range r {
        if i+1 < len(r) && rn[v] < rn[r[i+1]] {
            result -= rn[v]
        } else {
            result += rn[v]
        }
    }

    return result
}
