package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	b := flag.Bool("bool", false, "this is bool flag")
	n := flag.Int("int", 0, "this is int flag")
	os.Args = []string{"progname", "-bool", "-int", "100"}

	flag.Parse()
	fmt.Println(*b)
	fmt.Println(*n)
}
