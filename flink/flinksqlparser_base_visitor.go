// Code generated from FlinkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package flink // FlinkSqlParser
import "github.com/antlr4-go/antlr/v4"


type BaseFlinkSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFlinkSqlParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSqlStatement(ctx *SqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitEmptyStatement(ctx *EmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDdlStatement(ctx *DdlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDescribeStatement(ctx *DescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitExplainStatement(ctx *ExplainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitExplainDetails(ctx *ExplainDetailsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitExplainDetail(ctx *ExplainDetailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUseStatement(ctx *UseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUseModuleStatement(ctx *UseModuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitShowStatement(ctx *ShowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUnloadStatement(ctx *UnloadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitResetStatement(ctx *ResetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitJarStatement(ctx *JarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDtAddStatement(ctx *DtAddStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDtFilePath(ctx *DtFilePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSimpleCreateTable(ctx *SimpleCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateTableAsSelect(ctx *CreateTableAsSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnOptionDefinition(ctx *ColumnOptionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPhysicalColumnDefinition(ctx *PhysicalColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnNameCreate(ctx *ColumnNameCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnNamePath(ctx *ColumnNamePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnNameList(ctx *ColumnNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnType(ctx *ColumnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLengthTwoStringDimension(ctx *LengthTwoStringDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLengthOneTypeDimension(ctx *LengthOneTypeDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMapTypeDimension(ctx *MapTypeDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitRowTypeDimension(ctx *RowTypeDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnConstraint(ctx *ColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMetadataColumnDefinition(ctx *MetadataColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMetadataKey(ctx *MetadataKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitComputedColumnExpression(ctx *ComputedColumnExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWatermarkDefinition(ctx *WatermarkDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTableConstraint(ctx *TableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSelfDefinitionClause(ctx *SelfDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTransformList(ctx *TransformListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIdentityTransform(ctx *IdentityTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitApplyTransform(ctx *ApplyTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTransformArgument(ctx *TransformArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLikeDefinition(ctx *LikeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLikeOption(ctx *LikeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateCatalog(ctx *CreateCatalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUsingClause(ctx *UsingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitJarFileName(ctx *JarFileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitRenameDefinition(ctx *RenameDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSetKeyValueDefinition(ctx *SetKeyValueDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAddConstraint(ctx *AddConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropConstraint(ctx *DropConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAddUnique(ctx *AddUniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitNotForced(ctx *NotForcedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAlterView(ctx *AlterViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAlterDatabase(ctx *AlterDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropCatalog(ctx *DropCatalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInsertSimpleStatement(ctx *InsertSimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInsertPartitionDefinition(ctx *InsertPartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitValuesDefinition(ctx *ValuesDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitValuesRowDefinition(ctx *ValuesRowDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitValueDefinition(ctx *ValueDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInsertMulStatementCompatibility(ctx *InsertMulStatementCompatibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInsertMulStatement(ctx *InsertMulStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitValuesClause(ctx *ValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWithItem(ctx *WithItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWithItemName(ctx *WithItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitProjectItemDefinition(ctx *ProjectItemDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitOverWindowItem(ctx *OverWindowItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTableExpression(ctx *TableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePrimary(ctx *TablePrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSystemTimePeriod(ctx *SystemTimePeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDateTimeExpression(ctx *DateTimeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitInlineDataValueClause(ctx *InlineDataValueClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowTVFClause(ctx *WindowTVFClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowTVFExpression(ctx *WindowTVFExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowTVFName(ctx *WindowTVFNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowTVFParam(ctx *WindowTVFParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimeIntervalParamName(ctx *TimeIntervalParamNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnDescriptor(ctx *ColumnDescriptorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitJoinCondition(ctx *JoinConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupItemDefinition(ctx *GroupItemDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupingSets(ctx *GroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupingSetsNotationName(ctx *GroupingSetsNotationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupWindowFunction(ctx *GroupWindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitGroupWindowFunctionName(ctx *GroupWindowFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimeAttrColumn(ctx *TimeAttrColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitNamedWindow(ctx *NamedWindowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowSpec(ctx *WindowSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMatchRecognizeClause(ctx *MatchRecognizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitOrderItemDefinition(ctx *OrderItemDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPartitionByClause(ctx *PartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitQuantifiers(ctx *QuantifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMeasuresClause(ctx *MeasuresClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPatternDefinition(ctx *PatternDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPatternVariable(ctx *PatternVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitOutputMode(ctx *OutputModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitAfterMatchStrategy(ctx *AfterMatchStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPatternVariablesDefinition(ctx *PatternVariablesDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFrameBound(ctx *FrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWithinClause(ctx *WithinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPredicated(ctx *PredicatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLogicalNested(ctx *LogicalNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLast(ctx *LastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitStar(ctx *StarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSubstring(ctx *SubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitConstantDefault(ctx *ConstantDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitPosition(ctx *PositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFirst(ctx *FirstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionNameCreate(ctx *FunctionNameCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionNameAndParams(ctx *FunctionNameAndParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionNameWithParams(ctx *FunctionNameWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitFunctionParam(ctx *FunctionParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDereferenceDefinition(ctx *DereferenceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCorrelationName(ctx *CorrelationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimeIntervalExpression(ctx *TimeIntervalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIntervalValue(ctx *IntervalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitErrorIdent(ctx *ErrorIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitRealIdent(ctx *RealIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIdentifierSeq(ctx *IdentifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitNonReservedKeywordsAlternative(ctx *NonReservedKeywordsAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCatalogPath(ctx *CatalogPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitCatalogPathCreate(ctx *CatalogPathCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDatabasePath(ctx *DatabasePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDatabasePathCreate(ctx *DatabasePathCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePathCreate(ctx *TablePathCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePath(ctx *TablePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitViewPath(ctx *ViewPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitViewPathCreate(ctx *ViewPathCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUid(ctx *UidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitWithOption(ctx *WithOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePropertyList(ctx *TablePropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTableProperty(ctx *TablePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePropertyKey(ctx *TablePropertyKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTablePropertyValue(ctx *TablePropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitBitOperator(ctx *BitOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimePointLiteral(ctx *TimePointLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimePointUnit(ctx *TimePointUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitTimeIntervalUnit(ctx *TimeIntervalUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitReservedKeywordsUsedAsFuncParam(ctx *ReservedKeywordsUsedAsFuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitReservedKeywordsNoParamsUsedAsFuncName(ctx *ReservedKeywordsNoParamsUsedAsFuncNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitReservedKeywordsFollowParamsUsedAsFuncName(ctx *ReservedKeywordsFollowParamsUsedAsFuncNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitReservedKeywordsUsedAsFuncName(ctx *ReservedKeywordsUsedAsFuncNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlinkSqlParserVisitor) VisitNonReservedKeywords(ctx *NonReservedKeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}
