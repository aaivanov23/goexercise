package main

import (
    _"fmt"
    "strings"
)

/*func main() {
    strs := []string{"flower","flow","flight"}
    fmt.Println(longestCommonPrefix(strs))
}*/

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0]
    for i, v := range strs {
        if i == 0 {
            continue
        }

        for !strings.HasPrefix(v, prefix) {
            prefix = prefix[:len([]rune(prefix))-1]
            if prefix == "" {
                return prefix
            }
        }
    }
    return prefix
}
