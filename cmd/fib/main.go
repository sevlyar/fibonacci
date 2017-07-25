// fib is an utility to calculate Fibonacci numbers. It prints Fibonacci number
// with the given index in decimal representation to stdout.
//
// Usage:
//
//		fib <n>
//			n - index of Fibonacci number.
//
package main

import (
	"flag"
	"fmt"
	"github.com/sevlyar/go-project-sample"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	arg := flag.Arg(0)
	num, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		log.Fatal("Invalid argument, number expected")
	}

	res := fibonacci.Calc(num)
	fmt.Println(res)
}
