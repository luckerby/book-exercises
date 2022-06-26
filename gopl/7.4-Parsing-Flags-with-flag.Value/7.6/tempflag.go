package main

import (
	"flag"
	"fmt"
	"github.com/luckerby/book-exercises/gopl/7.4-Parsing-Flags-with-flag.Value/7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	fmt.Println("In main()")
	flag.Parse()
	fmt.Println(*temp)
}
