// Code generated from SparkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spark // SparkSqlParser
import "github.com/antlr4-go/antlr/v4"


type BaseSparkSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSparkSqlParserVisitor) VisitCompoundOrSingleStatements(ctx *CompoundOrSingleStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCompoundOrSingleStatement(ctx *CompoundOrSingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleCompoundStatement(ctx *SingleCompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBeginEndCompoundBlock(ctx *BeginEndCompoundBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCompoundBody(ctx *CompoundBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetVariableInsideSqlScript(ctx *SetVariableInsideSqlScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSqlStateValue(ctx *SqlStateValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDeclareConditionStatement(ctx *DeclareConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConditionValue(ctx *ConditionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConditionValues(ctx *ConditionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLeaveStatement(ctx *LeaveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIterateStatement(ctx *IterateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSearchedCaseStatement(ctx *SearchedCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSimpleCaseStatement(ctx *SimpleCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBeginLabel(ctx *BeginLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitEndLabel(ctx *EndLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleExpression(ctx *SingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleTableIdentifier(ctx *SingleTableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleMultipartIdentifier(ctx *SingleMultipartIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleFunctionIdentifier(ctx *SingleFunctionIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleDataType(ctx *SingleDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleTableSchema(ctx *SingleTableSchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleRoutineParamList(ctx *SingleRoutineParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStatementDefault(ctx *StatementDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitVisitExecuteImmediate(ctx *VisitExecuteImmediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUse(ctx *UseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUseNamespace(ctx *UseNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetCatalog(ctx *SetCatalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateNamespace(ctx *CreateNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetNamespaceProperties(ctx *SetNamespacePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnsetNamespaceProperties(ctx *UnsetNamespacePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetNamespaceCollation(ctx *SetNamespaceCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetNamespaceLocation(ctx *SetNamespaceLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropNamespace(ctx *DropNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowNamespaces(ctx *ShowNamespacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateTableLike(ctx *CreateTableLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitReplaceTable(ctx *ReplaceTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAnalyze(ctx *AnalyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAnalyzeTables(ctx *AnalyzeTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAddTableColumns(ctx *AddTableColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRenameTableColumn(ctx *RenameTableColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropTableColumns(ctx *DropTableColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRenameTable(ctx *RenameTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetTableProperties(ctx *SetTablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnsetTableProperties(ctx *UnsetTablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterTableAlterColumn(ctx *AlterTableAlterColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitHiveChangeColumn(ctx *HiveChangeColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitHiveReplaceColumns(ctx *HiveReplaceColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetTableSerDe(ctx *SetTableSerDeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAddTablePartition(ctx *AddTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRenameTablePartition(ctx *RenameTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropTablePartitions(ctx *DropTablePartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetTableLocation(ctx *SetTableLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRecoverPartitions(ctx *RecoverPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterClusterBy(ctx *AlterClusterByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterTableCollation(ctx *AlterTableCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAddTableConstraint(ctx *AddTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropTableConstraint(ctx *DropTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateMetricView(ctx *CreateMetricViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateTempViewUsing(ctx *CreateTempViewUsingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterViewQuery(ctx *AlterViewQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterViewSchemaBinding(ctx *AlterViewSchemaBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateUserDefinedFunction(ctx *CreateUserDefinedFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateVariable(ctx *CreateVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropVariable(ctx *DropVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExplain(ctx *ExplainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowTables(ctx *ShowTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowTableExtended(ctx *ShowTableExtendedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowTblProperties(ctx *ShowTblPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowColumns(ctx *ShowColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowViews(ctx *ShowViewsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowPartitions(ctx *ShowPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowFunctions(ctx *ShowFunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowProcedures(ctx *ShowProceduresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowCreateTable(ctx *ShowCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowCurrentNamespace(ctx *ShowCurrentNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShowCatalogs(ctx *ShowCatalogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeFunction(ctx *DescribeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeProcedure(ctx *DescribeProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeNamespace(ctx *DescribeNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeRelation(ctx *DescribeRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeQuery(ctx *DescribeQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCommentNamespace(ctx *CommentNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCommentTable(ctx *CommentTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRefreshTable(ctx *RefreshTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRefreshFunction(ctx *RefreshFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRefreshResource(ctx *RefreshResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCacheTable(ctx *CacheTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUncacheTable(ctx *UncacheTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitClearCache(ctx *ClearCacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLoadData(ctx *LoadDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTruncateTable(ctx *TruncateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRepairTable(ctx *RepairTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitManageResource(ctx *ManageResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFailNativeCommand(ctx *FailNativeCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreatePipelineDataset(ctx *CreatePipelineDatasetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreatePipelineInsertIntoFlow(ctx *CreatePipelineInsertIntoFlowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMaterializedView(ctx *MaterializedViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStreamingTable(ctx *StreamingTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreatePipelineDatasetHeader(ctx *CreatePipelineDatasetHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStreamTableName(ctx *StreamTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFailSetRole(ctx *FailSetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetTimeZone(ctx *SetTimeZoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetVariable(ctx *SetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetQuotedConfiguration(ctx *SetQuotedConfigurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetConfiguration(ctx *SetConfigurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitResetQuotedConfiguration(ctx *ResetQuotedConfigurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitResetConfiguration(ctx *ResetConfigurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExecuteImmediate(ctx *ExecuteImmediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExecuteImmediateUsing(ctx *ExecuteImmediateUsingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTimezone(ctx *TimezoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConfigKey(ctx *ConfigKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConfigValue(ctx *ConfigValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnsupportedHiveNativeCommands(ctx *UnsupportedHiveNativeCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateTableHeader(ctx *CreateTableHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitReplaceTableHeader(ctx *ReplaceTableHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitClusterBySpec(ctx *ClusterBySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBucketSpec(ctx *BucketSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSkewSpec(ctx *SkewSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLocationSpec(ctx *LocationSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSchemaBinding(ctx *SchemaBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCommentSpec(ctx *CommentSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleQuery(ctx *SingleQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInsertOverwriteTable(ctx *InsertOverwriteTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInsertIntoTable(ctx *InsertIntoTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInsertIntoReplaceWhere(ctx *InsertIntoReplaceWhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInsertOverwriteHiveDir(ctx *InsertOverwriteHiveDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInsertOverwriteDir(ctx *InsertOverwriteDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionSpecLocation(ctx *PartitionSpecLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionSpec(ctx *PartitionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionVal(ctx *PartitionValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreatePipelineFlowHeader(ctx *CreatePipelineFlowHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamespaces(ctx *NamespacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeFuncName(ctx *DescribeFuncNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDescribeColName(ctx *DescribeColNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCtes(ctx *CtesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedQuery(ctx *NamedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableProvider(ctx *TableProviderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateTableClauses(ctx *CreateTableClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyWithKeyAndEquals(ctx *PropertyWithKeyAndEqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyWithKeyNoEquals(ctx *PropertyWithKeyNoEqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyKey(ctx *PropertyKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyKeyOrStringLit(ctx *PropertyKeyOrStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyKeyOrStringLitNoCoalesce(ctx *PropertyKeyOrStringLitNoCoalesceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPropertyValue(ctx *PropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExpressionPropertyList(ctx *ExpressionPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExpressionPropertyWithKeyAndEquals(ctx *ExpressionPropertyWithKeyAndEqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExpressionPropertyWithKeyNoEquals(ctx *ExpressionPropertyWithKeyNoEqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConstantList(ctx *ConstantListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNestedConstantList(ctx *NestedConstantListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCreateFileFormat(ctx *CreateFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableFileFormat(ctx *TableFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGenericFileFormat(ctx *GenericFileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStorageHandler(ctx *StorageHandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleInsertQuery(ctx *SingleInsertQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultiInsertQuery(ctx *MultiInsertQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDeleteFromTable(ctx *DeleteFromTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUpdateTable(ctx *UpdateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMergeIntoTable(ctx *MergeIntoTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierReference(ctx *IdentifierReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCatalogIdentifierReference(ctx *CatalogIdentifierReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQueryOrganization(ctx *QueryOrganizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultiInsertQueryBody(ctx *MultiInsertQueryBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOperatorPipeStatement(ctx *OperatorPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQueryTermDefault(ctx *QueryTermDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFromStmt(ctx *FromStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInlineTableDefault1(ctx *InlineTableDefault1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFromStatement(ctx *FromStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFromStatementBody(ctx *FromStatementBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTransformQuerySpecification(ctx *TransformQuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRegularQuerySpecification(ctx *RegularQuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTransformClause(ctx *TransformClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetClause(ctx *SetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMatchedClause(ctx *MatchedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNotMatchedClause(ctx *NotMatchedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNotMatchedBySourceClause(ctx *NotMatchedBySourceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMatchedAction(ctx *MatchedActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNotMatchedAction(ctx *NotMatchedActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNotMatchedBySourceAction(ctx *NotMatchedBySourceActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExceptClause(ctx *ExceptClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitHint(ctx *HintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitHintStatement(ctx *HintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTemporalClause(ctx *TemporalClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAggregationClause(ctx *AggregationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGroupingAnalytics(ctx *GroupingAnalyticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGroupingElement(ctx *GroupingElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPivotColumn(ctx *PivotColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPivotValue(ctx *PivotValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotClause(ctx *UnpivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotNullClause(ctx *UnpivotNullClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotOperator(ctx *UnpivotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotSingleValueColumnClause(ctx *UnpivotSingleValueColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotMultiValueColumnClause(ctx *UnpivotMultiValueColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotColumnSet(ctx *UnpivotColumnSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotValueColumn(ctx *UnpivotValueColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotNameColumn(ctx *UnpivotNameColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotColumnAndAlias(ctx *UnpivotColumnAndAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotColumn(ctx *UnpivotColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnpivotAlias(ctx *UnpivotAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLateralView(ctx *LateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWatermarkClause(ctx *WatermarkClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRelationExtension(ctx *RelationExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJoinType(ctx *JoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSample(ctx *SampleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSampleByPercentile(ctx *SampleByPercentileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSampleByRows(ctx *SampleByRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSampleByBucket(ctx *SampleByBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSampleByBytes(ctx *SampleByBytesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierSeq(ctx *IdentifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOrderedIdentifierList(ctx *OrderedIdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOrderedIdentifier(ctx *OrderedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierCommentList(ctx *IdentifierCommentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierComment(ctx *IdentifierCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStreamRelation(ctx *StreamRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAliasedQuery(ctx *AliasedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInlineTableDefault2(ctx *InlineTableDefault2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableValuedFunction(ctx *TableValuedFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOptionsClause(ctx *OptionsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionTableSubqueryArgument(ctx *FunctionTableSubqueryArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableArgumentPartitioning(ctx *TableArgumentPartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionTableNamedArgumentExpression(ctx *FunctionTableNamedArgumentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionTableReferenceArgument(ctx *FunctionTableReferenceArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionTableArgument(ctx *FunctionTableArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionTable(ctx *FunctionTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRowFormatSerde(ctx *RowFormatSerdeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRowFormatDelimited(ctx *RowFormatDelimitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultipartIdentifierList(ctx *MultipartIdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultipartIdentifier(ctx *MultipartIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultipartIdentifierPropertyList(ctx *MultipartIdentifierPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultipartIdentifierProperty(ctx *MultipartIdentifierPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableIdentifier(ctx *TableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionIdentifier(ctx *FunctionIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedExpression(ctx *NamedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedExpressionSeq(ctx *NamedExpressionSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionFieldList(ctx *PartitionFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionTransform(ctx *PartitionTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPartitionColumn(ctx *PartitionColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentityTransform(ctx *IdentityTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitApplyTransform(ctx *ApplyTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTransformArgument(ctx *TransformArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedArgumentExpression(ctx *NamedArgumentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionArgument(ctx *FunctionArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExpressionSeq(ctx *ExpressionSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPredicated(ctx *PredicatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitErrorCapturingNot(ctx *ErrorCapturingNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitShiftOperator(ctx *ShiftOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDatetimeUnit(ctx *DatetimeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCastByColon(ctx *CastByColonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTimestampadd(ctx *TimestampaddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSubstring(ctx *SubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAny_value(ctx *Any_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTrim(ctx *TrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSemiStructuredExtract(ctx *SemiStructuredExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCurrentLike(ctx *CurrentLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLast(ctx *LastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStar(ctx *StarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOverlay(ctx *OverlayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTimestampdiff(ctx *TimestampdiffContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConstantDefault(ctx *ConstantDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPosition(ctx *PositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFirst(ctx *FirstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSemiStructuredExtractionPath(ctx *SemiStructuredExtractionPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJsonPathIdentifier(ctx *JsonPathIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJsonPathBracketedIdentifier(ctx *JsonPathBracketedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJsonPathFirstPart(ctx *JsonPathFirstPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitJsonPathParts(ctx *JsonPathPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLiteralType(ctx *LiteralTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPosParameterLiteral(ctx *PosParameterLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedParameterLiteral(ctx *NamedParameterLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTypeConstructor(ctx *TypeConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedParameterMarker(ctx *NamedParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitArithmeticOperator(ctx *ArithmeticOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPredicateOperator(ctx *PredicateOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIntervalValue(ctx *IntervalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnitInMultiUnits(ctx *UnitInMultiUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnitInUnitToUnit(ctx *UnitInUnitToUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColPosition(ctx *ColPositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCollationSpec(ctx *CollationSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCollateClause(ctx *CollateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNonTrivialPrimitiveType(ctx *NonTrivialPrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTrivialPrimitiveType(ctx *TrivialPrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComplexDataType(ctx *ComplexDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPrimitiveDataType(ctx *PrimitiveDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQualifiedColTypeWithPositionList(ctx *QualifiedColTypeWithPositionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQualifiedColTypeWithPosition(ctx *QualifiedColTypeWithPositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColDefinitionDescriptorWithPosition(ctx *ColDefinitionDescriptorWithPositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDefaultExpression(ctx *DefaultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitVariableDefaultExpression(ctx *VariableDefaultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColTypeList(ctx *ColTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColType(ctx *ColTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableElementList(ctx *TableElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableElement(ctx *TableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColDefinitionList(ctx *ColDefinitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColDefinition(ctx *ColDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColDefinitionOption(ctx *ColDefinitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitGeneratedColumn(ctx *GeneratedColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentityColumn(ctx *IdentityColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentityColSpec(ctx *IdentityColSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSequenceGeneratorOption(ctx *SequenceGeneratorOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSequenceGeneratorStartOrStep(ctx *SequenceGeneratorStartOrStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComplexColTypeList(ctx *ComplexColTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComplexColType(ctx *ComplexColTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCodeLiteral(ctx *CodeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRoutineCharacteristics(ctx *RoutineCharacteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRoutineLanguage(ctx *RoutineLanguageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSpecificName(ctx *SpecificNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDeterministic(ctx *DeterministicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSqlDataAccess(ctx *SqlDataAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNullCall(ctx *NullCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRightsClause(ctx *RightsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedWindow(ctx *NamedWindowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWindowRef(ctx *WindowRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWindowDef(ctx *WindowDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFrameBound(ctx *FrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitErrorIdent(ctx *ErrorIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRealIdent(ctx *RealIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIdentifierLiteral(ctx *IdentifierLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSimpleUnquotedIdentifier(ctx *SimpleUnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSimpleQuotedIdentifierAlternative(ctx *SimpleQuotedIdentifierAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitExponentLiteral(ctx *ExponentLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBigIntLiteral(ctx *BigIntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSmallIntLiteral(ctx *SmallIntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTinyIntLiteral(ctx *TinyIntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitDoubleLiteral(ctx *DoubleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitBigDecimalLiteral(ctx *BigDecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitIntegerVal(ctx *IntegerValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitParameterIntegerValue(ctx *ParameterIntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColumnConstraintDefinition(ctx *ColumnConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitColumnConstraint(ctx *ColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableConstraintDefinition(ctx *TableConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitTableConstraint(ctx *TableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUniqueSpec(ctx *UniqueSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitUniqueConstraint(ctx *UniqueConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitReferenceSpec(ctx *ReferenceSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitConstraintCharacteristic(ctx *ConstraintCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitEnforcedCharacteristic(ctx *EnforcedCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitRelyCharacteristic(ctx *RelyCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterColumnSpecList(ctx *AlterColumnSpecListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterColumnSpec(ctx *AlterColumnSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAlterColumnAction(ctx *AlterColumnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleStringLiteralValue(ctx *SingleStringLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleDoubleQuotedStringLiteralValue(ctx *SingleDoubleQuotedStringLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitSingleStringLit(ctx *SingleStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNamedParameterMarkerRule(ctx *NamedParameterMarkerRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitPositionalParameterMarkerRule(ctx *PositionalParameterMarkerRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStringLit(ctx *StringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOperatorPipeRightSide(ctx *OperatorPipeRightSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitOperatorPipeSetAssignmentSeq(ctx *OperatorPipeSetAssignmentSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitAnsiNonReserved(ctx *AnsiNonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitStrictNonReserved(ctx *StrictNonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparkSqlParserVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
