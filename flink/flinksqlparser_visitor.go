// Code generated from FlinkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package flink // FlinkSqlParser
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by FlinkSqlParser.
type FlinkSqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FlinkSqlParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#sqlStatement.
	VisitSqlStatement(ctx *SqlStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#ddlStatement.
	VisitDdlStatement(ctx *DdlStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#describeStatement.
	VisitDescribeStatement(ctx *DescribeStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#explainStatement.
	VisitExplainStatement(ctx *ExplainStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#explainDetails.
	VisitExplainDetails(ctx *ExplainDetailsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#explainDetail.
	VisitExplainDetail(ctx *ExplainDetailContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#useModuleStatement.
	VisitUseModuleStatement(ctx *UseModuleStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#showStatement.
	VisitShowStatement(ctx *ShowStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#unloadStatement.
	VisitUnloadStatement(ctx *UnloadStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#resetStatement.
	VisitResetStatement(ctx *ResetStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#jarStatement.
	VisitJarStatement(ctx *JarStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dtAddStatement.
	VisitDtAddStatement(ctx *DtAddStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dtFilePath.
	VisitDtFilePath(ctx *DtFilePathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#simpleCreateTable.
	VisitSimpleCreateTable(ctx *SimpleCreateTableContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createTableAsSelect.
	VisitCreateTableAsSelect(ctx *CreateTableAsSelectContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnOptionDefinition.
	VisitColumnOptionDefinition(ctx *ColumnOptionDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#physicalColumnDefinition.
	VisitPhysicalColumnDefinition(ctx *PhysicalColumnDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnNameCreate.
	VisitColumnNameCreate(ctx *ColumnNameCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnNamePath.
	VisitColumnNamePath(ctx *ColumnNamePathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnNameList.
	VisitColumnNameList(ctx *ColumnNameListContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnType.
	VisitColumnType(ctx *ColumnTypeContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#lengthOneDimension.
	VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#lengthTwoOptionalDimension.
	VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#lengthTwoStringDimension.
	VisitLengthTwoStringDimension(ctx *LengthTwoStringDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#lengthOneTypeDimension.
	VisitLengthOneTypeDimension(ctx *LengthOneTypeDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#mapTypeDimension.
	VisitMapTypeDimension(ctx *MapTypeDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#rowTypeDimension.
	VisitRowTypeDimension(ctx *RowTypeDimensionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnConstraint.
	VisitColumnConstraint(ctx *ColumnConstraintContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#metadataColumnDefinition.
	VisitMetadataColumnDefinition(ctx *MetadataColumnDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#metadataKey.
	VisitMetadataKey(ctx *MetadataKeyContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#computedColumnDefinition.
	VisitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#computedColumnExpression.
	VisitComputedColumnExpression(ctx *ComputedColumnExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#watermarkDefinition.
	VisitWatermarkDefinition(ctx *WatermarkDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tableConstraint.
	VisitTableConstraint(ctx *TableConstraintContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#selfDefinitionClause.
	VisitSelfDefinitionClause(ctx *SelfDefinitionClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#partitionDefinition.
	VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#transformList.
	VisitTransformList(ctx *TransformListContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#identityTransform.
	VisitIdentityTransform(ctx *IdentityTransformContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#applyTransform.
	VisitApplyTransform(ctx *ApplyTransformContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#transformArgument.
	VisitTransformArgument(ctx *TransformArgumentContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#likeDefinition.
	VisitLikeDefinition(ctx *LikeDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#likeOption.
	VisitLikeOption(ctx *LikeOptionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createCatalog.
	VisitCreateCatalog(ctx *CreateCatalogContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#usingClause.
	VisitUsingClause(ctx *UsingClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#jarFileName.
	VisitJarFileName(ctx *JarFileNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#renameDefinition.
	VisitRenameDefinition(ctx *RenameDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#setKeyValueDefinition.
	VisitSetKeyValueDefinition(ctx *SetKeyValueDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#addConstraint.
	VisitAddConstraint(ctx *AddConstraintContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropConstraint.
	VisitDropConstraint(ctx *DropConstraintContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#addUnique.
	VisitAddUnique(ctx *AddUniqueContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#notForced.
	VisitNotForced(ctx *NotForcedContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#alterView.
	VisitAlterView(ctx *AlterViewContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#alterDatabase.
	VisitAlterDatabase(ctx *AlterDatabaseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropCatalog.
	VisitDropCatalog(ctx *DropCatalogContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#insertSimpleStatement.
	VisitInsertSimpleStatement(ctx *InsertSimpleStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#insertPartitionDefinition.
	VisitInsertPartitionDefinition(ctx *InsertPartitionDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#valuesDefinition.
	VisitValuesDefinition(ctx *ValuesDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#valuesRowDefinition.
	VisitValuesRowDefinition(ctx *ValuesRowDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#valueDefinition.
	VisitValueDefinition(ctx *ValueDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#insertMulStatementCompatibility.
	VisitInsertMulStatementCompatibility(ctx *InsertMulStatementCompatibilityContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#insertMulStatement.
	VisitInsertMulStatement(ctx *InsertMulStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#queryStatement.
	VisitQueryStatement(ctx *QueryStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#valuesClause.
	VisitValuesClause(ctx *ValuesClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#withItem.
	VisitWithItem(ctx *WithItemContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#withItemName.
	VisitWithItemName(ctx *WithItemNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#projectItemDefinition.
	VisitProjectItemDefinition(ctx *ProjectItemDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#overWindowItem.
	VisitOverWindowItem(ctx *OverWindowItemContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tableExpression.
	VisitTableExpression(ctx *TableExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePrimary.
	VisitTablePrimary(ctx *TablePrimaryContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#systemTimePeriod.
	VisitSystemTimePeriod(ctx *SystemTimePeriodContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dateTimeExpression.
	VisitDateTimeExpression(ctx *DateTimeExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#inlineDataValueClause.
	VisitInlineDataValueClause(ctx *InlineDataValueClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowTVFClause.
	VisitWindowTVFClause(ctx *WindowTVFClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowTVFExpression.
	VisitWindowTVFExpression(ctx *WindowTVFExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowTVFName.
	VisitWindowTVFName(ctx *WindowTVFNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowTVFParam.
	VisitWindowTVFParam(ctx *WindowTVFParamContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timeIntervalParamName.
	VisitTimeIntervalParamName(ctx *TimeIntervalParamNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnDescriptor.
	VisitColumnDescriptor(ctx *ColumnDescriptorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#joinCondition.
	VisitJoinCondition(ctx *JoinConditionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupItemDefinition.
	VisitGroupItemDefinition(ctx *GroupItemDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupingSets.
	VisitGroupingSets(ctx *GroupingSetsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupingSetsNotationName.
	VisitGroupingSetsNotationName(ctx *GroupingSetsNotationNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupWindowFunction.
	VisitGroupWindowFunction(ctx *GroupWindowFunctionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#groupWindowFunctionName.
	VisitGroupWindowFunctionName(ctx *GroupWindowFunctionNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timeAttrColumn.
	VisitTimeAttrColumn(ctx *TimeAttrColumnContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#namedWindow.
	VisitNamedWindow(ctx *NamedWindowContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowSpec.
	VisitWindowSpec(ctx *WindowSpecContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#matchRecognizeClause.
	VisitMatchRecognizeClause(ctx *MatchRecognizeClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#orderItemDefinition.
	VisitOrderItemDefinition(ctx *OrderItemDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#partitionByClause.
	VisitPartitionByClause(ctx *PartitionByClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#quantifiers.
	VisitQuantifiers(ctx *QuantifiersContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#measuresClause.
	VisitMeasuresClause(ctx *MeasuresClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#patternDefinition.
	VisitPatternDefinition(ctx *PatternDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#patternVariable.
	VisitPatternVariable(ctx *PatternVariableContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#outputMode.
	VisitOutputMode(ctx *OutputModeContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#afterMatchStrategy.
	VisitAfterMatchStrategy(ctx *AfterMatchStrategyContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#patternVariablesDefinition.
	VisitPatternVariablesDefinition(ctx *PatternVariablesDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#windowFrame.
	VisitWindowFrame(ctx *WindowFrameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#frameBound.
	VisitFrameBound(ctx *FrameBoundContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#withinClause.
	VisitWithinClause(ctx *WithinClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#logicalNot.
	VisitLogicalNot(ctx *LogicalNotContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#predicated.
	VisitPredicated(ctx *PredicatedContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#exists.
	VisitExists(ctx *ExistsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#logicalNested.
	VisitLogicalNested(ctx *LogicalNestedContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#logicalBinary.
	VisitLogicalBinary(ctx *LogicalBinaryContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#valueExpressionDefault.
	VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#arithmeticBinary.
	VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#arithmeticUnary.
	VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dereference.
	VisitDereference(ctx *DereferenceContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#simpleCase.
	VisitSimpleCase(ctx *SimpleCaseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#last.
	VisitLast(ctx *LastContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#star.
	VisitStar(ctx *StarContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#subqueryExpression.
	VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#substring.
	VisitSubstring(ctx *SubstringContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#constantDefault.
	VisitConstantDefault(ctx *ConstantDefaultContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#searchedCase.
	VisitSearchedCase(ctx *SearchedCaseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#position.
	VisitPosition(ctx *PositionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#first.
	VisitFirst(ctx *FirstContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionNameCreate.
	VisitFunctionNameCreate(ctx *FunctionNameCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionNameAndParams.
	VisitFunctionNameAndParams(ctx *FunctionNameAndParamsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionNameWithParams.
	VisitFunctionNameWithParams(ctx *FunctionNameWithParamsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#functionParam.
	VisitFunctionParam(ctx *FunctionParamContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#dereferenceDefinition.
	VisitDereferenceDefinition(ctx *DereferenceDefinitionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#correlationName.
	VisitCorrelationName(ctx *CorrelationNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timeIntervalExpression.
	VisitTimeIntervalExpression(ctx *TimeIntervalExpressionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#errorCapturingMultiUnitsInterval.
	VisitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#multiUnitsInterval.
	VisitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#errorCapturingUnitToUnitInterval.
	VisitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#unitToUnitInterval.
	VisitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#intervalValue.
	VisitIntervalValue(ctx *IntervalValueContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#errorCapturingIdentifier.
	VisitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#errorIdent.
	VisitErrorIdent(ctx *ErrorIdentContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#realIdent.
	VisitRealIdent(ctx *RealIdentContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#identifierSeq.
	VisitIdentifierSeq(ctx *IdentifierSeqContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#unquotedIdentifierAlternative.
	VisitUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#quotedIdentifierAlternative.
	VisitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#nonReservedKeywordsAlternative.
	VisitNonReservedKeywordsAlternative(ctx *NonReservedKeywordsAlternativeContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#quotedIdentifier.
	VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#catalogPath.
	VisitCatalogPath(ctx *CatalogPathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#catalogPathCreate.
	VisitCatalogPathCreate(ctx *CatalogPathCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#databasePath.
	VisitDatabasePath(ctx *DatabasePathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#databasePathCreate.
	VisitDatabasePathCreate(ctx *DatabasePathCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePathCreate.
	VisitTablePathCreate(ctx *TablePathCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePath.
	VisitTablePath(ctx *TablePathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#viewPath.
	VisitViewPath(ctx *ViewPathContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#viewPathCreate.
	VisitViewPathCreate(ctx *ViewPathCreateContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#uid.
	VisitUid(ctx *UidContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#withOption.
	VisitWithOption(ctx *WithOptionContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePropertyList.
	VisitTablePropertyList(ctx *TablePropertyListContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tableProperty.
	VisitTableProperty(ctx *TablePropertyContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePropertyKey.
	VisitTablePropertyKey(ctx *TablePropertyKeyContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#tablePropertyValue.
	VisitTablePropertyValue(ctx *TablePropertyValueContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#bitOperator.
	VisitBitOperator(ctx *BitOperatorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timePointLiteral.
	VisitTimePointLiteral(ctx *TimePointLiteralContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timePointUnit.
	VisitTimePointUnit(ctx *TimePointUnitContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#timeIntervalUnit.
	VisitTimeIntervalUnit(ctx *TimeIntervalUnitContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#reservedKeywordsUsedAsFuncParam.
	VisitReservedKeywordsUsedAsFuncParam(ctx *ReservedKeywordsUsedAsFuncParamContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#reservedKeywordsNoParamsUsedAsFuncName.
	VisitReservedKeywordsNoParamsUsedAsFuncName(ctx *ReservedKeywordsNoParamsUsedAsFuncNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#reservedKeywordsFollowParamsUsedAsFuncName.
	VisitReservedKeywordsFollowParamsUsedAsFuncName(ctx *ReservedKeywordsFollowParamsUsedAsFuncNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#reservedKeywordsUsedAsFuncName.
	VisitReservedKeywordsUsedAsFuncName(ctx *ReservedKeywordsUsedAsFuncNameContext) interface{}

	// Visit a parse tree produced by FlinkSqlParser#nonReservedKeywords.
	VisitNonReservedKeywords(ctx *NonReservedKeywordsContext) interface{}

}