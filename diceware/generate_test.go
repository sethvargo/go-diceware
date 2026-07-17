package diceware

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

const (
	N = 10000
)

func testUnique(tb testing.TB, list []string) {
	tb.Helper()

	seen := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := seen[v]; ok {
			tb.Errorf("found duplicate: %q", list)
		}
		seen[v] = struct{}{}
	}
}

func TestGenerator_Generate(t *testing.T) {
	t.Parallel()

	gen, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	for range N {
		list, err := gen.Generate(16)
		if err != nil {
			t.Fatal(err)
		}
		testUnique(t, list)
	}
}

func TestGenerator_GenerateNegative(t *testing.T) {
	t.Parallel()

	gen, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := gen.Generate(-1); !errors.Is(err, ErrNumWordsNegative) {
		t.Errorf("expected %q to be %q", err, ErrNumWordsNegative)
	}
}

func TestGenerator_GenerateWithReader(t *testing.T) {
	t.Parallel()

	var firstList []string

	for i := range 3 {
		gen, err := NewGenerator(&GeneratorInput{RandReader: bytes.NewBufferString(strings.Repeat("foopityboopityflippityfloppity", 16))})
		if err != nil {
			t.Fatal(err)
		}
		list, err := gen.Generate(16)
		if err != nil {
			t.Fatal(err)
		}
		if i == 0 {
			firstList = list
		} else if !reflect.DeepEqual(list, firstList) {
			t.Fatalf("mismatched values from custom rand: %v vs %v", firstList, list)
		}
	}
}

func TestGenerateWordList(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		wordList WordList
	}{
		{
			"eff_large",
			WordListEffLarge(),
		},
		{
			"eff_small",
			WordListEffSmall(),
		},
		{
			"original",
			WordListOriginal(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for range N {
				list, err := GenerateWithWordList(16, tc.wordList)
				if err != nil {
					t.Fatal(err)
				}
				testUnique(t, list)
			}
		})
	}
}

func TestRollDie(t *testing.T) {
	t.Parallel()

	for range N {
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

	for range N {
		r, err := RollWord(5)
		if err != nil {
			t.Fatal(err)
		}

		if r < 11111 || r > 66666 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}
