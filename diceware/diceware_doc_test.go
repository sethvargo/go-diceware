package diceware_test

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-diceware/diceware"
)

func ExampleGenerate() {
	words, err := diceware.Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}

func ExampleMustGenerate() {
	words := diceware.MustGenerate(6)
	log.Printf("%q", words)
}

func ExampleGenerator_Generate() {
	gen, err := diceware.NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	words, err := gen.Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}

func ExampleNewGenerator_nil() {
	// This is the same as calling Generate directly at the package level, but you
	// can share the generator across functions.
	gen, err := diceware.NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}

func ExampleNewGenerator_custom() {
	gen, err := diceware.NewGenerator(&diceware.GeneratorInput{
		WordList: diceware.WordListOriginal(),
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}

func ExampleNewMockGenerator_testing() {
	// Accept a diceware.DicewareGenerator interface instead of a
	// diceware.Generator struct.
	f := func(g diceware.DicewareGenerator) []string {
		// These values don't matter
		return g.MustGenerate(1)
	}

	// In tests
	gen := diceware.NewMockGenerator([]string{"canned", "response"}, nil)

	fmt.Printf("%s", f(gen))
	// Output: [canned response]
}
