// Program to print run times of
// different implementation of the
// echo program

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time Taken by parsing over array = %f seconds\n", time.Since(start).Seconds())

	start2 := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	fmt.Printf("Time Taken by optimizing parsing over array = %f seconds\n", time.Since(start2).Seconds())

	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Time Taken using strings.Join() method = %f seconds\n", time.Since(start3).Seconds())
}

// Sample Output
// $ ./ex13 hello testing the run times of different echo implementations
// hello testing the run times of different echo implementations
// Time Taken by parsing over array = 0.000638 seconds
// hello testing the run times of different echo implementations
// Time Taken by optimizing parsing over array = 0.000458 seconds
// hello testing the run times of different echo implementations
// Time Taken using strings.Join() method = 0.000440 seconds
