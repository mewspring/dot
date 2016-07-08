// Package astx implements utility functions for generating abstract syntax
// trees of Graphviz DOT graphs.
package astx

import (
	"strings"

	"github.com/graphism/dot/gocc/ast"
	"github.com/graphism/dot/gocc/token"
	"github.com/mewkiz/pkg/errutil"
)

// === [ Graphs ] ==============================================================

// NewGraph returns a new graph based on the given graph strictness, direction,
// optional ID and optional statements.
func NewGraph(strict, directed, optID, optStmts interface{}) (*ast.Graph, error) {
	s, ok := strict.(bool)
	if !ok {
		return nil, errutil.Newf("invalid strictness type; expected bool, got %T", strict)
	}
	d, ok := directed.(bool)
	if !ok {
		return nil, errutil.Newf("invalid direction type; expected bool, got %T", directed)
	}
	i, ok := optID.(string)
	if optID != nil && !ok {
		return nil, errutil.Newf("invalid ID type; expected string, got %T", optID)
	}
	ss, ok := optStmts.([]ast.Stmt)
	if optStmts != nil && !ok {
		return nil, errutil.Newf("invalid statements type; expected []ast.Stmt, got %T", optStmts)
	}
	return &ast.Graph{Strict: s, Directed: d, ID: i, Stmts: ss}, nil
}

// === [ Statements ] ==========================================================

// NewStmtList returns a new statement list based on the given statement.
func NewStmtList(stmt interface{}) ([]ast.Stmt, error) {
	s, ok := stmt.(ast.Stmt)
	if !ok {
		return nil, errutil.Newf("invalid statement type; expected ast.Stmt, got %T", stmt)
	}
	return []ast.Stmt{s}, nil
}

// AppendStmt appends stmt to the given statement list.
func AppendStmt(list, stmt interface{}) ([]ast.Stmt, error) {
	l, ok := list.([]ast.Stmt)
	if !ok {
		return nil, errutil.Newf("invalid statement list type; expected []ast.Stmt, got %T", list)
	}
	s, ok := stmt.(ast.Stmt)
	if !ok {
		return nil, errutil.Newf("invalid statement type; expected ast.Stmt, got %T", stmt)
	}
	return append(l, s), nil
}

// --- [ Node statement ] ------------------------------------------------------

// NewNodeStmt returns a new node statement based on the given node ID and
// optional attributes.
func NewNodeStmt(nodeID, optAttrs interface{}) (*ast.NodeStmt, error) {
	n, ok := nodeID.(*ast.NodeID)
	if !ok {
		return nil, errutil.Newf("invalid node ID type; expected *ast.NodeID, got %T", nodeID)
	}
	as, ok := optAttrs.([]*ast.Attr)
	if optAttrs != nil && !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", optAttrs)
	}
	return &ast.NodeStmt{NodeID: n, Attrs: as}, nil
}

// --- [ Edge statement ] ------------------------------------------------------

// NewEdgeStmt returns a new edge statement based on the given source vertex,
// outgoing edge and optional attributes.
func NewEdgeStmt(from, to, optAttrs interface{}) (*ast.EdgeStmt, error) {
	f, ok := from.(ast.Vertex)
	if !ok {
		return nil, errutil.Newf("invalid source vertex type; expected ast.Vertex, got %T", from)
	}
	t, ok := to.(*ast.Edge)
	if !ok {
		return nil, errutil.Newf("invalid outgoing edge type; expected *ast.Edge, got %T", to)
	}
	as, ok := optAttrs.([]*ast.Attr)
	if optAttrs != nil && !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", optAttrs)
	}
	return &ast.EdgeStmt{From: f, To: t, Attrs: as}, nil
}

// NewEdge returns a new edge based on the given edge direction, destination
// vertex and optional outgoing edge.
func NewEdge(directed, vertex, optTo interface{}) (*ast.Edge, error) {
	d, ok := directed.(bool)
	if !ok {
		return nil, errutil.Newf("invalid direction type; expected bool, got %T", directed)
	}
	v, ok := vertex.(ast.Vertex)
	if !ok {
		return nil, errutil.Newf("invalid destination vertex type; expected ast.Vertex, got %T", vertex)
	}
	t, ok := optTo.(*ast.Edge)
	if optTo != nil && !ok {
		return nil, errutil.Newf("invalid outgoing edge type; expected *ast.Edge, got %T", optTo)
	}
	return &ast.Edge{Directed: d, Vertex: v, To: t}, nil
}

// --- [ Attribute statement ] -------------------------------------------------

// NewAttrStmt returns a new attribute statement based on the given graph
// component kind and attributes.
func NewAttrStmt(kind, attrs interface{}) (*ast.AttrStmt, error) {
	k, ok := kind.(ast.Kind)
	if !ok {
		return nil, errutil.Newf("invalid graph component kind type; expected ast.Kind, got %T", kind)
	}
	as, ok := attrs.([]*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", attrs)
	}
	return &ast.AttrStmt{Kind: k, Attrs: as}, nil
}

// NewAttrList returns a new attribute list based on the given attribute.
func NewAttrList(attr interface{}) ([]*ast.Attr, error) {
	a, ok := attr.(*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute type; expected *ast.Attr, got %T", attr)
	}
	return []*ast.Attr{a}, nil
}

// AppendAttr appends attr to the given attribute list.
func AppendAttr(list, attr interface{}) ([]*ast.Attr, error) {
	l, ok := list.([]*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute list type; expected []*ast.Attr, got %T", list)
	}
	a, ok := attr.(*ast.Attr)
	if !ok {
		return nil, errutil.Newf("invalid attribute type; expected *ast.Attr, got %T", attr)
	}
	return append(l, a), nil
}

// AppendAttrList appends the optional attrs to the given optional attribute
// list.
func AppendAttrList(optList, optAttrs interface{}) ([]*ast.Attr, error) {
	l, ok := optList.([]*ast.Attr)
	if optList != nil && !ok {
		return nil, errutil.Newf("invalid attribute list type; expected []*ast.Attr, got %T", optList)
	}
	as, ok := optAttrs.([]*ast.Attr)
	if optAttrs != nil && !ok {
		return nil, errutil.Newf("invalid attributes type; expected []*ast.Attr, got %T", optAttrs)
	}
	return append(l, as...), nil
}

// --- [ Attribute ] -----------------------------------------------------------

// NewAttr returns a new attribute based on the given key-value pair.
func NewAttr(key, val interface{}) (*ast.Attr, error) {
	k, ok := key.(string)
	if !ok {
		return nil, errutil.Newf("invalid key type; expected string, got %T", key)
	}
	v, ok := val.(string)
	if !ok {
		return nil, errutil.Newf("invalid value type; expected string, got %T", val)
	}
	return &ast.Attr{Key: k, Val: v}, nil
}

// --- [ Subgraph ] ------------------------------------------------------------

// NewSubgraph returns a new subgraph based on the given optional subgraph ID
// and optional statements.
func NewSubgraph(optID, optStmts interface{}) (*ast.Subgraph, error) {
	i, ok := optID.(string)
	if optID != nil && !ok {
		return nil, errutil.Newf("invalid ID type; expected string, got %T", optID)
	}
	ss, ok := optStmts.([]ast.Stmt)
	if optStmts != nil && !ok {
		return nil, errutil.Newf("invalid statements type; expected []ast.Stmt, got %T", optStmts)
	}
	return &ast.Subgraph{ID: i, Stmts: ss}, nil
}

// === [ Vertices ] ============================================================

// --- [ Node identifier ] -----------------------------------------------------

// NewNodeID returns a new node ID based on the given node id and optional port.
func NewNodeID(id, optPort interface{}) (*ast.NodeID, error) {
	i, ok := id.(string)
	if !ok {
		return nil, errutil.Newf("invalid ID type; expected string, got %T", id)
	}
	p, ok := optPort.(*ast.Port)
	if optPort != nil && !ok {
		return nil, errutil.Newf("invalid port type; expected *ast.Port, got %T", optPort)
	}
	return &ast.NodeID{ID: i, Port: p}, nil
}

// NewPort returns a new port based on the given id and optional compass point.
func NewPort(id, optCompassPoint interface{}) (*ast.Port, error) {
	// NOTE: If optCompassPoint is nil, id may be either an identifier or a
	// compass point.
	//
	// The following strings are valid compass points:
	//
	//    "n", "ne", "e", "se", "s", "sw", "w", "nw", "c" and "_"
	i, ok := id.(string)
	if !ok {
		return nil, errutil.Newf("invalid ID type; expected string, got %T", id)
	}

	// Early return if optional compass point is absent and ID is a valid compass
	// point.
	if optCompassPoint == nil {
		if compassPoint, ok := getCompassPoint(i); ok {
			return &ast.Port{CompassPoint: compassPoint}, nil
		}
	}

	c, ok := optCompassPoint.(string)
	if optCompassPoint != nil && !ok {
		return nil, errutil.Newf("invalid compass point type; expected string, got %T", optCompassPoint)
	}
	compassPoint, _ := getCompassPoint(c)
	return &ast.Port{ID: i, CompassPoint: compassPoint}, nil
}

// getCompassPoint returns the corresponding compass point to the given string,
// and a boolean value indicating if such a compass point exists.
func getCompassPoint(s string) (ast.CompassPoint, bool) {
	switch s {
	case "_":
		return ast.CompassPointDefault, true
	case "n":
		return ast.CompassPointNorth, true
	case "ne":
		return ast.CompassPointNorthEast, true
	case "e":
		return ast.CompassPointEast, true
	case "se":
		return ast.CompassPointSouthEast, true
	case "s":
		return ast.CompassPointSouth, true
	case "sw":
		return ast.CompassPointSouthWest, true
	case "w":
		return ast.CompassPointWest, true
	case "nw":
		return ast.CompassPointNorthWest, true
	case "c":
		return ast.CompassPointCenter, true
	}
	return ast.CompassPointDefault, false
}

// === [ Identifiers ] =========================================================

// NewID returns a new identifier based on the given ID token.
func NewID(id interface{}) (string, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return "", errutil.Newf("invalid identifier type; expected *token.Token, got %T", id)
	}
	s := string(i.Lit)

	// In quoted strings in DOT, the only escaped character is double-quote (").
	// That is, in quoted strings, the dyad \" is converted to "; all other
	// characters are left unchanged. In particular, \\ remains \\.
	//
	// Convert \" to "
	s = strings.Replace(s, `\"`, `"`, -1)

	// As another aid for readability, dot allows double-quoted strings to span
	// multiple physical lines using the standard C convention of a backslash
	// immediately preceding a newline character.
	//
	// Strip "\\\n" sequences.
	s = strings.Replace(s, "\\\n", "", -1)

	// TODO: Add support for concatenated using a '+' operator.

	return s, nil
}
