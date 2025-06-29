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
	if s == "clear" {
		fmt.Print("\033[2J\033[H") // ANSI sequences to clear screen and move cursor to home position
		return
	}
	if s == "help" {
		PrintHelp()
		return
	}

	calc := New()

	val, err := calc.Calculate(s)
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// completer provides autocomplete suggestions
func completer(d prompt.Document) []prompt.Suggest {
	var suggestions []prompt.Suggest

	word := d.GetWordBeforeCursor()
	if word == "" {
		return suggestions
	}

	commands := []prompt.Suggest{
		{Text: "help", Description: "Show help"},
		{Text: "exit", Description: "Exit the calculator"},
		{Text: "quit", Description: "Exit the calculator"},
		{Text: "clear", Description: "Clear the screen"},
	}

	mathFunctions := [...]string{"abs", "acos", "acosh", "asin", "asinh", "atan", "atan2", "atanh", "cbrt", "ceil",
		"copysign", "cos", "cosh", "dim", "erf", "erfc", "erfcinv", "erfinv", "exp", "exp2", "expm1", "fma", "floor",
		"gamma", "hypot", "j0", "j1", "log", "log10", "log1p", "log2", "logb", "max", "min", "mod", "nan", "nextafter",
		"pow", "remainder", "round", "roundtoeven", "sin", "sinh", "sqrt", "tan", "tanh", "trunc", "y0", "y1"}

	mathConstants := [...]string{"e", "pi", "phi", "sqrt2", "sqrte", "sqrtpi", "sqrtphi", "ln2", "log2e", "ln10",
		"log10e"}

	for _, fn := range mathFunctions {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        fn + "(",
			Description: "Function",
		})
	}

	for _, c := range mathConstants {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        c,
			Description: "Constant",
		})
	}

	suggestions = append(suggestions, commands...)

	return prompt.FilterHasPrefix(suggestions, word, true)
}

// PrintHelp displays help information
func PrintHelp() {
	fmt.Println("\nCommands:")
	fmt.Println()
	fmt.Println("help	- Show this help")
	fmt.Println("clear	- Clear the screen")
	fmt.Println("exit	- Exit the calculator")
	fmt.Println()
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Println("Usage:")
		fmt.Println("  Command mode:")
		fmt.Fprintf(w, "\t%s \"math expression\"\n", os.Args[0])
		fmt.Println()
		fmt.Println("  REPL usage:")
		fmt.Fprintf(w, "\t%s\n", os.Args[0])
		flag.PrintDefaults()
	}

	switch {
	case len(os.Args) == 1:
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix("calc> "),
			prompt.OptionMaxSuggestion(6),
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
