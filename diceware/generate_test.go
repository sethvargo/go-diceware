package diceware

import (
	"log"
	"testing"
)

const (
	N = 10000
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	unique := func(list []string) bool {
		seen := make(map[string]struct{}, len(list))
		for _, v := range list {
			if _, ok := seen[v]; ok {
				return false
			}
			seen[v] = struct{}{}
		}
		return true
	}

	for i := 0; i < N; i++ {
		list, err := Generate(16)
		if err != nil {
			t.Fatal(err)
		}
		if !unique(list) {
			t.Errorf("contains duplicate words: %q", list)
		}
	}
}

func TestRollDie(t *testing.T) {
	t.Parallel()

	for i := 0; i < N; i++ {
		r, err := RollDie()
		if err != nil {
			t.Fatal(err)
		}

		if r < 1 || r > 6 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func TestRollWord(t *testing.T) {
	t.Parallel()

	for i := 0; i < N; i++ {
		r, err := RollWord(5)
		if err != nil {
			t.Fatal(err)
		}

		if r < 11111 || r > 66666 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func ExampleGenerate() {
	words, err := Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}
