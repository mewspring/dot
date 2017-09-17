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
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/graphism/simple"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
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
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.WithStack(err)
	}
	g := NewGraph()
	if err := dot.Unmarshal(input, g); err != nil {
		return errors.WithStack(err)
	}

	// Convert back and forth to gonum, thus stripping all non-essential
	// information.
	// TODO: Strip non-essential information.

	// Output graph.
	buf, err := dot.Marshal(g, g.DOTID(), "", "\t", false)
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

// Graph represents a control flow graph.
type Graph struct {
	*simple.DirectedGraph
	// Function name.
	name string
}

// NewGraph returns a new control flow graph.
func NewGraph() *Graph {
	return &Graph{
		DirectedGraph: simple.NewDirectedGraph(),
	}
}

// DOTID returns the DOT ID of the graph.
func (g *Graph) DOTID() string {
	return g.name
}

// SetDOTID sets the DOT ID of the graph.
func (g *Graph) SetDOTID(id string) {
	g.name = id
}

// Node represents a node of a control flow graph.
type Node struct {
	graph.Node
	// Basic block label.
	name string
	// Entry basic block.
	entry bool
}

// NewNode returns a new graph node with a unique arbitrary ID.
func (g *Graph) NewNode() graph.Node {
	return &Node{
		Node: g.DirectedGraph.NewNode(),
	}
}

// DOTID returns the DOT ID of the node.
func (n *Node) DOTID() string {
	return n.name
}

// SetDOTID sets the DOT ID of the node.
func (n *Node) SetDOTID(id string) {
	if id == "0" {
		n.entry = true
	}
	n.name = id
}

// Attributes returns the DOT attributes of the node.
func (n *Node) Attributes() []encoding.Attribute {
	if n.entry {
		return []encoding.Attribute{{Key: "label", Value: "entry"}}
	}
	return nil
}

// Edge represents an edge of a control flow graph.
type Edge struct {
	graph.Edge
	// Edge label.
	Label string
}

// NewEdge returns a new Edge from the source to the destination node.
func (g *Graph) NewEdge(from, to graph.Node) graph.Edge {
	return &Edge{
		Edge: g.DirectedGraph.NewEdge(from, to),
	}
}

// Attributes returns the attributes of the edge.
func (e *Edge) Attributes() []encoding.Attribute {
	if len(e.Label) > 0 {
		val := e.Label
		if !(strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`)) && strings.ContainsAny(val, "\t ") {
			val = strconv.Quote(val)
		}
		var attrs []encoding.Attribute
		attrs = append(attrs, encoding.Attribute{Key: "label", Value: val})
		switch val {
		case "true":
			attrs = append(attrs, encoding.Attribute{Key: "color", Value: "darkgreen"})
		case "false":
			attrs = append(attrs, encoding.Attribute{Key: "color", Value: "red"})
		default:
			// nothing to do.
		}
		return attrs
	}
	return nil
}

// SetAttribute sets the attribute of the edge.
func (e *Edge) SetAttribute(attr encoding.Attribute) error {
	switch attr.Key {
	case "label":
		val := attr.Value
		if strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`) {
			s, err := strconv.Unquote(val)
			if err != nil {
				return errors.WithStack(err)
			}
			val = s
		}
		e.Label = val
	default:
		// ignore attribute.
	}
	return nil
}
