package main

import (
	"flag"
	"fmt"
)

func main() {
	env := flag.String("env", "", "")
	name := flag.String("name", "", "")
	flag.Parse()
	fmt.Printf("Running file 'gen.go' with environment: \"%s\", written by \"%s\"\n", *env, *name)
}
