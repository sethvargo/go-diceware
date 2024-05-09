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
