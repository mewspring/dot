package dot

import (
	"sort"
	"strconv"

	"github.com/gonum/graph"
	"github.com/graphism/dot/ast"
)

// TODO: Add support for undirected graphs.

// NewGraph returns a new dot AST graph based on the nodes and edges of the
// given gonum graph. The dot AST graph may later be manually annotated with
// additional attributes.
func NewGraph(src graph.Directed) *ast.Graph {
	gen := &dotGen{ids: make(map[int]*ast.NodeID)}
	dst := &ast.Graph{Directed: true}

	// Add nodes.
	nodes := src.Nodes()
	var keys []int
	for _, n := range nodes {
		keys = append(keys, n.ID())
	}
	sort.Ints(keys)
	for _, key := range keys {
		stmt := &ast.NodeStmt{NodeID: gen.node(key)}
		dst.Stmts = append(dst.Stmts, stmt)
	}

	// Add edges.
	for _, to := range nodes {
		for _, from := range src.To(to) {
			stmt := &ast.EdgeStmt{
				From: gen.node(from.ID()),
				To: &ast.Edge{
					Directed: true,
					Vertex:   gen.node(to.ID()),
				},
			}
			dst.Stmts = append(dst.Stmts, stmt)
		}
	}

	return dst
}

// A dotGen keeps track of the information required for generating a dot AST
// graph from a gonum graph.
type dotGen struct {
	// ids maps from gonum node ID to dot AST node ID.
	ids map[int]*ast.NodeID
}

// node returns the dot AST node ID corresponding to the given gonum node ID,
// generating a new such node ID if none exist.
func (gen *dotGen) node(id int) *ast.NodeID {
	if n, ok := gen.ids[id]; ok {
		return n
	}
	n := &ast.NodeID{ID: strconv.Itoa(id)}
	gen.ids[id] = n
	return n
}
