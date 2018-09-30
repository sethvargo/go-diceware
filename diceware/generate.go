// Package diceware provides a library for generating random words via the
// diceware algorithm by rolling five six-sided dice to randomly select a word
// from a list of english words.
//
// Read more about the diceware algorithm here: https://en.wikipedia.org/wiki/Diceware.
//
//    list, err := diceware.Generate(6)
//    if err != nil  {
//      log.Fatal(err)
//    }
//    log.Printf(strings.Join(list, "-"))
//
package diceware

import (
	"crypto/rand"
	"math"
	"math/big"
)

// sides is the number of sides on a die
const sides = 6

// Generate generates a list of the given number of words from EFF's large
// wordlist.
func Generate(numWords int) ([]string, error) {
	return GenerateFromWordlist(numWords, WordlistEFFBig)
}

// Generate generates a list of the given number of words from the given word
// list.
func GenerateFromWordlist(numWords int, wordlist WordListT) ([]string, error) {
	list := make([]string, 0, numWords)
	seen := make(map[string]struct{}, numWords)

	for i := 0; i < numWords; i++ {
		n, err := RollWord(wordlist.digits)
		if err != nil {
			return nil, err
		}

		word := WordAtWordlist(n, wordlist)
		if _, ok := seen[word]; ok {
			i--
			continue
		}

		list = append(list, word)
		seen[word] = struct{}{}
	}

	return list, nil
}

// MustGenerate behaves like Generate, but panics on error.
func MustGenerate(numWords int) []string {
	res, err := Generate(numWords)
	if err != nil {
		panic(err)
	}
	return res
}

// WordAt retrieves the word at the given index from EFF's large wordlist.
func WordAt(i int) string {
	return WordAtWordlist(i, WordlistEFFBig)
}

// WordAt retrieves the word at the given index from the given wordlist.
func WordAtWordlist(i int, wordlist WordListT) string {
	return wordlist.words[i]
}

// RollDie rolls a single 6-sided die and returns a value between [1,6].
func RollDie() (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(sides))
	if err != nil {
		return 0, err
	}
	return int(r.Int64()) + 1, nil
}

// RollWord rolls and aggregates dice to represent one word in the list. The
// result is the index of the word in the list.
func RollWord(d int) (int, error) {
	var final int

	for i := d; i > 0; i-- {
		res, err := RollDie()
		if err != nil {
			return 0, err
		}

		final += res * int(math.Pow(10, float64(i-1)))
	}

	return final, nil
}
