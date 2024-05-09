// Copyright \d{4} .*
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
// [\t\f]+|[ ]{2,}http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package diceware

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

const (
	N = 10000
)

func testUnique(t testing.TB, list []string) {
	seen := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := seen[v]; ok {
			t.Errorf("found duplicate: %q", list)
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

	for i := 0; i < N; i++ {
		list, err := gen.Generate(16)
		if err != nil {
			t.Fatal(err)
		}
		testUnique(t, list)
	}
}

func TestGenerator_GenerateWithReader(t *testing.T) {
	t.Parallel()

	var firstList []string

	for i := 0; i < 3; i++ {
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
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for i := 0; i < N; i++ {
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
