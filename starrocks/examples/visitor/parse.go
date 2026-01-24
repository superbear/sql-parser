package main

import (
	"fmt"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
	"github.com/superbear/sql-parser/starrocks"
)

type ColumnExtractorVisitor struct {
	starrocks.BaseStarRocksParserVisitor

	colNames []string
}

func (s *ColumnExtractorVisitor) GetColNames() []string {
	return s.colNames
}

func (v *ColumnExtractorVisitor) Visit(node antlr.ParseTree) any {
	if node == nil {
		return nil
	}

	switch node := node.(type) {
	case *starrocks.SingleStatementContext:
		return v.VisitSingleStatement(node)

	case *starrocks.StatementContext:
		return v.VisitStatement(node)

	case *starrocks.QueryStatementContext:
		return v.VisitQueryStatement(node)

	case *starrocks.QueryRelationContext:
		return v.VisitQueryRelation(node)

	case *starrocks.QueryNoWithContext:
		return v.VisitQueryNoWith(node)

	case *starrocks.QueryPrimaryDefaultContext:
		return v.VisitQueryPrimaryDefault(node)

	case *starrocks.QueryWithParenthesesContext:
		return v.VisitQueryWithParentheses(node)

	case *starrocks.QuerySpecificationContext:
		return v.VisitQuerySpecification(node)

	case *starrocks.SetOperationContext:
		return v.VisitSetOperation(node)

	case *antlr.TerminalNodeImpl:
		return node.GetText()

	default:
		return node.Accept(v)
	}
}

func (v *ColumnExtractorVisitor) VisitSqlStatements(ctx *starrocks.SqlStatementsContext) any {
	for _, stmt := range ctx.AllSingleStatement() {
		v.Visit(stmt)
	}
	return nil
}

func (v *ColumnExtractorVisitor) VisitSingleStatement(ctx *starrocks.SingleStatementContext) any {
	return v.Visit(ctx.Statement())
}

func (v *ColumnExtractorVisitor) VisitStatement(ctx *starrocks.StatementContext) any {
	return v.Visit(ctx.QueryStatement())
}

func (v *ColumnExtractorVisitor) VisitQueryStatement(ctx *starrocks.QueryStatementContext) any {
	return v.Visit(ctx.QueryRelation())
}

func (v *ColumnExtractorVisitor) VisitQueryRelation(ctx *starrocks.QueryRelationContext) any {
	return v.Visit(ctx.QueryNoWith())
}

func (v *ColumnExtractorVisitor) VisitQueryNoWith(ctx *starrocks.QueryNoWithContext) any {
	return v.Visit(ctx.QueryPrimary())
}

func (v *ColumnExtractorVisitor) VisitQueryWithParentheses(ctx *starrocks.QueryWithParenthesesContext) any {
	// ignore subquery
	return nil
}

func (v *ColumnExtractorVisitor) VisitSetOperation(ctx *starrocks.SetOperationContext) any {
	// union get left
	return v.Visit(ctx.GetLeft())
}

func (v *ColumnExtractorVisitor) VisitQueryPrimaryDefault(ctx *starrocks.QueryPrimaryDefaultContext) any {
	return v.Visit(ctx.QuerySpecification())
}

func (v *ColumnExtractorVisitor) VisitQuerySpecification(ctx *starrocks.QuerySpecificationContext) any {
	if len(v.colNames) > 0 {
		return nil
	}

	for _, item := range ctx.AllSelectItem() {
		var name string
		switch v := item.(type) {
		case *starrocks.SelectSingleContext:
			if v.Identifier() == nil {
				if v.String_() == nil {
					name = v.GetText()
				} else {
					// alias select xxx as 'xxx' or select xxx as "xxx"
					name = v.String_().GetText()
				}
			} else {
				// as
				name = v.Identifier().GetText()
			}

		default:
			// select * or select * exclude
			name = v.GetText()
		}

		v.colNames = append(v.colNames, unquote(name))
	}
	return nil
}

// CustomErrorListener 自定义错误监听器
type CustomErrorListener struct {
	antlr.DefaultErrorListener
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{}
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("Syntax error at line %d, column %d: %s\n", line, column, msg)
}

func unquote(s string) string {
	chars := []string{"\"", "'", "`"}
	for _, char := range chars {
		if strings.HasPrefix(s, char) && strings.HasSuffix(s, char) {
			s = strings.TrimPrefix(s, char)
			s = strings.TrimSuffix(s, char)
			break
		}
	}
	return s
}

func parse(sql string) ([]string, error) {
	input := antlr.NewInputStream(sql)
	lexer := starrocks.NewStarRocksLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := starrocks.NewStarRocksParser(stream)

	errListener := NewCustomErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errListener)

	var vistor ColumnExtractorVisitor
	parser.SqlStatements().Accept(&vistor)

	return vistor.GetColNames(), nil
}
