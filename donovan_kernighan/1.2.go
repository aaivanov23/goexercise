// 1.2 prints the index and value of each of its arguments, one per line.
package main

import (
    "fmt"
    "os"
)

func main() {
    for i, v := range os.Args[1:] {
        fmt.Printf("%d %s\n", i+1, v)
    }
}
