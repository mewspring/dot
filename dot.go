// Package dot provides access to Graphviz DOT files.
package dot

import (
	"io"
	"io/ioutil"

	"github.com/graphism/dot/ast"
	"github.com/graphism/dot/gocc/lexer"
	"github.com/graphism/dot/gocc/parser"
	"github.com/mewkiz/pkg/errutil"
)

// ParseFile parses the given Graphviz DOT file.
func ParseFile(path string) (*ast.File, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return ParseBytes(buf)
}

// Parse parses the given Graphviz DOT file, reading from r.
func Parse(r io.Reader) (*ast.File, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return ParseBytes(buf)
}

// ParseBytes parses the given Graphviz DOT file, reading from b.
func ParseBytes(b []byte) (*ast.File, error) {
	l := lexer.NewLexer(b)
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

// ParseString parses the given Graphviz DOT file, reading from s.
func ParseString(s string) (*ast.File, error) {
	return ParseBytes([]byte(s))
}
