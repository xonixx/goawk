package cover

import (
	"errors"
	"fmt"
	"github.com/benhoyt/goawk/internal/ast"
	. "github.com/benhoyt/goawk/parser"
	"os"
)

type Annotator struct {
	covermode     string
	annotationIdx int
	boundaries    map[int]ast.Boundary
	stmtsCnt      map[int]int
}

func NewAnnotator(covermode string) *Annotator {
	return &Annotator{covermode, 0, map[int]ast.Boundary{}, map[int]int{}}
}

func (annotator *Annotator) Annotate(prog *Program) {
	IDX_COVER = len(prog.Arrays)
	IDX_COVER_DATA = IDX_COVER + 1
	prog.Begin = annotator.annotateStmtsList(prog.Begin)
	prog.Actions = annotator.annotateActions(prog.Actions)
	prog.End = annotator.annotateStmtsList(prog.End)
	prog.Functions = annotator.annotateFunctions(prog.Functions)
	prog.Arrays[ARR_COVER] = IDX_COVER
	prog.Arrays[ARR_COVER_DATA] = IDX_COVER_DATA
}

func (annotator *Annotator) AppendCoverData(coverprofile string, coverData map[int]int64) error {
	//fmt.Printf("Cover data: %v", coverData)
	//return nil

	// 1a. If file doesn't exist - create and write covermode line
	// 1b. If file exists - open it for writing in append mode
	// 2.  Write all coverData lines

	var f *os.File
	if _, err := os.Stat(coverprofile); errors.Is(err, os.ErrNotExist) { // TODO error if exists and is folder
		f, err = os.OpenFile(coverprofile, os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		_, err := f.WriteString("mode: " + annotator.covermode + "\n")
		if err != nil {
			return err
		}
	} else if err == nil {
		// file exists
		f, err = os.OpenFile(coverprofile, os.O_APPEND, 0644)
		if err != nil {
			return err
		}
	} else {
		panic(err)
	}
	for i := 1; i <= annotator.annotationIdx; i++ {
		_, err := f.WriteString(renderCoverDataLine(annotator.boundaries[i], annotator.stmtsCnt[i], coverData[i]))
		if err != nil {
			return err
		}
	}
	return nil
}

func (annotator *Annotator) annotateActions(actions []ast.Action) (res []ast.Action) {
	for _, action := range actions {
		action.Stmts = annotator.annotateStmts(action.Stmts)
		res = append(res, action)
	}
	return
}

func (annotator *Annotator) annotateFunctions(functions []ast.Function) (res []ast.Function) {
	for _, function := range functions {
		function.Body = annotator.annotateStmts(function.Body)
		res = append(res, function)
	}
	return
}

func (annotator *Annotator) annotateStmtsList(stmtsList []ast.Stmts) (res []ast.Stmts) {
	for _, stmts := range stmtsList {
		res = append(res, annotator.annotateStmts(stmts))
	}
	return
}

// annotateStmts takes a list of statements and adds counters to the beginning of
// each basic block at the top level of that list. For instance, given
//
//	S1
//	if cond {
//		S2
//	}
//	S3
//
// counters will be added before S1,S2,S3.
func (annotator *Annotator) annotateStmts(stmts ast.Stmts) (res ast.Stmts) {
	var simpleStatements []ast.Stmt
	for _, stmt := range stmts {
		wasBlock := true
		switch s := stmt.(type) {
		case *ast.IfStmt:
			s.Body = annotator.annotateStmts(s.Body)
			s.Else = annotator.annotateStmts(s.Else)
		case *ast.ForStmt:
			s.Body = annotator.annotateStmts(s.Body) // TODO should we do smth with pre & post?
		case *ast.ForInStmt:
			s.Body = annotator.annotateStmts(s.Body)
		case *ast.WhileStmt:
			s.Body = annotator.annotateStmts(s.Body)
		case *ast.DoWhileStmt:
			s.Body = annotator.annotateStmts(s.Body)
		case *ast.BlockStmt:
			s.Body = annotator.annotateStmts(s.Body)
		default:
			wasBlock = false
		}
		if wasBlock {
			if len(simpleStatements) > 0 {
				res = append(res, annotator.trackStatement(simpleStatements))
				res = append(res, simpleStatements...)
			}
			res = append(res, stmt)
			simpleStatements = []ast.Stmt{}
		} else {
			simpleStatements = append(simpleStatements, stmt)
		}
	}
	if len(simpleStatements) > 0 {
		res = append(res, annotator.trackStatement(simpleStatements))
		res = append(res, simpleStatements...)
	}
	return
	// TODO complete handling of if/else/else if
}
func (annotator *Annotator) trackStatement(statements []ast.Stmt) ast.Stmt {
	op := "=1"
	if annotator.covermode == "count" {
		op = "++"
	}
	annotator.annotationIdx++
	firstStmtBoundary := statements[0].(ast.SimpleStmt).GetBoundary()
	lastStmtBoundary := statements[len(statements)-1].(ast.SimpleStmt).GetBoundary()
	annotator.boundaries[annotator.annotationIdx] = ast.Boundary{
		Start:    firstStmtBoundary.Start,
		End:      lastStmtBoundary.End,
		FileName: firstStmtBoundary.FileName,
	}
	annotator.stmtsCnt[annotator.annotationIdx] = len(statements)
	return parseProg(fmt.Sprintf(`BEGIN { %s[%d]%s }`, ARR_COVER, annotator.annotationIdx, op)).Begin[0][0]
}

func parseProg(code string) *Program {
	prog, err := ParseProgram([]byte(code), nil)
	if err != nil {
		panic(err)
	}
	return prog
}

/*func (Annotator *Annotator) addCoverageEnd(prog *Program) {
	var code strings.Builder
	code.WriteString("END {")
	for i := 1; i <= Annotator.annotationIdx; i++ {
		code.WriteString(fmt.Sprintf("%s[%d]=\"%s\"\n", ARR_COVER_DATA, i, renderCoverData(Annotator.boundaries[i], Annotator.stmtsCnt[i])))
	}

	code.WriteString(fmt.Sprintf("for(i=1;i<=%d;i++){\n", Annotator.annotationIdx))
	code.WriteString("  printf \"%s %s\\n\", " + ARR_COVER_DATA + "[i], +" + ARR_COVER + "[i] >> \"" + Annotator.coverpofile + "\"\n")
	code.WriteString("}\n")
	//code.WriteString("fflush(\"" + Annotator.coverpofile + "\")\n")
	//code.WriteString("close(\"" + Annotator.coverpofile + "\")\n")

	code.WriteString("}\n")
	prog.End = append(prog.End, parseProg(code.String()).End...)
}*/

func renderCoverDataLine(boundary ast.Boundary, stmtsCnt int, cnt int64) string {
	return fmt.Sprintf("%s:%d.%d,%d.%d %d %d",
		boundary.FileName,
		boundary.Start.Line, boundary.Start.Column,
		boundary.End.Line, boundary.End.Column,
		stmtsCnt, cnt,
	)
}
