// TODO: Consider which types are required as pointers (to be used as
// SOME/NONE), and which may be defined as values. Decide whether to define
// those types as values, or keep them as pointers for consistency.

// Package ast declares the types used to represent abstract syntax trees of
// Graphviz DOT graphs.
package ast

// === [ File ] ================================================================

// A File represents a DOT file.
type File struct {
	// Graphs.
	Graphs []*Graph
}

// === [ Graphs ] ==============================================================

// A Graph represents a directed or an undirected graph.
//
// Examples.
//
//    digraph G {
//       A -> {B C}
//       B -> C
//    }
type Graph struct {
	// Strict graph; multi-edges forbidden.
	Strict bool
	// Directed graph.
	Directed bool
	// Graph ID; or empty if anonymous.
	ID string
	// Graph statements.
	Stmts []Stmt
}

// === [ Statements ] ==========================================================

// A Stmt represents a statement, and has one of the following underlying types.
//
//    *NodeStmt
//    *EdgeStmt
//    *AttrStmt
//    *Attr
//    *Subgraph
type Stmt interface {
	// isStmt ensures that only statements can be assigned to the Stmt interface.
	isStmt()
}

// --- [ Node statement ] ------------------------------------------------------

// A NodeStmt represents a node statement.
//
// Examples.
//
//    A [ color=blue ]
type NodeStmt struct {
	// Node ID.
	NodeID *NodeID
	// Node attributes.
	Attrs []*Attr
}

// --- [ Edge statement ] ------------------------------------------------------

// An EdgeStmt represents an edge statement.
//
// Examples.
//
//    A -> B
//    A -> {B C}
//    A -> B -> C
type EdgeStmt struct {
	// Source vertex.
	From Vertex
	// Outgoing edge.
	To *Edge
	// Edge attributes.
	Attrs []*Attr
}

// An Edge represents an edge between two vertices.
type Edge struct {
	// Directed edge.
	Directed bool
	// Destination vertex.
	Vertex Vertex
	// Outgoing edge; or nil if none.
	To *Edge
}

// --- [ Attribute statement ] -------------------------------------------------

// An AttrStmt represents an attribute statement.
//
// Examples.
//
//    graph [ rankdir=LR ]
//    node [ color=blue ]
//    edge [ minlen=1 ]
type AttrStmt struct {
	// Graph component kind to which the attributes are assigned.
	Kind Kind
	// Attributes.
	Attrs []*Attr
}

// Kind specifies the set of graph components to which attribute statements may
// be assigned.
type Kind uint

// Graph component kinds.
const (
	KindGraph Kind = iota
	KindNode
	KindEdge
)

// --- [ Attribute ] -----------------------------------------------------------

// An Attr represents an attribute.
//
// Examples.
//
//    rank=same
type Attr struct {
	// Attribute key.
	Key string
	// Attribute value.
	Val string
}

// --- [ Subgraph ] ------------------------------------------------------------

// A Subgraph represents a subgraph vertex.
//
// Examples.
//
//    subgraph S {A B C}
type Subgraph struct {
	// Subgraph ID; or empty if none.
	ID string
	// Subgraph statements.
	Stmts []Stmt
}

// isStmt ensures that only statements can be assigned to the Stmt interface.
func (*NodeStmt) isStmt() {}
func (*EdgeStmt) isStmt() {}
func (*AttrStmt) isStmt() {}
func (*Attr) isStmt()     {}
func (*Subgraph) isStmt() {}

// === [ Vertices ] ============================================================

// A Vertex represents a vertex, and has one of the following underlying types.
//
//    *NodeID
//    *Subgraph
type Vertex interface {
	// isVertex ensures that only vertices can be assigned to the Vertex
	// interface.
	isVertex()
}

// --- [ Node identifier ] -----------------------------------------------------

// A NodeID represents a node vertex.
//
// Examples.
//
//    A
//    A:nw
type NodeID struct {
	// Node ID.
	ID string
	// Node port; or nil if none.
	Port *Port
}

// A Port specifies where on a node an edge should be aimed.
type Port struct {
	// Port ID; or empty if none.
	ID string
	// Compass point.
	CompassPoint CompassPoint
}

// CompassPoint specifies the set of compass points.
type CompassPoint uint

// Compass points.
const (
	CompassPointDefault   CompassPoint = iota // _
	CompassPointNorth                         // n
	CompassPointNorthEast                     // ne
	CompassPointEast                          // e
	CompassPointSouthEast                     // se
	CompassPointSouth                         // s
	CompassPointSouthWest                     // sw
	CompassPointWest                          // w
	CompassPointNorthWest                     // nw
	CompassPointCenter                        // c
)

// isVertex ensures that only vertices can be assigned to the Vertex interface.
func (*NodeID) isVertex()   {}
func (*Subgraph) isVertex() {}
