// flow is a tool which simplifies Graphviz DOT files, keeping only control flow
// information.
//
// Usage: flow [OPTION]... FILE...
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

	"github.com/gonum/graph/simple"
	"github.com/graphism/dot"
	"github.com/graphism/dot/ast"
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
		if err := flow(path, output, inplace); err != nil {
			log.Fatal(err)
		}
	}
}

// flow simplifies the given Graphviz DOT file, keeping only control flow
// information.
func flow(path, output string, inplace bool) error {
	// Parse input file.
	file, err := dot.ParseFile(path)
	if err != nil {
		return errutil.Err(err)
	}
	if len(file.Graphs) != 1 {
		return errutil.Newf("invalid number of graphs in %q; expected 1, got %d", path, len(file.Graphs))
	}

	// Convert back and forth to gonum, thus stripping all non-essential
	// information.
	src := file.Graphs[0]
	dst := simple.NewDirectedGraph(1, 1)
	dot.CopyDirected(dst, src)
	graph := dot.NewGraph(dst)
	file = &ast.File{Graphs: []*ast.Graph{graph}}

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
