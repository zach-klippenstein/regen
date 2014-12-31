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
	"math/rand"
	"os"
	"path"
	"time"
)

var NumberOfGenerations = flag.Int("n", 1, "number of strings to generate.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [args] pattern\n", path.Base(os.Args[0]))
		fmt.Fprintln(os.Stderr, "Generates random strings from regular expressions.")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: must specify a pattern")
		flag.Usage()
	}

	rand.Seed(time.Now().UTC().UnixNano())

	pattern := flag.Arg(0)

	generator, err := regen.NewGenerator(pattern, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for i := 0; i < *NumberOfGenerations; i++ {
		fmt.Println(generator.Generate())
	}
}
