//using println default formating, ideal for quick debug. Similar output but retains the square buckets.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:], " ")
}
