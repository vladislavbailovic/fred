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
}

func parseArgs(args []string) options {
	set := flag.NewFlagSet("options", flag.ContinueOnError)

	var list topics
	set.Var(&list, "topic", "Topic to focus on (accumulated by repeating)")
	set.Var(&list, "t", "Topic to focus on (accumulated by repeating)")

	var days int
	set.IntVar(&days, "days", 0, "Last however many days")
	set.IntVar(&days, "d", 0, "Last however many days")

	set.Parse(args)
	return options{topics: list, days: days}
}
