package cover

import (
	"github.com/benhoyt/goawk/internal/ast"
	"github.com/benhoyt/goawk/parser"
)

func Annotate(prog *parser.Program) {
	prog.Begin = annotateStmtsList(prog.Begin)
	prog.Actions = annotateActions(prog.Actions)
	prog.End = annotateStmtsList(prog.End)
	prog.Functions = annotateFunctions(prog.Functions)
}

func annotateActions(actions []ast.Action) (res []ast.Action) {
	for _, action := range actions {
		action.Stmts = annotateStmts(action.Stmts)
		res = append(res, action)
	}
	return
}

func annotateFunctions(functions []ast.Function) (res []ast.Function) {
	for _, function := range functions {
		function.Body = annotateStmts(function.Body)
		res = append(res, function)
	}
	return
}

func annotateStmtsList(stmtsList []ast.Stmts) (res []ast.Stmts) {
	for _, stmts := range stmtsList {
		res = append(res, annotateStmts(stmts))
	}
	return
}
func annotateStmts(stmts ast.Stmts) (res ast.Stmts) {
	trackProg, err := parser.ParseProgram([]byte(`BEGIN { __COVER[0]++ }`), nil)
	if err != nil {
		panic(err)
	}
	res = append(res, trackProg.Begin[0][0])
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.IfStmt:
			s.Body = annotateStmts(s.Body)
			s.Else = annotateStmts(s.Else)
		case *ast.ForStmt:
			s.Body = annotateStmts(s.Body) // TODO should we do smth with pre & post?
		case *ast.ForInStmt:
			s.Body = annotateStmts(s.Body)
		case *ast.WhileStmt:
			s.Body = annotateStmts(s.Body)
		case *ast.DoWhileStmt:
			s.Body = annotateStmts(s.Body)
		case *ast.BlockStmt:
			s.Body = annotateStmts(s.Body)
		}
		res = append(res, stmt)
	}
	return
	// TODO complete handling of if/else/else if
}
