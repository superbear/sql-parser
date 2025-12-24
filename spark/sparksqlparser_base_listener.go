// Code generated from SparkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spark // SparkSqlParser
import "github.com/antlr4-go/antlr/v4"

// BaseSparkSqlParserListener is a complete listener for a parse tree produced by SparkSqlParser.
type BaseSparkSqlParserListener struct{}

var _ SparkSqlParserListener = &BaseSparkSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSparkSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSparkSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSparkSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSparkSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompoundOrSingleStatements is called when production compoundOrSingleStatements is entered.
func (s *BaseSparkSqlParserListener) EnterCompoundOrSingleStatements(ctx *CompoundOrSingleStatementsContext) {}

// ExitCompoundOrSingleStatements is called when production compoundOrSingleStatements is exited.
func (s *BaseSparkSqlParserListener) ExitCompoundOrSingleStatements(ctx *CompoundOrSingleStatementsContext) {}

// EnterCompoundOrSingleStatement is called when production compoundOrSingleStatement is entered.
func (s *BaseSparkSqlParserListener) EnterCompoundOrSingleStatement(ctx *CompoundOrSingleStatementContext) {}

// ExitCompoundOrSingleStatement is called when production compoundOrSingleStatement is exited.
func (s *BaseSparkSqlParserListener) ExitCompoundOrSingleStatement(ctx *CompoundOrSingleStatementContext) {}

// EnterSingleCompoundStatement is called when production singleCompoundStatement is entered.
func (s *BaseSparkSqlParserListener) EnterSingleCompoundStatement(ctx *SingleCompoundStatementContext) {}

// ExitSingleCompoundStatement is called when production singleCompoundStatement is exited.
func (s *BaseSparkSqlParserListener) ExitSingleCompoundStatement(ctx *SingleCompoundStatementContext) {}

// EnterBeginEndCompoundBlock is called when production beginEndCompoundBlock is entered.
func (s *BaseSparkSqlParserListener) EnterBeginEndCompoundBlock(ctx *BeginEndCompoundBlockContext) {}

// ExitBeginEndCompoundBlock is called when production beginEndCompoundBlock is exited.
func (s *BaseSparkSqlParserListener) ExitBeginEndCompoundBlock(ctx *BeginEndCompoundBlockContext) {}

// EnterCompoundBody is called when production compoundBody is entered.
func (s *BaseSparkSqlParserListener) EnterCompoundBody(ctx *CompoundBodyContext) {}

// ExitCompoundBody is called when production compoundBody is exited.
func (s *BaseSparkSqlParserListener) ExitCompoundBody(ctx *CompoundBodyContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseSparkSqlParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseSparkSqlParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterSetVariableInsideSqlScript is called when production setVariableInsideSqlScript is entered.
func (s *BaseSparkSqlParserListener) EnterSetVariableInsideSqlScript(ctx *SetVariableInsideSqlScriptContext) {}

// ExitSetVariableInsideSqlScript is called when production setVariableInsideSqlScript is exited.
func (s *BaseSparkSqlParserListener) ExitSetVariableInsideSqlScript(ctx *SetVariableInsideSqlScriptContext) {}

// EnterSqlStateValue is called when production sqlStateValue is entered.
func (s *BaseSparkSqlParserListener) EnterSqlStateValue(ctx *SqlStateValueContext) {}

// ExitSqlStateValue is called when production sqlStateValue is exited.
func (s *BaseSparkSqlParserListener) ExitSqlStateValue(ctx *SqlStateValueContext) {}

// EnterDeclareConditionStatement is called when production declareConditionStatement is entered.
func (s *BaseSparkSqlParserListener) EnterDeclareConditionStatement(ctx *DeclareConditionStatementContext) {}

// ExitDeclareConditionStatement is called when production declareConditionStatement is exited.
func (s *BaseSparkSqlParserListener) ExitDeclareConditionStatement(ctx *DeclareConditionStatementContext) {}

// EnterConditionValue is called when production conditionValue is entered.
func (s *BaseSparkSqlParserListener) EnterConditionValue(ctx *ConditionValueContext) {}

// ExitConditionValue is called when production conditionValue is exited.
func (s *BaseSparkSqlParserListener) ExitConditionValue(ctx *ConditionValueContext) {}

// EnterConditionValues is called when production conditionValues is entered.
func (s *BaseSparkSqlParserListener) EnterConditionValues(ctx *ConditionValuesContext) {}

// ExitConditionValues is called when production conditionValues is exited.
func (s *BaseSparkSqlParserListener) ExitConditionValues(ctx *ConditionValuesContext) {}

// EnterDeclareHandlerStatement is called when production declareHandlerStatement is entered.
func (s *BaseSparkSqlParserListener) EnterDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) {}

// ExitDeclareHandlerStatement is called when production declareHandlerStatement is exited.
func (s *BaseSparkSqlParserListener) ExitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseSparkSqlParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseSparkSqlParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterIfElseStatement is called when production ifElseStatement is entered.
func (s *BaseSparkSqlParserListener) EnterIfElseStatement(ctx *IfElseStatementContext) {}

// ExitIfElseStatement is called when production ifElseStatement is exited.
func (s *BaseSparkSqlParserListener) ExitIfElseStatement(ctx *IfElseStatementContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BaseSparkSqlParserListener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BaseSparkSqlParserListener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterLeaveStatement is called when production leaveStatement is entered.
func (s *BaseSparkSqlParserListener) EnterLeaveStatement(ctx *LeaveStatementContext) {}

// ExitLeaveStatement is called when production leaveStatement is exited.
func (s *BaseSparkSqlParserListener) ExitLeaveStatement(ctx *LeaveStatementContext) {}

// EnterIterateStatement is called when production iterateStatement is entered.
func (s *BaseSparkSqlParserListener) EnterIterateStatement(ctx *IterateStatementContext) {}

// ExitIterateStatement is called when production iterateStatement is exited.
func (s *BaseSparkSqlParserListener) ExitIterateStatement(ctx *IterateStatementContext) {}

// EnterSearchedCaseStatement is called when production searchedCaseStatement is entered.
func (s *BaseSparkSqlParserListener) EnterSearchedCaseStatement(ctx *SearchedCaseStatementContext) {}

// ExitSearchedCaseStatement is called when production searchedCaseStatement is exited.
func (s *BaseSparkSqlParserListener) ExitSearchedCaseStatement(ctx *SearchedCaseStatementContext) {}

// EnterSimpleCaseStatement is called when production simpleCaseStatement is entered.
func (s *BaseSparkSqlParserListener) EnterSimpleCaseStatement(ctx *SimpleCaseStatementContext) {}

// ExitSimpleCaseStatement is called when production simpleCaseStatement is exited.
func (s *BaseSparkSqlParserListener) ExitSimpleCaseStatement(ctx *SimpleCaseStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseSparkSqlParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseSparkSqlParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseSparkSqlParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseSparkSqlParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseSparkSqlParserListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseSparkSqlParserListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterBeginLabel is called when production beginLabel is entered.
func (s *BaseSparkSqlParserListener) EnterBeginLabel(ctx *BeginLabelContext) {}

// ExitBeginLabel is called when production beginLabel is exited.
func (s *BaseSparkSqlParserListener) ExitBeginLabel(ctx *BeginLabelContext) {}

// EnterEndLabel is called when production endLabel is entered.
func (s *BaseSparkSqlParserListener) EnterEndLabel(ctx *EndLabelContext) {}

// ExitEndLabel is called when production endLabel is exited.
func (s *BaseSparkSqlParserListener) ExitEndLabel(ctx *EndLabelContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseSparkSqlParserListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseSparkSqlParserListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterSingleTableIdentifier is called when production singleTableIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterSingleTableIdentifier(ctx *SingleTableIdentifierContext) {}

// ExitSingleTableIdentifier is called when production singleTableIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitSingleTableIdentifier(ctx *SingleTableIdentifierContext) {}

// EnterSingleMultipartIdentifier is called when production singleMultipartIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterSingleMultipartIdentifier(ctx *SingleMultipartIdentifierContext) {}

// ExitSingleMultipartIdentifier is called when production singleMultipartIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitSingleMultipartIdentifier(ctx *SingleMultipartIdentifierContext) {}

// EnterSingleFunctionIdentifier is called when production singleFunctionIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterSingleFunctionIdentifier(ctx *SingleFunctionIdentifierContext) {}

// ExitSingleFunctionIdentifier is called when production singleFunctionIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitSingleFunctionIdentifier(ctx *SingleFunctionIdentifierContext) {}

// EnterSingleDataType is called when production singleDataType is entered.
func (s *BaseSparkSqlParserListener) EnterSingleDataType(ctx *SingleDataTypeContext) {}

// ExitSingleDataType is called when production singleDataType is exited.
func (s *BaseSparkSqlParserListener) ExitSingleDataType(ctx *SingleDataTypeContext) {}

// EnterSingleTableSchema is called when production singleTableSchema is entered.
func (s *BaseSparkSqlParserListener) EnterSingleTableSchema(ctx *SingleTableSchemaContext) {}

// ExitSingleTableSchema is called when production singleTableSchema is exited.
func (s *BaseSparkSqlParserListener) ExitSingleTableSchema(ctx *SingleTableSchemaContext) {}

// EnterSingleRoutineParamList is called when production singleRoutineParamList is entered.
func (s *BaseSparkSqlParserListener) EnterSingleRoutineParamList(ctx *SingleRoutineParamListContext) {}

// ExitSingleRoutineParamList is called when production singleRoutineParamList is exited.
func (s *BaseSparkSqlParserListener) ExitSingleRoutineParamList(ctx *SingleRoutineParamListContext) {}

// EnterStatementDefault is called when production statementDefault is entered.
func (s *BaseSparkSqlParserListener) EnterStatementDefault(ctx *StatementDefaultContext) {}

// ExitStatementDefault is called when production statementDefault is exited.
func (s *BaseSparkSqlParserListener) ExitStatementDefault(ctx *StatementDefaultContext) {}

// EnterVisitExecuteImmediate is called when production visitExecuteImmediate is entered.
func (s *BaseSparkSqlParserListener) EnterVisitExecuteImmediate(ctx *VisitExecuteImmediateContext) {}

// ExitVisitExecuteImmediate is called when production visitExecuteImmediate is exited.
func (s *BaseSparkSqlParserListener) ExitVisitExecuteImmediate(ctx *VisitExecuteImmediateContext) {}

// EnterDmlStatement is called when production dmlStatement is entered.
func (s *BaseSparkSqlParserListener) EnterDmlStatement(ctx *DmlStatementContext) {}

// ExitDmlStatement is called when production dmlStatement is exited.
func (s *BaseSparkSqlParserListener) ExitDmlStatement(ctx *DmlStatementContext) {}

// EnterUse is called when production use is entered.
func (s *BaseSparkSqlParserListener) EnterUse(ctx *UseContext) {}

// ExitUse is called when production use is exited.
func (s *BaseSparkSqlParserListener) ExitUse(ctx *UseContext) {}

// EnterUseNamespace is called when production useNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterUseNamespace(ctx *UseNamespaceContext) {}

// ExitUseNamespace is called when production useNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitUseNamespace(ctx *UseNamespaceContext) {}

// EnterSetCatalog is called when production setCatalog is entered.
func (s *BaseSparkSqlParserListener) EnterSetCatalog(ctx *SetCatalogContext) {}

// ExitSetCatalog is called when production setCatalog is exited.
func (s *BaseSparkSqlParserListener) ExitSetCatalog(ctx *SetCatalogContext) {}

// EnterCreateNamespace is called when production createNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterCreateNamespace(ctx *CreateNamespaceContext) {}

// ExitCreateNamespace is called when production createNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitCreateNamespace(ctx *CreateNamespaceContext) {}

// EnterSetNamespaceProperties is called when production setNamespaceProperties is entered.
func (s *BaseSparkSqlParserListener) EnterSetNamespaceProperties(ctx *SetNamespacePropertiesContext) {}

// ExitSetNamespaceProperties is called when production setNamespaceProperties is exited.
func (s *BaseSparkSqlParserListener) ExitSetNamespaceProperties(ctx *SetNamespacePropertiesContext) {}

// EnterUnsetNamespaceProperties is called when production unsetNamespaceProperties is entered.
func (s *BaseSparkSqlParserListener) EnterUnsetNamespaceProperties(ctx *UnsetNamespacePropertiesContext) {}

// ExitUnsetNamespaceProperties is called when production unsetNamespaceProperties is exited.
func (s *BaseSparkSqlParserListener) ExitUnsetNamespaceProperties(ctx *UnsetNamespacePropertiesContext) {}

// EnterSetNamespaceCollation is called when production setNamespaceCollation is entered.
func (s *BaseSparkSqlParserListener) EnterSetNamespaceCollation(ctx *SetNamespaceCollationContext) {}

// ExitSetNamespaceCollation is called when production setNamespaceCollation is exited.
func (s *BaseSparkSqlParserListener) ExitSetNamespaceCollation(ctx *SetNamespaceCollationContext) {}

// EnterSetNamespaceLocation is called when production setNamespaceLocation is entered.
func (s *BaseSparkSqlParserListener) EnterSetNamespaceLocation(ctx *SetNamespaceLocationContext) {}

// ExitSetNamespaceLocation is called when production setNamespaceLocation is exited.
func (s *BaseSparkSqlParserListener) ExitSetNamespaceLocation(ctx *SetNamespaceLocationContext) {}

// EnterDropNamespace is called when production dropNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterDropNamespace(ctx *DropNamespaceContext) {}

// ExitDropNamespace is called when production dropNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitDropNamespace(ctx *DropNamespaceContext) {}

// EnterShowNamespaces is called when production showNamespaces is entered.
func (s *BaseSparkSqlParserListener) EnterShowNamespaces(ctx *ShowNamespacesContext) {}

// ExitShowNamespaces is called when production showNamespaces is exited.
func (s *BaseSparkSqlParserListener) ExitShowNamespaces(ctx *ShowNamespacesContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSparkSqlParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSparkSqlParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateTableLike is called when production createTableLike is entered.
func (s *BaseSparkSqlParserListener) EnterCreateTableLike(ctx *CreateTableLikeContext) {}

// ExitCreateTableLike is called when production createTableLike is exited.
func (s *BaseSparkSqlParserListener) ExitCreateTableLike(ctx *CreateTableLikeContext) {}

// EnterReplaceTable is called when production replaceTable is entered.
func (s *BaseSparkSqlParserListener) EnterReplaceTable(ctx *ReplaceTableContext) {}

// ExitReplaceTable is called when production replaceTable is exited.
func (s *BaseSparkSqlParserListener) ExitReplaceTable(ctx *ReplaceTableContext) {}

// EnterAnalyze is called when production analyze is entered.
func (s *BaseSparkSqlParserListener) EnterAnalyze(ctx *AnalyzeContext) {}

// ExitAnalyze is called when production analyze is exited.
func (s *BaseSparkSqlParserListener) ExitAnalyze(ctx *AnalyzeContext) {}

// EnterAnalyzeTables is called when production analyzeTables is entered.
func (s *BaseSparkSqlParserListener) EnterAnalyzeTables(ctx *AnalyzeTablesContext) {}

// ExitAnalyzeTables is called when production analyzeTables is exited.
func (s *BaseSparkSqlParserListener) ExitAnalyzeTables(ctx *AnalyzeTablesContext) {}

// EnterAddTableColumns is called when production addTableColumns is entered.
func (s *BaseSparkSqlParserListener) EnterAddTableColumns(ctx *AddTableColumnsContext) {}

// ExitAddTableColumns is called when production addTableColumns is exited.
func (s *BaseSparkSqlParserListener) ExitAddTableColumns(ctx *AddTableColumnsContext) {}

// EnterRenameTableColumn is called when production renameTableColumn is entered.
func (s *BaseSparkSqlParserListener) EnterRenameTableColumn(ctx *RenameTableColumnContext) {}

// ExitRenameTableColumn is called when production renameTableColumn is exited.
func (s *BaseSparkSqlParserListener) ExitRenameTableColumn(ctx *RenameTableColumnContext) {}

// EnterDropTableColumns is called when production dropTableColumns is entered.
func (s *BaseSparkSqlParserListener) EnterDropTableColumns(ctx *DropTableColumnsContext) {}

// ExitDropTableColumns is called when production dropTableColumns is exited.
func (s *BaseSparkSqlParserListener) ExitDropTableColumns(ctx *DropTableColumnsContext) {}

// EnterRenameTable is called when production renameTable is entered.
func (s *BaseSparkSqlParserListener) EnterRenameTable(ctx *RenameTableContext) {}

// ExitRenameTable is called when production renameTable is exited.
func (s *BaseSparkSqlParserListener) ExitRenameTable(ctx *RenameTableContext) {}

// EnterSetTableProperties is called when production setTableProperties is entered.
func (s *BaseSparkSqlParserListener) EnterSetTableProperties(ctx *SetTablePropertiesContext) {}

// ExitSetTableProperties is called when production setTableProperties is exited.
func (s *BaseSparkSqlParserListener) ExitSetTableProperties(ctx *SetTablePropertiesContext) {}

// EnterUnsetTableProperties is called when production unsetTableProperties is entered.
func (s *BaseSparkSqlParserListener) EnterUnsetTableProperties(ctx *UnsetTablePropertiesContext) {}

// ExitUnsetTableProperties is called when production unsetTableProperties is exited.
func (s *BaseSparkSqlParserListener) ExitUnsetTableProperties(ctx *UnsetTablePropertiesContext) {}

// EnterAlterTableAlterColumn is called when production alterTableAlterColumn is entered.
func (s *BaseSparkSqlParserListener) EnterAlterTableAlterColumn(ctx *AlterTableAlterColumnContext) {}

// ExitAlterTableAlterColumn is called when production alterTableAlterColumn is exited.
func (s *BaseSparkSqlParserListener) ExitAlterTableAlterColumn(ctx *AlterTableAlterColumnContext) {}

// EnterHiveChangeColumn is called when production hiveChangeColumn is entered.
func (s *BaseSparkSqlParserListener) EnterHiveChangeColumn(ctx *HiveChangeColumnContext) {}

// ExitHiveChangeColumn is called when production hiveChangeColumn is exited.
func (s *BaseSparkSqlParserListener) ExitHiveChangeColumn(ctx *HiveChangeColumnContext) {}

// EnterHiveReplaceColumns is called when production hiveReplaceColumns is entered.
func (s *BaseSparkSqlParserListener) EnterHiveReplaceColumns(ctx *HiveReplaceColumnsContext) {}

// ExitHiveReplaceColumns is called when production hiveReplaceColumns is exited.
func (s *BaseSparkSqlParserListener) ExitHiveReplaceColumns(ctx *HiveReplaceColumnsContext) {}

// EnterSetTableSerDe is called when production setTableSerDe is entered.
func (s *BaseSparkSqlParserListener) EnterSetTableSerDe(ctx *SetTableSerDeContext) {}

// ExitSetTableSerDe is called when production setTableSerDe is exited.
func (s *BaseSparkSqlParserListener) ExitSetTableSerDe(ctx *SetTableSerDeContext) {}

// EnterAddTablePartition is called when production addTablePartition is entered.
func (s *BaseSparkSqlParserListener) EnterAddTablePartition(ctx *AddTablePartitionContext) {}

// ExitAddTablePartition is called when production addTablePartition is exited.
func (s *BaseSparkSqlParserListener) ExitAddTablePartition(ctx *AddTablePartitionContext) {}

// EnterRenameTablePartition is called when production renameTablePartition is entered.
func (s *BaseSparkSqlParserListener) EnterRenameTablePartition(ctx *RenameTablePartitionContext) {}

// ExitRenameTablePartition is called when production renameTablePartition is exited.
func (s *BaseSparkSqlParserListener) ExitRenameTablePartition(ctx *RenameTablePartitionContext) {}

// EnterDropTablePartitions is called when production dropTablePartitions is entered.
func (s *BaseSparkSqlParserListener) EnterDropTablePartitions(ctx *DropTablePartitionsContext) {}

// ExitDropTablePartitions is called when production dropTablePartitions is exited.
func (s *BaseSparkSqlParserListener) ExitDropTablePartitions(ctx *DropTablePartitionsContext) {}

// EnterSetTableLocation is called when production setTableLocation is entered.
func (s *BaseSparkSqlParserListener) EnterSetTableLocation(ctx *SetTableLocationContext) {}

// ExitSetTableLocation is called when production setTableLocation is exited.
func (s *BaseSparkSqlParserListener) ExitSetTableLocation(ctx *SetTableLocationContext) {}

// EnterRecoverPartitions is called when production recoverPartitions is entered.
func (s *BaseSparkSqlParserListener) EnterRecoverPartitions(ctx *RecoverPartitionsContext) {}

// ExitRecoverPartitions is called when production recoverPartitions is exited.
func (s *BaseSparkSqlParserListener) ExitRecoverPartitions(ctx *RecoverPartitionsContext) {}

// EnterAlterClusterBy is called when production alterClusterBy is entered.
func (s *BaseSparkSqlParserListener) EnterAlterClusterBy(ctx *AlterClusterByContext) {}

// ExitAlterClusterBy is called when production alterClusterBy is exited.
func (s *BaseSparkSqlParserListener) ExitAlterClusterBy(ctx *AlterClusterByContext) {}

// EnterAlterTableCollation is called when production alterTableCollation is entered.
func (s *BaseSparkSqlParserListener) EnterAlterTableCollation(ctx *AlterTableCollationContext) {}

// ExitAlterTableCollation is called when production alterTableCollation is exited.
func (s *BaseSparkSqlParserListener) ExitAlterTableCollation(ctx *AlterTableCollationContext) {}

// EnterAddTableConstraint is called when production addTableConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterAddTableConstraint(ctx *AddTableConstraintContext) {}

// ExitAddTableConstraint is called when production addTableConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitAddTableConstraint(ctx *AddTableConstraintContext) {}

// EnterDropTableConstraint is called when production dropTableConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterDropTableConstraint(ctx *DropTableConstraintContext) {}

// ExitDropTableConstraint is called when production dropTableConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitDropTableConstraint(ctx *DropTableConstraintContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseSparkSqlParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseSparkSqlParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseSparkSqlParserListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseSparkSqlParserListener) ExitDropView(ctx *DropViewContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseSparkSqlParserListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseSparkSqlParserListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterCreateMetricView is called when production createMetricView is entered.
func (s *BaseSparkSqlParserListener) EnterCreateMetricView(ctx *CreateMetricViewContext) {}

// ExitCreateMetricView is called when production createMetricView is exited.
func (s *BaseSparkSqlParserListener) ExitCreateMetricView(ctx *CreateMetricViewContext) {}

// EnterCreateTempViewUsing is called when production createTempViewUsing is entered.
func (s *BaseSparkSqlParserListener) EnterCreateTempViewUsing(ctx *CreateTempViewUsingContext) {}

// ExitCreateTempViewUsing is called when production createTempViewUsing is exited.
func (s *BaseSparkSqlParserListener) ExitCreateTempViewUsing(ctx *CreateTempViewUsingContext) {}

// EnterAlterViewQuery is called when production alterViewQuery is entered.
func (s *BaseSparkSqlParserListener) EnterAlterViewQuery(ctx *AlterViewQueryContext) {}

// ExitAlterViewQuery is called when production alterViewQuery is exited.
func (s *BaseSparkSqlParserListener) ExitAlterViewQuery(ctx *AlterViewQueryContext) {}

// EnterAlterViewSchemaBinding is called when production alterViewSchemaBinding is entered.
func (s *BaseSparkSqlParserListener) EnterAlterViewSchemaBinding(ctx *AlterViewSchemaBindingContext) {}

// ExitAlterViewSchemaBinding is called when production alterViewSchemaBinding is exited.
func (s *BaseSparkSqlParserListener) ExitAlterViewSchemaBinding(ctx *AlterViewSchemaBindingContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseSparkSqlParserListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseSparkSqlParserListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterCreateUserDefinedFunction is called when production createUserDefinedFunction is entered.
func (s *BaseSparkSqlParserListener) EnterCreateUserDefinedFunction(ctx *CreateUserDefinedFunctionContext) {}

// ExitCreateUserDefinedFunction is called when production createUserDefinedFunction is exited.
func (s *BaseSparkSqlParserListener) ExitCreateUserDefinedFunction(ctx *CreateUserDefinedFunctionContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseSparkSqlParserListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseSparkSqlParserListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterCreateVariable is called when production createVariable is entered.
func (s *BaseSparkSqlParserListener) EnterCreateVariable(ctx *CreateVariableContext) {}

// ExitCreateVariable is called when production createVariable is exited.
func (s *BaseSparkSqlParserListener) ExitCreateVariable(ctx *CreateVariableContext) {}

// EnterDropVariable is called when production dropVariable is entered.
func (s *BaseSparkSqlParserListener) EnterDropVariable(ctx *DropVariableContext) {}

// ExitDropVariable is called when production dropVariable is exited.
func (s *BaseSparkSqlParserListener) ExitDropVariable(ctx *DropVariableContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseSparkSqlParserListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseSparkSqlParserListener) ExitExplain(ctx *ExplainContext) {}

// EnterShowTables is called when production showTables is entered.
func (s *BaseSparkSqlParserListener) EnterShowTables(ctx *ShowTablesContext) {}

// ExitShowTables is called when production showTables is exited.
func (s *BaseSparkSqlParserListener) ExitShowTables(ctx *ShowTablesContext) {}

// EnterShowTableExtended is called when production showTableExtended is entered.
func (s *BaseSparkSqlParserListener) EnterShowTableExtended(ctx *ShowTableExtendedContext) {}

// ExitShowTableExtended is called when production showTableExtended is exited.
func (s *BaseSparkSqlParserListener) ExitShowTableExtended(ctx *ShowTableExtendedContext) {}

// EnterShowTblProperties is called when production showTblProperties is entered.
func (s *BaseSparkSqlParserListener) EnterShowTblProperties(ctx *ShowTblPropertiesContext) {}

// ExitShowTblProperties is called when production showTblProperties is exited.
func (s *BaseSparkSqlParserListener) ExitShowTblProperties(ctx *ShowTblPropertiesContext) {}

// EnterShowColumns is called when production showColumns is entered.
func (s *BaseSparkSqlParserListener) EnterShowColumns(ctx *ShowColumnsContext) {}

// ExitShowColumns is called when production showColumns is exited.
func (s *BaseSparkSqlParserListener) ExitShowColumns(ctx *ShowColumnsContext) {}

// EnterShowViews is called when production showViews is entered.
func (s *BaseSparkSqlParserListener) EnterShowViews(ctx *ShowViewsContext) {}

// ExitShowViews is called when production showViews is exited.
func (s *BaseSparkSqlParserListener) ExitShowViews(ctx *ShowViewsContext) {}

// EnterShowPartitions is called when production showPartitions is entered.
func (s *BaseSparkSqlParserListener) EnterShowPartitions(ctx *ShowPartitionsContext) {}

// ExitShowPartitions is called when production showPartitions is exited.
func (s *BaseSparkSqlParserListener) ExitShowPartitions(ctx *ShowPartitionsContext) {}

// EnterShowFunctions is called when production showFunctions is entered.
func (s *BaseSparkSqlParserListener) EnterShowFunctions(ctx *ShowFunctionsContext) {}

// ExitShowFunctions is called when production showFunctions is exited.
func (s *BaseSparkSqlParserListener) ExitShowFunctions(ctx *ShowFunctionsContext) {}

// EnterShowProcedures is called when production showProcedures is entered.
func (s *BaseSparkSqlParserListener) EnterShowProcedures(ctx *ShowProceduresContext) {}

// ExitShowProcedures is called when production showProcedures is exited.
func (s *BaseSparkSqlParserListener) ExitShowProcedures(ctx *ShowProceduresContext) {}

// EnterShowCreateTable is called when production showCreateTable is entered.
func (s *BaseSparkSqlParserListener) EnterShowCreateTable(ctx *ShowCreateTableContext) {}

// ExitShowCreateTable is called when production showCreateTable is exited.
func (s *BaseSparkSqlParserListener) ExitShowCreateTable(ctx *ShowCreateTableContext) {}

// EnterShowCurrentNamespace is called when production showCurrentNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterShowCurrentNamespace(ctx *ShowCurrentNamespaceContext) {}

// ExitShowCurrentNamespace is called when production showCurrentNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitShowCurrentNamespace(ctx *ShowCurrentNamespaceContext) {}

// EnterShowCatalogs is called when production showCatalogs is entered.
func (s *BaseSparkSqlParserListener) EnterShowCatalogs(ctx *ShowCatalogsContext) {}

// ExitShowCatalogs is called when production showCatalogs is exited.
func (s *BaseSparkSqlParserListener) ExitShowCatalogs(ctx *ShowCatalogsContext) {}

// EnterDescribeFunction is called when production describeFunction is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeFunction(ctx *DescribeFunctionContext) {}

// ExitDescribeFunction is called when production describeFunction is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeFunction(ctx *DescribeFunctionContext) {}

// EnterDescribeProcedure is called when production describeProcedure is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeProcedure(ctx *DescribeProcedureContext) {}

// ExitDescribeProcedure is called when production describeProcedure is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeProcedure(ctx *DescribeProcedureContext) {}

// EnterDescribeNamespace is called when production describeNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeNamespace(ctx *DescribeNamespaceContext) {}

// ExitDescribeNamespace is called when production describeNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeNamespace(ctx *DescribeNamespaceContext) {}

// EnterDescribeRelation is called when production describeRelation is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeRelation(ctx *DescribeRelationContext) {}

// ExitDescribeRelation is called when production describeRelation is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeRelation(ctx *DescribeRelationContext) {}

// EnterDescribeQuery is called when production describeQuery is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeQuery(ctx *DescribeQueryContext) {}

// ExitDescribeQuery is called when production describeQuery is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeQuery(ctx *DescribeQueryContext) {}

// EnterCommentNamespace is called when production commentNamespace is entered.
func (s *BaseSparkSqlParserListener) EnterCommentNamespace(ctx *CommentNamespaceContext) {}

// ExitCommentNamespace is called when production commentNamespace is exited.
func (s *BaseSparkSqlParserListener) ExitCommentNamespace(ctx *CommentNamespaceContext) {}

// EnterCommentTable is called when production commentTable is entered.
func (s *BaseSparkSqlParserListener) EnterCommentTable(ctx *CommentTableContext) {}

// ExitCommentTable is called when production commentTable is exited.
func (s *BaseSparkSqlParserListener) ExitCommentTable(ctx *CommentTableContext) {}

// EnterRefreshTable is called when production refreshTable is entered.
func (s *BaseSparkSqlParserListener) EnterRefreshTable(ctx *RefreshTableContext) {}

// ExitRefreshTable is called when production refreshTable is exited.
func (s *BaseSparkSqlParserListener) ExitRefreshTable(ctx *RefreshTableContext) {}

// EnterRefreshFunction is called when production refreshFunction is entered.
func (s *BaseSparkSqlParserListener) EnterRefreshFunction(ctx *RefreshFunctionContext) {}

// ExitRefreshFunction is called when production refreshFunction is exited.
func (s *BaseSparkSqlParserListener) ExitRefreshFunction(ctx *RefreshFunctionContext) {}

// EnterRefreshResource is called when production refreshResource is entered.
func (s *BaseSparkSqlParserListener) EnterRefreshResource(ctx *RefreshResourceContext) {}

// ExitRefreshResource is called when production refreshResource is exited.
func (s *BaseSparkSqlParserListener) ExitRefreshResource(ctx *RefreshResourceContext) {}

// EnterCacheTable is called when production cacheTable is entered.
func (s *BaseSparkSqlParserListener) EnterCacheTable(ctx *CacheTableContext) {}

// ExitCacheTable is called when production cacheTable is exited.
func (s *BaseSparkSqlParserListener) ExitCacheTable(ctx *CacheTableContext) {}

// EnterUncacheTable is called when production uncacheTable is entered.
func (s *BaseSparkSqlParserListener) EnterUncacheTable(ctx *UncacheTableContext) {}

// ExitUncacheTable is called when production uncacheTable is exited.
func (s *BaseSparkSqlParserListener) ExitUncacheTable(ctx *UncacheTableContext) {}

// EnterClearCache is called when production clearCache is entered.
func (s *BaseSparkSqlParserListener) EnterClearCache(ctx *ClearCacheContext) {}

// ExitClearCache is called when production clearCache is exited.
func (s *BaseSparkSqlParserListener) ExitClearCache(ctx *ClearCacheContext) {}

// EnterLoadData is called when production loadData is entered.
func (s *BaseSparkSqlParserListener) EnterLoadData(ctx *LoadDataContext) {}

// ExitLoadData is called when production loadData is exited.
func (s *BaseSparkSqlParserListener) ExitLoadData(ctx *LoadDataContext) {}

// EnterTruncateTable is called when production truncateTable is entered.
func (s *BaseSparkSqlParserListener) EnterTruncateTable(ctx *TruncateTableContext) {}

// ExitTruncateTable is called when production truncateTable is exited.
func (s *BaseSparkSqlParserListener) ExitTruncateTable(ctx *TruncateTableContext) {}

// EnterRepairTable is called when production repairTable is entered.
func (s *BaseSparkSqlParserListener) EnterRepairTable(ctx *RepairTableContext) {}

// ExitRepairTable is called when production repairTable is exited.
func (s *BaseSparkSqlParserListener) ExitRepairTable(ctx *RepairTableContext) {}

// EnterManageResource is called when production manageResource is entered.
func (s *BaseSparkSqlParserListener) EnterManageResource(ctx *ManageResourceContext) {}

// ExitManageResource is called when production manageResource is exited.
func (s *BaseSparkSqlParserListener) ExitManageResource(ctx *ManageResourceContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseSparkSqlParserListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseSparkSqlParserListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseSparkSqlParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseSparkSqlParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterCall is called when production call is entered.
func (s *BaseSparkSqlParserListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseSparkSqlParserListener) ExitCall(ctx *CallContext) {}

// EnterFailNativeCommand is called when production failNativeCommand is entered.
func (s *BaseSparkSqlParserListener) EnterFailNativeCommand(ctx *FailNativeCommandContext) {}

// ExitFailNativeCommand is called when production failNativeCommand is exited.
func (s *BaseSparkSqlParserListener) ExitFailNativeCommand(ctx *FailNativeCommandContext) {}

// EnterCreatePipelineDataset is called when production createPipelineDataset is entered.
func (s *BaseSparkSqlParserListener) EnterCreatePipelineDataset(ctx *CreatePipelineDatasetContext) {}

// ExitCreatePipelineDataset is called when production createPipelineDataset is exited.
func (s *BaseSparkSqlParserListener) ExitCreatePipelineDataset(ctx *CreatePipelineDatasetContext) {}

// EnterCreatePipelineInsertIntoFlow is called when production createPipelineInsertIntoFlow is entered.
func (s *BaseSparkSqlParserListener) EnterCreatePipelineInsertIntoFlow(ctx *CreatePipelineInsertIntoFlowContext) {}

// ExitCreatePipelineInsertIntoFlow is called when production createPipelineInsertIntoFlow is exited.
func (s *BaseSparkSqlParserListener) ExitCreatePipelineInsertIntoFlow(ctx *CreatePipelineInsertIntoFlowContext) {}

// EnterMaterializedView is called when production materializedView is entered.
func (s *BaseSparkSqlParserListener) EnterMaterializedView(ctx *MaterializedViewContext) {}

// ExitMaterializedView is called when production materializedView is exited.
func (s *BaseSparkSqlParserListener) ExitMaterializedView(ctx *MaterializedViewContext) {}

// EnterStreamingTable is called when production streamingTable is entered.
func (s *BaseSparkSqlParserListener) EnterStreamingTable(ctx *StreamingTableContext) {}

// ExitStreamingTable is called when production streamingTable is exited.
func (s *BaseSparkSqlParserListener) ExitStreamingTable(ctx *StreamingTableContext) {}

// EnterCreatePipelineDatasetHeader is called when production createPipelineDatasetHeader is entered.
func (s *BaseSparkSqlParserListener) EnterCreatePipelineDatasetHeader(ctx *CreatePipelineDatasetHeaderContext) {}

// ExitCreatePipelineDatasetHeader is called when production createPipelineDatasetHeader is exited.
func (s *BaseSparkSqlParserListener) ExitCreatePipelineDatasetHeader(ctx *CreatePipelineDatasetHeaderContext) {}

// EnterStreamTableName is called when production streamTableName is entered.
func (s *BaseSparkSqlParserListener) EnterStreamTableName(ctx *StreamTableNameContext) {}

// ExitStreamTableName is called when production streamTableName is exited.
func (s *BaseSparkSqlParserListener) ExitStreamTableName(ctx *StreamTableNameContext) {}

// EnterFailSetRole is called when production failSetRole is entered.
func (s *BaseSparkSqlParserListener) EnterFailSetRole(ctx *FailSetRoleContext) {}

// ExitFailSetRole is called when production failSetRole is exited.
func (s *BaseSparkSqlParserListener) ExitFailSetRole(ctx *FailSetRoleContext) {}

// EnterSetTimeZone is called when production setTimeZone is entered.
func (s *BaseSparkSqlParserListener) EnterSetTimeZone(ctx *SetTimeZoneContext) {}

// ExitSetTimeZone is called when production setTimeZone is exited.
func (s *BaseSparkSqlParserListener) ExitSetTimeZone(ctx *SetTimeZoneContext) {}

// EnterSetVariable is called when production setVariable is entered.
func (s *BaseSparkSqlParserListener) EnterSetVariable(ctx *SetVariableContext) {}

// ExitSetVariable is called when production setVariable is exited.
func (s *BaseSparkSqlParserListener) ExitSetVariable(ctx *SetVariableContext) {}

// EnterSetQuotedConfiguration is called when production setQuotedConfiguration is entered.
func (s *BaseSparkSqlParserListener) EnterSetQuotedConfiguration(ctx *SetQuotedConfigurationContext) {}

// ExitSetQuotedConfiguration is called when production setQuotedConfiguration is exited.
func (s *BaseSparkSqlParserListener) ExitSetQuotedConfiguration(ctx *SetQuotedConfigurationContext) {}

// EnterSetConfiguration is called when production setConfiguration is entered.
func (s *BaseSparkSqlParserListener) EnterSetConfiguration(ctx *SetConfigurationContext) {}

// ExitSetConfiguration is called when production setConfiguration is exited.
func (s *BaseSparkSqlParserListener) ExitSetConfiguration(ctx *SetConfigurationContext) {}

// EnterResetQuotedConfiguration is called when production resetQuotedConfiguration is entered.
func (s *BaseSparkSqlParserListener) EnterResetQuotedConfiguration(ctx *ResetQuotedConfigurationContext) {}

// ExitResetQuotedConfiguration is called when production resetQuotedConfiguration is exited.
func (s *BaseSparkSqlParserListener) ExitResetQuotedConfiguration(ctx *ResetQuotedConfigurationContext) {}

// EnterResetConfiguration is called when production resetConfiguration is entered.
func (s *BaseSparkSqlParserListener) EnterResetConfiguration(ctx *ResetConfigurationContext) {}

// ExitResetConfiguration is called when production resetConfiguration is exited.
func (s *BaseSparkSqlParserListener) ExitResetConfiguration(ctx *ResetConfigurationContext) {}

// EnterExecuteImmediate is called when production executeImmediate is entered.
func (s *BaseSparkSqlParserListener) EnterExecuteImmediate(ctx *ExecuteImmediateContext) {}

// ExitExecuteImmediate is called when production executeImmediate is exited.
func (s *BaseSparkSqlParserListener) ExitExecuteImmediate(ctx *ExecuteImmediateContext) {}

// EnterExecuteImmediateUsing is called when production executeImmediateUsing is entered.
func (s *BaseSparkSqlParserListener) EnterExecuteImmediateUsing(ctx *ExecuteImmediateUsingContext) {}

// ExitExecuteImmediateUsing is called when production executeImmediateUsing is exited.
func (s *BaseSparkSqlParserListener) ExitExecuteImmediateUsing(ctx *ExecuteImmediateUsingContext) {}

// EnterTimezone is called when production timezone is entered.
func (s *BaseSparkSqlParserListener) EnterTimezone(ctx *TimezoneContext) {}

// ExitTimezone is called when production timezone is exited.
func (s *BaseSparkSqlParserListener) ExitTimezone(ctx *TimezoneContext) {}

// EnterConfigKey is called when production configKey is entered.
func (s *BaseSparkSqlParserListener) EnterConfigKey(ctx *ConfigKeyContext) {}

// ExitConfigKey is called when production configKey is exited.
func (s *BaseSparkSqlParserListener) ExitConfigKey(ctx *ConfigKeyContext) {}

// EnterConfigValue is called when production configValue is entered.
func (s *BaseSparkSqlParserListener) EnterConfigValue(ctx *ConfigValueContext) {}

// ExitConfigValue is called when production configValue is exited.
func (s *BaseSparkSqlParserListener) ExitConfigValue(ctx *ConfigValueContext) {}

// EnterUnsupportedHiveNativeCommands is called when production unsupportedHiveNativeCommands is entered.
func (s *BaseSparkSqlParserListener) EnterUnsupportedHiveNativeCommands(ctx *UnsupportedHiveNativeCommandsContext) {}

// ExitUnsupportedHiveNativeCommands is called when production unsupportedHiveNativeCommands is exited.
func (s *BaseSparkSqlParserListener) ExitUnsupportedHiveNativeCommands(ctx *UnsupportedHiveNativeCommandsContext) {}

// EnterCreateTableHeader is called when production createTableHeader is entered.
func (s *BaseSparkSqlParserListener) EnterCreateTableHeader(ctx *CreateTableHeaderContext) {}

// ExitCreateTableHeader is called when production createTableHeader is exited.
func (s *BaseSparkSqlParserListener) ExitCreateTableHeader(ctx *CreateTableHeaderContext) {}

// EnterReplaceTableHeader is called when production replaceTableHeader is entered.
func (s *BaseSparkSqlParserListener) EnterReplaceTableHeader(ctx *ReplaceTableHeaderContext) {}

// ExitReplaceTableHeader is called when production replaceTableHeader is exited.
func (s *BaseSparkSqlParserListener) ExitReplaceTableHeader(ctx *ReplaceTableHeaderContext) {}

// EnterClusterBySpec is called when production clusterBySpec is entered.
func (s *BaseSparkSqlParserListener) EnterClusterBySpec(ctx *ClusterBySpecContext) {}

// ExitClusterBySpec is called when production clusterBySpec is exited.
func (s *BaseSparkSqlParserListener) ExitClusterBySpec(ctx *ClusterBySpecContext) {}

// EnterBucketSpec is called when production bucketSpec is entered.
func (s *BaseSparkSqlParserListener) EnterBucketSpec(ctx *BucketSpecContext) {}

// ExitBucketSpec is called when production bucketSpec is exited.
func (s *BaseSparkSqlParserListener) ExitBucketSpec(ctx *BucketSpecContext) {}

// EnterSkewSpec is called when production skewSpec is entered.
func (s *BaseSparkSqlParserListener) EnterSkewSpec(ctx *SkewSpecContext) {}

// ExitSkewSpec is called when production skewSpec is exited.
func (s *BaseSparkSqlParserListener) ExitSkewSpec(ctx *SkewSpecContext) {}

// EnterLocationSpec is called when production locationSpec is entered.
func (s *BaseSparkSqlParserListener) EnterLocationSpec(ctx *LocationSpecContext) {}

// ExitLocationSpec is called when production locationSpec is exited.
func (s *BaseSparkSqlParserListener) ExitLocationSpec(ctx *LocationSpecContext) {}

// EnterSchemaBinding is called when production schemaBinding is entered.
func (s *BaseSparkSqlParserListener) EnterSchemaBinding(ctx *SchemaBindingContext) {}

// ExitSchemaBinding is called when production schemaBinding is exited.
func (s *BaseSparkSqlParserListener) ExitSchemaBinding(ctx *SchemaBindingContext) {}

// EnterCommentSpec is called when production commentSpec is entered.
func (s *BaseSparkSqlParserListener) EnterCommentSpec(ctx *CommentSpecContext) {}

// ExitCommentSpec is called when production commentSpec is exited.
func (s *BaseSparkSqlParserListener) ExitCommentSpec(ctx *CommentSpecContext) {}

// EnterSingleQuery is called when production singleQuery is entered.
func (s *BaseSparkSqlParserListener) EnterSingleQuery(ctx *SingleQueryContext) {}

// ExitSingleQuery is called when production singleQuery is exited.
func (s *BaseSparkSqlParserListener) ExitSingleQuery(ctx *SingleQueryContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSparkSqlParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSparkSqlParserListener) ExitQuery(ctx *QueryContext) {}

// EnterInsertOverwriteTable is called when production insertOverwriteTable is entered.
func (s *BaseSparkSqlParserListener) EnterInsertOverwriteTable(ctx *InsertOverwriteTableContext) {}

// ExitInsertOverwriteTable is called when production insertOverwriteTable is exited.
func (s *BaseSparkSqlParserListener) ExitInsertOverwriteTable(ctx *InsertOverwriteTableContext) {}

// EnterInsertIntoTable is called when production insertIntoTable is entered.
func (s *BaseSparkSqlParserListener) EnterInsertIntoTable(ctx *InsertIntoTableContext) {}

// ExitInsertIntoTable is called when production insertIntoTable is exited.
func (s *BaseSparkSqlParserListener) ExitInsertIntoTable(ctx *InsertIntoTableContext) {}

// EnterInsertIntoReplaceWhere is called when production insertIntoReplaceWhere is entered.
func (s *BaseSparkSqlParserListener) EnterInsertIntoReplaceWhere(ctx *InsertIntoReplaceWhereContext) {}

// ExitInsertIntoReplaceWhere is called when production insertIntoReplaceWhere is exited.
func (s *BaseSparkSqlParserListener) ExitInsertIntoReplaceWhere(ctx *InsertIntoReplaceWhereContext) {}

// EnterInsertOverwriteHiveDir is called when production insertOverwriteHiveDir is entered.
func (s *BaseSparkSqlParserListener) EnterInsertOverwriteHiveDir(ctx *InsertOverwriteHiveDirContext) {}

// ExitInsertOverwriteHiveDir is called when production insertOverwriteHiveDir is exited.
func (s *BaseSparkSqlParserListener) ExitInsertOverwriteHiveDir(ctx *InsertOverwriteHiveDirContext) {}

// EnterInsertOverwriteDir is called when production insertOverwriteDir is entered.
func (s *BaseSparkSqlParserListener) EnterInsertOverwriteDir(ctx *InsertOverwriteDirContext) {}

// ExitInsertOverwriteDir is called when production insertOverwriteDir is exited.
func (s *BaseSparkSqlParserListener) ExitInsertOverwriteDir(ctx *InsertOverwriteDirContext) {}

// EnterPartitionSpecLocation is called when production partitionSpecLocation is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionSpecLocation(ctx *PartitionSpecLocationContext) {}

// ExitPartitionSpecLocation is called when production partitionSpecLocation is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionSpecLocation(ctx *PartitionSpecLocationContext) {}

// EnterPartitionSpec is called when production partitionSpec is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionSpec(ctx *PartitionSpecContext) {}

// ExitPartitionSpec is called when production partitionSpec is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionSpec(ctx *PartitionSpecContext) {}

// EnterPartitionVal is called when production partitionVal is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionVal(ctx *PartitionValContext) {}

// ExitPartitionVal is called when production partitionVal is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionVal(ctx *PartitionValContext) {}

// EnterCreatePipelineFlowHeader is called when production createPipelineFlowHeader is entered.
func (s *BaseSparkSqlParserListener) EnterCreatePipelineFlowHeader(ctx *CreatePipelineFlowHeaderContext) {}

// ExitCreatePipelineFlowHeader is called when production createPipelineFlowHeader is exited.
func (s *BaseSparkSqlParserListener) ExitCreatePipelineFlowHeader(ctx *CreatePipelineFlowHeaderContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseSparkSqlParserListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseSparkSqlParserListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterNamespaces is called when production namespaces is entered.
func (s *BaseSparkSqlParserListener) EnterNamespaces(ctx *NamespacesContext) {}

// ExitNamespaces is called when production namespaces is exited.
func (s *BaseSparkSqlParserListener) ExitNamespaces(ctx *NamespacesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSparkSqlParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSparkSqlParserListener) ExitVariable(ctx *VariableContext) {}

// EnterDescribeFuncName is called when production describeFuncName is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeFuncName(ctx *DescribeFuncNameContext) {}

// ExitDescribeFuncName is called when production describeFuncName is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeFuncName(ctx *DescribeFuncNameContext) {}

// EnterDescribeColName is called when production describeColName is entered.
func (s *BaseSparkSqlParserListener) EnterDescribeColName(ctx *DescribeColNameContext) {}

// ExitDescribeColName is called when production describeColName is exited.
func (s *BaseSparkSqlParserListener) ExitDescribeColName(ctx *DescribeColNameContext) {}

// EnterCtes is called when production ctes is entered.
func (s *BaseSparkSqlParserListener) EnterCtes(ctx *CtesContext) {}

// ExitCtes is called when production ctes is exited.
func (s *BaseSparkSqlParserListener) ExitCtes(ctx *CtesContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseSparkSqlParserListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseSparkSqlParserListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterTableProvider is called when production tableProvider is entered.
func (s *BaseSparkSqlParserListener) EnterTableProvider(ctx *TableProviderContext) {}

// ExitTableProvider is called when production tableProvider is exited.
func (s *BaseSparkSqlParserListener) ExitTableProvider(ctx *TableProviderContext) {}

// EnterCreateTableClauses is called when production createTableClauses is entered.
func (s *BaseSparkSqlParserListener) EnterCreateTableClauses(ctx *CreateTableClausesContext) {}

// ExitCreateTableClauses is called when production createTableClauses is exited.
func (s *BaseSparkSqlParserListener) ExitCreateTableClauses(ctx *CreateTableClausesContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterPropertyWithKeyAndEquals is called when production propertyWithKeyAndEquals is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyWithKeyAndEquals(ctx *PropertyWithKeyAndEqualsContext) {}

// ExitPropertyWithKeyAndEquals is called when production propertyWithKeyAndEquals is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyWithKeyAndEquals(ctx *PropertyWithKeyAndEqualsContext) {}

// EnterPropertyWithKeyNoEquals is called when production propertyWithKeyNoEquals is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyWithKeyNoEquals(ctx *PropertyWithKeyNoEqualsContext) {}

// ExitPropertyWithKeyNoEquals is called when production propertyWithKeyNoEquals is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyWithKeyNoEquals(ctx *PropertyWithKeyNoEqualsContext) {}

// EnterPropertyKey is called when production propertyKey is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyKey(ctx *PropertyKeyContext) {}

// ExitPropertyKey is called when production propertyKey is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyKey(ctx *PropertyKeyContext) {}

// EnterPropertyKeyOrStringLit is called when production propertyKeyOrStringLit is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyKeyOrStringLit(ctx *PropertyKeyOrStringLitContext) {}

// ExitPropertyKeyOrStringLit is called when production propertyKeyOrStringLit is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyKeyOrStringLit(ctx *PropertyKeyOrStringLitContext) {}

// EnterPropertyKeyOrStringLitNoCoalesce is called when production propertyKeyOrStringLitNoCoalesce is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyKeyOrStringLitNoCoalesce(ctx *PropertyKeyOrStringLitNoCoalesceContext) {}

// ExitPropertyKeyOrStringLitNoCoalesce is called when production propertyKeyOrStringLitNoCoalesce is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyKeyOrStringLitNoCoalesce(ctx *PropertyKeyOrStringLitNoCoalesceContext) {}

// EnterPropertyValue is called when production propertyValue is entered.
func (s *BaseSparkSqlParserListener) EnterPropertyValue(ctx *PropertyValueContext) {}

// ExitPropertyValue is called when production propertyValue is exited.
func (s *BaseSparkSqlParserListener) ExitPropertyValue(ctx *PropertyValueContext) {}

// EnterExpressionPropertyList is called when production expressionPropertyList is entered.
func (s *BaseSparkSqlParserListener) EnterExpressionPropertyList(ctx *ExpressionPropertyListContext) {}

// ExitExpressionPropertyList is called when production expressionPropertyList is exited.
func (s *BaseSparkSqlParserListener) ExitExpressionPropertyList(ctx *ExpressionPropertyListContext) {}

// EnterExpressionPropertyWithKeyAndEquals is called when production expressionPropertyWithKeyAndEquals is entered.
func (s *BaseSparkSqlParserListener) EnterExpressionPropertyWithKeyAndEquals(ctx *ExpressionPropertyWithKeyAndEqualsContext) {}

// ExitExpressionPropertyWithKeyAndEquals is called when production expressionPropertyWithKeyAndEquals is exited.
func (s *BaseSparkSqlParserListener) ExitExpressionPropertyWithKeyAndEquals(ctx *ExpressionPropertyWithKeyAndEqualsContext) {}

// EnterExpressionPropertyWithKeyNoEquals is called when production expressionPropertyWithKeyNoEquals is entered.
func (s *BaseSparkSqlParserListener) EnterExpressionPropertyWithKeyNoEquals(ctx *ExpressionPropertyWithKeyNoEqualsContext) {}

// ExitExpressionPropertyWithKeyNoEquals is called when production expressionPropertyWithKeyNoEquals is exited.
func (s *BaseSparkSqlParserListener) ExitExpressionPropertyWithKeyNoEquals(ctx *ExpressionPropertyWithKeyNoEqualsContext) {}

// EnterConstantList is called when production constantList is entered.
func (s *BaseSparkSqlParserListener) EnterConstantList(ctx *ConstantListContext) {}

// ExitConstantList is called when production constantList is exited.
func (s *BaseSparkSqlParserListener) ExitConstantList(ctx *ConstantListContext) {}

// EnterNestedConstantList is called when production nestedConstantList is entered.
func (s *BaseSparkSqlParserListener) EnterNestedConstantList(ctx *NestedConstantListContext) {}

// ExitNestedConstantList is called when production nestedConstantList is exited.
func (s *BaseSparkSqlParserListener) ExitNestedConstantList(ctx *NestedConstantListContext) {}

// EnterCreateFileFormat is called when production createFileFormat is entered.
func (s *BaseSparkSqlParserListener) EnterCreateFileFormat(ctx *CreateFileFormatContext) {}

// ExitCreateFileFormat is called when production createFileFormat is exited.
func (s *BaseSparkSqlParserListener) ExitCreateFileFormat(ctx *CreateFileFormatContext) {}

// EnterTableFileFormat is called when production tableFileFormat is entered.
func (s *BaseSparkSqlParserListener) EnterTableFileFormat(ctx *TableFileFormatContext) {}

// ExitTableFileFormat is called when production tableFileFormat is exited.
func (s *BaseSparkSqlParserListener) ExitTableFileFormat(ctx *TableFileFormatContext) {}

// EnterGenericFileFormat is called when production genericFileFormat is entered.
func (s *BaseSparkSqlParserListener) EnterGenericFileFormat(ctx *GenericFileFormatContext) {}

// ExitGenericFileFormat is called when production genericFileFormat is exited.
func (s *BaseSparkSqlParserListener) ExitGenericFileFormat(ctx *GenericFileFormatContext) {}

// EnterStorageHandler is called when production storageHandler is entered.
func (s *BaseSparkSqlParserListener) EnterStorageHandler(ctx *StorageHandlerContext) {}

// ExitStorageHandler is called when production storageHandler is exited.
func (s *BaseSparkSqlParserListener) ExitStorageHandler(ctx *StorageHandlerContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseSparkSqlParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseSparkSqlParserListener) ExitResource(ctx *ResourceContext) {}

// EnterSingleInsertQuery is called when production singleInsertQuery is entered.
func (s *BaseSparkSqlParserListener) EnterSingleInsertQuery(ctx *SingleInsertQueryContext) {}

// ExitSingleInsertQuery is called when production singleInsertQuery is exited.
func (s *BaseSparkSqlParserListener) ExitSingleInsertQuery(ctx *SingleInsertQueryContext) {}

// EnterMultiInsertQuery is called when production multiInsertQuery is entered.
func (s *BaseSparkSqlParserListener) EnterMultiInsertQuery(ctx *MultiInsertQueryContext) {}

// ExitMultiInsertQuery is called when production multiInsertQuery is exited.
func (s *BaseSparkSqlParserListener) ExitMultiInsertQuery(ctx *MultiInsertQueryContext) {}

// EnterDeleteFromTable is called when production deleteFromTable is entered.
func (s *BaseSparkSqlParserListener) EnterDeleteFromTable(ctx *DeleteFromTableContext) {}

// ExitDeleteFromTable is called when production deleteFromTable is exited.
func (s *BaseSparkSqlParserListener) ExitDeleteFromTable(ctx *DeleteFromTableContext) {}

// EnterUpdateTable is called when production updateTable is entered.
func (s *BaseSparkSqlParserListener) EnterUpdateTable(ctx *UpdateTableContext) {}

// ExitUpdateTable is called when production updateTable is exited.
func (s *BaseSparkSqlParserListener) ExitUpdateTable(ctx *UpdateTableContext) {}

// EnterMergeIntoTable is called when production mergeIntoTable is entered.
func (s *BaseSparkSqlParserListener) EnterMergeIntoTable(ctx *MergeIntoTableContext) {}

// ExitMergeIntoTable is called when production mergeIntoTable is exited.
func (s *BaseSparkSqlParserListener) ExitMergeIntoTable(ctx *MergeIntoTableContext) {}

// EnterIdentifierReference is called when production identifierReference is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierReference(ctx *IdentifierReferenceContext) {}

// ExitIdentifierReference is called when production identifierReference is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierReference(ctx *IdentifierReferenceContext) {}

// EnterCatalogIdentifierReference is called when production catalogIdentifierReference is entered.
func (s *BaseSparkSqlParserListener) EnterCatalogIdentifierReference(ctx *CatalogIdentifierReferenceContext) {}

// ExitCatalogIdentifierReference is called when production catalogIdentifierReference is exited.
func (s *BaseSparkSqlParserListener) ExitCatalogIdentifierReference(ctx *CatalogIdentifierReferenceContext) {}

// EnterQueryOrganization is called when production queryOrganization is entered.
func (s *BaseSparkSqlParserListener) EnterQueryOrganization(ctx *QueryOrganizationContext) {}

// ExitQueryOrganization is called when production queryOrganization is exited.
func (s *BaseSparkSqlParserListener) ExitQueryOrganization(ctx *QueryOrganizationContext) {}

// EnterMultiInsertQueryBody is called when production multiInsertQueryBody is entered.
func (s *BaseSparkSqlParserListener) EnterMultiInsertQueryBody(ctx *MultiInsertQueryBodyContext) {}

// ExitMultiInsertQueryBody is called when production multiInsertQueryBody is exited.
func (s *BaseSparkSqlParserListener) ExitMultiInsertQueryBody(ctx *MultiInsertQueryBodyContext) {}

// EnterOperatorPipeStatement is called when production operatorPipeStatement is entered.
func (s *BaseSparkSqlParserListener) EnterOperatorPipeStatement(ctx *OperatorPipeStatementContext) {}

// ExitOperatorPipeStatement is called when production operatorPipeStatement is exited.
func (s *BaseSparkSqlParserListener) ExitOperatorPipeStatement(ctx *OperatorPipeStatementContext) {}

// EnterQueryTermDefault is called when production queryTermDefault is entered.
func (s *BaseSparkSqlParserListener) EnterQueryTermDefault(ctx *QueryTermDefaultContext) {}

// ExitQueryTermDefault is called when production queryTermDefault is exited.
func (s *BaseSparkSqlParserListener) ExitQueryTermDefault(ctx *QueryTermDefaultContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseSparkSqlParserListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseSparkSqlParserListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseSparkSqlParserListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseSparkSqlParserListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterFromStmt is called when production fromStmt is entered.
func (s *BaseSparkSqlParserListener) EnterFromStmt(ctx *FromStmtContext) {}

// ExitFromStmt is called when production fromStmt is exited.
func (s *BaseSparkSqlParserListener) ExitFromStmt(ctx *FromStmtContext) {}

// EnterTable is called when production table is entered.
func (s *BaseSparkSqlParserListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSparkSqlParserListener) ExitTable(ctx *TableContext) {}

// EnterInlineTableDefault1 is called when production inlineTableDefault1 is entered.
func (s *BaseSparkSqlParserListener) EnterInlineTableDefault1(ctx *InlineTableDefault1Context) {}

// ExitInlineTableDefault1 is called when production inlineTableDefault1 is exited.
func (s *BaseSparkSqlParserListener) ExitInlineTableDefault1(ctx *InlineTableDefault1Context) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseSparkSqlParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseSparkSqlParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseSparkSqlParserListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseSparkSqlParserListener) ExitSortItem(ctx *SortItemContext) {}

// EnterFromStatement is called when production fromStatement is entered.
func (s *BaseSparkSqlParserListener) EnterFromStatement(ctx *FromStatementContext) {}

// ExitFromStatement is called when production fromStatement is exited.
func (s *BaseSparkSqlParserListener) ExitFromStatement(ctx *FromStatementContext) {}

// EnterFromStatementBody is called when production fromStatementBody is entered.
func (s *BaseSparkSqlParserListener) EnterFromStatementBody(ctx *FromStatementBodyContext) {}

// ExitFromStatementBody is called when production fromStatementBody is exited.
func (s *BaseSparkSqlParserListener) ExitFromStatementBody(ctx *FromStatementBodyContext) {}

// EnterTransformQuerySpecification is called when production transformQuerySpecification is entered.
func (s *BaseSparkSqlParserListener) EnterTransformQuerySpecification(ctx *TransformQuerySpecificationContext) {}

// ExitTransformQuerySpecification is called when production transformQuerySpecification is exited.
func (s *BaseSparkSqlParserListener) ExitTransformQuerySpecification(ctx *TransformQuerySpecificationContext) {}

// EnterRegularQuerySpecification is called when production regularQuerySpecification is entered.
func (s *BaseSparkSqlParserListener) EnterRegularQuerySpecification(ctx *RegularQuerySpecificationContext) {}

// ExitRegularQuerySpecification is called when production regularQuerySpecification is exited.
func (s *BaseSparkSqlParserListener) ExitRegularQuerySpecification(ctx *RegularQuerySpecificationContext) {}

// EnterTransformClause is called when production transformClause is entered.
func (s *BaseSparkSqlParserListener) EnterTransformClause(ctx *TransformClauseContext) {}

// ExitTransformClause is called when production transformClause is exited.
func (s *BaseSparkSqlParserListener) ExitTransformClause(ctx *TransformClauseContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseSparkSqlParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseSparkSqlParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSetClause is called when production setClause is entered.
func (s *BaseSparkSqlParserListener) EnterSetClause(ctx *SetClauseContext) {}

// ExitSetClause is called when production setClause is exited.
func (s *BaseSparkSqlParserListener) ExitSetClause(ctx *SetClauseContext) {}

// EnterMatchedClause is called when production matchedClause is entered.
func (s *BaseSparkSqlParserListener) EnterMatchedClause(ctx *MatchedClauseContext) {}

// ExitMatchedClause is called when production matchedClause is exited.
func (s *BaseSparkSqlParserListener) ExitMatchedClause(ctx *MatchedClauseContext) {}

// EnterNotMatchedClause is called when production notMatchedClause is entered.
func (s *BaseSparkSqlParserListener) EnterNotMatchedClause(ctx *NotMatchedClauseContext) {}

// ExitNotMatchedClause is called when production notMatchedClause is exited.
func (s *BaseSparkSqlParserListener) ExitNotMatchedClause(ctx *NotMatchedClauseContext) {}

// EnterNotMatchedBySourceClause is called when production notMatchedBySourceClause is entered.
func (s *BaseSparkSqlParserListener) EnterNotMatchedBySourceClause(ctx *NotMatchedBySourceClauseContext) {}

// ExitNotMatchedBySourceClause is called when production notMatchedBySourceClause is exited.
func (s *BaseSparkSqlParserListener) ExitNotMatchedBySourceClause(ctx *NotMatchedBySourceClauseContext) {}

// EnterMatchedAction is called when production matchedAction is entered.
func (s *BaseSparkSqlParserListener) EnterMatchedAction(ctx *MatchedActionContext) {}

// ExitMatchedAction is called when production matchedAction is exited.
func (s *BaseSparkSqlParserListener) ExitMatchedAction(ctx *MatchedActionContext) {}

// EnterNotMatchedAction is called when production notMatchedAction is entered.
func (s *BaseSparkSqlParserListener) EnterNotMatchedAction(ctx *NotMatchedActionContext) {}

// ExitNotMatchedAction is called when production notMatchedAction is exited.
func (s *BaseSparkSqlParserListener) ExitNotMatchedAction(ctx *NotMatchedActionContext) {}

// EnterNotMatchedBySourceAction is called when production notMatchedBySourceAction is entered.
func (s *BaseSparkSqlParserListener) EnterNotMatchedBySourceAction(ctx *NotMatchedBySourceActionContext) {}

// ExitNotMatchedBySourceAction is called when production notMatchedBySourceAction is exited.
func (s *BaseSparkSqlParserListener) ExitNotMatchedBySourceAction(ctx *NotMatchedBySourceActionContext) {}

// EnterExceptClause is called when production exceptClause is entered.
func (s *BaseSparkSqlParserListener) EnterExceptClause(ctx *ExceptClauseContext) {}

// ExitExceptClause is called when production exceptClause is exited.
func (s *BaseSparkSqlParserListener) ExitExceptClause(ctx *ExceptClauseContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseSparkSqlParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseSparkSqlParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSparkSqlParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSparkSqlParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSparkSqlParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSparkSqlParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSparkSqlParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSparkSqlParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseSparkSqlParserListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseSparkSqlParserListener) ExitHint(ctx *HintContext) {}

// EnterHintStatement is called when production hintStatement is entered.
func (s *BaseSparkSqlParserListener) EnterHintStatement(ctx *HintStatementContext) {}

// ExitHintStatement is called when production hintStatement is exited.
func (s *BaseSparkSqlParserListener) ExitHintStatement(ctx *HintStatementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSparkSqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSparkSqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTemporalClause is called when production temporalClause is entered.
func (s *BaseSparkSqlParserListener) EnterTemporalClause(ctx *TemporalClauseContext) {}

// ExitTemporalClause is called when production temporalClause is exited.
func (s *BaseSparkSqlParserListener) ExitTemporalClause(ctx *TemporalClauseContext) {}

// EnterAggregationClause is called when production aggregationClause is entered.
func (s *BaseSparkSqlParserListener) EnterAggregationClause(ctx *AggregationClauseContext) {}

// ExitAggregationClause is called when production aggregationClause is exited.
func (s *BaseSparkSqlParserListener) ExitAggregationClause(ctx *AggregationClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSparkSqlParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSparkSqlParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupingAnalytics is called when production groupingAnalytics is entered.
func (s *BaseSparkSqlParserListener) EnterGroupingAnalytics(ctx *GroupingAnalyticsContext) {}

// ExitGroupingAnalytics is called when production groupingAnalytics is exited.
func (s *BaseSparkSqlParserListener) ExitGroupingAnalytics(ctx *GroupingAnalyticsContext) {}

// EnterGroupingElement is called when production groupingElement is entered.
func (s *BaseSparkSqlParserListener) EnterGroupingElement(ctx *GroupingElementContext) {}

// ExitGroupingElement is called when production groupingElement is exited.
func (s *BaseSparkSqlParserListener) ExitGroupingElement(ctx *GroupingElementContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseSparkSqlParserListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseSparkSqlParserListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseSparkSqlParserListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseSparkSqlParserListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotColumn is called when production pivotColumn is entered.
func (s *BaseSparkSqlParserListener) EnterPivotColumn(ctx *PivotColumnContext) {}

// ExitPivotColumn is called when production pivotColumn is exited.
func (s *BaseSparkSqlParserListener) ExitPivotColumn(ctx *PivotColumnContext) {}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseSparkSqlParserListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseSparkSqlParserListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterUnpivotClause is called when production unpivotClause is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotClause(ctx *UnpivotClauseContext) {}

// ExitUnpivotClause is called when production unpivotClause is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotClause(ctx *UnpivotClauseContext) {}

// EnterUnpivotNullClause is called when production unpivotNullClause is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotNullClause(ctx *UnpivotNullClauseContext) {}

// ExitUnpivotNullClause is called when production unpivotNullClause is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotNullClause(ctx *UnpivotNullClauseContext) {}

// EnterUnpivotOperator is called when production unpivotOperator is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotOperator(ctx *UnpivotOperatorContext) {}

// ExitUnpivotOperator is called when production unpivotOperator is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotOperator(ctx *UnpivotOperatorContext) {}

// EnterUnpivotSingleValueColumnClause is called when production unpivotSingleValueColumnClause is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotSingleValueColumnClause(ctx *UnpivotSingleValueColumnClauseContext) {}

// ExitUnpivotSingleValueColumnClause is called when production unpivotSingleValueColumnClause is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotSingleValueColumnClause(ctx *UnpivotSingleValueColumnClauseContext) {}

// EnterUnpivotMultiValueColumnClause is called when production unpivotMultiValueColumnClause is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotMultiValueColumnClause(ctx *UnpivotMultiValueColumnClauseContext) {}

// ExitUnpivotMultiValueColumnClause is called when production unpivotMultiValueColumnClause is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotMultiValueColumnClause(ctx *UnpivotMultiValueColumnClauseContext) {}

// EnterUnpivotColumnSet is called when production unpivotColumnSet is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotColumnSet(ctx *UnpivotColumnSetContext) {}

// ExitUnpivotColumnSet is called when production unpivotColumnSet is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotColumnSet(ctx *UnpivotColumnSetContext) {}

// EnterUnpivotValueColumn is called when production unpivotValueColumn is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotValueColumn(ctx *UnpivotValueColumnContext) {}

// ExitUnpivotValueColumn is called when production unpivotValueColumn is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotValueColumn(ctx *UnpivotValueColumnContext) {}

// EnterUnpivotNameColumn is called when production unpivotNameColumn is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotNameColumn(ctx *UnpivotNameColumnContext) {}

// ExitUnpivotNameColumn is called when production unpivotNameColumn is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotNameColumn(ctx *UnpivotNameColumnContext) {}

// EnterUnpivotColumnAndAlias is called when production unpivotColumnAndAlias is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotColumnAndAlias(ctx *UnpivotColumnAndAliasContext) {}

// ExitUnpivotColumnAndAlias is called when production unpivotColumnAndAlias is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotColumnAndAlias(ctx *UnpivotColumnAndAliasContext) {}

// EnterUnpivotColumn is called when production unpivotColumn is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotColumn(ctx *UnpivotColumnContext) {}

// ExitUnpivotColumn is called when production unpivotColumn is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotColumn(ctx *UnpivotColumnContext) {}

// EnterUnpivotAlias is called when production unpivotAlias is entered.
func (s *BaseSparkSqlParserListener) EnterUnpivotAlias(ctx *UnpivotAliasContext) {}

// ExitUnpivotAlias is called when production unpivotAlias is exited.
func (s *BaseSparkSqlParserListener) ExitUnpivotAlias(ctx *UnpivotAliasContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *BaseSparkSqlParserListener) EnterLateralView(ctx *LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *BaseSparkSqlParserListener) ExitLateralView(ctx *LateralViewContext) {}

// EnterWatermarkClause is called when production watermarkClause is entered.
func (s *BaseSparkSqlParserListener) EnterWatermarkClause(ctx *WatermarkClauseContext) {}

// ExitWatermarkClause is called when production watermarkClause is exited.
func (s *BaseSparkSqlParserListener) ExitWatermarkClause(ctx *WatermarkClauseContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseSparkSqlParserListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseSparkSqlParserListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseSparkSqlParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseSparkSqlParserListener) ExitRelation(ctx *RelationContext) {}

// EnterRelationExtension is called when production relationExtension is entered.
func (s *BaseSparkSqlParserListener) EnterRelationExtension(ctx *RelationExtensionContext) {}

// ExitRelationExtension is called when production relationExtension is exited.
func (s *BaseSparkSqlParserListener) ExitRelationExtension(ctx *RelationExtensionContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseSparkSqlParserListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseSparkSqlParserListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseSparkSqlParserListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseSparkSqlParserListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseSparkSqlParserListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseSparkSqlParserListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterSample is called when production sample is entered.
func (s *BaseSparkSqlParserListener) EnterSample(ctx *SampleContext) {}

// ExitSample is called when production sample is exited.
func (s *BaseSparkSqlParserListener) ExitSample(ctx *SampleContext) {}

// EnterSampleByPercentile is called when production sampleByPercentile is entered.
func (s *BaseSparkSqlParserListener) EnterSampleByPercentile(ctx *SampleByPercentileContext) {}

// ExitSampleByPercentile is called when production sampleByPercentile is exited.
func (s *BaseSparkSqlParserListener) ExitSampleByPercentile(ctx *SampleByPercentileContext) {}

// EnterSampleByRows is called when production sampleByRows is entered.
func (s *BaseSparkSqlParserListener) EnterSampleByRows(ctx *SampleByRowsContext) {}

// ExitSampleByRows is called when production sampleByRows is exited.
func (s *BaseSparkSqlParserListener) ExitSampleByRows(ctx *SampleByRowsContext) {}

// EnterSampleByBucket is called when production sampleByBucket is entered.
func (s *BaseSparkSqlParserListener) EnterSampleByBucket(ctx *SampleByBucketContext) {}

// ExitSampleByBucket is called when production sampleByBucket is exited.
func (s *BaseSparkSqlParserListener) ExitSampleByBucket(ctx *SampleByBucketContext) {}

// EnterSampleByBytes is called when production sampleByBytes is entered.
func (s *BaseSparkSqlParserListener) EnterSampleByBytes(ctx *SampleByBytesContext) {}

// ExitSampleByBytes is called when production sampleByBytes is exited.
func (s *BaseSparkSqlParserListener) ExitSampleByBytes(ctx *SampleByBytesContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierSeq is called when production identifierSeq is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierSeq(ctx *IdentifierSeqContext) {}

// ExitIdentifierSeq is called when production identifierSeq is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierSeq(ctx *IdentifierSeqContext) {}

// EnterOrderedIdentifierList is called when production orderedIdentifierList is entered.
func (s *BaseSparkSqlParserListener) EnterOrderedIdentifierList(ctx *OrderedIdentifierListContext) {}

// ExitOrderedIdentifierList is called when production orderedIdentifierList is exited.
func (s *BaseSparkSqlParserListener) ExitOrderedIdentifierList(ctx *OrderedIdentifierListContext) {}

// EnterOrderedIdentifier is called when production orderedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterOrderedIdentifier(ctx *OrderedIdentifierContext) {}

// ExitOrderedIdentifier is called when production orderedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitOrderedIdentifier(ctx *OrderedIdentifierContext) {}

// EnterIdentifierCommentList is called when production identifierCommentList is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierCommentList(ctx *IdentifierCommentListContext) {}

// ExitIdentifierCommentList is called when production identifierCommentList is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierCommentList(ctx *IdentifierCommentListContext) {}

// EnterIdentifierComment is called when production identifierComment is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierComment(ctx *IdentifierCommentContext) {}

// ExitIdentifierComment is called when production identifierComment is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierComment(ctx *IdentifierCommentContext) {}

// EnterStreamRelation is called when production streamRelation is entered.
func (s *BaseSparkSqlParserListener) EnterStreamRelation(ctx *StreamRelationContext) {}

// ExitStreamRelation is called when production streamRelation is exited.
func (s *BaseSparkSqlParserListener) ExitStreamRelation(ctx *StreamRelationContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSparkSqlParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSparkSqlParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterAliasedQuery is called when production aliasedQuery is entered.
func (s *BaseSparkSqlParserListener) EnterAliasedQuery(ctx *AliasedQueryContext) {}

// ExitAliasedQuery is called when production aliasedQuery is exited.
func (s *BaseSparkSqlParserListener) ExitAliasedQuery(ctx *AliasedQueryContext) {}

// EnterAliasedRelation is called when production aliasedRelation is entered.
func (s *BaseSparkSqlParserListener) EnterAliasedRelation(ctx *AliasedRelationContext) {}

// ExitAliasedRelation is called when production aliasedRelation is exited.
func (s *BaseSparkSqlParserListener) ExitAliasedRelation(ctx *AliasedRelationContext) {}

// EnterInlineTableDefault2 is called when production inlineTableDefault2 is entered.
func (s *BaseSparkSqlParserListener) EnterInlineTableDefault2(ctx *InlineTableDefault2Context) {}

// ExitInlineTableDefault2 is called when production inlineTableDefault2 is exited.
func (s *BaseSparkSqlParserListener) ExitInlineTableDefault2(ctx *InlineTableDefault2Context) {}

// EnterTableValuedFunction is called when production tableValuedFunction is entered.
func (s *BaseSparkSqlParserListener) EnterTableValuedFunction(ctx *TableValuedFunctionContext) {}

// ExitTableValuedFunction is called when production tableValuedFunction is exited.
func (s *BaseSparkSqlParserListener) ExitTableValuedFunction(ctx *TableValuedFunctionContext) {}

// EnterOptionsClause is called when production optionsClause is entered.
func (s *BaseSparkSqlParserListener) EnterOptionsClause(ctx *OptionsClauseContext) {}

// ExitOptionsClause is called when production optionsClause is exited.
func (s *BaseSparkSqlParserListener) ExitOptionsClause(ctx *OptionsClauseContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseSparkSqlParserListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseSparkSqlParserListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterFunctionTableSubqueryArgument is called when production functionTableSubqueryArgument is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionTableSubqueryArgument(ctx *FunctionTableSubqueryArgumentContext) {}

// ExitFunctionTableSubqueryArgument is called when production functionTableSubqueryArgument is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionTableSubqueryArgument(ctx *FunctionTableSubqueryArgumentContext) {}

// EnterTableArgumentPartitioning is called when production tableArgumentPartitioning is entered.
func (s *BaseSparkSqlParserListener) EnterTableArgumentPartitioning(ctx *TableArgumentPartitioningContext) {}

// ExitTableArgumentPartitioning is called when production tableArgumentPartitioning is exited.
func (s *BaseSparkSqlParserListener) ExitTableArgumentPartitioning(ctx *TableArgumentPartitioningContext) {}

// EnterFunctionTableNamedArgumentExpression is called when production functionTableNamedArgumentExpression is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionTableNamedArgumentExpression(ctx *FunctionTableNamedArgumentExpressionContext) {}

// ExitFunctionTableNamedArgumentExpression is called when production functionTableNamedArgumentExpression is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionTableNamedArgumentExpression(ctx *FunctionTableNamedArgumentExpressionContext) {}

// EnterFunctionTableReferenceArgument is called when production functionTableReferenceArgument is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionTableReferenceArgument(ctx *FunctionTableReferenceArgumentContext) {}

// ExitFunctionTableReferenceArgument is called when production functionTableReferenceArgument is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionTableReferenceArgument(ctx *FunctionTableReferenceArgumentContext) {}

// EnterFunctionTableArgument is called when production functionTableArgument is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionTableArgument(ctx *FunctionTableArgumentContext) {}

// ExitFunctionTableArgument is called when production functionTableArgument is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionTableArgument(ctx *FunctionTableArgumentContext) {}

// EnterFunctionTable is called when production functionTable is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionTable(ctx *FunctionTableContext) {}

// ExitFunctionTable is called when production functionTable is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionTable(ctx *FunctionTableContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseSparkSqlParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseSparkSqlParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterRowFormatSerde is called when production rowFormatSerde is entered.
func (s *BaseSparkSqlParserListener) EnterRowFormatSerde(ctx *RowFormatSerdeContext) {}

// ExitRowFormatSerde is called when production rowFormatSerde is exited.
func (s *BaseSparkSqlParserListener) ExitRowFormatSerde(ctx *RowFormatSerdeContext) {}

// EnterRowFormatDelimited is called when production rowFormatDelimited is entered.
func (s *BaseSparkSqlParserListener) EnterRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// ExitRowFormatDelimited is called when production rowFormatDelimited is exited.
func (s *BaseSparkSqlParserListener) ExitRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// EnterMultipartIdentifierList is called when production multipartIdentifierList is entered.
func (s *BaseSparkSqlParserListener) EnterMultipartIdentifierList(ctx *MultipartIdentifierListContext) {}

// ExitMultipartIdentifierList is called when production multipartIdentifierList is exited.
func (s *BaseSparkSqlParserListener) ExitMultipartIdentifierList(ctx *MultipartIdentifierListContext) {}

// EnterMultipartIdentifier is called when production multipartIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterMultipartIdentifier(ctx *MultipartIdentifierContext) {}

// ExitMultipartIdentifier is called when production multipartIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitMultipartIdentifier(ctx *MultipartIdentifierContext) {}

// EnterMultipartIdentifierPropertyList is called when production multipartIdentifierPropertyList is entered.
func (s *BaseSparkSqlParserListener) EnterMultipartIdentifierPropertyList(ctx *MultipartIdentifierPropertyListContext) {}

// ExitMultipartIdentifierPropertyList is called when production multipartIdentifierPropertyList is exited.
func (s *BaseSparkSqlParserListener) ExitMultipartIdentifierPropertyList(ctx *MultipartIdentifierPropertyListContext) {}

// EnterMultipartIdentifierProperty is called when production multipartIdentifierProperty is entered.
func (s *BaseSparkSqlParserListener) EnterMultipartIdentifierProperty(ctx *MultipartIdentifierPropertyContext) {}

// ExitMultipartIdentifierProperty is called when production multipartIdentifierProperty is exited.
func (s *BaseSparkSqlParserListener) ExitMultipartIdentifierProperty(ctx *MultipartIdentifierPropertyContext) {}

// EnterTableIdentifier is called when production tableIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterTableIdentifier(ctx *TableIdentifierContext) {}

// ExitTableIdentifier is called when production tableIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitTableIdentifier(ctx *TableIdentifierContext) {}

// EnterFunctionIdentifier is called when production functionIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionIdentifier(ctx *FunctionIdentifierContext) {}

// ExitFunctionIdentifier is called when production functionIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionIdentifier(ctx *FunctionIdentifierContext) {}

// EnterNamedExpression is called when production namedExpression is entered.
func (s *BaseSparkSqlParserListener) EnterNamedExpression(ctx *NamedExpressionContext) {}

// ExitNamedExpression is called when production namedExpression is exited.
func (s *BaseSparkSqlParserListener) ExitNamedExpression(ctx *NamedExpressionContext) {}

// EnterNamedExpressionSeq is called when production namedExpressionSeq is entered.
func (s *BaseSparkSqlParserListener) EnterNamedExpressionSeq(ctx *NamedExpressionSeqContext) {}

// ExitNamedExpressionSeq is called when production namedExpressionSeq is exited.
func (s *BaseSparkSqlParserListener) ExitNamedExpressionSeq(ctx *NamedExpressionSeqContext) {}

// EnterPartitionFieldList is called when production partitionFieldList is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionFieldList(ctx *PartitionFieldListContext) {}

// ExitPartitionFieldList is called when production partitionFieldList is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionFieldList(ctx *PartitionFieldListContext) {}

// EnterPartitionTransform is called when production partitionTransform is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionTransform(ctx *PartitionTransformContext) {}

// ExitPartitionTransform is called when production partitionTransform is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionTransform(ctx *PartitionTransformContext) {}

// EnterPartitionColumn is called when production partitionColumn is entered.
func (s *BaseSparkSqlParserListener) EnterPartitionColumn(ctx *PartitionColumnContext) {}

// ExitPartitionColumn is called when production partitionColumn is exited.
func (s *BaseSparkSqlParserListener) ExitPartitionColumn(ctx *PartitionColumnContext) {}

// EnterIdentityTransform is called when production identityTransform is entered.
func (s *BaseSparkSqlParserListener) EnterIdentityTransform(ctx *IdentityTransformContext) {}

// ExitIdentityTransform is called when production identityTransform is exited.
func (s *BaseSparkSqlParserListener) ExitIdentityTransform(ctx *IdentityTransformContext) {}

// EnterApplyTransform is called when production applyTransform is entered.
func (s *BaseSparkSqlParserListener) EnterApplyTransform(ctx *ApplyTransformContext) {}

// ExitApplyTransform is called when production applyTransform is exited.
func (s *BaseSparkSqlParserListener) ExitApplyTransform(ctx *ApplyTransformContext) {}

// EnterTransformArgument is called when production transformArgument is entered.
func (s *BaseSparkSqlParserListener) EnterTransformArgument(ctx *TransformArgumentContext) {}

// ExitTransformArgument is called when production transformArgument is exited.
func (s *BaseSparkSqlParserListener) ExitTransformArgument(ctx *TransformArgumentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSparkSqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSparkSqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNamedArgumentExpression is called when production namedArgumentExpression is entered.
func (s *BaseSparkSqlParserListener) EnterNamedArgumentExpression(ctx *NamedArgumentExpressionContext) {}

// ExitNamedArgumentExpression is called when production namedArgumentExpression is exited.
func (s *BaseSparkSqlParserListener) ExitNamedArgumentExpression(ctx *NamedArgumentExpressionContext) {}

// EnterFunctionArgument is called when production functionArgument is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionArgument(ctx *FunctionArgumentContext) {}

// ExitFunctionArgument is called when production functionArgument is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionArgument(ctx *FunctionArgumentContext) {}

// EnterExpressionSeq is called when production expressionSeq is entered.
func (s *BaseSparkSqlParserListener) EnterExpressionSeq(ctx *ExpressionSeqContext) {}

// ExitExpressionSeq is called when production expressionSeq is exited.
func (s *BaseSparkSqlParserListener) ExitExpressionSeq(ctx *ExpressionSeqContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseSparkSqlParserListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseSparkSqlParserListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseSparkSqlParserListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseSparkSqlParserListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseSparkSqlParserListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseSparkSqlParserListener) ExitExists(ctx *ExistsContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseSparkSqlParserListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseSparkSqlParserListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSparkSqlParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSparkSqlParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterErrorCapturingNot is called when production errorCapturingNot is entered.
func (s *BaseSparkSqlParserListener) EnterErrorCapturingNot(ctx *ErrorCapturingNotContext) {}

// ExitErrorCapturingNot is called when production errorCapturingNot is exited.
func (s *BaseSparkSqlParserListener) ExitErrorCapturingNot(ctx *ErrorCapturingNotContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseSparkSqlParserListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseSparkSqlParserListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSparkSqlParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSparkSqlParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseSparkSqlParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseSparkSqlParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseSparkSqlParserListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseSparkSqlParserListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseSparkSqlParserListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseSparkSqlParserListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *BaseSparkSqlParserListener) EnterShiftOperator(ctx *ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *BaseSparkSqlParserListener) ExitShiftOperator(ctx *ShiftOperatorContext) {}

// EnterDatetimeUnit is called when production datetimeUnit is entered.
func (s *BaseSparkSqlParserListener) EnterDatetimeUnit(ctx *DatetimeUnitContext) {}

// ExitDatetimeUnit is called when production datetimeUnit is exited.
func (s *BaseSparkSqlParserListener) ExitDatetimeUnit(ctx *DatetimeUnitContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseSparkSqlParserListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseSparkSqlParserListener) ExitStruct(ctx *StructContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseSparkSqlParserListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseSparkSqlParserListener) ExitDereference(ctx *DereferenceContext) {}

// EnterCastByColon is called when production castByColon is entered.
func (s *BaseSparkSqlParserListener) EnterCastByColon(ctx *CastByColonContext) {}

// ExitCastByColon is called when production castByColon is exited.
func (s *BaseSparkSqlParserListener) ExitCastByColon(ctx *CastByColonContext) {}

// EnterTimestampadd is called when production timestampadd is entered.
func (s *BaseSparkSqlParserListener) EnterTimestampadd(ctx *TimestampaddContext) {}

// ExitTimestampadd is called when production timestampadd is exited.
func (s *BaseSparkSqlParserListener) ExitTimestampadd(ctx *TimestampaddContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BaseSparkSqlParserListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BaseSparkSqlParserListener) ExitSubstring(ctx *SubstringContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseSparkSqlParserListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseSparkSqlParserListener) ExitCast(ctx *CastContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseSparkSqlParserListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseSparkSqlParserListener) ExitLambda(ctx *LambdaContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseSparkSqlParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseSparkSqlParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterAny_value is called when production any_value is entered.
func (s *BaseSparkSqlParserListener) EnterAny_value(ctx *Any_valueContext) {}

// ExitAny_value is called when production any_value is exited.
func (s *BaseSparkSqlParserListener) ExitAny_value(ctx *Any_valueContext) {}

// EnterTrim is called when production trim is entered.
func (s *BaseSparkSqlParserListener) EnterTrim(ctx *TrimContext) {}

// ExitTrim is called when production trim is exited.
func (s *BaseSparkSqlParserListener) ExitTrim(ctx *TrimContext) {}

// EnterSemiStructuredExtract is called when production semiStructuredExtract is entered.
func (s *BaseSparkSqlParserListener) EnterSemiStructuredExtract(ctx *SemiStructuredExtractContext) {}

// ExitSemiStructuredExtract is called when production semiStructuredExtract is exited.
func (s *BaseSparkSqlParserListener) ExitSemiStructuredExtract(ctx *SemiStructuredExtractContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseSparkSqlParserListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseSparkSqlParserListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterCurrentLike is called when production currentLike is entered.
func (s *BaseSparkSqlParserListener) EnterCurrentLike(ctx *CurrentLikeContext) {}

// ExitCurrentLike is called when production currentLike is exited.
func (s *BaseSparkSqlParserListener) ExitCurrentLike(ctx *CurrentLikeContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseSparkSqlParserListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseSparkSqlParserListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseSparkSqlParserListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseSparkSqlParserListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterLast is called when production last is entered.
func (s *BaseSparkSqlParserListener) EnterLast(ctx *LastContext) {}

// ExitLast is called when production last is exited.
func (s *BaseSparkSqlParserListener) ExitLast(ctx *LastContext) {}

// EnterStar is called when production star is entered.
func (s *BaseSparkSqlParserListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production star is exited.
func (s *BaseSparkSqlParserListener) ExitStar(ctx *StarContext) {}

// EnterOverlay is called when production overlay is entered.
func (s *BaseSparkSqlParserListener) EnterOverlay(ctx *OverlayContext) {}

// ExitOverlay is called when production overlay is exited.
func (s *BaseSparkSqlParserListener) ExitOverlay(ctx *OverlayContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseSparkSqlParserListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseSparkSqlParserListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterTimestampdiff is called when production timestampdiff is entered.
func (s *BaseSparkSqlParserListener) EnterTimestampdiff(ctx *TimestampdiffContext) {}

// ExitTimestampdiff is called when production timestampdiff is exited.
func (s *BaseSparkSqlParserListener) ExitTimestampdiff(ctx *TimestampdiffContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseSparkSqlParserListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseSparkSqlParserListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseSparkSqlParserListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseSparkSqlParserListener) ExitCollate(ctx *CollateContext) {}

// EnterConstantDefault is called when production constantDefault is entered.
func (s *BaseSparkSqlParserListener) EnterConstantDefault(ctx *ConstantDefaultContext) {}

// ExitConstantDefault is called when production constantDefault is exited.
func (s *BaseSparkSqlParserListener) ExitConstantDefault(ctx *ConstantDefaultContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseSparkSqlParserListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseSparkSqlParserListener) ExitExtract(ctx *ExtractContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseSparkSqlParserListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseSparkSqlParserListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterPosition is called when production position is entered.
func (s *BaseSparkSqlParserListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BaseSparkSqlParserListener) ExitPosition(ctx *PositionContext) {}

// EnterFirst is called when production first is entered.
func (s *BaseSparkSqlParserListener) EnterFirst(ctx *FirstContext) {}

// ExitFirst is called when production first is exited.
func (s *BaseSparkSqlParserListener) ExitFirst(ctx *FirstContext) {}

// EnterSemiStructuredExtractionPath is called when production semiStructuredExtractionPath is entered.
func (s *BaseSparkSqlParserListener) EnterSemiStructuredExtractionPath(ctx *SemiStructuredExtractionPathContext) {}

// ExitSemiStructuredExtractionPath is called when production semiStructuredExtractionPath is exited.
func (s *BaseSparkSqlParserListener) ExitSemiStructuredExtractionPath(ctx *SemiStructuredExtractionPathContext) {}

// EnterJsonPathIdentifier is called when production jsonPathIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterJsonPathIdentifier(ctx *JsonPathIdentifierContext) {}

// ExitJsonPathIdentifier is called when production jsonPathIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitJsonPathIdentifier(ctx *JsonPathIdentifierContext) {}

// EnterJsonPathBracketedIdentifier is called when production jsonPathBracketedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterJsonPathBracketedIdentifier(ctx *JsonPathBracketedIdentifierContext) {}

// ExitJsonPathBracketedIdentifier is called when production jsonPathBracketedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitJsonPathBracketedIdentifier(ctx *JsonPathBracketedIdentifierContext) {}

// EnterJsonPathFirstPart is called when production jsonPathFirstPart is entered.
func (s *BaseSparkSqlParserListener) EnterJsonPathFirstPart(ctx *JsonPathFirstPartContext) {}

// ExitJsonPathFirstPart is called when production jsonPathFirstPart is exited.
func (s *BaseSparkSqlParserListener) ExitJsonPathFirstPart(ctx *JsonPathFirstPartContext) {}

// EnterJsonPathParts is called when production jsonPathParts is entered.
func (s *BaseSparkSqlParserListener) EnterJsonPathParts(ctx *JsonPathPartsContext) {}

// ExitJsonPathParts is called when production jsonPathParts is exited.
func (s *BaseSparkSqlParserListener) ExitJsonPathParts(ctx *JsonPathPartsContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *BaseSparkSqlParserListener) EnterLiteralType(ctx *LiteralTypeContext) {}

// ExitLiteralType is called when production literalType is exited.
func (s *BaseSparkSqlParserListener) ExitLiteralType(ctx *LiteralTypeContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterPosParameterLiteral is called when production posParameterLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterPosParameterLiteral(ctx *PosParameterLiteralContext) {}

// ExitPosParameterLiteral is called when production posParameterLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitPosParameterLiteral(ctx *PosParameterLiteralContext) {}

// EnterNamedParameterLiteral is called when production namedParameterLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterNamedParameterLiteral(ctx *NamedParameterLiteralContext) {}

// ExitNamedParameterLiteral is called when production namedParameterLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitNamedParameterLiteral(ctx *NamedParameterLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterTypeConstructor is called when production typeConstructor is entered.
func (s *BaseSparkSqlParserListener) EnterTypeConstructor(ctx *TypeConstructorContext) {}

// ExitTypeConstructor is called when production typeConstructor is exited.
func (s *BaseSparkSqlParserListener) ExitTypeConstructor(ctx *TypeConstructorContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNamedParameterMarker is called when production namedParameterMarker is entered.
func (s *BaseSparkSqlParserListener) EnterNamedParameterMarker(ctx *NamedParameterMarkerContext) {}

// ExitNamedParameterMarker is called when production namedParameterMarker is exited.
func (s *BaseSparkSqlParserListener) ExitNamedParameterMarker(ctx *NamedParameterMarkerContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSparkSqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSparkSqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterArithmeticOperator is called when production arithmeticOperator is entered.
func (s *BaseSparkSqlParserListener) EnterArithmeticOperator(ctx *ArithmeticOperatorContext) {}

// ExitArithmeticOperator is called when production arithmeticOperator is exited.
func (s *BaseSparkSqlParserListener) ExitArithmeticOperator(ctx *ArithmeticOperatorContext) {}

// EnterPredicateOperator is called when production predicateOperator is entered.
func (s *BaseSparkSqlParserListener) EnterPredicateOperator(ctx *PredicateOperatorContext) {}

// ExitPredicateOperator is called when production predicateOperator is exited.
func (s *BaseSparkSqlParserListener) ExitPredicateOperator(ctx *PredicateOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseSparkSqlParserListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseSparkSqlParserListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseSparkSqlParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseSparkSqlParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterErrorCapturingMultiUnitsInterval is called when production errorCapturingMultiUnitsInterval is entered.
func (s *BaseSparkSqlParserListener) EnterErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) {}

// ExitErrorCapturingMultiUnitsInterval is called when production errorCapturingMultiUnitsInterval is exited.
func (s *BaseSparkSqlParserListener) ExitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) {}

// EnterMultiUnitsInterval is called when production multiUnitsInterval is entered.
func (s *BaseSparkSqlParserListener) EnterMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// ExitMultiUnitsInterval is called when production multiUnitsInterval is exited.
func (s *BaseSparkSqlParserListener) ExitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// EnterErrorCapturingUnitToUnitInterval is called when production errorCapturingUnitToUnitInterval is entered.
func (s *BaseSparkSqlParserListener) EnterErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) {}

// ExitErrorCapturingUnitToUnitInterval is called when production errorCapturingUnitToUnitInterval is exited.
func (s *BaseSparkSqlParserListener) ExitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) {}

// EnterUnitToUnitInterval is called when production unitToUnitInterval is entered.
func (s *BaseSparkSqlParserListener) EnterUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// ExitUnitToUnitInterval is called when production unitToUnitInterval is exited.
func (s *BaseSparkSqlParserListener) ExitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// EnterIntervalValue is called when production intervalValue is entered.
func (s *BaseSparkSqlParserListener) EnterIntervalValue(ctx *IntervalValueContext) {}

// ExitIntervalValue is called when production intervalValue is exited.
func (s *BaseSparkSqlParserListener) ExitIntervalValue(ctx *IntervalValueContext) {}

// EnterUnitInMultiUnits is called when production unitInMultiUnits is entered.
func (s *BaseSparkSqlParserListener) EnterUnitInMultiUnits(ctx *UnitInMultiUnitsContext) {}

// ExitUnitInMultiUnits is called when production unitInMultiUnits is exited.
func (s *BaseSparkSqlParserListener) ExitUnitInMultiUnits(ctx *UnitInMultiUnitsContext) {}

// EnterUnitInUnitToUnit is called when production unitInUnitToUnit is entered.
func (s *BaseSparkSqlParserListener) EnterUnitInUnitToUnit(ctx *UnitInUnitToUnitContext) {}

// ExitUnitInUnitToUnit is called when production unitInUnitToUnit is exited.
func (s *BaseSparkSqlParserListener) ExitUnitInUnitToUnit(ctx *UnitInUnitToUnitContext) {}

// EnterColPosition is called when production colPosition is entered.
func (s *BaseSparkSqlParserListener) EnterColPosition(ctx *ColPositionContext) {}

// ExitColPosition is called when production colPosition is exited.
func (s *BaseSparkSqlParserListener) ExitColPosition(ctx *ColPositionContext) {}

// EnterCollationSpec is called when production collationSpec is entered.
func (s *BaseSparkSqlParserListener) EnterCollationSpec(ctx *CollationSpecContext) {}

// ExitCollationSpec is called when production collationSpec is exited.
func (s *BaseSparkSqlParserListener) ExitCollationSpec(ctx *CollationSpecContext) {}

// EnterCollateClause is called when production collateClause is entered.
func (s *BaseSparkSqlParserListener) EnterCollateClause(ctx *CollateClauseContext) {}

// ExitCollateClause is called when production collateClause is exited.
func (s *BaseSparkSqlParserListener) ExitCollateClause(ctx *CollateClauseContext) {}

// EnterNonTrivialPrimitiveType is called when production nonTrivialPrimitiveType is entered.
func (s *BaseSparkSqlParserListener) EnterNonTrivialPrimitiveType(ctx *NonTrivialPrimitiveTypeContext) {}

// ExitNonTrivialPrimitiveType is called when production nonTrivialPrimitiveType is exited.
func (s *BaseSparkSqlParserListener) ExitNonTrivialPrimitiveType(ctx *NonTrivialPrimitiveTypeContext) {}

// EnterTrivialPrimitiveType is called when production trivialPrimitiveType is entered.
func (s *BaseSparkSqlParserListener) EnterTrivialPrimitiveType(ctx *TrivialPrimitiveTypeContext) {}

// ExitTrivialPrimitiveType is called when production trivialPrimitiveType is exited.
func (s *BaseSparkSqlParserListener) ExitTrivialPrimitiveType(ctx *TrivialPrimitiveTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseSparkSqlParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseSparkSqlParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterComplexDataType is called when production complexDataType is entered.
func (s *BaseSparkSqlParserListener) EnterComplexDataType(ctx *ComplexDataTypeContext) {}

// ExitComplexDataType is called when production complexDataType is exited.
func (s *BaseSparkSqlParserListener) ExitComplexDataType(ctx *ComplexDataTypeContext) {}

// EnterPrimitiveDataType is called when production primitiveDataType is entered.
func (s *BaseSparkSqlParserListener) EnterPrimitiveDataType(ctx *PrimitiveDataTypeContext) {}

// ExitPrimitiveDataType is called when production primitiveDataType is exited.
func (s *BaseSparkSqlParserListener) ExitPrimitiveDataType(ctx *PrimitiveDataTypeContext) {}

// EnterQualifiedColTypeWithPositionList is called when production qualifiedColTypeWithPositionList is entered.
func (s *BaseSparkSqlParserListener) EnterQualifiedColTypeWithPositionList(ctx *QualifiedColTypeWithPositionListContext) {}

// ExitQualifiedColTypeWithPositionList is called when production qualifiedColTypeWithPositionList is exited.
func (s *BaseSparkSqlParserListener) ExitQualifiedColTypeWithPositionList(ctx *QualifiedColTypeWithPositionListContext) {}

// EnterQualifiedColTypeWithPosition is called when production qualifiedColTypeWithPosition is entered.
func (s *BaseSparkSqlParserListener) EnterQualifiedColTypeWithPosition(ctx *QualifiedColTypeWithPositionContext) {}

// ExitQualifiedColTypeWithPosition is called when production qualifiedColTypeWithPosition is exited.
func (s *BaseSparkSqlParserListener) ExitQualifiedColTypeWithPosition(ctx *QualifiedColTypeWithPositionContext) {}

// EnterColDefinitionDescriptorWithPosition is called when production colDefinitionDescriptorWithPosition is entered.
func (s *BaseSparkSqlParserListener) EnterColDefinitionDescriptorWithPosition(ctx *ColDefinitionDescriptorWithPositionContext) {}

// ExitColDefinitionDescriptorWithPosition is called when production colDefinitionDescriptorWithPosition is exited.
func (s *BaseSparkSqlParserListener) ExitColDefinitionDescriptorWithPosition(ctx *ColDefinitionDescriptorWithPositionContext) {}

// EnterDefaultExpression is called when production defaultExpression is entered.
func (s *BaseSparkSqlParserListener) EnterDefaultExpression(ctx *DefaultExpressionContext) {}

// ExitDefaultExpression is called when production defaultExpression is exited.
func (s *BaseSparkSqlParserListener) ExitDefaultExpression(ctx *DefaultExpressionContext) {}

// EnterVariableDefaultExpression is called when production variableDefaultExpression is entered.
func (s *BaseSparkSqlParserListener) EnterVariableDefaultExpression(ctx *VariableDefaultExpressionContext) {}

// ExitVariableDefaultExpression is called when production variableDefaultExpression is exited.
func (s *BaseSparkSqlParserListener) ExitVariableDefaultExpression(ctx *VariableDefaultExpressionContext) {}

// EnterColTypeList is called when production colTypeList is entered.
func (s *BaseSparkSqlParserListener) EnterColTypeList(ctx *ColTypeListContext) {}

// ExitColTypeList is called when production colTypeList is exited.
func (s *BaseSparkSqlParserListener) ExitColTypeList(ctx *ColTypeListContext) {}

// EnterColType is called when production colType is entered.
func (s *BaseSparkSqlParserListener) EnterColType(ctx *ColTypeContext) {}

// ExitColType is called when production colType is exited.
func (s *BaseSparkSqlParserListener) ExitColType(ctx *ColTypeContext) {}

// EnterTableElementList is called when production tableElementList is entered.
func (s *BaseSparkSqlParserListener) EnterTableElementList(ctx *TableElementListContext) {}

// ExitTableElementList is called when production tableElementList is exited.
func (s *BaseSparkSqlParserListener) ExitTableElementList(ctx *TableElementListContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseSparkSqlParserListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseSparkSqlParserListener) ExitTableElement(ctx *TableElementContext) {}

// EnterColDefinitionList is called when production colDefinitionList is entered.
func (s *BaseSparkSqlParserListener) EnterColDefinitionList(ctx *ColDefinitionListContext) {}

// ExitColDefinitionList is called when production colDefinitionList is exited.
func (s *BaseSparkSqlParserListener) ExitColDefinitionList(ctx *ColDefinitionListContext) {}

// EnterColDefinition is called when production colDefinition is entered.
func (s *BaseSparkSqlParserListener) EnterColDefinition(ctx *ColDefinitionContext) {}

// ExitColDefinition is called when production colDefinition is exited.
func (s *BaseSparkSqlParserListener) ExitColDefinition(ctx *ColDefinitionContext) {}

// EnterColDefinitionOption is called when production colDefinitionOption is entered.
func (s *BaseSparkSqlParserListener) EnterColDefinitionOption(ctx *ColDefinitionOptionContext) {}

// ExitColDefinitionOption is called when production colDefinitionOption is exited.
func (s *BaseSparkSqlParserListener) ExitColDefinitionOption(ctx *ColDefinitionOptionContext) {}

// EnterGeneratedColumn is called when production generatedColumn is entered.
func (s *BaseSparkSqlParserListener) EnterGeneratedColumn(ctx *GeneratedColumnContext) {}

// ExitGeneratedColumn is called when production generatedColumn is exited.
func (s *BaseSparkSqlParserListener) ExitGeneratedColumn(ctx *GeneratedColumnContext) {}

// EnterIdentityColumn is called when production identityColumn is entered.
func (s *BaseSparkSqlParserListener) EnterIdentityColumn(ctx *IdentityColumnContext) {}

// ExitIdentityColumn is called when production identityColumn is exited.
func (s *BaseSparkSqlParserListener) ExitIdentityColumn(ctx *IdentityColumnContext) {}

// EnterIdentityColSpec is called when production identityColSpec is entered.
func (s *BaseSparkSqlParserListener) EnterIdentityColSpec(ctx *IdentityColSpecContext) {}

// ExitIdentityColSpec is called when production identityColSpec is exited.
func (s *BaseSparkSqlParserListener) ExitIdentityColSpec(ctx *IdentityColSpecContext) {}

// EnterSequenceGeneratorOption is called when production sequenceGeneratorOption is entered.
func (s *BaseSparkSqlParserListener) EnterSequenceGeneratorOption(ctx *SequenceGeneratorOptionContext) {}

// ExitSequenceGeneratorOption is called when production sequenceGeneratorOption is exited.
func (s *BaseSparkSqlParserListener) ExitSequenceGeneratorOption(ctx *SequenceGeneratorOptionContext) {}

// EnterSequenceGeneratorStartOrStep is called when production sequenceGeneratorStartOrStep is entered.
func (s *BaseSparkSqlParserListener) EnterSequenceGeneratorStartOrStep(ctx *SequenceGeneratorStartOrStepContext) {}

// ExitSequenceGeneratorStartOrStep is called when production sequenceGeneratorStartOrStep is exited.
func (s *BaseSparkSqlParserListener) ExitSequenceGeneratorStartOrStep(ctx *SequenceGeneratorStartOrStepContext) {}

// EnterComplexColTypeList is called when production complexColTypeList is entered.
func (s *BaseSparkSqlParserListener) EnterComplexColTypeList(ctx *ComplexColTypeListContext) {}

// ExitComplexColTypeList is called when production complexColTypeList is exited.
func (s *BaseSparkSqlParserListener) ExitComplexColTypeList(ctx *ComplexColTypeListContext) {}

// EnterComplexColType is called when production complexColType is entered.
func (s *BaseSparkSqlParserListener) EnterComplexColType(ctx *ComplexColTypeContext) {}

// ExitComplexColType is called when production complexColType is exited.
func (s *BaseSparkSqlParserListener) ExitComplexColType(ctx *ComplexColTypeContext) {}

// EnterCodeLiteral is called when production codeLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterCodeLiteral(ctx *CodeLiteralContext) {}

// ExitCodeLiteral is called when production codeLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitCodeLiteral(ctx *CodeLiteralContext) {}

// EnterRoutineCharacteristics is called when production routineCharacteristics is entered.
func (s *BaseSparkSqlParserListener) EnterRoutineCharacteristics(ctx *RoutineCharacteristicsContext) {}

// ExitRoutineCharacteristics is called when production routineCharacteristics is exited.
func (s *BaseSparkSqlParserListener) ExitRoutineCharacteristics(ctx *RoutineCharacteristicsContext) {}

// EnterRoutineLanguage is called when production routineLanguage is entered.
func (s *BaseSparkSqlParserListener) EnterRoutineLanguage(ctx *RoutineLanguageContext) {}

// ExitRoutineLanguage is called when production routineLanguage is exited.
func (s *BaseSparkSqlParserListener) ExitRoutineLanguage(ctx *RoutineLanguageContext) {}

// EnterSpecificName is called when production specificName is entered.
func (s *BaseSparkSqlParserListener) EnterSpecificName(ctx *SpecificNameContext) {}

// ExitSpecificName is called when production specificName is exited.
func (s *BaseSparkSqlParserListener) ExitSpecificName(ctx *SpecificNameContext) {}

// EnterDeterministic is called when production deterministic is entered.
func (s *BaseSparkSqlParserListener) EnterDeterministic(ctx *DeterministicContext) {}

// ExitDeterministic is called when production deterministic is exited.
func (s *BaseSparkSqlParserListener) ExitDeterministic(ctx *DeterministicContext) {}

// EnterSqlDataAccess is called when production sqlDataAccess is entered.
func (s *BaseSparkSqlParserListener) EnterSqlDataAccess(ctx *SqlDataAccessContext) {}

// ExitSqlDataAccess is called when production sqlDataAccess is exited.
func (s *BaseSparkSqlParserListener) ExitSqlDataAccess(ctx *SqlDataAccessContext) {}

// EnterNullCall is called when production nullCall is entered.
func (s *BaseSparkSqlParserListener) EnterNullCall(ctx *NullCallContext) {}

// ExitNullCall is called when production nullCall is exited.
func (s *BaseSparkSqlParserListener) ExitNullCall(ctx *NullCallContext) {}

// EnterRightsClause is called when production rightsClause is entered.
func (s *BaseSparkSqlParserListener) EnterRightsClause(ctx *RightsClauseContext) {}

// ExitRightsClause is called when production rightsClause is exited.
func (s *BaseSparkSqlParserListener) ExitRightsClause(ctx *RightsClauseContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseSparkSqlParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseSparkSqlParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseSparkSqlParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseSparkSqlParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterNamedWindow is called when production namedWindow is entered.
func (s *BaseSparkSqlParserListener) EnterNamedWindow(ctx *NamedWindowContext) {}

// ExitNamedWindow is called when production namedWindow is exited.
func (s *BaseSparkSqlParserListener) ExitNamedWindow(ctx *NamedWindowContext) {}

// EnterWindowRef is called when production windowRef is entered.
func (s *BaseSparkSqlParserListener) EnterWindowRef(ctx *WindowRefContext) {}

// ExitWindowRef is called when production windowRef is exited.
func (s *BaseSparkSqlParserListener) ExitWindowRef(ctx *WindowRefContext) {}

// EnterWindowDef is called when production windowDef is entered.
func (s *BaseSparkSqlParserListener) EnterWindowDef(ctx *WindowDefContext) {}

// ExitWindowDef is called when production windowDef is exited.
func (s *BaseSparkSqlParserListener) ExitWindowDef(ctx *WindowDefContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseSparkSqlParserListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseSparkSqlParserListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterFrameBound is called when production frameBound is entered.
func (s *BaseSparkSqlParserListener) EnterFrameBound(ctx *FrameBoundContext) {}

// ExitFrameBound is called when production frameBound is exited.
func (s *BaseSparkSqlParserListener) ExitFrameBound(ctx *FrameBoundContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BaseSparkSqlParserListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BaseSparkSqlParserListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseSparkSqlParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseSparkSqlParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseSparkSqlParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseSparkSqlParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterErrorCapturingIdentifier is called when production errorCapturingIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) {}

// ExitErrorCapturingIdentifier is called when production errorCapturingIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) {}

// EnterErrorIdent is called when production errorIdent is entered.
func (s *BaseSparkSqlParserListener) EnterErrorIdent(ctx *ErrorIdentContext) {}

// ExitErrorIdent is called when production errorIdent is exited.
func (s *BaseSparkSqlParserListener) ExitErrorIdent(ctx *ErrorIdentContext) {}

// EnterRealIdent is called when production realIdent is entered.
func (s *BaseSparkSqlParserListener) EnterRealIdent(ctx *RealIdentContext) {}

// ExitRealIdent is called when production realIdent is exited.
func (s *BaseSparkSqlParserListener) ExitRealIdent(ctx *RealIdentContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is entered.
func (s *BaseSparkSqlParserListener) EnterQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {}

// ExitQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is exited.
func (s *BaseSparkSqlParserListener) ExitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {}

// EnterIdentifierLiteral is called when production identifierLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterIdentifierLiteral(ctx *IdentifierLiteralContext) {}

// ExitIdentifierLiteral is called when production identifierLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitIdentifierLiteral(ctx *IdentifierLiteralContext) {}

// EnterSimpleUnquotedIdentifier is called when production simpleUnquotedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterSimpleUnquotedIdentifier(ctx *SimpleUnquotedIdentifierContext) {}

// ExitSimpleUnquotedIdentifier is called when production simpleUnquotedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitSimpleUnquotedIdentifier(ctx *SimpleUnquotedIdentifierContext) {}

// EnterSimpleQuotedIdentifierAlternative is called when production simpleQuotedIdentifierAlternative is entered.
func (s *BaseSparkSqlParserListener) EnterSimpleQuotedIdentifierAlternative(ctx *SimpleQuotedIdentifierAlternativeContext) {}

// ExitSimpleQuotedIdentifierAlternative is called when production simpleQuotedIdentifierAlternative is exited.
func (s *BaseSparkSqlParserListener) ExitSimpleQuotedIdentifierAlternative(ctx *SimpleQuotedIdentifierAlternativeContext) {}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseSparkSqlParserListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseSparkSqlParserListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterExponentLiteral is called when production exponentLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterExponentLiteral(ctx *ExponentLiteralContext) {}

// ExitExponentLiteral is called when production exponentLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitExponentLiteral(ctx *ExponentLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterLegacyDecimalLiteral is called when production legacyDecimalLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) {}

// ExitLegacyDecimalLiteral is called when production legacyDecimalLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterBigIntLiteral is called when production bigIntLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterBigIntLiteral(ctx *BigIntLiteralContext) {}

// ExitBigIntLiteral is called when production bigIntLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitBigIntLiteral(ctx *BigIntLiteralContext) {}

// EnterSmallIntLiteral is called when production smallIntLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterSmallIntLiteral(ctx *SmallIntLiteralContext) {}

// ExitSmallIntLiteral is called when production smallIntLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitSmallIntLiteral(ctx *SmallIntLiteralContext) {}

// EnterTinyIntLiteral is called when production tinyIntLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterTinyIntLiteral(ctx *TinyIntLiteralContext) {}

// ExitTinyIntLiteral is called when production tinyIntLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitTinyIntLiteral(ctx *TinyIntLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterBigDecimalLiteral is called when production bigDecimalLiteral is entered.
func (s *BaseSparkSqlParserListener) EnterBigDecimalLiteral(ctx *BigDecimalLiteralContext) {}

// ExitBigDecimalLiteral is called when production bigDecimalLiteral is exited.
func (s *BaseSparkSqlParserListener) ExitBigDecimalLiteral(ctx *BigDecimalLiteralContext) {}

// EnterIntegerVal is called when production integerVal is entered.
func (s *BaseSparkSqlParserListener) EnterIntegerVal(ctx *IntegerValContext) {}

// ExitIntegerVal is called when production integerVal is exited.
func (s *BaseSparkSqlParserListener) ExitIntegerVal(ctx *IntegerValContext) {}

// EnterParameterIntegerValue is called when production parameterIntegerValue is entered.
func (s *BaseSparkSqlParserListener) EnterParameterIntegerValue(ctx *ParameterIntegerValueContext) {}

// ExitParameterIntegerValue is called when production parameterIntegerValue is exited.
func (s *BaseSparkSqlParserListener) ExitParameterIntegerValue(ctx *ParameterIntegerValueContext) {}

// EnterColumnConstraintDefinition is called when production columnConstraintDefinition is entered.
func (s *BaseSparkSqlParserListener) EnterColumnConstraintDefinition(ctx *ColumnConstraintDefinitionContext) {}

// ExitColumnConstraintDefinition is called when production columnConstraintDefinition is exited.
func (s *BaseSparkSqlParserListener) ExitColumnConstraintDefinition(ctx *ColumnConstraintDefinitionContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterTableConstraintDefinition is called when production tableConstraintDefinition is entered.
func (s *BaseSparkSqlParserListener) EnterTableConstraintDefinition(ctx *TableConstraintDefinitionContext) {}

// ExitTableConstraintDefinition is called when production tableConstraintDefinition is exited.
func (s *BaseSparkSqlParserListener) ExitTableConstraintDefinition(ctx *TableConstraintDefinitionContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterUniqueSpec is called when production uniqueSpec is entered.
func (s *BaseSparkSqlParserListener) EnterUniqueSpec(ctx *UniqueSpecContext) {}

// ExitUniqueSpec is called when production uniqueSpec is exited.
func (s *BaseSparkSqlParserListener) ExitUniqueSpec(ctx *UniqueSpecContext) {}

// EnterUniqueConstraint is called when production uniqueConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterUniqueConstraint(ctx *UniqueConstraintContext) {}

// ExitUniqueConstraint is called when production uniqueConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitUniqueConstraint(ctx *UniqueConstraintContext) {}

// EnterReferenceSpec is called when production referenceSpec is entered.
func (s *BaseSparkSqlParserListener) EnterReferenceSpec(ctx *ReferenceSpecContext) {}

// ExitReferenceSpec is called when production referenceSpec is exited.
func (s *BaseSparkSqlParserListener) ExitReferenceSpec(ctx *ReferenceSpecContext) {}

// EnterForeignKeyConstraint is called when production foreignKeyConstraint is entered.
func (s *BaseSparkSqlParserListener) EnterForeignKeyConstraint(ctx *ForeignKeyConstraintContext) {}

// ExitForeignKeyConstraint is called when production foreignKeyConstraint is exited.
func (s *BaseSparkSqlParserListener) ExitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) {}

// EnterConstraintCharacteristic is called when production constraintCharacteristic is entered.
func (s *BaseSparkSqlParserListener) EnterConstraintCharacteristic(ctx *ConstraintCharacteristicContext) {}

// ExitConstraintCharacteristic is called when production constraintCharacteristic is exited.
func (s *BaseSparkSqlParserListener) ExitConstraintCharacteristic(ctx *ConstraintCharacteristicContext) {}

// EnterEnforcedCharacteristic is called when production enforcedCharacteristic is entered.
func (s *BaseSparkSqlParserListener) EnterEnforcedCharacteristic(ctx *EnforcedCharacteristicContext) {}

// ExitEnforcedCharacteristic is called when production enforcedCharacteristic is exited.
func (s *BaseSparkSqlParserListener) ExitEnforcedCharacteristic(ctx *EnforcedCharacteristicContext) {}

// EnterRelyCharacteristic is called when production relyCharacteristic is entered.
func (s *BaseSparkSqlParserListener) EnterRelyCharacteristic(ctx *RelyCharacteristicContext) {}

// ExitRelyCharacteristic is called when production relyCharacteristic is exited.
func (s *BaseSparkSqlParserListener) ExitRelyCharacteristic(ctx *RelyCharacteristicContext) {}

// EnterAlterColumnSpecList is called when production alterColumnSpecList is entered.
func (s *BaseSparkSqlParserListener) EnterAlterColumnSpecList(ctx *AlterColumnSpecListContext) {}

// ExitAlterColumnSpecList is called when production alterColumnSpecList is exited.
func (s *BaseSparkSqlParserListener) ExitAlterColumnSpecList(ctx *AlterColumnSpecListContext) {}

// EnterAlterColumnSpec is called when production alterColumnSpec is entered.
func (s *BaseSparkSqlParserListener) EnterAlterColumnSpec(ctx *AlterColumnSpecContext) {}

// ExitAlterColumnSpec is called when production alterColumnSpec is exited.
func (s *BaseSparkSqlParserListener) ExitAlterColumnSpec(ctx *AlterColumnSpecContext) {}

// EnterAlterColumnAction is called when production alterColumnAction is entered.
func (s *BaseSparkSqlParserListener) EnterAlterColumnAction(ctx *AlterColumnActionContext) {}

// ExitAlterColumnAction is called when production alterColumnAction is exited.
func (s *BaseSparkSqlParserListener) ExitAlterColumnAction(ctx *AlterColumnActionContext) {}

// EnterSingleStringLiteralValue is called when production singleStringLiteralValue is entered.
func (s *BaseSparkSqlParserListener) EnterSingleStringLiteralValue(ctx *SingleStringLiteralValueContext) {}

// ExitSingleStringLiteralValue is called when production singleStringLiteralValue is exited.
func (s *BaseSparkSqlParserListener) ExitSingleStringLiteralValue(ctx *SingleStringLiteralValueContext) {}

// EnterSingleDoubleQuotedStringLiteralValue is called when production singleDoubleQuotedStringLiteralValue is entered.
func (s *BaseSparkSqlParserListener) EnterSingleDoubleQuotedStringLiteralValue(ctx *SingleDoubleQuotedStringLiteralValueContext) {}

// ExitSingleDoubleQuotedStringLiteralValue is called when production singleDoubleQuotedStringLiteralValue is exited.
func (s *BaseSparkSqlParserListener) ExitSingleDoubleQuotedStringLiteralValue(ctx *SingleDoubleQuotedStringLiteralValueContext) {}

// EnterSingleStringLit is called when production singleStringLit is entered.
func (s *BaseSparkSqlParserListener) EnterSingleStringLit(ctx *SingleStringLitContext) {}

// ExitSingleStringLit is called when production singleStringLit is exited.
func (s *BaseSparkSqlParserListener) ExitSingleStringLit(ctx *SingleStringLitContext) {}

// EnterNamedParameterMarkerRule is called when production namedParameterMarkerRule is entered.
func (s *BaseSparkSqlParserListener) EnterNamedParameterMarkerRule(ctx *NamedParameterMarkerRuleContext) {}

// ExitNamedParameterMarkerRule is called when production namedParameterMarkerRule is exited.
func (s *BaseSparkSqlParserListener) ExitNamedParameterMarkerRule(ctx *NamedParameterMarkerRuleContext) {}

// EnterPositionalParameterMarkerRule is called when production positionalParameterMarkerRule is entered.
func (s *BaseSparkSqlParserListener) EnterPositionalParameterMarkerRule(ctx *PositionalParameterMarkerRuleContext) {}

// ExitPositionalParameterMarkerRule is called when production positionalParameterMarkerRule is exited.
func (s *BaseSparkSqlParserListener) ExitPositionalParameterMarkerRule(ctx *PositionalParameterMarkerRuleContext) {}

// EnterStringLit is called when production stringLit is entered.
func (s *BaseSparkSqlParserListener) EnterStringLit(ctx *StringLitContext) {}

// ExitStringLit is called when production stringLit is exited.
func (s *BaseSparkSqlParserListener) ExitStringLit(ctx *StringLitContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseSparkSqlParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseSparkSqlParserListener) ExitComment(ctx *CommentContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseSparkSqlParserListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseSparkSqlParserListener) ExitVersion(ctx *VersionContext) {}

// EnterOperatorPipeRightSide is called when production operatorPipeRightSide is entered.
func (s *BaseSparkSqlParserListener) EnterOperatorPipeRightSide(ctx *OperatorPipeRightSideContext) {}

// ExitOperatorPipeRightSide is called when production operatorPipeRightSide is exited.
func (s *BaseSparkSqlParserListener) ExitOperatorPipeRightSide(ctx *OperatorPipeRightSideContext) {}

// EnterOperatorPipeSetAssignmentSeq is called when production operatorPipeSetAssignmentSeq is entered.
func (s *BaseSparkSqlParserListener) EnterOperatorPipeSetAssignmentSeq(ctx *OperatorPipeSetAssignmentSeqContext) {}

// ExitOperatorPipeSetAssignmentSeq is called when production operatorPipeSetAssignmentSeq is exited.
func (s *BaseSparkSqlParserListener) ExitOperatorPipeSetAssignmentSeq(ctx *OperatorPipeSetAssignmentSeqContext) {}

// EnterAnsiNonReserved is called when production ansiNonReserved is entered.
func (s *BaseSparkSqlParserListener) EnterAnsiNonReserved(ctx *AnsiNonReservedContext) {}

// ExitAnsiNonReserved is called when production ansiNonReserved is exited.
func (s *BaseSparkSqlParserListener) ExitAnsiNonReserved(ctx *AnsiNonReservedContext) {}

// EnterStrictNonReserved is called when production strictNonReserved is entered.
func (s *BaseSparkSqlParserListener) EnterStrictNonReserved(ctx *StrictNonReservedContext) {}

// ExitStrictNonReserved is called when production strictNonReserved is exited.
func (s *BaseSparkSqlParserListener) ExitStrictNonReserved(ctx *StrictNonReservedContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseSparkSqlParserListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseSparkSqlParserListener) ExitNonReserved(ctx *NonReservedContext) {}
