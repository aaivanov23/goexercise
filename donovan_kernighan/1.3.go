// 1.3 measure the time of different echo implementations
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    echo1()
    echo2()
    echo3()
}

func echo1() {
    defer timeTracker("echo1")()
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

func echo2() {
    defer timeTracker("echo2")()
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func echo3() {
    defer timeTracker("echo3")()
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func timeTracker(funcName string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s result: %d\n", funcName, time.Since(start).Nanoseconds())
    }

}






