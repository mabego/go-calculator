package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	if s == "exit" || s == "quit" {
		os.Exit(0)
	}

	calc := New()

	val, err := calc.Calculate(s)
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Command mode usage: %s \"expression\"\n", os.Args[0])
		fmt.Fprintf(w, "Example: %s \"(2.5 - 1.35) * 2.0\"\n", os.Args[0])
		fmt.Fprintf(w, "REPL usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	switch {
	case len(os.Args) == 1:
		p := prompt.New(
			executor,
			func(_ prompt.Document) []prompt.Suggest { return []prompt.Suggest{} },
			prompt.OptionPrefix("calculator> "),
		)
		p.Run()
	case os.Args[1] == "-h" || os.Args[1] == "--help":
		// Run flag.Parse only for help flags to allow command mode expressions to begin with a negative `-`.
		flag.Parse()
	default:
		calc := New()
		expression := os.Args[1]
		result, err := calc.Calculate(expression)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	}
}
