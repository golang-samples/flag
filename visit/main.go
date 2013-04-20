package main

import (
	"flag"
	"fmt"
)

func main() {
	f := flag.NewFlagSet("flag", flag.ExitOnError)
	f.Bool("bool", false, "this is bool flag")
	f.Int("int", 0, "this is int flag")

	visitor := func(a *flag.Flag) {
		fmt.Println(">", a.Name, "value=", a.Value)
	}

	fmt.Println("Visit()")
	f.Visit(visitor)
	fmt.Println("VisitAll()")
	f.VisitAll(visitor)

	// set flags
	f.Parse([]string{"-bool", "-int", "100"})

	fmt.Println("Visit() after Parse()")
	f.Visit(visitor)
	fmt.Println("VisitAll() after Parse()")
	f.VisitAll(visitor)
}
