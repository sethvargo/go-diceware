package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sethvargo/go-diceware/diceware"
)

var (
	flagWords = flag.Int("words", 6,
		"number of words to generate")
	flagSeparator = flag.String("separator", "-",
		"character to use between words")

	stdout, stderr = os.Stdout, os.Stderr
)

func main() {
	flag.Parse()

	list, err := diceware.Generate(*flagWords)
	if err != nil {
		fmt.Fprintf(stderr, "error: %s\n", err)
		os.Exit(2)
	}

	for i, w := range list {
		fmt.Fprint(stdout, w)
		if i < len(list)-1 {
			fmt.Fprint(stdout, *flagSeparator)
		}
	}

	if fi, _ := stdout.Stat(); fi == nil || (fi.Mode()&os.ModeCharDevice) != 0 {
		fmt.Fprintln(stdout)
	}
}
