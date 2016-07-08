// dotfmt is a tool which formats Graphviz DOT files.
//
// Usage: dotfmt [OPTION]... FILE...
//
//   -i    edit file in place
//   -o string
//         output path
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/graphism/dot"
	"github.com/mewkiz/pkg/errutil"
)

func main() {
	// Parse command line flags.
	var (
		// inplace specifies whether to edit file in place.
		inplace bool
		// output specifies the output path.
		output string
	)
	flag.BoolVar(&inplace, "i", false, "edit file in place")
	flag.StringVar(&output, "o", "", "output path")
	flag.Parse()
	if inplace && len(output) > 0 {
		log.Fatal("invalid combination of -i and -o flags; only one may be set")
	}

	// Format input files.
	for _, path := range flag.Args() {
		if err := dotfmt(path, output, inplace); err != nil {
			log.Fatal(err)
		}
	}
}

// dotfmt formats the given Graphviz DOT file.
func dotfmt(path, output string, inplace bool) error {
	// Parse input file.
	file, err := dot.ParseFile(path)
	if err != nil {
		return errutil.Err(err)
	}

	// Write to standard output.
	w := os.Stdout

	// Edit file in place.
	if inplace {
		output = path
	}

	// Write to output file.
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			return errutil.Err(err)
		}
		defer f.Close()
		w = f
	}

	// Write to output stream.
	if _, err := fmt.Fprintln(w, file); err != nil {
		return errutil.Err(err)
	}

	return nil
}
