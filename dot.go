// Package dot provides access to Graphviz DOT files.
package dot

import (
	"github.com/graphism/dot/gocc/ast"
	"github.com/graphism/dot/gocc/lexer"
	"github.com/graphism/dot/gocc/parser"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Rethink API to allow for arbitrary graphs; the return type will not be
// AST related in the future.

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
