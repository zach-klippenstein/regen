/*
Copyright 2014 Zachary Klippenstein

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/zach-klippenstein/goregen"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

var NumberOfGenerations = flag.Int("n", 1, "number of strings to generate. A value of 0 will keep generating strings until the process is killed.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [args] [pattern]\n", path.Base(os.Args[0]))
		fmt.Fprintln(os.Stderr, "Generates random strings from regular expressions.")
		fmt.Fprintln(os.Stderr, "If no pattern is given on the command line, it will be read from stdin. If the last character is \\n, it will be removed.")
		fmt.Fprintln(os.Stderr, "\nargs:")
		flag.PrintDefaults()
	}
	flag.Parse()

	var pattern string

	if flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, "error: too many arguments")
		flag.Usage()
	} else if flag.NArg() == 1 {
		pattern = flag.Arg(0)
	} else {
		pattern = ReadPatternFromStdin()
	}

	rand.Seed(time.Now().UTC().UnixNano())

	generator, err := regen.NewGenerator(pattern, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for i := 0; *NumberOfGenerations == 0 || i < *NumberOfGenerations; i++ {
		fmt.Println(generator.Generate())
	}
}

func ReadPatternFromStdin() string {
	patternBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading from stdin: %s\n", err)
		os.Exit(1)
	}
	pattern := string(patternBytes)

	// if entered interactively, there will be an extra \n at the end that
	// is probably not intended to be part of the regex.
	pattern = strings.TrimSuffix(pattern, "\n")

	return pattern
}
