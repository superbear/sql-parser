package main

import (
	"fmt"
	"log"
	"strings"

	parser "github.com/superbear/sql-parser/starrocks"

	antlr "github.com/antlr4-go/antlr/v4"
)

type ColumnExtractor struct {
	parser.BaseStarRocksParserListener

	colNames []string
	skip     bool
}

func (l *ColumnExtractor) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// log.Printf("enter: %T %s\n", ctx, ctx.GetText())
}

func (l *ColumnExtractor) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// log.Println("exit: ", ctx.GetText())
}

// EnterQuerySpecification is called when production querySpecification is entered.
func (l *ColumnExtractor) EnterQuerySpecification(ctx *parser.QuerySpecificationContext) {
	if l.skip || len(l.colNames) > 0 {
		return
	}

	for _, item := range ctx.AllSelectItem() {
		var name string
		switch v := item.(type) {
		case *parser.SelectSingleContext:
			if v.Identifier() == nil {
				if v.String_() == nil {
					name = v.GetText()
				} else {
					name = v.String_().GetText()
				}
			} else {
				name = v.Identifier().GetText()
			}

		default:
			name = v.GetText()
		}

		l.colNames = append(l.colNames, unquote(name))
	}
}

// EnterSubquery is called when production subquery is entered.
func (l *ColumnExtractor) EnterSubquery(ctx *parser.SubqueryContext) {
	l.skip = true
}

// ExitSubquery is called when production subquery is exited.
func (l *ColumnExtractor) ExitSubquery(ctx *parser.SubqueryContext) {
	l.skip = false
}

// EnterWithClause is called when production withClause is entered.
func (l *ColumnExtractor) EnterWithClause(ctx *parser.WithClauseContext) {
	l.skip = true
}

// ExitWithClause is called when production withClause is exited.
func (l *ColumnExtractor) ExitWithClause(ctx *parser.WithClauseContext) {
	l.skip = false
}

func unquote(s string) string {
	specialChars := []string{"\"", "'", "`"}
	for _, char := range specialChars {
		if strings.HasPrefix(s, char) && strings.HasSuffix(s, char) {
			s = strings.TrimPrefix(s, char)
			s = strings.TrimSuffix(s, char)
			break
		}
	}
	return s
}

// CustomErrorListener 自定义错误监听器
type CustomErrorListener struct {
	antlr.DefaultErrorListener
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("Syntax error at line %d, column %d: %s\n", line, column, msg)
}

func main() {
	sql := `WITH avg_salary AS (
		SELECT AVG(salary) AS average FROM employees
	)
	SELECT name, salary
	FROM employees
	WHERE salary > (SELECT average FROM avg_salary);
	`

	log.Println("sql: ", sql)
	input := antlr.NewInputStream(sql)
	lexer := parser.NewStarRocksLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	starRocksSQLParser := parser.NewStarRocksParser(stream)
	starRocksSQLParser.RemoveErrorListeners()
	starRocksSQLParser.AddErrorListener(&CustomErrorListener{})

	var listener ColumnExtractor
	tree := starRocksSQLParser.SqlStatements()
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)

	log.Println("colNames: ", listener.colNames)
	// log.Println("tree: ", tree.ToStringTree(nil, starRocksSQLParser))
}
