package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)

	shaPtr := flag.String("sha", "256", "SHA to compute for the standard input")
	flag.Parse()
	fmt.Println("SHA flag specified: ", *shaPtr)

	bytes, err := ioutil.ReadAll(os.Stdin)
	log.Println(err, string(bytes))
	if *shaPtr == "256" {
		fmt.Println(sha256.Sum256(bytes))
	}
	if *shaPtr == "512" {
		fmt.Println(sha512.Sum512(bytes))
	}
	if *shaPtr == "384" {
		fmt.Println(sha512.Sum384(bytes))
	}

}
