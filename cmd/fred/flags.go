package main

import (
	"flag"
	"fmt"
)

type topics []string

func (x *topics) String() string {
	return fmt.Sprint(*x)
}

func (x *topics) Set(value string) error {
	*x = append(*x, value)
	return nil
}

type options struct {
	topics topics
	days   int
	read   bool
	fname  string
}

func parseArgs(args []string) options {
	set := flag.NewFlagSet("options", flag.ExitOnError)

	var list topics
	set.Var(&list, "topic", "Topic to focus on (accumulated by repeating)")
	set.Var(&list, "t", "Topic to focus on (accumulated by repeating)")

	var days int
	set.IntVar(&days, "days", 0, "Last however many days")
	set.IntVar(&days, "d", 0, "Last however many days")

	var read bool
	set.BoolVar(&read, "read", false, "Also open up in vim for consuming")
	set.BoolVar(&read, "r", false, "Also open up in vim for consuming")

	var fname string
	set.StringVar(&fname, "out", "", "Output to this file instead")
	set.StringVar(&fname, "o", "", "Output to this file instead")

	set.Parse(args)
	return options{topics: list, days: days, read: read, fname: fname}
}
