//using println default formating, ideal for quick debug. Similar output but retains the square buckets.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
