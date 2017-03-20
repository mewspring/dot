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
	"log"
	"os"

	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"github.com/gonum/graph/simple"
	dotparser "github.com/graphism/dot"
	"github.com/pkg/errors"
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
	file, err := dotparser.ParseFile(path)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(file.Graphs) != 1 {
		return errors.Errorf("invalid number of graphs in %q; expected 1, got %d", path, len(file.Graphs))
	}

	// Convert back and forth to gonum, thus stripping all non-essential
	// information.
	src := file.Graphs[0]
	dst := newDirectedGraph()
	if dot.Copy(dst, src); err != nil {
		return errors.WithStack(err)
	}
	buf, err := dot.Marshal(dst, src.ID, "", "\t", false)
	if err != nil {
		return errors.WithStack(err)
	}
	buf = append(buf, '\n')

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
			return errors.WithStack(err)
		}
		defer f.Close()
		w = f
	}

	// Write to output stream.
	if _, err := w.Write(buf); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// directedGraph extends simple.DirectedGraph with NewNode and NewEdge methods
// to provide DOT decoding support.
type directedGraph struct {
	*simple.DirectedGraph
}

func newDirectedGraph() *directedGraph {
	return &directedGraph{
		DirectedGraph: simple.NewDirectedGraph(0, 0),
	}
}

// NewNode adds a new node with a unique node ID to the graph.
func (g *directedGraph) NewNode() graph.Node {
	n := simple.Node(g.NewNodeID())
	g.AddNode(n)
	return n
}

// NewEdge adds a new edge from the source to the destination node to the graph,
// or returns the existing edge if already present.
func (g *directedGraph) NewEdge(from, to graph.Node) graph.Edge {
	if e := g.Edge(from, to); e != nil {
		return e
	}
	e := simple.Edge{F: from, T: to}
	g.SetEdge(e)
	return e
}
