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

// WordList is an interface that must be implemented to be considered a word
// list for use in the diceware algorithm. This interface can be implemented by
// other libraries.
type WordList interface {
	// Digits is the number of digits for indexes in the word list. This
	// determines the number of dice rolls.
	Digits() int

	// WordAt returns the word at the given integer in the word list.
	WordAt(int) string
}

var _ WordList = (*wordListInternal)(nil)

type wordListInternal struct {
	digits int
	words  map[int]string
}

func (w *wordListInternal) Digits() int {
	return w.digits
}

func (w *wordListInternal) WordAt(i int) string {
	return w.words[i]
}
