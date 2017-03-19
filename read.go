package dot

import (
	"fmt"
	"log"

	"github.com/gonum/graph"
	"github.com/graphism/dot/ast"
)

// DirectedBuilder is a directed graph builder capable of handling user-defined
// node and edge types.
type DirectedBuilder interface {
	graph.DirectedBuilder
	// NewNode returns a new node with a unique node ID in the graph.
	NewNode() graph.Node
	// NewEdge returns a new edge from the source to the destination node in the
	// graph, or the existing edge if already present.
	NewEdge(from, to graph.Node) graph.Edge
}

// UnmarshalerAttr is the interface implemented by objects that can unmarshal a
// DOT attribute description of themselves.
type UnmarshalerAttr interface {
	// UnmarshalDOTAttr decodes a single DOT attribute.
	UnmarshalDOTAttr(attr *ast.Attr) error
}

// CopyDirected copies the nodes and edges from the AST of the directed source
// graph to the destination graph.
func CopyDirected(dst DirectedBuilder, src *ast.Graph) {
	if !src.Directed {
		panic("invalid source graph; undirected")
	}
	gen := &generator{
		ids: make(map[string]graph.Node),
	}
	for _, stmt := range src.Stmts {
		gen.addStmt(dst, stmt)
	}
}

// A generator keeps track of the information required for generating a gonum
// graph from a dot AST graph.
type generator struct {
	// ids maps from dot AST node ID to gonum node.
	ids map[string]graph.Node
	// Nodes processed within the context of a subgraph, that is to be used as a
	// vertex of an edge.
	subNodes []graph.Node
	// Stack of start indices into the subgraph node slice. The top element
	// corresponds to the start index of the active (or inner-most) subgraph.
	subStart []int
}

// node returns the gonum node corresponding to the given dot AST node ID,
// generating a new such node if none exist.
func (gen *generator) node(dst DirectedBuilder, id string) graph.Node {
	if n, ok := gen.ids[id]; ok {
		return n
	}
	n := dst.NewNode()
	gen.ids[id] = n
	// Check if within the context of a subgraph, that is to be used as a vertex
	// of an edge.
	if gen.isInSubgraph() {
		// Append node processed within the context of a subgraph, that is to be
		// used as a vertex of an edge
		gen.appendSubgraphNode(n)
	}
	return n
}

// edge returns the gonum edge from the source to the destination node,
// generating a new such edge if none exist.
func (gen *generator) edge(dst DirectedBuilder, from, to graph.Node) graph.Edge {
	if e := dst.Edge(from, to); e != nil {
		return e
	}
	e := dst.NewEdge(from, to)
	dst.SetEdge(e)
	return e
}

// addStmt adds the given statement to the graph.
func (gen *generator) addStmt(dst DirectedBuilder, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.NodeStmt:
		n := gen.node(dst, stmt.Node.ID)
		if n, ok := n.(UnmarshalerAttr); ok {
			for _, attr := range stmt.Attrs {
				if err := n.UnmarshalDOTAttr(attr); err != nil {
					log.Printf("unable to unmarshal DOT attribute %#v; %v", attr, err)
				}
			}
		}
	case *ast.EdgeStmt:
		gen.addEdgeStmt(dst, stmt)
	case *ast.AttrStmt:
		// ignore.
	case *ast.Attr:
		// ignore.
	case *ast.Subgraph:
		for _, stmt := range stmt.Stmts {
			gen.addStmt(dst, stmt)
		}
	default:
		panic(fmt.Sprintf("unknown statement type %T", stmt))
	}
}

// addEdgeStmt adds the given edge statement to the graph.
func (gen *generator) addEdgeStmt(dst DirectedBuilder, e *ast.EdgeStmt) {
	fs := gen.addVertex(dst, e.From)
	ts := gen.addEdge(dst, e.To)
	for _, f := range fs {
		for _, t := range ts {
			edge := gen.edge(dst, f, t)
			if edge, ok := edge.(UnmarshalerAttr); ok {
				for _, attr := range e.Attrs {
					if err := edge.UnmarshalDOTAttr(attr); err != nil {
						log.Printf("unable to unmarshal DOT attribute %#v; %v", attr, err)
					}
				}
			}

		}
	}
}

// addVertex adds the given vertex to the graph, and returns its set of nodes.
func (gen *generator) addVertex(dst DirectedBuilder, v ast.Vertex) []graph.Node {
	switch v := v.(type) {
	case *ast.Node:
		n := gen.node(dst, v.ID)
		return []graph.Node{n}
	case *ast.Subgraph:
		gen.pushSubgraph()
		for _, stmt := range v.Stmts {
			gen.addStmt(dst, stmt)
		}
		return gen.popSubgraph()
	default:
		panic(fmt.Sprintf("unknown vertex type %T", v))
	}
}

// addEdge adds the given edge to the graph, and returns its set of nodes.
func (gen *generator) addEdge(dst DirectedBuilder, to *ast.Edge) []graph.Node {
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

// pushSubgraph pushes the node start index of the active subgraph onto the
// stack.
func (gen *generator) pushSubgraph() {
	gen.subStart = append(gen.subStart, len(gen.subNodes))
}

// pushSubgraph pops the node start index of the active subgraph from the stack,
// and returns the nodes processed since.
func (gen *generator) popSubgraph() []graph.Node {
	// Get nodes processed since the subgraph became active.
	start := gen.subStart[len(gen.subStart)-1]
	// TODO: Figure out a better way to store subgraph nodes, so that duplicates
	// may not occur.
	nodes := unique(gen.subNodes[start:])
	// Remove subgraph from stack.
	gen.subStart = gen.subStart[:len(gen.subStart)-1]
	if len(gen.subStart) == 0 {
		// Remove subgraph nodes when the bottom-most subgraph has been processed.
		gen.subNodes = gen.subNodes[:0]
	}
	return nodes
}

// unique returns the set of unique nodes contained within ns.
func unique(ns []graph.Node) []graph.Node {
	var nodes []graph.Node
	m := make(map[int]bool)
	for _, n := range ns {
		id := n.ID()
		if m[id] {
			// skip duplicate node
			continue
		}
		m[id] = true
		nodes = append(nodes, n)
	}
	return nodes
}

// isInSubgraph reports whether the active context is within a subgraph, that is
// to be used as a vertex of an edge.
func (gen *generator) isInSubgraph() bool {
	return len(gen.subStart) > 0
}

// appendSubgraphNode appends the given node to the slice of nodes processed
// within the context of a subgraph.
func (gen *generator) appendSubgraphNode(n graph.Node) {
	gen.subNodes = append(gen.subNodes, n)
}
