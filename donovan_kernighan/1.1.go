// 1.1 prints itself and its command-line arguments.
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args, " "))
}
