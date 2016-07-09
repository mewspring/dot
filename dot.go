// Package dot provides access to Graphviz DOT files.
package dot

import (
	"fmt"

	"github.com/gonum/graph"
	"github.com/gonum/graph/simple"
	"github.com/graphism/dot/gocc/ast"
	"github.com/graphism/dot/gocc/lexer"
	"github.com/graphism/dot/gocc/parser"
	"github.com/mewkiz/pkg/errutil"
)

// ParseFile parses the given Graphviz DOT file.
func ParseFile(path string) (*ast.File, error) {
	l, err := lexer.NewLexerFile(path)
	if err != nil {
		return nil, errutil.Err(err)
	}
	p := parser.NewParser()
	file, err := p.Parse(l)
	if err != nil {
		return nil, errutil.Err(err)
	}
	f, ok := file.(*ast.File)
	if !ok {
		return nil, errutil.Newf("invalid file type; expected *ast.File, got %T", file)
	}
	return f, nil
}

// CopyDirected copies the nodes and edges from the AST of the directed source
// graph to the destination graph.
func CopyDirected(dst graph.DirectedBuilder, src *ast.Graph) {
	if !src.Directed {
		panic("invalid source graph; undirected")
	}
	gen := &generator{
		ids: make(map[string]graph.Node),
	}
	for _, stmt := range src.Stmts {
		switch stmt := stmt.(type) {
		case *ast.NodeStmt:
			panic("not yet implemented")
		case *ast.EdgeStmt:
			gen.addEdgeStmt(dst, stmt)
		case *ast.AttrStmt:
			panic("not yet implemented")
		case *ast.Attr:
			panic("not yet implemented")
		case *ast.Subgraph:
			panic("not yet implemented")
		default:
			panic(fmt.Sprintf("unknown statement type %T", stmt))
		}
	}
}

// A generator keeps track of the information required for generating a gonum
// graph from a dot AST graph.
type generator struct {
	// ids maps from dot AST node ID to gonum node.
	ids map[string]graph.Node
}

// node returns the gonum node corresponding to the given dot AST node ID,
// generating a new such node if none exist.
func (gen *generator) node(dst graph.DirectedBuilder, id string) graph.Node {
	if n, ok := gen.ids[id]; ok {
		return n
	}
	n := simple.Node(dst.NewNodeID())
	gen.ids[id] = n
	dst.AddNode(n)
	return n
}

// edge returns the gonum edge from the source to the destination node,
// generating a new such edge if none exist.
func (gen *generator) edge(dst graph.DirectedBuilder, from, to graph.Node) graph.Edge {
	if e := dst.Edge(from, to); e != nil {
		return e
	}
	// TODO: Figure out if weight is needed.
	e := simple.Edge{F: from, T: to, W: 1}
	dst.SetEdge(e)
	return e
}

// addEdgeStmt adds the given edge statement to the graph.
func (gen *generator) addEdgeStmt(dst graph.DirectedBuilder, e *ast.EdgeStmt) {
	fs := gen.addVertex(dst, e.From)
	ts := gen.addEdge(dst, e.To)
	for _, f := range fs {
		for _, t := range ts {
			gen.edge(dst, f, t)
		}
	}
}

// addVertex adds the given vertex to the graph, and returns its set of nodes.
func (gen *generator) addVertex(dst graph.DirectedBuilder, v ast.Vertex) []graph.Node {
	switch v := v.(type) {
	case *ast.NodeID:
		n := gen.node(dst, v.ID)
		return []graph.Node{n}
	case *ast.Subgraph:
		panic("not yet implemented")
	default:
		panic(fmt.Sprintf("unknown vertex type %T", v))
	}
}

// addEdge adds the given edge to the graph, and returns its set of nodes.
func (gen *generator) addEdge(dst graph.DirectedBuilder, to *ast.Edge) []graph.Node {
	// TODO: Handle to.Directed.
	fs := gen.addVertex(dst, to.Vertex)
	if to.To != nil {
		ts := gen.addEdge(dst, to.To)
		for _, f := range fs {
			for _, t := range ts {
				gen.edge(dst, f, t)
			}
		}
	}
	return fs
}
