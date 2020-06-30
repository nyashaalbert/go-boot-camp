//using a join for the echo statement as it would improve on memory usage as compared to storing the variable repeatedly.package Tutorials

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
