// Print the index and value of each
// argument, one per line

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var printstring string
	for index, value := range os.Args[:] {
		printstring = strconv.Itoa(index) + " " + value
		fmt.Println(printstring)
	}
}