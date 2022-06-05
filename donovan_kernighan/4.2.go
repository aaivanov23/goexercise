// usage: go run 4.2.go [-sha512|-sha384] hello

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s384 = flag.Bool("sha384", false, "sha384 hash")
var s512 = flag.Bool("sha512", false, "sha512 hash")

func main() {
	flag.Parse()
	data := os.Args[1]

	switch {
	case *s384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(data)))
	case *s512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(data)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(data)))
	}
}
