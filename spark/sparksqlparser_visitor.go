// Code generated from SparkSqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spark // SparkSqlParser
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by SparkSqlParser.
type SparkSqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SparkSqlParser#compoundOrSingleStatements.
	VisitCompoundOrSingleStatements(ctx *CompoundOrSingleStatementsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#compoundOrSingleStatement.
	VisitCompoundOrSingleStatement(ctx *CompoundOrSingleStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleCompoundStatement.
	VisitSingleCompoundStatement(ctx *SingleCompoundStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#beginEndCompoundBlock.
	VisitBeginEndCompoundBlock(ctx *BeginEndCompoundBlockContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#compoundBody.
	VisitCompoundBody(ctx *CompoundBodyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setVariableInsideSqlScript.
	VisitSetVariableInsideSqlScript(ctx *SetVariableInsideSqlScriptContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sqlStateValue.
	VisitSqlStateValue(ctx *SqlStateValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#declareConditionStatement.
	VisitDeclareConditionStatement(ctx *DeclareConditionStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#conditionValue.
	VisitConditionValue(ctx *ConditionValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#conditionValues.
	VisitConditionValues(ctx *ConditionValuesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#declareHandlerStatement.
	VisitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#ifElseStatement.
	VisitIfElseStatement(ctx *IfElseStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#leaveStatement.
	VisitLeaveStatement(ctx *LeaveStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#iterateStatement.
	VisitIterateStatement(ctx *IterateStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#searchedCaseStatement.
	VisitSearchedCaseStatement(ctx *SearchedCaseStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#simpleCaseStatement.
	VisitSimpleCaseStatement(ctx *SimpleCaseStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#beginLabel.
	VisitBeginLabel(ctx *BeginLabelContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#endLabel.
	VisitEndLabel(ctx *EndLabelContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleExpression.
	VisitSingleExpression(ctx *SingleExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleTableIdentifier.
	VisitSingleTableIdentifier(ctx *SingleTableIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleMultipartIdentifier.
	VisitSingleMultipartIdentifier(ctx *SingleMultipartIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleFunctionIdentifier.
	VisitSingleFunctionIdentifier(ctx *SingleFunctionIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleDataType.
	VisitSingleDataType(ctx *SingleDataTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleTableSchema.
	VisitSingleTableSchema(ctx *SingleTableSchemaContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleRoutineParamList.
	VisitSingleRoutineParamList(ctx *SingleRoutineParamListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#statementDefault.
	VisitStatementDefault(ctx *StatementDefaultContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#visitExecuteImmediate.
	VisitVisitExecuteImmediate(ctx *VisitExecuteImmediateContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#use.
	VisitUse(ctx *UseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#useNamespace.
	VisitUseNamespace(ctx *UseNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setCatalog.
	VisitSetCatalog(ctx *SetCatalogContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createNamespace.
	VisitCreateNamespace(ctx *CreateNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setNamespaceProperties.
	VisitSetNamespaceProperties(ctx *SetNamespacePropertiesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unsetNamespaceProperties.
	VisitUnsetNamespaceProperties(ctx *UnsetNamespacePropertiesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setNamespaceCollation.
	VisitSetNamespaceCollation(ctx *SetNamespaceCollationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setNamespaceLocation.
	VisitSetNamespaceLocation(ctx *SetNamespaceLocationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropNamespace.
	VisitDropNamespace(ctx *DropNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showNamespaces.
	VisitShowNamespaces(ctx *ShowNamespacesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createTableLike.
	VisitCreateTableLike(ctx *CreateTableLikeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#replaceTable.
	VisitReplaceTable(ctx *ReplaceTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#analyze.
	VisitAnalyze(ctx *AnalyzeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#analyzeTables.
	VisitAnalyzeTables(ctx *AnalyzeTablesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#addTableColumns.
	VisitAddTableColumns(ctx *AddTableColumnsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#renameTableColumn.
	VisitRenameTableColumn(ctx *RenameTableColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropTableColumns.
	VisitDropTableColumns(ctx *DropTableColumnsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#renameTable.
	VisitRenameTable(ctx *RenameTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setTableProperties.
	VisitSetTableProperties(ctx *SetTablePropertiesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unsetTableProperties.
	VisitUnsetTableProperties(ctx *UnsetTablePropertiesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterTableAlterColumn.
	VisitAlterTableAlterColumn(ctx *AlterTableAlterColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#hiveChangeColumn.
	VisitHiveChangeColumn(ctx *HiveChangeColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#hiveReplaceColumns.
	VisitHiveReplaceColumns(ctx *HiveReplaceColumnsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setTableSerDe.
	VisitSetTableSerDe(ctx *SetTableSerDeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#addTablePartition.
	VisitAddTablePartition(ctx *AddTablePartitionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#renameTablePartition.
	VisitRenameTablePartition(ctx *RenameTablePartitionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropTablePartitions.
	VisitDropTablePartitions(ctx *DropTablePartitionsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setTableLocation.
	VisitSetTableLocation(ctx *SetTableLocationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#recoverPartitions.
	VisitRecoverPartitions(ctx *RecoverPartitionsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterClusterBy.
	VisitAlterClusterBy(ctx *AlterClusterByContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterTableCollation.
	VisitAlterTableCollation(ctx *AlterTableCollationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#addTableConstraint.
	VisitAddTableConstraint(ctx *AddTableConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropTableConstraint.
	VisitDropTableConstraint(ctx *DropTableConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createMetricView.
	VisitCreateMetricView(ctx *CreateMetricViewContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createTempViewUsing.
	VisitCreateTempViewUsing(ctx *CreateTempViewUsingContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterViewQuery.
	VisitAlterViewQuery(ctx *AlterViewQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterViewSchemaBinding.
	VisitAlterViewSchemaBinding(ctx *AlterViewSchemaBindingContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createUserDefinedFunction.
	VisitCreateUserDefinedFunction(ctx *CreateUserDefinedFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createVariable.
	VisitCreateVariable(ctx *CreateVariableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropVariable.
	VisitDropVariable(ctx *DropVariableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#explain.
	VisitExplain(ctx *ExplainContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showTables.
	VisitShowTables(ctx *ShowTablesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showTableExtended.
	VisitShowTableExtended(ctx *ShowTableExtendedContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showTblProperties.
	VisitShowTblProperties(ctx *ShowTblPropertiesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showColumns.
	VisitShowColumns(ctx *ShowColumnsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showViews.
	VisitShowViews(ctx *ShowViewsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showPartitions.
	VisitShowPartitions(ctx *ShowPartitionsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showFunctions.
	VisitShowFunctions(ctx *ShowFunctionsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showProcedures.
	VisitShowProcedures(ctx *ShowProceduresContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showCreateTable.
	VisitShowCreateTable(ctx *ShowCreateTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showCurrentNamespace.
	VisitShowCurrentNamespace(ctx *ShowCurrentNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#showCatalogs.
	VisitShowCatalogs(ctx *ShowCatalogsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeFunction.
	VisitDescribeFunction(ctx *DescribeFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeProcedure.
	VisitDescribeProcedure(ctx *DescribeProcedureContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeNamespace.
	VisitDescribeNamespace(ctx *DescribeNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeRelation.
	VisitDescribeRelation(ctx *DescribeRelationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeQuery.
	VisitDescribeQuery(ctx *DescribeQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#commentNamespace.
	VisitCommentNamespace(ctx *CommentNamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#commentTable.
	VisitCommentTable(ctx *CommentTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#refreshTable.
	VisitRefreshTable(ctx *RefreshTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#refreshFunction.
	VisitRefreshFunction(ctx *RefreshFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#refreshResource.
	VisitRefreshResource(ctx *RefreshResourceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#cacheTable.
	VisitCacheTable(ctx *CacheTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#uncacheTable.
	VisitUncacheTable(ctx *UncacheTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#clearCache.
	VisitClearCache(ctx *ClearCacheContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#loadData.
	VisitLoadData(ctx *LoadDataContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#repairTable.
	VisitRepairTable(ctx *RepairTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#manageResource.
	VisitManageResource(ctx *ManageResourceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#failNativeCommand.
	VisitFailNativeCommand(ctx *FailNativeCommandContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createPipelineDataset.
	VisitCreatePipelineDataset(ctx *CreatePipelineDatasetContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createPipelineInsertIntoFlow.
	VisitCreatePipelineInsertIntoFlow(ctx *CreatePipelineInsertIntoFlowContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#materializedView.
	VisitMaterializedView(ctx *MaterializedViewContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#streamingTable.
	VisitStreamingTable(ctx *StreamingTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createPipelineDatasetHeader.
	VisitCreatePipelineDatasetHeader(ctx *CreatePipelineDatasetHeaderContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#streamTableName.
	VisitStreamTableName(ctx *StreamTableNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#failSetRole.
	VisitFailSetRole(ctx *FailSetRoleContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setTimeZone.
	VisitSetTimeZone(ctx *SetTimeZoneContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setVariable.
	VisitSetVariable(ctx *SetVariableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setQuotedConfiguration.
	VisitSetQuotedConfiguration(ctx *SetQuotedConfigurationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setConfiguration.
	VisitSetConfiguration(ctx *SetConfigurationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#resetQuotedConfiguration.
	VisitResetQuotedConfiguration(ctx *ResetQuotedConfigurationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#resetConfiguration.
	VisitResetConfiguration(ctx *ResetConfigurationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#executeImmediate.
	VisitExecuteImmediate(ctx *ExecuteImmediateContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#executeImmediateUsing.
	VisitExecuteImmediateUsing(ctx *ExecuteImmediateUsingContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#timezone.
	VisitTimezone(ctx *TimezoneContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#configKey.
	VisitConfigKey(ctx *ConfigKeyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#configValue.
	VisitConfigValue(ctx *ConfigValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unsupportedHiveNativeCommands.
	VisitUnsupportedHiveNativeCommands(ctx *UnsupportedHiveNativeCommandsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createTableHeader.
	VisitCreateTableHeader(ctx *CreateTableHeaderContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#replaceTableHeader.
	VisitReplaceTableHeader(ctx *ReplaceTableHeaderContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#clusterBySpec.
	VisitClusterBySpec(ctx *ClusterBySpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#bucketSpec.
	VisitBucketSpec(ctx *BucketSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#skewSpec.
	VisitSkewSpec(ctx *SkewSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#locationSpec.
	VisitLocationSpec(ctx *LocationSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#schemaBinding.
	VisitSchemaBinding(ctx *SchemaBindingContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#commentSpec.
	VisitCommentSpec(ctx *CommentSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleQuery.
	VisitSingleQuery(ctx *SingleQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#insertOverwriteTable.
	VisitInsertOverwriteTable(ctx *InsertOverwriteTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#insertIntoTable.
	VisitInsertIntoTable(ctx *InsertIntoTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#insertIntoReplaceWhere.
	VisitInsertIntoReplaceWhere(ctx *InsertIntoReplaceWhereContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#insertOverwriteHiveDir.
	VisitInsertOverwriteHiveDir(ctx *InsertOverwriteHiveDirContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#insertOverwriteDir.
	VisitInsertOverwriteDir(ctx *InsertOverwriteDirContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionSpecLocation.
	VisitPartitionSpecLocation(ctx *PartitionSpecLocationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionSpec.
	VisitPartitionSpec(ctx *PartitionSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionVal.
	VisitPartitionVal(ctx *PartitionValContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createPipelineFlowHeader.
	VisitCreatePipelineFlowHeader(ctx *CreatePipelineFlowHeaderContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namespaces.
	VisitNamespaces(ctx *NamespacesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeFuncName.
	VisitDescribeFuncName(ctx *DescribeFuncNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#describeColName.
	VisitDescribeColName(ctx *DescribeColNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#ctes.
	VisitCtes(ctx *CtesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedQuery.
	VisitNamedQuery(ctx *NamedQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableProvider.
	VisitTableProvider(ctx *TableProviderContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createTableClauses.
	VisitCreateTableClauses(ctx *CreateTableClausesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyList.
	VisitPropertyList(ctx *PropertyListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyWithKeyAndEquals.
	VisitPropertyWithKeyAndEquals(ctx *PropertyWithKeyAndEqualsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyWithKeyNoEquals.
	VisitPropertyWithKeyNoEquals(ctx *PropertyWithKeyNoEqualsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyKey.
	VisitPropertyKey(ctx *PropertyKeyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyKeyOrStringLit.
	VisitPropertyKeyOrStringLit(ctx *PropertyKeyOrStringLitContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyKeyOrStringLitNoCoalesce.
	VisitPropertyKeyOrStringLitNoCoalesce(ctx *PropertyKeyOrStringLitNoCoalesceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#propertyValue.
	VisitPropertyValue(ctx *PropertyValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#expressionPropertyList.
	VisitExpressionPropertyList(ctx *ExpressionPropertyListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#expressionPropertyWithKeyAndEquals.
	VisitExpressionPropertyWithKeyAndEquals(ctx *ExpressionPropertyWithKeyAndEqualsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#expressionPropertyWithKeyNoEquals.
	VisitExpressionPropertyWithKeyNoEquals(ctx *ExpressionPropertyWithKeyNoEqualsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#constantList.
	VisitConstantList(ctx *ConstantListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#nestedConstantList.
	VisitNestedConstantList(ctx *NestedConstantListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#createFileFormat.
	VisitCreateFileFormat(ctx *CreateFileFormatContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableFileFormat.
	VisitTableFileFormat(ctx *TableFileFormatContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#genericFileFormat.
	VisitGenericFileFormat(ctx *GenericFileFormatContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#storageHandler.
	VisitStorageHandler(ctx *StorageHandlerContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleInsertQuery.
	VisitSingleInsertQuery(ctx *SingleInsertQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multiInsertQuery.
	VisitMultiInsertQuery(ctx *MultiInsertQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#deleteFromTable.
	VisitDeleteFromTable(ctx *DeleteFromTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#updateTable.
	VisitUpdateTable(ctx *UpdateTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#mergeIntoTable.
	VisitMergeIntoTable(ctx *MergeIntoTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierReference.
	VisitIdentifierReference(ctx *IdentifierReferenceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#catalogIdentifierReference.
	VisitCatalogIdentifierReference(ctx *CatalogIdentifierReferenceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#queryOrganization.
	VisitQueryOrganization(ctx *QueryOrganizationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multiInsertQueryBody.
	VisitMultiInsertQueryBody(ctx *MultiInsertQueryBodyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#operatorPipeStatement.
	VisitOperatorPipeStatement(ctx *OperatorPipeStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#queryTermDefault.
	VisitQueryTermDefault(ctx *QueryTermDefaultContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setOperation.
	VisitSetOperation(ctx *SetOperationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#queryPrimaryDefault.
	VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#fromStmt.
	VisitFromStmt(ctx *FromStmtContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#inlineTableDefault1.
	VisitInlineTableDefault1(ctx *InlineTableDefault1Context) interface{}

	// Visit a parse tree produced by SparkSqlParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#fromStatement.
	VisitFromStatement(ctx *FromStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#fromStatementBody.
	VisitFromStatementBody(ctx *FromStatementBodyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#transformQuerySpecification.
	VisitTransformQuerySpecification(ctx *TransformQuerySpecificationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#regularQuerySpecification.
	VisitRegularQuerySpecification(ctx *RegularQuerySpecificationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#transformClause.
	VisitTransformClause(ctx *TransformClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setClause.
	VisitSetClause(ctx *SetClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#matchedClause.
	VisitMatchedClause(ctx *MatchedClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#notMatchedClause.
	VisitNotMatchedClause(ctx *NotMatchedClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#notMatchedBySourceClause.
	VisitNotMatchedBySourceClause(ctx *NotMatchedBySourceClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#matchedAction.
	VisitMatchedAction(ctx *MatchedActionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#notMatchedAction.
	VisitNotMatchedAction(ctx *NotMatchedActionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#notMatchedBySourceAction.
	VisitNotMatchedBySourceAction(ctx *NotMatchedBySourceActionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#exceptClause.
	VisitExceptClause(ctx *ExceptClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#hint.
	VisitHint(ctx *HintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#hintStatement.
	VisitHintStatement(ctx *HintStatementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#temporalClause.
	VisitTemporalClause(ctx *TemporalClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#aggregationClause.
	VisitAggregationClause(ctx *AggregationClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#groupingAnalytics.
	VisitGroupingAnalytics(ctx *GroupingAnalyticsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#groupingElement.
	VisitGroupingElement(ctx *GroupingElementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#groupingSet.
	VisitGroupingSet(ctx *GroupingSetContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#pivotClause.
	VisitPivotClause(ctx *PivotClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#pivotColumn.
	VisitPivotColumn(ctx *PivotColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#pivotValue.
	VisitPivotValue(ctx *PivotValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotClause.
	VisitUnpivotClause(ctx *UnpivotClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotNullClause.
	VisitUnpivotNullClause(ctx *UnpivotNullClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotOperator.
	VisitUnpivotOperator(ctx *UnpivotOperatorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotSingleValueColumnClause.
	VisitUnpivotSingleValueColumnClause(ctx *UnpivotSingleValueColumnClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotMultiValueColumnClause.
	VisitUnpivotMultiValueColumnClause(ctx *UnpivotMultiValueColumnClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotColumnSet.
	VisitUnpivotColumnSet(ctx *UnpivotColumnSetContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotValueColumn.
	VisitUnpivotValueColumn(ctx *UnpivotValueColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotNameColumn.
	VisitUnpivotNameColumn(ctx *UnpivotNameColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotColumnAndAlias.
	VisitUnpivotColumnAndAlias(ctx *UnpivotColumnAndAliasContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotColumn.
	VisitUnpivotColumn(ctx *UnpivotColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unpivotAlias.
	VisitUnpivotAlias(ctx *UnpivotAliasContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#lateralView.
	VisitLateralView(ctx *LateralViewContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#watermarkClause.
	VisitWatermarkClause(ctx *WatermarkClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#relationExtension.
	VisitRelationExtension(ctx *RelationExtensionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#joinType.
	VisitJoinType(ctx *JoinTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#joinCriteria.
	VisitJoinCriteria(ctx *JoinCriteriaContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sample.
	VisitSample(ctx *SampleContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sampleByPercentile.
	VisitSampleByPercentile(ctx *SampleByPercentileContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sampleByRows.
	VisitSampleByRows(ctx *SampleByRowsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sampleByBucket.
	VisitSampleByBucket(ctx *SampleByBucketContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sampleByBytes.
	VisitSampleByBytes(ctx *SampleByBytesContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierSeq.
	VisitIdentifierSeq(ctx *IdentifierSeqContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#orderedIdentifierList.
	VisitOrderedIdentifierList(ctx *OrderedIdentifierListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#orderedIdentifier.
	VisitOrderedIdentifier(ctx *OrderedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierCommentList.
	VisitIdentifierCommentList(ctx *IdentifierCommentListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierComment.
	VisitIdentifierComment(ctx *IdentifierCommentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#streamRelation.
	VisitStreamRelation(ctx *StreamRelationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#aliasedQuery.
	VisitAliasedQuery(ctx *AliasedQueryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#aliasedRelation.
	VisitAliasedRelation(ctx *AliasedRelationContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#inlineTableDefault2.
	VisitInlineTableDefault2(ctx *InlineTableDefault2Context) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableValuedFunction.
	VisitTableValuedFunction(ctx *TableValuedFunctionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#optionsClause.
	VisitOptionsClause(ctx *OptionsClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionTableSubqueryArgument.
	VisitFunctionTableSubqueryArgument(ctx *FunctionTableSubqueryArgumentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableArgumentPartitioning.
	VisitTableArgumentPartitioning(ctx *TableArgumentPartitioningContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionTableNamedArgumentExpression.
	VisitFunctionTableNamedArgumentExpression(ctx *FunctionTableNamedArgumentExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionTableReferenceArgument.
	VisitFunctionTableReferenceArgument(ctx *FunctionTableReferenceArgumentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionTableArgument.
	VisitFunctionTableArgument(ctx *FunctionTableArgumentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionTable.
	VisitFunctionTable(ctx *FunctionTableContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#rowFormatSerde.
	VisitRowFormatSerde(ctx *RowFormatSerdeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#rowFormatDelimited.
	VisitRowFormatDelimited(ctx *RowFormatDelimitedContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multipartIdentifierList.
	VisitMultipartIdentifierList(ctx *MultipartIdentifierListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multipartIdentifier.
	VisitMultipartIdentifier(ctx *MultipartIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multipartIdentifierPropertyList.
	VisitMultipartIdentifierPropertyList(ctx *MultipartIdentifierPropertyListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multipartIdentifierProperty.
	VisitMultipartIdentifierProperty(ctx *MultipartIdentifierPropertyContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableIdentifier.
	VisitTableIdentifier(ctx *TableIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionIdentifier.
	VisitFunctionIdentifier(ctx *FunctionIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedExpression.
	VisitNamedExpression(ctx *NamedExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedExpressionSeq.
	VisitNamedExpressionSeq(ctx *NamedExpressionSeqContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionFieldList.
	VisitPartitionFieldList(ctx *PartitionFieldListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionTransform.
	VisitPartitionTransform(ctx *PartitionTransformContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#partitionColumn.
	VisitPartitionColumn(ctx *PartitionColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identityTransform.
	VisitIdentityTransform(ctx *IdentityTransformContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#applyTransform.
	VisitApplyTransform(ctx *ApplyTransformContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#transformArgument.
	VisitTransformArgument(ctx *TransformArgumentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedArgumentExpression.
	VisitNamedArgumentExpression(ctx *NamedArgumentExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionArgument.
	VisitFunctionArgument(ctx *FunctionArgumentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#expressionSeq.
	VisitExpressionSeq(ctx *ExpressionSeqContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#logicalNot.
	VisitLogicalNot(ctx *LogicalNotContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#predicated.
	VisitPredicated(ctx *PredicatedContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#exists.
	VisitExists(ctx *ExistsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#logicalBinary.
	VisitLogicalBinary(ctx *LogicalBinaryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#errorCapturingNot.
	VisitErrorCapturingNot(ctx *ErrorCapturingNotContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#valueExpressionDefault.
	VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#arithmeticBinary.
	VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#arithmeticUnary.
	VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#shiftOperator.
	VisitShiftOperator(ctx *ShiftOperatorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#datetimeUnit.
	VisitDatetimeUnit(ctx *DatetimeUnitContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#dereference.
	VisitDereference(ctx *DereferenceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#castByColon.
	VisitCastByColon(ctx *CastByColonContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#timestampadd.
	VisitTimestampadd(ctx *TimestampaddContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#substring.
	VisitSubstring(ctx *SubstringContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#any_value.
	VisitAny_value(ctx *Any_valueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#trim.
	VisitTrim(ctx *TrimContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#semiStructuredExtract.
	VisitSemiStructuredExtract(ctx *SemiStructuredExtractContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#simpleCase.
	VisitSimpleCase(ctx *SimpleCaseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#currentLike.
	VisitCurrentLike(ctx *CurrentLikeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#rowConstructor.
	VisitRowConstructor(ctx *RowConstructorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#last.
	VisitLast(ctx *LastContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#star.
	VisitStar(ctx *StarContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#overlay.
	VisitOverlay(ctx *OverlayContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#timestampdiff.
	VisitTimestampdiff(ctx *TimestampdiffContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#subqueryExpression.
	VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#constantDefault.
	VisitConstantDefault(ctx *ConstantDefaultContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#extract.
	VisitExtract(ctx *ExtractContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#searchedCase.
	VisitSearchedCase(ctx *SearchedCaseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#position.
	VisitPosition(ctx *PositionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#first.
	VisitFirst(ctx *FirstContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#semiStructuredExtractionPath.
	VisitSemiStructuredExtractionPath(ctx *SemiStructuredExtractionPathContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#jsonPathIdentifier.
	VisitJsonPathIdentifier(ctx *JsonPathIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#jsonPathBracketedIdentifier.
	VisitJsonPathBracketedIdentifier(ctx *JsonPathBracketedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#jsonPathFirstPart.
	VisitJsonPathFirstPart(ctx *JsonPathFirstPartContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#jsonPathParts.
	VisitJsonPathParts(ctx *JsonPathPartsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#literalType.
	VisitLiteralType(ctx *LiteralTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#posParameterLiteral.
	VisitPosParameterLiteral(ctx *PosParameterLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedParameterLiteral.
	VisitNamedParameterLiteral(ctx *NamedParameterLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#intervalLiteral.
	VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#typeConstructor.
	VisitTypeConstructor(ctx *TypeConstructorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedParameterMarker.
	VisitNamedParameterMarker(ctx *NamedParameterMarkerContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#arithmeticOperator.
	VisitArithmeticOperator(ctx *ArithmeticOperatorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#predicateOperator.
	VisitPredicateOperator(ctx *PredicateOperatorContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#errorCapturingMultiUnitsInterval.
	VisitErrorCapturingMultiUnitsInterval(ctx *ErrorCapturingMultiUnitsIntervalContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#multiUnitsInterval.
	VisitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#errorCapturingUnitToUnitInterval.
	VisitErrorCapturingUnitToUnitInterval(ctx *ErrorCapturingUnitToUnitIntervalContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unitToUnitInterval.
	VisitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#intervalValue.
	VisitIntervalValue(ctx *IntervalValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unitInMultiUnits.
	VisitUnitInMultiUnits(ctx *UnitInMultiUnitsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unitInUnitToUnit.
	VisitUnitInUnitToUnit(ctx *UnitInUnitToUnitContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colPosition.
	VisitColPosition(ctx *ColPositionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#collationSpec.
	VisitCollationSpec(ctx *CollationSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#collateClause.
	VisitCollateClause(ctx *CollateClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#nonTrivialPrimitiveType.
	VisitNonTrivialPrimitiveType(ctx *NonTrivialPrimitiveTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#trivialPrimitiveType.
	VisitTrivialPrimitiveType(ctx *TrivialPrimitiveTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#complexDataType.
	VisitComplexDataType(ctx *ComplexDataTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#primitiveDataType.
	VisitPrimitiveDataType(ctx *PrimitiveDataTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#qualifiedColTypeWithPositionList.
	VisitQualifiedColTypeWithPositionList(ctx *QualifiedColTypeWithPositionListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#qualifiedColTypeWithPosition.
	VisitQualifiedColTypeWithPosition(ctx *QualifiedColTypeWithPositionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colDefinitionDescriptorWithPosition.
	VisitColDefinitionDescriptorWithPosition(ctx *ColDefinitionDescriptorWithPositionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#defaultExpression.
	VisitDefaultExpression(ctx *DefaultExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#variableDefaultExpression.
	VisitVariableDefaultExpression(ctx *VariableDefaultExpressionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colTypeList.
	VisitColTypeList(ctx *ColTypeListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colType.
	VisitColType(ctx *ColTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableElementList.
	VisitTableElementList(ctx *TableElementListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableElement.
	VisitTableElement(ctx *TableElementContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colDefinitionList.
	VisitColDefinitionList(ctx *ColDefinitionListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colDefinition.
	VisitColDefinition(ctx *ColDefinitionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#colDefinitionOption.
	VisitColDefinitionOption(ctx *ColDefinitionOptionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#generatedColumn.
	VisitGeneratedColumn(ctx *GeneratedColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identityColumn.
	VisitIdentityColumn(ctx *IdentityColumnContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identityColSpec.
	VisitIdentityColSpec(ctx *IdentityColSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sequenceGeneratorOption.
	VisitSequenceGeneratorOption(ctx *SequenceGeneratorOptionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sequenceGeneratorStartOrStep.
	VisitSequenceGeneratorStartOrStep(ctx *SequenceGeneratorStartOrStepContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#complexColTypeList.
	VisitComplexColTypeList(ctx *ComplexColTypeListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#complexColType.
	VisitComplexColType(ctx *ComplexColTypeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#codeLiteral.
	VisitCodeLiteral(ctx *CodeLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#routineCharacteristics.
	VisitRoutineCharacteristics(ctx *RoutineCharacteristicsContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#routineLanguage.
	VisitRoutineLanguage(ctx *RoutineLanguageContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#specificName.
	VisitSpecificName(ctx *SpecificNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#deterministic.
	VisitDeterministic(ctx *DeterministicContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#sqlDataAccess.
	VisitSqlDataAccess(ctx *SqlDataAccessContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#nullCall.
	VisitNullCall(ctx *NullCallContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#rightsClause.
	VisitRightsClause(ctx *RightsClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedWindow.
	VisitNamedWindow(ctx *NamedWindowContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#windowRef.
	VisitWindowRef(ctx *WindowRefContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#windowDef.
	VisitWindowDef(ctx *WindowDefContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#windowFrame.
	VisitWindowFrame(ctx *WindowFrameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#frameBound.
	VisitFrameBound(ctx *FrameBoundContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#errorCapturingIdentifier.
	VisitErrorCapturingIdentifier(ctx *ErrorCapturingIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#errorIdent.
	VisitErrorIdent(ctx *ErrorIdentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#realIdent.
	VisitRealIdent(ctx *RealIdentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#simpleIdentifier.
	VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#quotedIdentifierAlternative.
	VisitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#identifierLiteral.
	VisitIdentifierLiteral(ctx *IdentifierLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#simpleUnquotedIdentifier.
	VisitSimpleUnquotedIdentifier(ctx *SimpleUnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#simpleQuotedIdentifierAlternative.
	VisitSimpleQuotedIdentifierAlternative(ctx *SimpleQuotedIdentifierAlternativeContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#quotedIdentifier.
	VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#backQuotedIdentifier.
	VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#exponentLiteral.
	VisitExponentLiteral(ctx *ExponentLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#legacyDecimalLiteral.
	VisitLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#bigIntLiteral.
	VisitBigIntLiteral(ctx *BigIntLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#smallIntLiteral.
	VisitSmallIntLiteral(ctx *SmallIntLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tinyIntLiteral.
	VisitTinyIntLiteral(ctx *TinyIntLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#doubleLiteral.
	VisitDoubleLiteral(ctx *DoubleLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#bigDecimalLiteral.
	VisitBigDecimalLiteral(ctx *BigDecimalLiteralContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#integerVal.
	VisitIntegerVal(ctx *IntegerValContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#parameterIntegerValue.
	VisitParameterIntegerValue(ctx *ParameterIntegerValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#columnConstraintDefinition.
	VisitColumnConstraintDefinition(ctx *ColumnConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#columnConstraint.
	VisitColumnConstraint(ctx *ColumnConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableConstraintDefinition.
	VisitTableConstraintDefinition(ctx *TableConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#tableConstraint.
	VisitTableConstraint(ctx *TableConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#uniqueSpec.
	VisitUniqueSpec(ctx *UniqueSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#uniqueConstraint.
	VisitUniqueConstraint(ctx *UniqueConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#referenceSpec.
	VisitReferenceSpec(ctx *ReferenceSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#foreignKeyConstraint.
	VisitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#constraintCharacteristic.
	VisitConstraintCharacteristic(ctx *ConstraintCharacteristicContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#enforcedCharacteristic.
	VisitEnforcedCharacteristic(ctx *EnforcedCharacteristicContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#relyCharacteristic.
	VisitRelyCharacteristic(ctx *RelyCharacteristicContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterColumnSpecList.
	VisitAlterColumnSpecList(ctx *AlterColumnSpecListContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterColumnSpec.
	VisitAlterColumnSpec(ctx *AlterColumnSpecContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#alterColumnAction.
	VisitAlterColumnAction(ctx *AlterColumnActionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleStringLiteralValue.
	VisitSingleStringLiteralValue(ctx *SingleStringLiteralValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleDoubleQuotedStringLiteralValue.
	VisitSingleDoubleQuotedStringLiteralValue(ctx *SingleDoubleQuotedStringLiteralValueContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#singleStringLit.
	VisitSingleStringLit(ctx *SingleStringLitContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#namedParameterMarkerRule.
	VisitNamedParameterMarkerRule(ctx *NamedParameterMarkerRuleContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#positionalParameterMarkerRule.
	VisitPositionalParameterMarkerRule(ctx *PositionalParameterMarkerRuleContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#stringLit.
	VisitStringLit(ctx *StringLitContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#operatorPipeRightSide.
	VisitOperatorPipeRightSide(ctx *OperatorPipeRightSideContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#operatorPipeSetAssignmentSeq.
	VisitOperatorPipeSetAssignmentSeq(ctx *OperatorPipeSetAssignmentSeqContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#ansiNonReserved.
	VisitAnsiNonReserved(ctx *AnsiNonReservedContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#strictNonReserved.
	VisitStrictNonReserved(ctx *StrictNonReservedContext) interface{}

	// Visit a parse tree produced by SparkSqlParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}

}