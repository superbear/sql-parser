// Code generated from StarRocksParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package starrocks // StarRocksParser
import "github.com/antlr4-go/antlr/v4"

// BaseStarRocksParserListener is a complete listener for a parse tree produced by StarRocksParser.
type BaseStarRocksParserListener struct{}

var _ StarRocksParserListener = &BaseStarRocksParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStarRocksParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStarRocksParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStarRocksParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStarRocksParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlStatements is called when production sqlStatements is entered.
func (s *BaseStarRocksParserListener) EnterSqlStatements(ctx *SqlStatementsContext) {}

// ExitSqlStatements is called when production sqlStatements is exited.
func (s *BaseStarRocksParserListener) ExitSqlStatements(ctx *SqlStatementsContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseStarRocksParserListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseStarRocksParserListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseStarRocksParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseStarRocksParserListener) ExitStatement(ctx *StatementContext) {}

// EnterUseDatabaseStatement is called when production useDatabaseStatement is entered.
func (s *BaseStarRocksParserListener) EnterUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// ExitUseDatabaseStatement is called when production useDatabaseStatement is exited.
func (s *BaseStarRocksParserListener) ExitUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// EnterUseCatalogStatement is called when production useCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// ExitUseCatalogStatement is called when production useCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// EnterSetCatalogStatement is called when production setCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// ExitSetCatalogStatement is called when production setCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// EnterShowDatabasesStatement is called when production showDatabasesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {
}

// ExitShowDatabasesStatement is called when production showDatabasesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {
}

// EnterAlterDbQuotaStatement is called when production alterDbQuotaStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// ExitAlterDbQuotaStatement is called when production alterDbQuotaStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// EnterCreateDbStatement is called when production createDbStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateDbStatement(ctx *CreateDbStatementContext) {}

// ExitCreateDbStatement is called when production createDbStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateDbStatement(ctx *CreateDbStatementContext) {}

// EnterDropDbStatement is called when production dropDbStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropDbStatement(ctx *DropDbStatementContext) {}

// ExitDropDbStatement is called when production dropDbStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropDbStatement(ctx *DropDbStatementContext) {}

// EnterShowCreateDbStatement is called when production showCreateDbStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// ExitShowCreateDbStatement is called when production showCreateDbStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// EnterAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// ExitAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// EnterRecoverDbStmt is called when production recoverDbStmt is entered.
func (s *BaseStarRocksParserListener) EnterRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// ExitRecoverDbStmt is called when production recoverDbStmt is exited.
func (s *BaseStarRocksParserListener) ExitRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// EnterShowDataStmt is called when production showDataStmt is entered.
func (s *BaseStarRocksParserListener) EnterShowDataStmt(ctx *ShowDataStmtContext) {}

// ExitShowDataStmt is called when production showDataStmt is exited.
func (s *BaseStarRocksParserListener) ExitShowDataStmt(ctx *ShowDataStmtContext) {}

// EnterShowDataDistributionStmt is called when production showDataDistributionStmt is entered.
func (s *BaseStarRocksParserListener) EnterShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// ExitShowDataDistributionStmt is called when production showDataDistributionStmt is exited.
func (s *BaseStarRocksParserListener) ExitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterColumnDesc is called when production columnDesc is entered.
func (s *BaseStarRocksParserListener) EnterColumnDesc(ctx *ColumnDescContext) {}

// ExitColumnDesc is called when production columnDesc is exited.
func (s *BaseStarRocksParserListener) ExitColumnDesc(ctx *ColumnDescContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseStarRocksParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseStarRocksParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterDefaultDesc is called when production defaultDesc is entered.
func (s *BaseStarRocksParserListener) EnterDefaultDesc(ctx *DefaultDescContext) {}

// ExitDefaultDesc is called when production defaultDesc is exited.
func (s *BaseStarRocksParserListener) ExitDefaultDesc(ctx *DefaultDescContext) {}

// EnterGeneratedColumnDesc is called when production generatedColumnDesc is entered.
func (s *BaseStarRocksParserListener) EnterGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// ExitGeneratedColumnDesc is called when production generatedColumnDesc is exited.
func (s *BaseStarRocksParserListener) ExitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// EnterIndexDesc is called when production indexDesc is entered.
func (s *BaseStarRocksParserListener) EnterIndexDesc(ctx *IndexDescContext) {}

// ExitIndexDesc is called when production indexDesc is exited.
func (s *BaseStarRocksParserListener) ExitIndexDesc(ctx *IndexDescContext) {}

// EnterEngineDesc is called when production engineDesc is entered.
func (s *BaseStarRocksParserListener) EnterEngineDesc(ctx *EngineDescContext) {}

// ExitEngineDesc is called when production engineDesc is exited.
func (s *BaseStarRocksParserListener) ExitEngineDesc(ctx *EngineDescContext) {}

// EnterCharsetDesc is called when production charsetDesc is entered.
func (s *BaseStarRocksParserListener) EnterCharsetDesc(ctx *CharsetDescContext) {}

// ExitCharsetDesc is called when production charsetDesc is exited.
func (s *BaseStarRocksParserListener) ExitCharsetDesc(ctx *CharsetDescContext) {}

// EnterCollateDesc is called when production collateDesc is entered.
func (s *BaseStarRocksParserListener) EnterCollateDesc(ctx *CollateDescContext) {}

// ExitCollateDesc is called when production collateDesc is exited.
func (s *BaseStarRocksParserListener) ExitCollateDesc(ctx *CollateDescContext) {}

// EnterKeyDesc is called when production keyDesc is entered.
func (s *BaseStarRocksParserListener) EnterKeyDesc(ctx *KeyDescContext) {}

// ExitKeyDesc is called when production keyDesc is exited.
func (s *BaseStarRocksParserListener) ExitKeyDesc(ctx *KeyDescContext) {}

// EnterOrderByDesc is called when production orderByDesc is entered.
func (s *BaseStarRocksParserListener) EnterOrderByDesc(ctx *OrderByDescContext) {}

// ExitOrderByDesc is called when production orderByDesc is exited.
func (s *BaseStarRocksParserListener) ExitOrderByDesc(ctx *OrderByDescContext) {}

// EnterColumnNullable is called when production columnNullable is entered.
func (s *BaseStarRocksParserListener) EnterColumnNullable(ctx *ColumnNullableContext) {}

// ExitColumnNullable is called when production columnNullable is exited.
func (s *BaseStarRocksParserListener) ExitColumnNullable(ctx *ColumnNullableContext) {}

// EnterTypeWithNullable is called when production typeWithNullable is entered.
func (s *BaseStarRocksParserListener) EnterTypeWithNullable(ctx *TypeWithNullableContext) {}

// ExitTypeWithNullable is called when production typeWithNullable is exited.
func (s *BaseStarRocksParserListener) ExitTypeWithNullable(ctx *TypeWithNullableContext) {}

// EnterAggStateDesc is called when production aggStateDesc is entered.
func (s *BaseStarRocksParserListener) EnterAggStateDesc(ctx *AggStateDescContext) {}

// ExitAggStateDesc is called when production aggStateDesc is exited.
func (s *BaseStarRocksParserListener) ExitAggStateDesc(ctx *AggStateDescContext) {}

// EnterAggDesc is called when production aggDesc is entered.
func (s *BaseStarRocksParserListener) EnterAggDesc(ctx *AggDescContext) {}

// ExitAggDesc is called when production aggDesc is exited.
func (s *BaseStarRocksParserListener) ExitAggDesc(ctx *AggDescContext) {}

// EnterRollupDesc is called when production rollupDesc is entered.
func (s *BaseStarRocksParserListener) EnterRollupDesc(ctx *RollupDescContext) {}

// ExitRollupDesc is called when production rollupDesc is exited.
func (s *BaseStarRocksParserListener) ExitRollupDesc(ctx *RollupDescContext) {}

// EnterRollupItem is called when production rollupItem is entered.
func (s *BaseStarRocksParserListener) EnterRollupItem(ctx *RollupItemContext) {}

// ExitRollupItem is called when production rollupItem is exited.
func (s *BaseStarRocksParserListener) ExitRollupItem(ctx *RollupItemContext) {}

// EnterDupKeys is called when production dupKeys is entered.
func (s *BaseStarRocksParserListener) EnterDupKeys(ctx *DupKeysContext) {}

// ExitDupKeys is called when production dupKeys is exited.
func (s *BaseStarRocksParserListener) ExitDupKeys(ctx *DupKeysContext) {}

// EnterFromRollup is called when production fromRollup is entered.
func (s *BaseStarRocksParserListener) EnterFromRollup(ctx *FromRollupContext) {}

// ExitFromRollup is called when production fromRollup is exited.
func (s *BaseStarRocksParserListener) ExitFromRollup(ctx *FromRollupContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseStarRocksParserListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseStarRocksParserListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseStarRocksParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseStarRocksParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateTableAsSelectStatement is called when production createTableAsSelectStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// ExitCreateTableAsSelectStatement is called when production createTableAsSelectStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// ExitCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterDropIndexStatement is called when production dropIndexStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropIndexStatement(ctx *DropIndexStatementContext) {}

// ExitDropIndexStatement is called when production dropIndexStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropIndexStatement(ctx *DropIndexStatementContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseStarRocksParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseStarRocksParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterShowTableStatement is called when production showTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTableStatement(ctx *ShowTableStatementContext) {}

// ExitShowTableStatement is called when production showTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTableStatement(ctx *ShowTableStatementContext) {}

// EnterShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// ExitShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// EnterShowCreateTableStatement is called when production showCreateTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// ExitShowCreateTableStatement is called when production showCreateTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// EnterShowColumnStatement is called when production showColumnStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowColumnStatement(ctx *ShowColumnStatementContext) {}

// ExitShowColumnStatement is called when production showColumnStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowColumnStatement(ctx *ShowColumnStatementContext) {}

// EnterShowTableStatusStatement is called when production showTableStatusStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// ExitShowTableStatusStatement is called when production showTableStatusStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// EnterRefreshTableStatement is called when production refreshTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// ExitRefreshTableStatement is called when production refreshTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// EnterShowAlterStatement is called when production showAlterStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowAlterStatement(ctx *ShowAlterStatementContext) {}

// ExitShowAlterStatement is called when production showAlterStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowAlterStatement(ctx *ShowAlterStatementContext) {}

// EnterDescTableStatement is called when production descTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterDescTableStatement(ctx *DescTableStatementContext) {}

// ExitDescTableStatement is called when production descTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitDescTableStatement(ctx *DescTableStatementContext) {}

// EnterCreateTableLikeStatement is called when production createTableLikeStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// ExitCreateTableLikeStatement is called when production createTableLikeStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// EnterShowIndexStatement is called when production showIndexStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowIndexStatement(ctx *ShowIndexStatementContext) {}

// ExitShowIndexStatement is called when production showIndexStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowIndexStatement(ctx *ShowIndexStatementContext) {}

// EnterRecoverTableStatement is called when production recoverTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// ExitRecoverTableStatement is called when production recoverTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {
}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {
}

// EnterCancelAlterTableStatement is called when production cancelAlterTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// ExitCancelAlterTableStatement is called when production cancelAlterTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// EnterShowPartitionsStatement is called when production showPartitionsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {
}

// ExitShowPartitionsStatement is called when production showPartitionsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {
}

// EnterRecoverPartitionStatement is called when production recoverPartitionStatement is entered.
func (s *BaseStarRocksParserListener) EnterRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// ExitRecoverPartitionStatement is called when production recoverPartitionStatement is exited.
func (s *BaseStarRocksParserListener) ExitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterAlterViewStatement is called when production alterViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterViewStatement(ctx *AlterViewStatementContext) {}

// ExitAlterViewStatement is called when production alterViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterViewStatement(ctx *AlterViewStatementContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterColumnNameWithComment is called when production columnNameWithComment is entered.
func (s *BaseStarRocksParserListener) EnterColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// ExitColumnNameWithComment is called when production columnNameWithComment is exited.
func (s *BaseStarRocksParserListener) ExitColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// EnterSubmitTaskStatement is called when production submitTaskStatement is entered.
func (s *BaseStarRocksParserListener) EnterSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// ExitSubmitTaskStatement is called when production submitTaskStatement is exited.
func (s *BaseStarRocksParserListener) ExitSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// EnterTaskClause is called when production taskClause is entered.
func (s *BaseStarRocksParserListener) EnterTaskClause(ctx *TaskClauseContext) {}

// ExitTaskClause is called when production taskClause is exited.
func (s *BaseStarRocksParserListener) ExitTaskClause(ctx *TaskClauseContext) {}

// EnterDropTaskStatement is called when production dropTaskStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropTaskStatement(ctx *DropTaskStatementContext) {}

// ExitDropTaskStatement is called when production dropTaskStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropTaskStatement(ctx *DropTaskStatementContext) {}

// EnterTaskScheduleDesc is called when production taskScheduleDesc is entered.
func (s *BaseStarRocksParserListener) EnterTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// ExitTaskScheduleDesc is called when production taskScheduleDesc is exited.
func (s *BaseStarRocksParserListener) ExitTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterMvPartitionExprs is called when production mvPartitionExprs is entered.
func (s *BaseStarRocksParserListener) EnterMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// ExitMvPartitionExprs is called when production mvPartitionExprs is exited.
func (s *BaseStarRocksParserListener) ExitMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// EnterMaterializedViewDesc is called when production materializedViewDesc is entered.
func (s *BaseStarRocksParserListener) EnterMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// ExitMaterializedViewDesc is called when production materializedViewDesc is exited.
func (s *BaseStarRocksParserListener) ExitMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// EnterShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// ExitShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// ExitAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// EnterRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// ExitRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// EnterCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// ExitCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// EnterAdminSetConfigStatement is called when production adminSetConfigStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {
}

// ExitAdminSetConfigStatement is called when production adminSetConfigStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {
}

// EnterAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// ExitAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// EnterAdminShowConfigStatement is called when production adminShowConfigStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// ExitAdminShowConfigStatement is called when production adminShowConfigStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// EnterAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// ExitAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// EnterAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// ExitAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// EnterAdminRepairTableStatement is called when production adminRepairTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// ExitAdminRepairTableStatement is called when production adminRepairTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// EnterAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// ExitAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// EnterAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// ExitAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// EnterAdminSetPartitionVersion is called when production adminSetPartitionVersion is entered.
func (s *BaseStarRocksParserListener) EnterAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// ExitAdminSetPartitionVersion is called when production adminSetPartitionVersion is exited.
func (s *BaseStarRocksParserListener) ExitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// EnterKillStatement is called when production killStatement is entered.
func (s *BaseStarRocksParserListener) EnterKillStatement(ctx *KillStatementContext) {}

// ExitKillStatement is called when production killStatement is exited.
func (s *BaseStarRocksParserListener) ExitKillStatement(ctx *KillStatementContext) {}

// EnterSyncStatement is called when production syncStatement is entered.
func (s *BaseStarRocksParserListener) EnterSyncStatement(ctx *SyncStatementContext) {}

// ExitSyncStatement is called when production syncStatement is exited.
func (s *BaseStarRocksParserListener) ExitSyncStatement(ctx *SyncStatementContext) {}

// EnterAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// EnterAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is entered.
func (s *BaseStarRocksParserListener) EnterAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is exited.
func (s *BaseStarRocksParserListener) ExitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// EnterAlterSystemStatement is called when production alterSystemStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// ExitAlterSystemStatement is called when production alterSystemStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// EnterCancelAlterSystemStatement is called when production cancelAlterSystemStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// ExitCancelAlterSystemStatement is called when production cancelAlterSystemStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// EnterShowComputeNodesStatement is called when production showComputeNodesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// ExitShowComputeNodesStatement is called when production showComputeNodesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// EnterCreateExternalCatalogStatement is called when production createExternalCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// ExitCreateExternalCatalogStatement is called when production createExternalCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// EnterShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// ExitShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// EnterDropExternalCatalogStatement is called when production dropExternalCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// ExitDropExternalCatalogStatement is called when production dropExternalCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// EnterShowCatalogsStatement is called when production showCatalogsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// ExitShowCatalogsStatement is called when production showCatalogsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// EnterAlterCatalogStatement is called when production alterCatalogStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// ExitAlterCatalogStatement is called when production alterCatalogStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// EnterCreateStorageVolumeStatement is called when production createStorageVolumeStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// ExitCreateStorageVolumeStatement is called when production createStorageVolumeStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// EnterTypeDesc is called when production typeDesc is entered.
func (s *BaseStarRocksParserListener) EnterTypeDesc(ctx *TypeDescContext) {}

// ExitTypeDesc is called when production typeDesc is exited.
func (s *BaseStarRocksParserListener) ExitTypeDesc(ctx *TypeDescContext) {}

// EnterLocationsDesc is called when production locationsDesc is entered.
func (s *BaseStarRocksParserListener) EnterLocationsDesc(ctx *LocationsDescContext) {}

// ExitLocationsDesc is called when production locationsDesc is exited.
func (s *BaseStarRocksParserListener) ExitLocationsDesc(ctx *LocationsDescContext) {}

// EnterShowStorageVolumesStatement is called when production showStorageVolumesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// ExitShowStorageVolumesStatement is called when production showStorageVolumesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// EnterDropStorageVolumeStatement is called when production dropStorageVolumeStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// ExitDropStorageVolumeStatement is called when production dropStorageVolumeStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// ExitAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeClause is called when production alterStorageVolumeClause is entered.
func (s *BaseStarRocksParserListener) EnterAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// ExitAlterStorageVolumeClause is called when production alterStorageVolumeClause is exited.
func (s *BaseStarRocksParserListener) ExitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// EnterModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// ExitModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// EnterModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// ExitModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// EnterDescStorageVolumeStatement is called when production descStorageVolumeStatement is entered.
func (s *BaseStarRocksParserListener) EnterDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// ExitDescStorageVolumeStatement is called when production descStorageVolumeStatement is exited.
func (s *BaseStarRocksParserListener) ExitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// EnterSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// ExitSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// EnterUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is entered.
func (s *BaseStarRocksParserListener) EnterUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// ExitUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is exited.
func (s *BaseStarRocksParserListener) ExitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// EnterShowFailPointStatement is called when production showFailPointStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowFailPointStatement(ctx *ShowFailPointStatementContext) {
}

// ExitShowFailPointStatement is called when production showFailPointStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowFailPointStatement(ctx *ShowFailPointStatementContext) {
}

// EnterCreateDictionaryStatement is called when production createDictionaryStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// ExitCreateDictionaryStatement is called when production createDictionaryStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// EnterDropDictionaryStatement is called when production dropDictionaryStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropDictionaryStatement(ctx *DropDictionaryStatementContext) {
}

// ExitDropDictionaryStatement is called when production dropDictionaryStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropDictionaryStatement(ctx *DropDictionaryStatementContext) {
}

// EnterRefreshDictionaryStatement is called when production refreshDictionaryStatement is entered.
func (s *BaseStarRocksParserListener) EnterRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// ExitRefreshDictionaryStatement is called when production refreshDictionaryStatement is exited.
func (s *BaseStarRocksParserListener) ExitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// EnterShowDictionaryStatement is called when production showDictionaryStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {
}

// ExitShowDictionaryStatement is called when production showDictionaryStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {
}

// EnterCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// ExitCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// EnterDictionaryColumnDesc is called when production dictionaryColumnDesc is entered.
func (s *BaseStarRocksParserListener) EnterDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// ExitDictionaryColumnDesc is called when production dictionaryColumnDesc is exited.
func (s *BaseStarRocksParserListener) ExitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// EnterDictionaryName is called when production dictionaryName is entered.
func (s *BaseStarRocksParserListener) EnterDictionaryName(ctx *DictionaryNameContext) {}

// ExitDictionaryName is called when production dictionaryName is exited.
func (s *BaseStarRocksParserListener) ExitDictionaryName(ctx *DictionaryNameContext) {}

// EnterAlterClause is called when production alterClause is entered.
func (s *BaseStarRocksParserListener) EnterAlterClause(ctx *AlterClauseContext) {}

// ExitAlterClause is called when production alterClause is exited.
func (s *BaseStarRocksParserListener) ExitAlterClause(ctx *AlterClauseContext) {}

// EnterAddFrontendClause is called when production addFrontendClause is entered.
func (s *BaseStarRocksParserListener) EnterAddFrontendClause(ctx *AddFrontendClauseContext) {}

// ExitAddFrontendClause is called when production addFrontendClause is exited.
func (s *BaseStarRocksParserListener) ExitAddFrontendClause(ctx *AddFrontendClauseContext) {}

// EnterDropFrontendClause is called when production dropFrontendClause is entered.
func (s *BaseStarRocksParserListener) EnterDropFrontendClause(ctx *DropFrontendClauseContext) {}

// ExitDropFrontendClause is called when production dropFrontendClause is exited.
func (s *BaseStarRocksParserListener) ExitDropFrontendClause(ctx *DropFrontendClauseContext) {}

// EnterModifyFrontendHostClause is called when production modifyFrontendHostClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// ExitModifyFrontendHostClause is called when production modifyFrontendHostClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// EnterAddBackendClause is called when production addBackendClause is entered.
func (s *BaseStarRocksParserListener) EnterAddBackendClause(ctx *AddBackendClauseContext) {}

// ExitAddBackendClause is called when production addBackendClause is exited.
func (s *BaseStarRocksParserListener) ExitAddBackendClause(ctx *AddBackendClauseContext) {}

// EnterDropBackendClause is called when production dropBackendClause is entered.
func (s *BaseStarRocksParserListener) EnterDropBackendClause(ctx *DropBackendClauseContext) {}

// ExitDropBackendClause is called when production dropBackendClause is exited.
func (s *BaseStarRocksParserListener) ExitDropBackendClause(ctx *DropBackendClauseContext) {}

// EnterDecommissionBackendClause is called when production decommissionBackendClause is entered.
func (s *BaseStarRocksParserListener) EnterDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// ExitDecommissionBackendClause is called when production decommissionBackendClause is exited.
func (s *BaseStarRocksParserListener) ExitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// EnterModifyBackendClause is called when production modifyBackendClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// ExitModifyBackendClause is called when production modifyBackendClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// EnterAddComputeNodeClause is called when production addComputeNodeClause is entered.
func (s *BaseStarRocksParserListener) EnterAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// ExitAddComputeNodeClause is called when production addComputeNodeClause is exited.
func (s *BaseStarRocksParserListener) ExitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// EnterDropComputeNodeClause is called when production dropComputeNodeClause is entered.
func (s *BaseStarRocksParserListener) EnterDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// ExitDropComputeNodeClause is called when production dropComputeNodeClause is exited.
func (s *BaseStarRocksParserListener) ExitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// EnterModifyBrokerClause is called when production modifyBrokerClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// ExitModifyBrokerClause is called when production modifyBrokerClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// EnterAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is entered.
func (s *BaseStarRocksParserListener) EnterAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {
}

// ExitAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is exited.
func (s *BaseStarRocksParserListener) ExitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {
}

// EnterCreateImageClause is called when production createImageClause is entered.
func (s *BaseStarRocksParserListener) EnterCreateImageClause(ctx *CreateImageClauseContext) {}

// ExitCreateImageClause is called when production createImageClause is exited.
func (s *BaseStarRocksParserListener) ExitCreateImageClause(ctx *CreateImageClauseContext) {}

// EnterCleanTabletSchedQClause is called when production cleanTabletSchedQClause is entered.
func (s *BaseStarRocksParserListener) EnterCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {
}

// ExitCleanTabletSchedQClause is called when production cleanTabletSchedQClause is exited.
func (s *BaseStarRocksParserListener) ExitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {
}

// EnterDecommissionDiskClause is called when production decommissionDiskClause is entered.
func (s *BaseStarRocksParserListener) EnterDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {
}

// ExitDecommissionDiskClause is called when production decommissionDiskClause is exited.
func (s *BaseStarRocksParserListener) ExitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {
}

// EnterCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is entered.
func (s *BaseStarRocksParserListener) EnterCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// ExitCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is exited.
func (s *BaseStarRocksParserListener) ExitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// EnterDisableDiskClause is called when production disableDiskClause is entered.
func (s *BaseStarRocksParserListener) EnterDisableDiskClause(ctx *DisableDiskClauseContext) {}

// ExitDisableDiskClause is called when production disableDiskClause is exited.
func (s *BaseStarRocksParserListener) ExitDisableDiskClause(ctx *DisableDiskClauseContext) {}

// EnterCancelDisableDiskClause is called when production cancelDisableDiskClause is entered.
func (s *BaseStarRocksParserListener) EnterCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {
}

// ExitCancelDisableDiskClause is called when production cancelDisableDiskClause is exited.
func (s *BaseStarRocksParserListener) ExitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {
}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseStarRocksParserListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseStarRocksParserListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterDropIndexClause is called when production dropIndexClause is entered.
func (s *BaseStarRocksParserListener) EnterDropIndexClause(ctx *DropIndexClauseContext) {}

// ExitDropIndexClause is called when production dropIndexClause is exited.
func (s *BaseStarRocksParserListener) ExitDropIndexClause(ctx *DropIndexClauseContext) {}

// EnterTableRenameClause is called when production tableRenameClause is entered.
func (s *BaseStarRocksParserListener) EnterTableRenameClause(ctx *TableRenameClauseContext) {}

// ExitTableRenameClause is called when production tableRenameClause is exited.
func (s *BaseStarRocksParserListener) ExitTableRenameClause(ctx *TableRenameClauseContext) {}

// EnterSwapTableClause is called when production swapTableClause is entered.
func (s *BaseStarRocksParserListener) EnterSwapTableClause(ctx *SwapTableClauseContext) {}

// ExitSwapTableClause is called when production swapTableClause is exited.
func (s *BaseStarRocksParserListener) ExitSwapTableClause(ctx *SwapTableClauseContext) {}

// EnterModifyPropertiesClause is called when production modifyPropertiesClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {
}

// ExitModifyPropertiesClause is called when production modifyPropertiesClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {
}

// EnterModifyCommentClause is called when production modifyCommentClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// ExitModifyCommentClause is called when production modifyCommentClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// EnterOptimizeRange is called when production optimizeRange is entered.
func (s *BaseStarRocksParserListener) EnterOptimizeRange(ctx *OptimizeRangeContext) {}

// ExitOptimizeRange is called when production optimizeRange is exited.
func (s *BaseStarRocksParserListener) ExitOptimizeRange(ctx *OptimizeRangeContext) {}

// EnterOptimizeClause is called when production optimizeClause is entered.
func (s *BaseStarRocksParserListener) EnterOptimizeClause(ctx *OptimizeClauseContext) {}

// ExitOptimizeClause is called when production optimizeClause is exited.
func (s *BaseStarRocksParserListener) ExitOptimizeClause(ctx *OptimizeClauseContext) {}

// EnterAddColumnClause is called when production addColumnClause is entered.
func (s *BaseStarRocksParserListener) EnterAddColumnClause(ctx *AddColumnClauseContext) {}

// ExitAddColumnClause is called when production addColumnClause is exited.
func (s *BaseStarRocksParserListener) ExitAddColumnClause(ctx *AddColumnClauseContext) {}

// EnterAddColumnsClause is called when production addColumnsClause is entered.
func (s *BaseStarRocksParserListener) EnterAddColumnsClause(ctx *AddColumnsClauseContext) {}

// ExitAddColumnsClause is called when production addColumnsClause is exited.
func (s *BaseStarRocksParserListener) ExitAddColumnsClause(ctx *AddColumnsClauseContext) {}

// EnterDropColumnClause is called when production dropColumnClause is entered.
func (s *BaseStarRocksParserListener) EnterDropColumnClause(ctx *DropColumnClauseContext) {}

// ExitDropColumnClause is called when production dropColumnClause is exited.
func (s *BaseStarRocksParserListener) ExitDropColumnClause(ctx *DropColumnClauseContext) {}

// EnterModifyColumnClause is called when production modifyColumnClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// ExitModifyColumnClause is called when production modifyColumnClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// EnterModifyColumnCommentClause is called when production modifyColumnCommentClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// ExitModifyColumnCommentClause is called when production modifyColumnCommentClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// EnterColumnRenameClause is called when production columnRenameClause is entered.
func (s *BaseStarRocksParserListener) EnterColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// ExitColumnRenameClause is called when production columnRenameClause is exited.
func (s *BaseStarRocksParserListener) ExitColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// EnterReorderColumnsClause is called when production reorderColumnsClause is entered.
func (s *BaseStarRocksParserListener) EnterReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// ExitReorderColumnsClause is called when production reorderColumnsClause is exited.
func (s *BaseStarRocksParserListener) ExitReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// EnterRollupRenameClause is called when production rollupRenameClause is entered.
func (s *BaseStarRocksParserListener) EnterRollupRenameClause(ctx *RollupRenameClauseContext) {}

// ExitRollupRenameClause is called when production rollupRenameClause is exited.
func (s *BaseStarRocksParserListener) ExitRollupRenameClause(ctx *RollupRenameClauseContext) {}

// EnterCompactionClause is called when production compactionClause is entered.
func (s *BaseStarRocksParserListener) EnterCompactionClause(ctx *CompactionClauseContext) {}

// ExitCompactionClause is called when production compactionClause is exited.
func (s *BaseStarRocksParserListener) ExitCompactionClause(ctx *CompactionClauseContext) {}

// EnterSubfieldName is called when production subfieldName is entered.
func (s *BaseStarRocksParserListener) EnterSubfieldName(ctx *SubfieldNameContext) {}

// ExitSubfieldName is called when production subfieldName is exited.
func (s *BaseStarRocksParserListener) ExitSubfieldName(ctx *SubfieldNameContext) {}

// EnterNestedFieldName is called when production nestedFieldName is entered.
func (s *BaseStarRocksParserListener) EnterNestedFieldName(ctx *NestedFieldNameContext) {}

// ExitNestedFieldName is called when production nestedFieldName is exited.
func (s *BaseStarRocksParserListener) ExitNestedFieldName(ctx *NestedFieldNameContext) {}

// EnterAddFieldClause is called when production addFieldClause is entered.
func (s *BaseStarRocksParserListener) EnterAddFieldClause(ctx *AddFieldClauseContext) {}

// ExitAddFieldClause is called when production addFieldClause is exited.
func (s *BaseStarRocksParserListener) ExitAddFieldClause(ctx *AddFieldClauseContext) {}

// EnterDropFieldClause is called when production dropFieldClause is entered.
func (s *BaseStarRocksParserListener) EnterDropFieldClause(ctx *DropFieldClauseContext) {}

// ExitDropFieldClause is called when production dropFieldClause is exited.
func (s *BaseStarRocksParserListener) ExitDropFieldClause(ctx *DropFieldClauseContext) {}

// EnterCreateOrReplaceTagClause is called when production createOrReplaceTagClause is entered.
func (s *BaseStarRocksParserListener) EnterCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// ExitCreateOrReplaceTagClause is called when production createOrReplaceTagClause is exited.
func (s *BaseStarRocksParserListener) ExitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// EnterCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is entered.
func (s *BaseStarRocksParserListener) EnterCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// ExitCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is exited.
func (s *BaseStarRocksParserListener) ExitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// EnterDropBranchClause is called when production dropBranchClause is entered.
func (s *BaseStarRocksParserListener) EnterDropBranchClause(ctx *DropBranchClauseContext) {}

// ExitDropBranchClause is called when production dropBranchClause is exited.
func (s *BaseStarRocksParserListener) ExitDropBranchClause(ctx *DropBranchClauseContext) {}

// EnterDropTagClause is called when production dropTagClause is entered.
func (s *BaseStarRocksParserListener) EnterDropTagClause(ctx *DropTagClauseContext) {}

// ExitDropTagClause is called when production dropTagClause is exited.
func (s *BaseStarRocksParserListener) ExitDropTagClause(ctx *DropTagClauseContext) {}

// EnterTableOperationClause is called when production tableOperationClause is entered.
func (s *BaseStarRocksParserListener) EnterTableOperationClause(ctx *TableOperationClauseContext) {}

// ExitTableOperationClause is called when production tableOperationClause is exited.
func (s *BaseStarRocksParserListener) ExitTableOperationClause(ctx *TableOperationClauseContext) {}

// EnterTagOptions is called when production tagOptions is entered.
func (s *BaseStarRocksParserListener) EnterTagOptions(ctx *TagOptionsContext) {}

// ExitTagOptions is called when production tagOptions is exited.
func (s *BaseStarRocksParserListener) ExitTagOptions(ctx *TagOptionsContext) {}

// EnterBranchOptions is called when production branchOptions is entered.
func (s *BaseStarRocksParserListener) EnterBranchOptions(ctx *BranchOptionsContext) {}

// ExitBranchOptions is called when production branchOptions is exited.
func (s *BaseStarRocksParserListener) ExitBranchOptions(ctx *BranchOptionsContext) {}

// EnterSnapshotRetention is called when production snapshotRetention is entered.
func (s *BaseStarRocksParserListener) EnterSnapshotRetention(ctx *SnapshotRetentionContext) {}

// ExitSnapshotRetention is called when production snapshotRetention is exited.
func (s *BaseStarRocksParserListener) ExitSnapshotRetention(ctx *SnapshotRetentionContext) {}

// EnterRefRetain is called when production refRetain is entered.
func (s *BaseStarRocksParserListener) EnterRefRetain(ctx *RefRetainContext) {}

// ExitRefRetain is called when production refRetain is exited.
func (s *BaseStarRocksParserListener) ExitRefRetain(ctx *RefRetainContext) {}

// EnterMaxSnapshotAge is called when production maxSnapshotAge is entered.
func (s *BaseStarRocksParserListener) EnterMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// ExitMaxSnapshotAge is called when production maxSnapshotAge is exited.
func (s *BaseStarRocksParserListener) ExitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// EnterMinSnapshotsToKeep is called when production minSnapshotsToKeep is entered.
func (s *BaseStarRocksParserListener) EnterMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// ExitMinSnapshotsToKeep is called when production minSnapshotsToKeep is exited.
func (s *BaseStarRocksParserListener) ExitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// EnterSnapshotId is called when production snapshotId is entered.
func (s *BaseStarRocksParserListener) EnterSnapshotId(ctx *SnapshotIdContext) {}

// ExitSnapshotId is called when production snapshotId is exited.
func (s *BaseStarRocksParserListener) ExitSnapshotId(ctx *SnapshotIdContext) {}

// EnterTimeUnit is called when production timeUnit is entered.
func (s *BaseStarRocksParserListener) EnterTimeUnit(ctx *TimeUnitContext) {}

// ExitTimeUnit is called when production timeUnit is exited.
func (s *BaseStarRocksParserListener) ExitTimeUnit(ctx *TimeUnitContext) {}

// EnterInteger_list is called when production integer_list is entered.
func (s *BaseStarRocksParserListener) EnterInteger_list(ctx *Integer_listContext) {}

// ExitInteger_list is called when production integer_list is exited.
func (s *BaseStarRocksParserListener) ExitInteger_list(ctx *Integer_listContext) {}

// EnterDropPersistentIndexClause is called when production dropPersistentIndexClause is entered.
func (s *BaseStarRocksParserListener) EnterDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// ExitDropPersistentIndexClause is called when production dropPersistentIndexClause is exited.
func (s *BaseStarRocksParserListener) ExitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// EnterSplitTabletClause is called when production splitTabletClause is entered.
func (s *BaseStarRocksParserListener) EnterSplitTabletClause(ctx *SplitTabletClauseContext) {}

// ExitSplitTabletClause is called when production splitTabletClause is exited.
func (s *BaseStarRocksParserListener) ExitSplitTabletClause(ctx *SplitTabletClauseContext) {}

// EnterAddPartitionClause is called when production addPartitionClause is entered.
func (s *BaseStarRocksParserListener) EnterAddPartitionClause(ctx *AddPartitionClauseContext) {}

// ExitAddPartitionClause is called when production addPartitionClause is exited.
func (s *BaseStarRocksParserListener) ExitAddPartitionClause(ctx *AddPartitionClauseContext) {}

// EnterDropPartitionClause is called when production dropPartitionClause is entered.
func (s *BaseStarRocksParserListener) EnterDropPartitionClause(ctx *DropPartitionClauseContext) {}

// ExitDropPartitionClause is called when production dropPartitionClause is exited.
func (s *BaseStarRocksParserListener) ExitDropPartitionClause(ctx *DropPartitionClauseContext) {}

// EnterTruncatePartitionClause is called when production truncatePartitionClause is entered.
func (s *BaseStarRocksParserListener) EnterTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {
}

// ExitTruncatePartitionClause is called when production truncatePartitionClause is exited.
func (s *BaseStarRocksParserListener) ExitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {
}

// EnterModifyPartitionClause is called when production modifyPartitionClause is entered.
func (s *BaseStarRocksParserListener) EnterModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// ExitModifyPartitionClause is called when production modifyPartitionClause is exited.
func (s *BaseStarRocksParserListener) ExitModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// EnterReplacePartitionClause is called when production replacePartitionClause is entered.
func (s *BaseStarRocksParserListener) EnterReplacePartitionClause(ctx *ReplacePartitionClauseContext) {
}

// ExitReplacePartitionClause is called when production replacePartitionClause is exited.
func (s *BaseStarRocksParserListener) ExitReplacePartitionClause(ctx *ReplacePartitionClauseContext) {
}

// EnterPartitionRenameClause is called when production partitionRenameClause is entered.
func (s *BaseStarRocksParserListener) EnterPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// ExitPartitionRenameClause is called when production partitionRenameClause is exited.
func (s *BaseStarRocksParserListener) ExitPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseStarRocksParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseStarRocksParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is entered.
func (s *BaseStarRocksParserListener) EnterInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// ExitInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is exited.
func (s *BaseStarRocksParserListener) ExitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// EnterColumnAliasesOrByName is called when production columnAliasesOrByName is entered.
func (s *BaseStarRocksParserListener) EnterColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// ExitColumnAliasesOrByName is called when production columnAliasesOrByName is exited.
func (s *BaseStarRocksParserListener) ExitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseStarRocksParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseStarRocksParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseStarRocksParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseStarRocksParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterCreateRoutineLoadStatement is called when production createRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// ExitCreateRoutineLoadStatement is called when production createRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// EnterAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// ExitAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// EnterDataSource is called when production dataSource is entered.
func (s *BaseStarRocksParserListener) EnterDataSource(ctx *DataSourceContext) {}

// ExitDataSource is called when production dataSource is exited.
func (s *BaseStarRocksParserListener) ExitDataSource(ctx *DataSourceContext) {}

// EnterLoadProperties is called when production loadProperties is entered.
func (s *BaseStarRocksParserListener) EnterLoadProperties(ctx *LoadPropertiesContext) {}

// ExitLoadProperties is called when production loadProperties is exited.
func (s *BaseStarRocksParserListener) ExitLoadProperties(ctx *LoadPropertiesContext) {}

// EnterColSeparatorProperty is called when production colSeparatorProperty is entered.
func (s *BaseStarRocksParserListener) EnterColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// ExitColSeparatorProperty is called when production colSeparatorProperty is exited.
func (s *BaseStarRocksParserListener) ExitColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// EnterRowDelimiterProperty is called when production rowDelimiterProperty is entered.
func (s *BaseStarRocksParserListener) EnterRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// ExitRowDelimiterProperty is called when production rowDelimiterProperty is exited.
func (s *BaseStarRocksParserListener) ExitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// EnterImportColumns is called when production importColumns is entered.
func (s *BaseStarRocksParserListener) EnterImportColumns(ctx *ImportColumnsContext) {}

// ExitImportColumns is called when production importColumns is exited.
func (s *BaseStarRocksParserListener) ExitImportColumns(ctx *ImportColumnsContext) {}

// EnterColumnProperties is called when production columnProperties is entered.
func (s *BaseStarRocksParserListener) EnterColumnProperties(ctx *ColumnPropertiesContext) {}

// ExitColumnProperties is called when production columnProperties is exited.
func (s *BaseStarRocksParserListener) ExitColumnProperties(ctx *ColumnPropertiesContext) {}

// EnterJobProperties is called when production jobProperties is entered.
func (s *BaseStarRocksParserListener) EnterJobProperties(ctx *JobPropertiesContext) {}

// ExitJobProperties is called when production jobProperties is exited.
func (s *BaseStarRocksParserListener) ExitJobProperties(ctx *JobPropertiesContext) {}

// EnterDataSourceProperties is called when production dataSourceProperties is entered.
func (s *BaseStarRocksParserListener) EnterDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// ExitDataSourceProperties is called when production dataSourceProperties is exited.
func (s *BaseStarRocksParserListener) ExitDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// EnterStopRoutineLoadStatement is called when production stopRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// ExitStopRoutineLoadStatement is called when production stopRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// EnterResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// ExitResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// EnterPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// ExitPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadStatement is called when production showRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// ExitShowRoutineLoadStatement is called when production showRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// ExitShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// EnterShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// ExitShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// EnterShowStreamLoadStatement is called when production showStreamLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {
}

// ExitShowStreamLoadStatement is called when production showStreamLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {
}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseStarRocksParserListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseStarRocksParserListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterRegularColumns is called when production regularColumns is entered.
func (s *BaseStarRocksParserListener) EnterRegularColumns(ctx *RegularColumnsContext) {}

// ExitRegularColumns is called when production regularColumns is exited.
func (s *BaseStarRocksParserListener) ExitRegularColumns(ctx *RegularColumnsContext) {}

// EnterAllColumns is called when production allColumns is entered.
func (s *BaseStarRocksParserListener) EnterAllColumns(ctx *AllColumnsContext) {}

// ExitAllColumns is called when production allColumns is exited.
func (s *BaseStarRocksParserListener) ExitAllColumns(ctx *AllColumnsContext) {}

// EnterPredicateColumns is called when production predicateColumns is entered.
func (s *BaseStarRocksParserListener) EnterPredicateColumns(ctx *PredicateColumnsContext) {}

// ExitPredicateColumns is called when production predicateColumns is exited.
func (s *BaseStarRocksParserListener) ExitPredicateColumns(ctx *PredicateColumnsContext) {}

// EnterMultiColumnSet is called when production multiColumnSet is entered.
func (s *BaseStarRocksParserListener) EnterMultiColumnSet(ctx *MultiColumnSetContext) {}

// ExitMultiColumnSet is called when production multiColumnSet is exited.
func (s *BaseStarRocksParserListener) ExitMultiColumnSet(ctx *MultiColumnSetContext) {}

// EnterDropStatsStatement is called when production dropStatsStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropStatsStatement(ctx *DropStatsStatementContext) {}

// ExitDropStatsStatement is called when production dropStatsStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropStatsStatement(ctx *DropStatsStatementContext) {}

// EnterHistogramStatement is called when production histogramStatement is entered.
func (s *BaseStarRocksParserListener) EnterHistogramStatement(ctx *HistogramStatementContext) {}

// ExitHistogramStatement is called when production histogramStatement is exited.
func (s *BaseStarRocksParserListener) ExitHistogramStatement(ctx *HistogramStatementContext) {}

// EnterAnalyzeHistogramStatement is called when production analyzeHistogramStatement is entered.
func (s *BaseStarRocksParserListener) EnterAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// ExitAnalyzeHistogramStatement is called when production analyzeHistogramStatement is exited.
func (s *BaseStarRocksParserListener) ExitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// EnterDropHistogramStatement is called when production dropHistogramStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropHistogramStatement(ctx *DropHistogramStatementContext) {
}

// ExitDropHistogramStatement is called when production dropHistogramStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropHistogramStatement(ctx *DropHistogramStatementContext) {
}

// EnterCreateAnalyzeStatement is called when production createAnalyzeStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {
}

// ExitCreateAnalyzeStatement is called when production createAnalyzeStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {
}

// EnterDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {
}

// ExitDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {
}

// EnterShowAnalyzeStatement is called when production showAnalyzeStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// ExitShowAnalyzeStatement is called when production showAnalyzeStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// EnterShowStatsMetaStatement is called when production showStatsMetaStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {
}

// ExitShowStatsMetaStatement is called when production showStatsMetaStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {
}

// EnterShowHistogramMetaStatement is called when production showHistogramMetaStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// ExitShowHistogramMetaStatement is called when production showHistogramMetaStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// EnterKillAnalyzeStatement is called when production killAnalyzeStatement is entered.
func (s *BaseStarRocksParserListener) EnterKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// ExitKillAnalyzeStatement is called when production killAnalyzeStatement is exited.
func (s *BaseStarRocksParserListener) ExitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// EnterAnalyzeProfileStatement is called when production analyzeProfileStatement is entered.
func (s *BaseStarRocksParserListener) EnterAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {
}

// ExitAnalyzeProfileStatement is called when production analyzeProfileStatement is exited.
func (s *BaseStarRocksParserListener) ExitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {
}

// EnterCreateBaselinePlanStatement is called when production createBaselinePlanStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// ExitCreateBaselinePlanStatement is called when production createBaselinePlanStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// EnterDropBaselinePlanStatement is called when production dropBaselinePlanStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// ExitDropBaselinePlanStatement is called when production dropBaselinePlanStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// EnterShowBaselinePlanStatement is called when production showBaselinePlanStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// ExitShowBaselinePlanStatement is called when production showBaselinePlanStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// EnterDisableBaselinePlanStatement is called when production disableBaselinePlanStatement is entered.
func (s *BaseStarRocksParserListener) EnterDisableBaselinePlanStatement(ctx *DisableBaselinePlanStatementContext) {
}

// ExitDisableBaselinePlanStatement is called when production disableBaselinePlanStatement is exited.
func (s *BaseStarRocksParserListener) ExitDisableBaselinePlanStatement(ctx *DisableBaselinePlanStatementContext) {
}

// EnterEnableBaselinePlanStatement is called when production enableBaselinePlanStatement is entered.
func (s *BaseStarRocksParserListener) EnterEnableBaselinePlanStatement(ctx *EnableBaselinePlanStatementContext) {
}

// ExitEnableBaselinePlanStatement is called when production enableBaselinePlanStatement is exited.
func (s *BaseStarRocksParserListener) ExitEnableBaselinePlanStatement(ctx *EnableBaselinePlanStatementContext) {
}

// EnterCreateResourceGroupStatement is called when production createResourceGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// ExitCreateResourceGroupStatement is called when production createResourceGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// EnterDropResourceGroupStatement is called when production dropResourceGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// ExitDropResourceGroupStatement is called when production dropResourceGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// EnterAlterResourceGroupStatement is called when production alterResourceGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// ExitAlterResourceGroupStatement is called when production alterResourceGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// EnterShowResourceGroupStatement is called when production showResourceGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// ExitShowResourceGroupStatement is called when production showResourceGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// EnterShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// ExitShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// EnterCreateResourceStatement is called when production createResourceStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateResourceStatement(ctx *CreateResourceStatementContext) {
}

// ExitCreateResourceStatement is called when production createResourceStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateResourceStatement(ctx *CreateResourceStatementContext) {
}

// EnterAlterResourceStatement is called when production alterResourceStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterResourceStatement(ctx *AlterResourceStatementContext) {
}

// ExitAlterResourceStatement is called when production alterResourceStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterResourceStatement(ctx *AlterResourceStatementContext) {
}

// EnterDropResourceStatement is called when production dropResourceStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropResourceStatement(ctx *DropResourceStatementContext) {}

// ExitDropResourceStatement is called when production dropResourceStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropResourceStatement(ctx *DropResourceStatementContext) {}

// EnterShowResourceStatement is called when production showResourceStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowResourceStatement(ctx *ShowResourceStatementContext) {}

// ExitShowResourceStatement is called when production showResourceStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowResourceStatement(ctx *ShowResourceStatementContext) {}

// EnterClassifier is called when production classifier is entered.
func (s *BaseStarRocksParserListener) EnterClassifier(ctx *ClassifierContext) {}

// ExitClassifier is called when production classifier is exited.
func (s *BaseStarRocksParserListener) ExitClassifier(ctx *ClassifierContext) {}

// EnterShowFunctionsStatement is called when production showFunctionsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {
}

// ExitShowFunctionsStatement is called when production showFunctionsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {
}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// EnterInlineFunction is called when production inlineFunction is entered.
func (s *BaseStarRocksParserListener) EnterInlineFunction(ctx *InlineFunctionContext) {}

// ExitInlineFunction is called when production inlineFunction is exited.
func (s *BaseStarRocksParserListener) ExitInlineFunction(ctx *InlineFunctionContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseStarRocksParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseStarRocksParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseStarRocksParserListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseStarRocksParserListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseStarRocksParserListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseStarRocksParserListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterDataDescList is called when production dataDescList is entered.
func (s *BaseStarRocksParserListener) EnterDataDescList(ctx *DataDescListContext) {}

// ExitDataDescList is called when production dataDescList is exited.
func (s *BaseStarRocksParserListener) ExitDataDescList(ctx *DataDescListContext) {}

// EnterDataDesc is called when production dataDesc is entered.
func (s *BaseStarRocksParserListener) EnterDataDesc(ctx *DataDescContext) {}

// ExitDataDesc is called when production dataDesc is exited.
func (s *BaseStarRocksParserListener) ExitDataDesc(ctx *DataDescContext) {}

// EnterFormatProps is called when production formatProps is entered.
func (s *BaseStarRocksParserListener) EnterFormatProps(ctx *FormatPropsContext) {}

// ExitFormatProps is called when production formatProps is exited.
func (s *BaseStarRocksParserListener) ExitFormatProps(ctx *FormatPropsContext) {}

// EnterBrokerDesc is called when production brokerDesc is entered.
func (s *BaseStarRocksParserListener) EnterBrokerDesc(ctx *BrokerDescContext) {}

// ExitBrokerDesc is called when production brokerDesc is exited.
func (s *BaseStarRocksParserListener) ExitBrokerDesc(ctx *BrokerDescContext) {}

// EnterResourceDesc is called when production resourceDesc is entered.
func (s *BaseStarRocksParserListener) EnterResourceDesc(ctx *ResourceDescContext) {}

// ExitResourceDesc is called when production resourceDesc is exited.
func (s *BaseStarRocksParserListener) ExitResourceDesc(ctx *ResourceDescContext) {}

// EnterShowLoadStatement is called when production showLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowLoadStatement(ctx *ShowLoadStatementContext) {}

// ExitShowLoadStatement is called when production showLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowLoadStatement(ctx *ShowLoadStatementContext) {}

// EnterShowLoadWarningsStatement is called when production showLoadWarningsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// ExitShowLoadWarningsStatement is called when production showLoadWarningsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// EnterCancelLoadStatement is called when production cancelLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// ExitCancelLoadStatement is called when production cancelLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// EnterAlterLoadStatement is called when production alterLoadStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// ExitAlterLoadStatement is called when production alterLoadStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// EnterCancelCompactionStatement is called when production cancelCompactionStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// ExitCancelCompactionStatement is called when production cancelCompactionStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// EnterShowAuthorStatement is called when production showAuthorStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// ExitShowAuthorStatement is called when production showAuthorStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// EnterShowBackendsStatement is called when production showBackendsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// ExitShowBackendsStatement is called when production showBackendsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// EnterShowBrokerStatement is called when production showBrokerStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// ExitShowBrokerStatement is called when production showBrokerStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// EnterShowCharsetStatement is called when production showCharsetStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// ExitShowCharsetStatement is called when production showCharsetStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// EnterShowCollationStatement is called when production showCollationStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCollationStatement(ctx *ShowCollationStatementContext) {
}

// ExitShowCollationStatement is called when production showCollationStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCollationStatement(ctx *ShowCollationStatementContext) {
}

// EnterShowDeleteStatement is called when production showDeleteStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// ExitShowDeleteStatement is called when production showDeleteStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// EnterShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// ExitShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// EnterShowEventsStatement is called when production showEventsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowEventsStatement(ctx *ShowEventsStatementContext) {}

// ExitShowEventsStatement is called when production showEventsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowEventsStatement(ctx *ShowEventsStatementContext) {}

// EnterShowEnginesStatement is called when production showEnginesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// ExitShowEnginesStatement is called when production showEnginesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// EnterShowFrontendsStatement is called when production showFrontendsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {
}

// ExitShowFrontendsStatement is called when production showFrontendsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {
}

// EnterShowPluginsStatement is called when production showPluginsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// ExitShowPluginsStatement is called when production showPluginsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// EnterShowRepositoriesStatement is called when production showRepositoriesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// ExitShowRepositoriesStatement is called when production showRepositoriesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// EnterShowOpenTableStatement is called when production showOpenTableStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {
}

// ExitShowOpenTableStatement is called when production showOpenTableStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {
}

// EnterShowPrivilegesStatement is called when production showPrivilegesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {
}

// ExitShowPrivilegesStatement is called when production showPrivilegesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {
}

// EnterShowProcedureStatement is called when production showProcedureStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowProcedureStatement(ctx *ShowProcedureStatementContext) {
}

// ExitShowProcedureStatement is called when production showProcedureStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowProcedureStatement(ctx *ShowProcedureStatementContext) {
}

// EnterShowProcStatement is called when production showProcStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowProcStatement(ctx *ShowProcStatementContext) {}

// ExitShowProcStatement is called when production showProcStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowProcStatement(ctx *ShowProcStatementContext) {}

// EnterShowProcesslistStatement is called when production showProcesslistStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// ExitShowProcesslistStatement is called when production showProcesslistStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// EnterShowProfilelistStatement is called when production showProfilelistStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// ExitShowProfilelistStatement is called when production showProfilelistStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// EnterShowRunningQueriesStatement is called when production showRunningQueriesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// ExitShowRunningQueriesStatement is called when production showRunningQueriesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// EnterShowStatusStatement is called when production showStatusStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowStatusStatement(ctx *ShowStatusStatementContext) {}

// ExitShowStatusStatement is called when production showStatusStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowStatusStatement(ctx *ShowStatusStatementContext) {}

// EnterShowTabletStatement is called when production showTabletStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTabletStatement(ctx *ShowTabletStatementContext) {}

// ExitShowTabletStatement is called when production showTabletStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTabletStatement(ctx *ShowTabletStatementContext) {}

// EnterShowTransactionStatement is called when production showTransactionStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// ExitShowTransactionStatement is called when production showTransactionStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// EnterShowTriggersStatement is called when production showTriggersStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// ExitShowTriggersStatement is called when production showTriggersStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// EnterShowUserPropertyStatement is called when production showUserPropertyStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// ExitShowUserPropertyStatement is called when production showUserPropertyStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// EnterShowVariablesStatement is called when production showVariablesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowVariablesStatement(ctx *ShowVariablesStatementContext) {
}

// ExitShowVariablesStatement is called when production showVariablesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowVariablesStatement(ctx *ShowVariablesStatementContext) {
}

// EnterShowWarningStatement is called when production showWarningStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowWarningStatement(ctx *ShowWarningStatementContext) {}

// ExitShowWarningStatement is called when production showWarningStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowWarningStatement(ctx *ShowWarningStatementContext) {}

// EnterHelpStatement is called when production helpStatement is entered.
func (s *BaseStarRocksParserListener) EnterHelpStatement(ctx *HelpStatementContext) {}

// ExitHelpStatement is called when production helpStatement is exited.
func (s *BaseStarRocksParserListener) ExitHelpStatement(ctx *HelpStatementContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterDropUserStatement is called when production dropUserStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropUserStatement(ctx *DropUserStatementContext) {}

// ExitDropUserStatement is called when production dropUserStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropUserStatement(ctx *DropUserStatementContext) {}

// EnterAlterUserStatement is called when production alterUserStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterUserStatement(ctx *AlterUserStatementContext) {}

// ExitAlterUserStatement is called when production alterUserStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterUserStatement(ctx *AlterUserStatementContext) {}

// EnterShowUserStatement is called when production showUserStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowUserStatement(ctx *ShowUserStatementContext) {}

// ExitShowUserStatement is called when production showUserStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowUserStatement(ctx *ShowUserStatementContext) {}

// EnterShowAllAuthentication is called when production showAllAuthentication is entered.
func (s *BaseStarRocksParserListener) EnterShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// ExitShowAllAuthentication is called when production showAllAuthentication is exited.
func (s *BaseStarRocksParserListener) ExitShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// EnterShowAuthenticationForUser is called when production showAuthenticationForUser is entered.
func (s *BaseStarRocksParserListener) EnterShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// ExitShowAuthenticationForUser is called when production showAuthenticationForUser is exited.
func (s *BaseStarRocksParserListener) ExitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// EnterExecuteAsStatement is called when production executeAsStatement is entered.
func (s *BaseStarRocksParserListener) EnterExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// ExitExecuteAsStatement is called when production executeAsStatement is exited.
func (s *BaseStarRocksParserListener) ExitExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterAlterRoleStatement is called when production alterRoleStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// ExitAlterRoleStatement is called when production alterRoleStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterShowRolesStatement is called when production showRolesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRolesStatement(ctx *ShowRolesStatementContext) {}

// ExitShowRolesStatement is called when production showRolesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRolesStatement(ctx *ShowRolesStatementContext) {}

// EnterGrantRoleToUser is called when production grantRoleToUser is entered.
func (s *BaseStarRocksParserListener) EnterGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// ExitGrantRoleToUser is called when production grantRoleToUser is exited.
func (s *BaseStarRocksParserListener) ExitGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// EnterGrantRoleToRole is called when production grantRoleToRole is entered.
func (s *BaseStarRocksParserListener) EnterGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// ExitGrantRoleToRole is called when production grantRoleToRole is exited.
func (s *BaseStarRocksParserListener) ExitGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// EnterRevokeRoleFromUser is called when production revokeRoleFromUser is entered.
func (s *BaseStarRocksParserListener) EnterRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// ExitRevokeRoleFromUser is called when production revokeRoleFromUser is exited.
func (s *BaseStarRocksParserListener) ExitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// EnterRevokeRoleFromRole is called when production revokeRoleFromRole is entered.
func (s *BaseStarRocksParserListener) EnterRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// ExitRevokeRoleFromRole is called when production revokeRoleFromRole is exited.
func (s *BaseStarRocksParserListener) ExitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// EnterSetRoleStatement is called when production setRoleStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetRoleStatement(ctx *SetRoleStatementContext) {}

// ExitSetRoleStatement is called when production setRoleStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetRoleStatement(ctx *SetRoleStatementContext) {}

// EnterSetDefaultRoleStatement is called when production setDefaultRoleStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {
}

// ExitSetDefaultRoleStatement is called when production setDefaultRoleStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {
}

// EnterGrantRevokeClause is called when production grantRevokeClause is entered.
func (s *BaseStarRocksParserListener) EnterGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// ExitGrantRevokeClause is called when production grantRevokeClause is exited.
func (s *BaseStarRocksParserListener) ExitGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// EnterGrantOnUser is called when production grantOnUser is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnUser(ctx *GrantOnUserContext) {}

// ExitGrantOnUser is called when production grantOnUser is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnUser(ctx *GrantOnUserContext) {}

// EnterGrantOnTableBrief is called when production grantOnTableBrief is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// ExitGrantOnTableBrief is called when production grantOnTableBrief is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// EnterGrantOnFunc is called when production grantOnFunc is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnFunc(ctx *GrantOnFuncContext) {}

// ExitGrantOnFunc is called when production grantOnFunc is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnFunc(ctx *GrantOnFuncContext) {}

// EnterGrantOnSystem is called when production grantOnSystem is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnSystem(ctx *GrantOnSystemContext) {}

// ExitGrantOnSystem is called when production grantOnSystem is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnSystem(ctx *GrantOnSystemContext) {}

// EnterGrantOnPrimaryObj is called when production grantOnPrimaryObj is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// ExitGrantOnPrimaryObj is called when production grantOnPrimaryObj is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// EnterGrantOnAll is called when production grantOnAll is entered.
func (s *BaseStarRocksParserListener) EnterGrantOnAll(ctx *GrantOnAllContext) {}

// ExitGrantOnAll is called when production grantOnAll is exited.
func (s *BaseStarRocksParserListener) ExitGrantOnAll(ctx *GrantOnAllContext) {}

// EnterRevokeOnUser is called when production revokeOnUser is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnUser(ctx *RevokeOnUserContext) {}

// ExitRevokeOnUser is called when production revokeOnUser is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnUser(ctx *RevokeOnUserContext) {}

// EnterRevokeOnTableBrief is called when production revokeOnTableBrief is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// ExitRevokeOnTableBrief is called when production revokeOnTableBrief is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// EnterRevokeOnFunc is called when production revokeOnFunc is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// ExitRevokeOnFunc is called when production revokeOnFunc is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// EnterRevokeOnSystem is called when production revokeOnSystem is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// ExitRevokeOnSystem is called when production revokeOnSystem is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// EnterRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// ExitRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// EnterRevokeOnAll is called when production revokeOnAll is entered.
func (s *BaseStarRocksParserListener) EnterRevokeOnAll(ctx *RevokeOnAllContext) {}

// ExitRevokeOnAll is called when production revokeOnAll is exited.
func (s *BaseStarRocksParserListener) ExitRevokeOnAll(ctx *RevokeOnAllContext) {}

// EnterShowGrantsStatement is called when production showGrantsStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// ExitShowGrantsStatement is called when production showGrantsStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// EnterAuthWithoutPlugin is called when production authWithoutPlugin is entered.
func (s *BaseStarRocksParserListener) EnterAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// ExitAuthWithoutPlugin is called when production authWithoutPlugin is exited.
func (s *BaseStarRocksParserListener) ExitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// EnterAuthWithPlugin is called when production authWithPlugin is entered.
func (s *BaseStarRocksParserListener) EnterAuthWithPlugin(ctx *AuthWithPluginContext) {}

// ExitAuthWithPlugin is called when production authWithPlugin is exited.
func (s *BaseStarRocksParserListener) ExitAuthWithPlugin(ctx *AuthWithPluginContext) {}

// EnterPrivObjectName is called when production privObjectName is entered.
func (s *BaseStarRocksParserListener) EnterPrivObjectName(ctx *PrivObjectNameContext) {}

// ExitPrivObjectName is called when production privObjectName is exited.
func (s *BaseStarRocksParserListener) ExitPrivObjectName(ctx *PrivObjectNameContext) {}

// EnterPrivObjectNameList is called when production privObjectNameList is entered.
func (s *BaseStarRocksParserListener) EnterPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// ExitPrivObjectNameList is called when production privObjectNameList is exited.
func (s *BaseStarRocksParserListener) ExitPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// EnterPrivFunctionObjectNameList is called when production privFunctionObjectNameList is entered.
func (s *BaseStarRocksParserListener) EnterPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// ExitPrivFunctionObjectNameList is called when production privFunctionObjectNameList is exited.
func (s *BaseStarRocksParserListener) ExitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// EnterPrivilegeTypeList is called when production privilegeTypeList is entered.
func (s *BaseStarRocksParserListener) EnterPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// ExitPrivilegeTypeList is called when production privilegeTypeList is exited.
func (s *BaseStarRocksParserListener) ExitPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseStarRocksParserListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseStarRocksParserListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterPrivObjectType is called when production privObjectType is entered.
func (s *BaseStarRocksParserListener) EnterPrivObjectType(ctx *PrivObjectTypeContext) {}

// ExitPrivObjectType is called when production privObjectType is exited.
func (s *BaseStarRocksParserListener) ExitPrivObjectType(ctx *PrivObjectTypeContext) {}

// EnterPrivObjectTypePlural is called when production privObjectTypePlural is entered.
func (s *BaseStarRocksParserListener) EnterPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// ExitPrivObjectTypePlural is called when production privObjectTypePlural is exited.
func (s *BaseStarRocksParserListener) ExitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// EnterCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// ExitCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// EnterAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// ExitAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// EnterDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// ExitDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// EnterShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// ExitShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// EnterShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// ExitShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// EnterCreateGroupProviderStatement is called when production createGroupProviderStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// ExitCreateGroupProviderStatement is called when production createGroupProviderStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// EnterDropGroupProviderStatement is called when production dropGroupProviderStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// ExitDropGroupProviderStatement is called when production dropGroupProviderStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// EnterShowGroupProvidersStatement is called when production showGroupProvidersStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// ExitShowGroupProvidersStatement is called when production showGroupProvidersStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// EnterShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// ExitShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// EnterBackupStatement is called when production backupStatement is entered.
func (s *BaseStarRocksParserListener) EnterBackupStatement(ctx *BackupStatementContext) {}

// ExitBackupStatement is called when production backupStatement is exited.
func (s *BaseStarRocksParserListener) ExitBackupStatement(ctx *BackupStatementContext) {}

// EnterCancelBackupStatement is called when production cancelBackupStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// ExitCancelBackupStatement is called when production cancelBackupStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// EnterShowBackupStatement is called when production showBackupStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowBackupStatement(ctx *ShowBackupStatementContext) {}

// ExitShowBackupStatement is called when production showBackupStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowBackupStatement(ctx *ShowBackupStatementContext) {}

// EnterRestoreStatement is called when production restoreStatement is entered.
func (s *BaseStarRocksParserListener) EnterRestoreStatement(ctx *RestoreStatementContext) {}

// ExitRestoreStatement is called when production restoreStatement is exited.
func (s *BaseStarRocksParserListener) ExitRestoreStatement(ctx *RestoreStatementContext) {}

// EnterCancelRestoreStatement is called when production cancelRestoreStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelRestoreStatement(ctx *CancelRestoreStatementContext) {
}

// ExitCancelRestoreStatement is called when production cancelRestoreStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelRestoreStatement(ctx *CancelRestoreStatementContext) {
}

// EnterShowRestoreStatement is called when production showRestoreStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// ExitShowRestoreStatement is called when production showRestoreStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// EnterShowSnapshotStatement is called when production showSnapshotStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// ExitShowSnapshotStatement is called when production showSnapshotStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// EnterCreateRepositoryStatement is called when production createRepositoryStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// ExitCreateRepositoryStatement is called when production createRepositoryStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// EnterDropRepositoryStatement is called when production dropRepositoryStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropRepositoryStatement(ctx *DropRepositoryStatementContext) {
}

// ExitDropRepositoryStatement is called when production dropRepositoryStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropRepositoryStatement(ctx *DropRepositoryStatementContext) {
}

// EnterAddSqlBlackListStatement is called when production addSqlBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// ExitAddSqlBlackListStatement is called when production addSqlBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// EnterDelSqlBlackListStatement is called when production delSqlBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// ExitDelSqlBlackListStatement is called when production delSqlBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// EnterShowSqlBlackListStatement is called when production showSqlBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// ExitShowSqlBlackListStatement is called when production showSqlBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// EnterShowWhiteListStatement is called when production showWhiteListStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {
}

// ExitShowWhiteListStatement is called when production showWhiteListStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {
}

// EnterAddBackendBlackListStatement is called when production addBackendBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// ExitAddBackendBlackListStatement is called when production addBackendBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// EnterDelBackendBlackListStatement is called when production delBackendBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// ExitDelBackendBlackListStatement is called when production delBackendBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// EnterShowBackendBlackListStatement is called when production showBackendBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// ExitShowBackendBlackListStatement is called when production showBackendBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// EnterAddComputeNodeBlackListStatement is called when production addComputeNodeBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterAddComputeNodeBlackListStatement(ctx *AddComputeNodeBlackListStatementContext) {
}

// ExitAddComputeNodeBlackListStatement is called when production addComputeNodeBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitAddComputeNodeBlackListStatement(ctx *AddComputeNodeBlackListStatementContext) {
}

// EnterDelComputeNodeBlackListStatement is called when production delComputeNodeBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterDelComputeNodeBlackListStatement(ctx *DelComputeNodeBlackListStatementContext) {
}

// ExitDelComputeNodeBlackListStatement is called when production delComputeNodeBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitDelComputeNodeBlackListStatement(ctx *DelComputeNodeBlackListStatementContext) {
}

// EnterShowComputeNodeBlackListStatement is called when production showComputeNodeBlackListStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowComputeNodeBlackListStatement(ctx *ShowComputeNodeBlackListStatementContext) {
}

// ExitShowComputeNodeBlackListStatement is called when production showComputeNodeBlackListStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowComputeNodeBlackListStatement(ctx *ShowComputeNodeBlackListStatementContext) {
}

// EnterDataCacheTarget is called when production dataCacheTarget is entered.
func (s *BaseStarRocksParserListener) EnterDataCacheTarget(ctx *DataCacheTargetContext) {}

// ExitDataCacheTarget is called when production dataCacheTarget is exited.
func (s *BaseStarRocksParserListener) ExitDataCacheTarget(ctx *DataCacheTargetContext) {}

// EnterCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// ExitCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// EnterShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// ExitShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// EnterDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// ExitDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// EnterClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is entered.
func (s *BaseStarRocksParserListener) EnterClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// ExitClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is exited.
func (s *BaseStarRocksParserListener) ExitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// EnterDataCacheSelectStatement is called when production dataCacheSelectStatement is entered.
func (s *BaseStarRocksParserListener) EnterDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// ExitDataCacheSelectStatement is called when production dataCacheSelectStatement is exited.
func (s *BaseStarRocksParserListener) ExitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseStarRocksParserListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseStarRocksParserListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterCancelExportStatement is called when production cancelExportStatement is entered.
func (s *BaseStarRocksParserListener) EnterCancelExportStatement(ctx *CancelExportStatementContext) {}

// ExitCancelExportStatement is called when production cancelExportStatement is exited.
func (s *BaseStarRocksParserListener) ExitCancelExportStatement(ctx *CancelExportStatementContext) {}

// EnterShowExportStatement is called when production showExportStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowExportStatement(ctx *ShowExportStatementContext) {}

// ExitShowExportStatement is called when production showExportStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowExportStatement(ctx *ShowExportStatementContext) {}

// EnterInstallPluginStatement is called when production installPluginStatement is entered.
func (s *BaseStarRocksParserListener) EnterInstallPluginStatement(ctx *InstallPluginStatementContext) {
}

// ExitInstallPluginStatement is called when production installPluginStatement is exited.
func (s *BaseStarRocksParserListener) ExitInstallPluginStatement(ctx *InstallPluginStatementContext) {
}

// EnterUninstallPluginStatement is called when production uninstallPluginStatement is entered.
func (s *BaseStarRocksParserListener) EnterUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// ExitUninstallPluginStatement is called when production uninstallPluginStatement is exited.
func (s *BaseStarRocksParserListener) ExitUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// EnterCreateFileStatement is called when production createFileStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateFileStatement(ctx *CreateFileStatementContext) {}

// ExitCreateFileStatement is called when production createFileStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateFileStatement(ctx *CreateFileStatementContext) {}

// EnterDropFileStatement is called when production dropFileStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropFileStatement(ctx *DropFileStatementContext) {}

// ExitDropFileStatement is called when production dropFileStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropFileStatement(ctx *DropFileStatementContext) {}

// EnterShowSmallFilesStatement is called when production showSmallFilesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {
}

// ExitShowSmallFilesStatement is called when production showSmallFilesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {
}

// EnterCreatePipeStatement is called when production createPipeStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// ExitCreatePipeStatement is called when production createPipeStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// EnterDropPipeStatement is called when production dropPipeStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropPipeStatement(ctx *DropPipeStatementContext) {}

// ExitDropPipeStatement is called when production dropPipeStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropPipeStatement(ctx *DropPipeStatementContext) {}

// EnterAlterPipeClause is called when production alterPipeClause is entered.
func (s *BaseStarRocksParserListener) EnterAlterPipeClause(ctx *AlterPipeClauseContext) {}

// ExitAlterPipeClause is called when production alterPipeClause is exited.
func (s *BaseStarRocksParserListener) ExitAlterPipeClause(ctx *AlterPipeClauseContext) {}

// EnterAlterPipeStatement is called when production alterPipeStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// ExitAlterPipeStatement is called when production alterPipeStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// EnterDescPipeStatement is called when production descPipeStatement is entered.
func (s *BaseStarRocksParserListener) EnterDescPipeStatement(ctx *DescPipeStatementContext) {}

// ExitDescPipeStatement is called when production descPipeStatement is exited.
func (s *BaseStarRocksParserListener) ExitDescPipeStatement(ctx *DescPipeStatementContext) {}

// EnterShowPipeStatement is called when production showPipeStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowPipeStatement(ctx *ShowPipeStatementContext) {}

// ExitShowPipeStatement is called when production showPipeStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowPipeStatement(ctx *ShowPipeStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterSetNames is called when production setNames is entered.
func (s *BaseStarRocksParserListener) EnterSetNames(ctx *SetNamesContext) {}

// ExitSetNames is called when production setNames is exited.
func (s *BaseStarRocksParserListener) ExitSetNames(ctx *SetNamesContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseStarRocksParserListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseStarRocksParserListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterSetUserVar is called when production setUserVar is entered.
func (s *BaseStarRocksParserListener) EnterSetUserVar(ctx *SetUserVarContext) {}

// ExitSetUserVar is called when production setUserVar is exited.
func (s *BaseStarRocksParserListener) ExitSetUserVar(ctx *SetUserVarContext) {}

// EnterSetSystemVar is called when production setSystemVar is entered.
func (s *BaseStarRocksParserListener) EnterSetSystemVar(ctx *SetSystemVarContext) {}

// ExitSetSystemVar is called when production setSystemVar is exited.
func (s *BaseStarRocksParserListener) ExitSetSystemVar(ctx *SetSystemVarContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseStarRocksParserListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseStarRocksParserListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterTransaction_characteristics is called when production transaction_characteristics is entered.
func (s *BaseStarRocksParserListener) EnterTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// ExitTransaction_characteristics is called when production transaction_characteristics is exited.
func (s *BaseStarRocksParserListener) ExitTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// EnterTransaction_access_mode is called when production transaction_access_mode is entered.
func (s *BaseStarRocksParserListener) EnterTransaction_access_mode(ctx *Transaction_access_modeContext) {
}

// ExitTransaction_access_mode is called when production transaction_access_mode is exited.
func (s *BaseStarRocksParserListener) ExitTransaction_access_mode(ctx *Transaction_access_modeContext) {
}

// EnterIsolation_level is called when production isolation_level is entered.
func (s *BaseStarRocksParserListener) EnterIsolation_level(ctx *Isolation_levelContext) {}

// ExitIsolation_level is called when production isolation_level is exited.
func (s *BaseStarRocksParserListener) ExitIsolation_level(ctx *Isolation_levelContext) {}

// EnterIsolation_types is called when production isolation_types is entered.
func (s *BaseStarRocksParserListener) EnterIsolation_types(ctx *Isolation_typesContext) {}

// ExitIsolation_types is called when production isolation_types is exited.
func (s *BaseStarRocksParserListener) ExitIsolation_types(ctx *Isolation_typesContext) {}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseStarRocksParserListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseStarRocksParserListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterSetUserPropertyStatement is called when production setUserPropertyStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// ExitSetUserPropertyStatement is called when production setUserPropertyStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// EnterRoleList is called when production roleList is entered.
func (s *BaseStarRocksParserListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseStarRocksParserListener) ExitRoleList(ctx *RoleListContext) {}

// EnterExecuteScriptStatement is called when production executeScriptStatement is entered.
func (s *BaseStarRocksParserListener) EnterExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {
}

// ExitExecuteScriptStatement is called when production executeScriptStatement is exited.
func (s *BaseStarRocksParserListener) ExitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {
}

// EnterUnsupportedStatement is called when production unsupportedStatement is entered.
func (s *BaseStarRocksParserListener) EnterUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// ExitUnsupportedStatement is called when production unsupportedStatement is exited.
func (s *BaseStarRocksParserListener) ExitUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// EnterLock_item is called when production lock_item is entered.
func (s *BaseStarRocksParserListener) EnterLock_item(ctx *Lock_itemContext) {}

// ExitLock_item is called when production lock_item is exited.
func (s *BaseStarRocksParserListener) ExitLock_item(ctx *Lock_itemContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseStarRocksParserListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseStarRocksParserListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// ExitAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// EnterTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is entered.
func (s *BaseStarRocksParserListener) EnterTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// ExitTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is exited.
func (s *BaseStarRocksParserListener) ExitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// EnterAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// ExitAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// EnterShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// ExitShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// EnterCreateWarehouseStatement is called when production createWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// ExitCreateWarehouseStatement is called when production createWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// EnterDropWarehouseStatement is called when production dropWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropWarehouseStatement(ctx *DropWarehouseStatementContext) {
}

// ExitDropWarehouseStatement is called when production dropWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropWarehouseStatement(ctx *DropWarehouseStatementContext) {
}

// EnterSuspendWarehouseStatement is called when production suspendWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// ExitSuspendWarehouseStatement is called when production suspendWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// EnterResumeWarehouseStatement is called when production resumeWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// ExitResumeWarehouseStatement is called when production resumeWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// EnterSetWarehouseStatement is called when production setWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// ExitSetWarehouseStatement is called when production setWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// EnterShowWarehousesStatement is called when production showWarehousesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {
}

// ExitShowWarehousesStatement is called when production showWarehousesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {
}

// EnterShowClustersStatement is called when production showClustersStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowClustersStatement(ctx *ShowClustersStatementContext) {}

// ExitShowClustersStatement is called when production showClustersStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowClustersStatement(ctx *ShowClustersStatementContext) {}

// EnterShowNodesStatement is called when production showNodesStatement is entered.
func (s *BaseStarRocksParserListener) EnterShowNodesStatement(ctx *ShowNodesStatementContext) {}

// ExitShowNodesStatement is called when production showNodesStatement is exited.
func (s *BaseStarRocksParserListener) ExitShowNodesStatement(ctx *ShowNodesStatementContext) {}

// EnterAlterWarehouseStatement is called when production alterWarehouseStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {
}

// ExitAlterWarehouseStatement is called when production alterWarehouseStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {
}

// EnterCreateCNGroupStatement is called when production createCNGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {
}

// ExitCreateCNGroupStatement is called when production createCNGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {
}

// EnterDropCNGroupStatement is called when production dropCNGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// ExitDropCNGroupStatement is called when production dropCNGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// EnterEnableCNGroupStatement is called when production enableCNGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {
}

// ExitEnableCNGroupStatement is called when production enableCNGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {
}

// EnterDisableCNGroupStatement is called when production disableCNGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {
}

// ExitDisableCNGroupStatement is called when production disableCNGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {
}

// EnterAlterCNGroupStatement is called when production alterCNGroupStatement is entered.
func (s *BaseStarRocksParserListener) EnterAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// ExitAlterCNGroupStatement is called when production alterCNGroupStatement is exited.
func (s *BaseStarRocksParserListener) ExitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseStarRocksParserListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseStarRocksParserListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseStarRocksParserListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseStarRocksParserListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseStarRocksParserListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseStarRocksParserListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterTranslateStatement is called when production translateStatement is entered.
func (s *BaseStarRocksParserListener) EnterTranslateStatement(ctx *TranslateStatementContext) {}

// ExitTranslateStatement is called when production translateStatement is exited.
func (s *BaseStarRocksParserListener) ExitTranslateStatement(ctx *TranslateStatementContext) {}

// EnterDialect is called when production dialect is entered.
func (s *BaseStarRocksParserListener) EnterDialect(ctx *DialectContext) {}

// ExitDialect is called when production dialect is exited.
func (s *BaseStarRocksParserListener) ExitDialect(ctx *DialectContext) {}

// EnterTranslateSQL is called when production translateSQL is entered.
func (s *BaseStarRocksParserListener) EnterTranslateSQL(ctx *TranslateSQLContext) {}

// ExitTranslateSQL is called when production translateSQL is exited.
func (s *BaseStarRocksParserListener) ExitTranslateSQL(ctx *TranslateSQLContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseStarRocksParserListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseStarRocksParserListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterQueryRelation is called when production queryRelation is entered.
func (s *BaseStarRocksParserListener) EnterQueryRelation(ctx *QueryRelationContext) {}

// ExitQueryRelation is called when production queryRelation is exited.
func (s *BaseStarRocksParserListener) ExitQueryRelation(ctx *QueryRelationContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseStarRocksParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseStarRocksParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterQueryNoWith is called when production queryNoWith is entered.
func (s *BaseStarRocksParserListener) EnterQueryNoWith(ctx *QueryNoWithContext) {}

// ExitQueryNoWith is called when production queryNoWith is exited.
func (s *BaseStarRocksParserListener) ExitQueryNoWith(ctx *QueryNoWithContext) {}

// EnterQueryPeriod is called when production queryPeriod is entered.
func (s *BaseStarRocksParserListener) EnterQueryPeriod(ctx *QueryPeriodContext) {}

// ExitQueryPeriod is called when production queryPeriod is exited.
func (s *BaseStarRocksParserListener) ExitQueryPeriod(ctx *QueryPeriodContext) {}

// EnterPeriodType is called when production periodType is entered.
func (s *BaseStarRocksParserListener) EnterPeriodType(ctx *PeriodTypeContext) {}

// ExitPeriodType is called when production periodType is exited.
func (s *BaseStarRocksParserListener) ExitPeriodType(ctx *PeriodTypeContext) {}

// EnterQueryWithParentheses is called when production queryWithParentheses is entered.
func (s *BaseStarRocksParserListener) EnterQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// ExitQueryWithParentheses is called when production queryWithParentheses is exited.
func (s *BaseStarRocksParserListener) ExitQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseStarRocksParserListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseStarRocksParserListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseStarRocksParserListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseStarRocksParserListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseStarRocksParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseStarRocksParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseStarRocksParserListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseStarRocksParserListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseStarRocksParserListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseStarRocksParserListener) ExitSortItem(ctx *SortItemContext) {}

// EnterLimitConstExpr is called when production limitConstExpr is entered.
func (s *BaseStarRocksParserListener) EnterLimitConstExpr(ctx *LimitConstExprContext) {}

// ExitLimitConstExpr is called when production limitConstExpr is exited.
func (s *BaseStarRocksParserListener) ExitLimitConstExpr(ctx *LimitConstExprContext) {}

// EnterLimitElement is called when production limitElement is entered.
func (s *BaseStarRocksParserListener) EnterLimitElement(ctx *LimitElementContext) {}

// ExitLimitElement is called when production limitElement is exited.
func (s *BaseStarRocksParserListener) ExitLimitElement(ctx *LimitElementContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseStarRocksParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseStarRocksParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseStarRocksParserListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseStarRocksParserListener) ExitFrom(ctx *FromContext) {}

// EnterDual is called when production dual is entered.
func (s *BaseStarRocksParserListener) EnterDual(ctx *DualContext) {}

// ExitDual is called when production dual is exited.
func (s *BaseStarRocksParserListener) ExitDual(ctx *DualContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseStarRocksParserListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseStarRocksParserListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseStarRocksParserListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseStarRocksParserListener) ExitCube(ctx *CubeContext) {}

// EnterMultipleGroupingSets is called when production multipleGroupingSets is entered.
func (s *BaseStarRocksParserListener) EnterMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// ExitMultipleGroupingSets is called when production multipleGroupingSets is exited.
func (s *BaseStarRocksParserListener) ExitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseStarRocksParserListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseStarRocksParserListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseStarRocksParserListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseStarRocksParserListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseStarRocksParserListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseStarRocksParserListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseStarRocksParserListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseStarRocksParserListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseStarRocksParserListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseStarRocksParserListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseStarRocksParserListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseStarRocksParserListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterExcludeClause is called when production excludeClause is entered.
func (s *BaseStarRocksParserListener) EnterExcludeClause(ctx *ExcludeClauseContext) {}

// ExitExcludeClause is called when production excludeClause is exited.
func (s *BaseStarRocksParserListener) ExitExcludeClause(ctx *ExcludeClauseContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseStarRocksParserListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseStarRocksParserListener) ExitRelations(ctx *RelationsContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseStarRocksParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseStarRocksParserListener) ExitRelation(ctx *RelationContext) {}

// EnterTableAtom is called when production tableAtom is entered.
func (s *BaseStarRocksParserListener) EnterTableAtom(ctx *TableAtomContext) {}

// ExitTableAtom is called when production tableAtom is exited.
func (s *BaseStarRocksParserListener) ExitTableAtom(ctx *TableAtomContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseStarRocksParserListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseStarRocksParserListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubqueryWithAlias is called when production subqueryWithAlias is entered.
func (s *BaseStarRocksParserListener) EnterSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// ExitSubqueryWithAlias is called when production subqueryWithAlias is exited.
func (s *BaseStarRocksParserListener) ExitSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseStarRocksParserListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseStarRocksParserListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterNormalizedTableFunction is called when production normalizedTableFunction is entered.
func (s *BaseStarRocksParserListener) EnterNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {
}

// ExitNormalizedTableFunction is called when production normalizedTableFunction is exited.
func (s *BaseStarRocksParserListener) ExitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {
}

// EnterFileTableFunction is called when production fileTableFunction is entered.
func (s *BaseStarRocksParserListener) EnterFileTableFunction(ctx *FileTableFunctionContext) {}

// ExitFileTableFunction is called when production fileTableFunction is exited.
func (s *BaseStarRocksParserListener) ExitFileTableFunction(ctx *FileTableFunctionContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseStarRocksParserListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseStarRocksParserListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseStarRocksParserListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseStarRocksParserListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotAggregationExpression is called when production pivotAggregationExpression is entered.
func (s *BaseStarRocksParserListener) EnterPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// ExitPivotAggregationExpression is called when production pivotAggregationExpression is exited.
func (s *BaseStarRocksParserListener) ExitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseStarRocksParserListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseStarRocksParserListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseStarRocksParserListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseStarRocksParserListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseStarRocksParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseStarRocksParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgumentList is called when production namedArgumentList is entered.
func (s *BaseStarRocksParserListener) EnterNamedArgumentList(ctx *NamedArgumentListContext) {}

// ExitNamedArgumentList is called when production namedArgumentList is exited.
func (s *BaseStarRocksParserListener) ExitNamedArgumentList(ctx *NamedArgumentListContext) {}

// EnterNamedArguments is called when production namedArguments is entered.
func (s *BaseStarRocksParserListener) EnterNamedArguments(ctx *NamedArgumentsContext) {}

// ExitNamedArguments is called when production namedArguments is exited.
func (s *BaseStarRocksParserListener) ExitNamedArguments(ctx *NamedArgumentsContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseStarRocksParserListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseStarRocksParserListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterCrossOrInnerJoinType is called when production crossOrInnerJoinType is entered.
func (s *BaseStarRocksParserListener) EnterCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// ExitCrossOrInnerJoinType is called when production crossOrInnerJoinType is exited.
func (s *BaseStarRocksParserListener) ExitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// EnterOuterAndSemiJoinType is called when production outerAndSemiJoinType is entered.
func (s *BaseStarRocksParserListener) EnterOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// ExitOuterAndSemiJoinType is called when production outerAndSemiJoinType is exited.
func (s *BaseStarRocksParserListener) ExitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// EnterBracketHint is called when production bracketHint is entered.
func (s *BaseStarRocksParserListener) EnterBracketHint(ctx *BracketHintContext) {}

// ExitBracketHint is called when production bracketHint is exited.
func (s *BaseStarRocksParserListener) ExitBracketHint(ctx *BracketHintContext) {}

// EnterHintMap is called when production hintMap is entered.
func (s *BaseStarRocksParserListener) EnterHintMap(ctx *HintMapContext) {}

// ExitHintMap is called when production hintMap is exited.
func (s *BaseStarRocksParserListener) ExitHintMap(ctx *HintMapContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseStarRocksParserListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseStarRocksParserListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseStarRocksParserListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseStarRocksParserListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterPartitionNames is called when production partitionNames is entered.
func (s *BaseStarRocksParserListener) EnterPartitionNames(ctx *PartitionNamesContext) {}

// ExitPartitionNames is called when production partitionNames is exited.
func (s *BaseStarRocksParserListener) ExitPartitionNames(ctx *PartitionNamesContext) {}

// EnterKeyPartitionList is called when production keyPartitionList is entered.
func (s *BaseStarRocksParserListener) EnterKeyPartitionList(ctx *KeyPartitionListContext) {}

// ExitKeyPartitionList is called when production keyPartitionList is exited.
func (s *BaseStarRocksParserListener) ExitKeyPartitionList(ctx *KeyPartitionListContext) {}

// EnterTabletList is called when production tabletList is entered.
func (s *BaseStarRocksParserListener) EnterTabletList(ctx *TabletListContext) {}

// ExitTabletList is called when production tabletList is exited.
func (s *BaseStarRocksParserListener) ExitTabletList(ctx *TabletListContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *BaseStarRocksParserListener) EnterPrepareStatement(ctx *PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *BaseStarRocksParserListener) ExitPrepareStatement(ctx *PrepareStatementContext) {}

// EnterPrepareSql is called when production prepareSql is entered.
func (s *BaseStarRocksParserListener) EnterPrepareSql(ctx *PrepareSqlContext) {}

// ExitPrepareSql is called when production prepareSql is exited.
func (s *BaseStarRocksParserListener) ExitPrepareSql(ctx *PrepareSqlContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseStarRocksParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseStarRocksParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterDeallocateStatement is called when production deallocateStatement is entered.
func (s *BaseStarRocksParserListener) EnterDeallocateStatement(ctx *DeallocateStatementContext) {}

// ExitDeallocateStatement is called when production deallocateStatement is exited.
func (s *BaseStarRocksParserListener) ExitDeallocateStatement(ctx *DeallocateStatementContext) {}

// EnterReplicaList is called when production replicaList is entered.
func (s *BaseStarRocksParserListener) EnterReplicaList(ctx *ReplicaListContext) {}

// ExitReplicaList is called when production replicaList is exited.
func (s *BaseStarRocksParserListener) ExitReplicaList(ctx *ReplicaListContext) {}

// EnterExpressionsWithDefault is called when production expressionsWithDefault is entered.
func (s *BaseStarRocksParserListener) EnterExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {
}

// ExitExpressionsWithDefault is called when production expressionsWithDefault is exited.
func (s *BaseStarRocksParserListener) ExitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {
}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *BaseStarRocksParserListener) EnterExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *BaseStarRocksParserListener) ExitExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// EnterMapExpressionList is called when production mapExpressionList is entered.
func (s *BaseStarRocksParserListener) EnterMapExpressionList(ctx *MapExpressionListContext) {}

// ExitMapExpressionList is called when production mapExpressionList is exited.
func (s *BaseStarRocksParserListener) ExitMapExpressionList(ctx *MapExpressionListContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseStarRocksParserListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseStarRocksParserListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterExpressionSingleton is called when production expressionSingleton is entered.
func (s *BaseStarRocksParserListener) EnterExpressionSingleton(ctx *ExpressionSingletonContext) {}

// ExitExpressionSingleton is called when production expressionSingleton is exited.
func (s *BaseStarRocksParserListener) ExitExpressionSingleton(ctx *ExpressionSingletonContext) {}

// EnterExpressionDefault is called when production expressionDefault is entered.
func (s *BaseStarRocksParserListener) EnterExpressionDefault(ctx *ExpressionDefaultContext) {}

// ExitExpressionDefault is called when production expressionDefault is exited.
func (s *BaseStarRocksParserListener) ExitExpressionDefault(ctx *ExpressionDefaultContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseStarRocksParserListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseStarRocksParserListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseStarRocksParserListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseStarRocksParserListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseStarRocksParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseStarRocksParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseStarRocksParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseStarRocksParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBooleanExpressionDefault is called when production booleanExpressionDefault is entered.
func (s *BaseStarRocksParserListener) EnterBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// ExitBooleanExpressionDefault is called when production booleanExpressionDefault is exited.
func (s *BaseStarRocksParserListener) ExitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// EnterIsNull is called when production isNull is entered.
func (s *BaseStarRocksParserListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production isNull is exited.
func (s *BaseStarRocksParserListener) ExitIsNull(ctx *IsNullContext) {}

// EnterScalarSubquery is called when production scalarSubquery is entered.
func (s *BaseStarRocksParserListener) EnterScalarSubquery(ctx *ScalarSubqueryContext) {}

// ExitScalarSubquery is called when production scalarSubquery is exited.
func (s *BaseStarRocksParserListener) ExitScalarSubquery(ctx *ScalarSubqueryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseStarRocksParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseStarRocksParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterTupleInSubquery is called when production tupleInSubquery is entered.
func (s *BaseStarRocksParserListener) EnterTupleInSubquery(ctx *TupleInSubqueryContext) {}

// ExitTupleInSubquery is called when production tupleInSubquery is exited.
func (s *BaseStarRocksParserListener) ExitTupleInSubquery(ctx *TupleInSubqueryContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseStarRocksParserListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseStarRocksParserListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseStarRocksParserListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseStarRocksParserListener) ExitInList(ctx *InListContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseStarRocksParserListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseStarRocksParserListener) ExitBetween(ctx *BetweenContext) {}

// EnterLike is called when production like is entered.
func (s *BaseStarRocksParserListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseStarRocksParserListener) ExitLike(ctx *LikeContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseStarRocksParserListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {
}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseStarRocksParserListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {
}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseStarRocksParserListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseStarRocksParserListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseStarRocksParserListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseStarRocksParserListener) ExitDereference(ctx *DereferenceContext) {}

// EnterOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is entered.
func (s *BaseStarRocksParserListener) EnterOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// ExitOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is exited.
func (s *BaseStarRocksParserListener) ExitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// EnterMatchExpr is called when production matchExpr is entered.
func (s *BaseStarRocksParserListener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production matchExpr is exited.
func (s *BaseStarRocksParserListener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseStarRocksParserListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseStarRocksParserListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseStarRocksParserListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseStarRocksParserListener) ExitConvert(ctx *ConvertContext) {}

// EnterCollectionSubscript is called when production collectionSubscript is entered.
func (s *BaseStarRocksParserListener) EnterCollectionSubscript(ctx *CollectionSubscriptContext) {}

// ExitCollectionSubscript is called when production collectionSubscript is exited.
func (s *BaseStarRocksParserListener) ExitCollectionSubscript(ctx *CollectionSubscriptContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseStarRocksParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseStarRocksParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseStarRocksParserListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseStarRocksParserListener) ExitCast(ctx *CastContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseStarRocksParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseStarRocksParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterUserVariableExpression is called when production userVariableExpression is entered.
func (s *BaseStarRocksParserListener) EnterUserVariableExpression(ctx *UserVariableExpressionContext) {
}

// ExitUserVariableExpression is called when production userVariableExpression is exited.
func (s *BaseStarRocksParserListener) ExitUserVariableExpression(ctx *UserVariableExpressionContext) {
}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseStarRocksParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseStarRocksParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseStarRocksParserListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseStarRocksParserListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterArrowExpression is called when production arrowExpression is entered.
func (s *BaseStarRocksParserListener) EnterArrowExpression(ctx *ArrowExpressionContext) {}

// ExitArrowExpression is called when production arrowExpression is exited.
func (s *BaseStarRocksParserListener) ExitArrowExpression(ctx *ArrowExpressionContext) {}

// EnterSystemVariableExpression is called when production systemVariableExpression is entered.
func (s *BaseStarRocksParserListener) EnterSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// ExitSystemVariableExpression is called when production systemVariableExpression is exited.
func (s *BaseStarRocksParserListener) ExitSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// EnterConcat is called when production concat is entered.
func (s *BaseStarRocksParserListener) EnterConcat(ctx *ConcatContext) {}

// ExitConcat is called when production concat is exited.
func (s *BaseStarRocksParserListener) ExitConcat(ctx *ConcatContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseStarRocksParserListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseStarRocksParserListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterLambdaFunctionExpr is called when production lambdaFunctionExpr is entered.
func (s *BaseStarRocksParserListener) EnterLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// ExitLambdaFunctionExpr is called when production lambdaFunctionExpr is exited.
func (s *BaseStarRocksParserListener) ExitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// EnterDictionaryGetExpr is called when production dictionaryGetExpr is entered.
func (s *BaseStarRocksParserListener) EnterDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// ExitDictionaryGetExpr is called when production dictionaryGetExpr is exited.
func (s *BaseStarRocksParserListener) ExitDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseStarRocksParserListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseStarRocksParserListener) ExitCollate(ctx *CollateContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseStarRocksParserListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseStarRocksParserListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterMapConstructor is called when production mapConstructor is entered.
func (s *BaseStarRocksParserListener) EnterMapConstructor(ctx *MapConstructorContext) {}

// ExitMapConstructor is called when production mapConstructor is exited.
func (s *BaseStarRocksParserListener) ExitMapConstructor(ctx *MapConstructorContext) {}

// EnterArraySlice is called when production arraySlice is entered.
func (s *BaseStarRocksParserListener) EnterArraySlice(ctx *ArraySliceContext) {}

// ExitArraySlice is called when production arraySlice is exited.
func (s *BaseStarRocksParserListener) ExitArraySlice(ctx *ArraySliceContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseStarRocksParserListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseStarRocksParserListener) ExitExists(ctx *ExistsContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseStarRocksParserListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseStarRocksParserListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseStarRocksParserListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseStarRocksParserListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseStarRocksParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseStarRocksParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseStarRocksParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseStarRocksParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseStarRocksParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseStarRocksParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseStarRocksParserListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseStarRocksParserListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseStarRocksParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseStarRocksParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseStarRocksParserListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseStarRocksParserListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterUnitBoundaryLiteral is called when production unitBoundaryLiteral is entered.
func (s *BaseStarRocksParserListener) EnterUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// ExitUnitBoundaryLiteral is called when production unitBoundaryLiteral is exited.
func (s *BaseStarRocksParserListener) ExitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseStarRocksParserListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseStarRocksParserListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterParameter is called when production Parameter is entered.
func (s *BaseStarRocksParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production Parameter is exited.
func (s *BaseStarRocksParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseStarRocksParserListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseStarRocksParserListener) ExitExtract(ctx *ExtractContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseStarRocksParserListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseStarRocksParserListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterInformationFunction is called when production informationFunction is entered.
func (s *BaseStarRocksParserListener) EnterInformationFunction(ctx *InformationFunctionContext) {}

// ExitInformationFunction is called when production informationFunction is exited.
func (s *BaseStarRocksParserListener) ExitInformationFunction(ctx *InformationFunctionContext) {}

// EnterSpecialDateTime is called when production specialDateTime is entered.
func (s *BaseStarRocksParserListener) EnterSpecialDateTime(ctx *SpecialDateTimeContext) {}

// ExitSpecialDateTime is called when production specialDateTime is exited.
func (s *BaseStarRocksParserListener) ExitSpecialDateTime(ctx *SpecialDateTimeContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseStarRocksParserListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseStarRocksParserListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterAggregationFunctionCall is called when production aggregationFunctionCall is entered.
func (s *BaseStarRocksParserListener) EnterAggregationFunctionCall(ctx *AggregationFunctionCallContext) {
}

// ExitAggregationFunctionCall is called when production aggregationFunctionCall is exited.
func (s *BaseStarRocksParserListener) ExitAggregationFunctionCall(ctx *AggregationFunctionCallContext) {
}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseStarRocksParserListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseStarRocksParserListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterTranslateFunctionCall is called when production translateFunctionCall is entered.
func (s *BaseStarRocksParserListener) EnterTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// ExitTranslateFunctionCall is called when production translateFunctionCall is exited.
func (s *BaseStarRocksParserListener) ExitTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// EnterSimpleFunctionCall is called when production simpleFunctionCall is entered.
func (s *BaseStarRocksParserListener) EnterSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// ExitSimpleFunctionCall is called when production simpleFunctionCall is exited.
func (s *BaseStarRocksParserListener) ExitSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseStarRocksParserListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseStarRocksParserListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseStarRocksParserListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseStarRocksParserListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseStarRocksParserListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseStarRocksParserListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseStarRocksParserListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseStarRocksParserListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterInformationFunctionExpression is called when production informationFunctionExpression is entered.
func (s *BaseStarRocksParserListener) EnterInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// ExitInformationFunctionExpression is called when production informationFunctionExpression is exited.
func (s *BaseStarRocksParserListener) ExitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// EnterSpecialDateTimeExpression is called when production specialDateTimeExpression is entered.
func (s *BaseStarRocksParserListener) EnterSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// ExitSpecialDateTimeExpression is called when production specialDateTimeExpression is exited.
func (s *BaseStarRocksParserListener) ExitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// EnterSpecialFunctionExpression is called when production specialFunctionExpression is entered.
func (s *BaseStarRocksParserListener) EnterSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// ExitSpecialFunctionExpression is called when production specialFunctionExpression is exited.
func (s *BaseStarRocksParserListener) ExitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseStarRocksParserListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseStarRocksParserListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseStarRocksParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseStarRocksParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseStarRocksParserListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseStarRocksParserListener) ExitFilter(ctx *FilterContext) {}

// EnterOver is called when production over is entered.
func (s *BaseStarRocksParserListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseStarRocksParserListener) ExitOver(ctx *OverContext) {}

// EnterIgnoreNulls is called when production ignoreNulls is entered.
func (s *BaseStarRocksParserListener) EnterIgnoreNulls(ctx *IgnoreNullsContext) {}

// ExitIgnoreNulls is called when production ignoreNulls is exited.
func (s *BaseStarRocksParserListener) ExitIgnoreNulls(ctx *IgnoreNullsContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseStarRocksParserListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseStarRocksParserListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterUnboundedFrame is called when production unboundedFrame is entered.
func (s *BaseStarRocksParserListener) EnterUnboundedFrame(ctx *UnboundedFrameContext) {}

// ExitUnboundedFrame is called when production unboundedFrame is exited.
func (s *BaseStarRocksParserListener) ExitUnboundedFrame(ctx *UnboundedFrameContext) {}

// EnterCurrentRowBound is called when production currentRowBound is entered.
func (s *BaseStarRocksParserListener) EnterCurrentRowBound(ctx *CurrentRowBoundContext) {}

// ExitCurrentRowBound is called when production currentRowBound is exited.
func (s *BaseStarRocksParserListener) ExitCurrentRowBound(ctx *CurrentRowBoundContext) {}

// EnterBoundedFrame is called when production boundedFrame is entered.
func (s *BaseStarRocksParserListener) EnterBoundedFrame(ctx *BoundedFrameContext) {}

// ExitBoundedFrame is called when production boundedFrame is exited.
func (s *BaseStarRocksParserListener) ExitBoundedFrame(ctx *BoundedFrameContext) {}

// EnterBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is entered.
func (s *BaseStarRocksParserListener) EnterBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {
}

// ExitBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is exited.
func (s *BaseStarRocksParserListener) ExitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {
}

// EnterTableDesc is called when production tableDesc is entered.
func (s *BaseStarRocksParserListener) EnterTableDesc(ctx *TableDescContext) {}

// ExitTableDesc is called when production tableDesc is exited.
func (s *BaseStarRocksParserListener) ExitTableDesc(ctx *TableDescContext) {}

// EnterBackupRestoreTableDesc is called when production backupRestoreTableDesc is entered.
func (s *BaseStarRocksParserListener) EnterBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {
}

// ExitBackupRestoreTableDesc is called when production backupRestoreTableDesc is exited.
func (s *BaseStarRocksParserListener) ExitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {
}

// EnterExplainDesc is called when production explainDesc is entered.
func (s *BaseStarRocksParserListener) EnterExplainDesc(ctx *ExplainDescContext) {}

// ExitExplainDesc is called when production explainDesc is exited.
func (s *BaseStarRocksParserListener) ExitExplainDesc(ctx *ExplainDescContext) {}

// EnterOptimizerTrace is called when production optimizerTrace is entered.
func (s *BaseStarRocksParserListener) EnterOptimizerTrace(ctx *OptimizerTraceContext) {}

// ExitOptimizerTrace is called when production optimizerTrace is exited.
func (s *BaseStarRocksParserListener) ExitOptimizerTrace(ctx *OptimizerTraceContext) {}

// EnterPartitionExpr is called when production partitionExpr is entered.
func (s *BaseStarRocksParserListener) EnterPartitionExpr(ctx *PartitionExprContext) {}

// ExitPartitionExpr is called when production partitionExpr is exited.
func (s *BaseStarRocksParserListener) ExitPartitionExpr(ctx *PartitionExprContext) {}

// EnterPartitionDesc is called when production partitionDesc is entered.
func (s *BaseStarRocksParserListener) EnterPartitionDesc(ctx *PartitionDescContext) {}

// ExitPartitionDesc is called when production partitionDesc is exited.
func (s *BaseStarRocksParserListener) ExitPartitionDesc(ctx *PartitionDescContext) {}

// EnterListPartitionDesc is called when production listPartitionDesc is entered.
func (s *BaseStarRocksParserListener) EnterListPartitionDesc(ctx *ListPartitionDescContext) {}

// ExitListPartitionDesc is called when production listPartitionDesc is exited.
func (s *BaseStarRocksParserListener) ExitListPartitionDesc(ctx *ListPartitionDescContext) {}

// EnterSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is entered.
func (s *BaseStarRocksParserListener) EnterSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// ExitSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is exited.
func (s *BaseStarRocksParserListener) ExitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// EnterMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is entered.
func (s *BaseStarRocksParserListener) EnterMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// ExitMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is exited.
func (s *BaseStarRocksParserListener) ExitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// EnterMultiListPartitionValues is called when production multiListPartitionValues is entered.
func (s *BaseStarRocksParserListener) EnterMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// ExitMultiListPartitionValues is called when production multiListPartitionValues is exited.
func (s *BaseStarRocksParserListener) ExitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// EnterSingleListPartitionValues is called when production singleListPartitionValues is entered.
func (s *BaseStarRocksParserListener) EnterSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// ExitSingleListPartitionValues is called when production singleListPartitionValues is exited.
func (s *BaseStarRocksParserListener) ExitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// EnterListPartitionValues is called when production listPartitionValues is entered.
func (s *BaseStarRocksParserListener) EnterListPartitionValues(ctx *ListPartitionValuesContext) {}

// ExitListPartitionValues is called when production listPartitionValues is exited.
func (s *BaseStarRocksParserListener) ExitListPartitionValues(ctx *ListPartitionValuesContext) {}

// EnterListPartitionValue is called when production listPartitionValue is entered.
func (s *BaseStarRocksParserListener) EnterListPartitionValue(ctx *ListPartitionValueContext) {}

// ExitListPartitionValue is called when production listPartitionValue is exited.
func (s *BaseStarRocksParserListener) ExitListPartitionValue(ctx *ListPartitionValueContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseStarRocksParserListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseStarRocksParserListener) ExitStringList(ctx *StringListContext) {}

// EnterLiteralExpressionList is called when production literalExpressionList is entered.
func (s *BaseStarRocksParserListener) EnterLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// ExitLiteralExpressionList is called when production literalExpressionList is exited.
func (s *BaseStarRocksParserListener) ExitLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// EnterRangePartitionDesc is called when production rangePartitionDesc is entered.
func (s *BaseStarRocksParserListener) EnterRangePartitionDesc(ctx *RangePartitionDescContext) {}

// ExitRangePartitionDesc is called when production rangePartitionDesc is exited.
func (s *BaseStarRocksParserListener) ExitRangePartitionDesc(ctx *RangePartitionDescContext) {}

// EnterSingleRangePartition is called when production singleRangePartition is entered.
func (s *BaseStarRocksParserListener) EnterSingleRangePartition(ctx *SingleRangePartitionContext) {}

// ExitSingleRangePartition is called when production singleRangePartition is exited.
func (s *BaseStarRocksParserListener) ExitSingleRangePartition(ctx *SingleRangePartitionContext) {}

// EnterMultiRangePartition is called when production multiRangePartition is entered.
func (s *BaseStarRocksParserListener) EnterMultiRangePartition(ctx *MultiRangePartitionContext) {}

// ExitMultiRangePartition is called when production multiRangePartition is exited.
func (s *BaseStarRocksParserListener) ExitMultiRangePartition(ctx *MultiRangePartitionContext) {}

// EnterPartitionRangeDesc is called when production partitionRangeDesc is entered.
func (s *BaseStarRocksParserListener) EnterPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// ExitPartitionRangeDesc is called when production partitionRangeDesc is exited.
func (s *BaseStarRocksParserListener) ExitPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// EnterPartitionKeyDesc is called when production partitionKeyDesc is entered.
func (s *BaseStarRocksParserListener) EnterPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// ExitPartitionKeyDesc is called when production partitionKeyDesc is exited.
func (s *BaseStarRocksParserListener) ExitPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// EnterPartitionValueList is called when production partitionValueList is entered.
func (s *BaseStarRocksParserListener) EnterPartitionValueList(ctx *PartitionValueListContext) {}

// ExitPartitionValueList is called when production partitionValueList is exited.
func (s *BaseStarRocksParserListener) ExitPartitionValueList(ctx *PartitionValueListContext) {}

// EnterKeyPartition is called when production keyPartition is entered.
func (s *BaseStarRocksParserListener) EnterKeyPartition(ctx *KeyPartitionContext) {}

// ExitKeyPartition is called when production keyPartition is exited.
func (s *BaseStarRocksParserListener) ExitKeyPartition(ctx *KeyPartitionContext) {}

// EnterPartitionValue is called when production partitionValue is entered.
func (s *BaseStarRocksParserListener) EnterPartitionValue(ctx *PartitionValueContext) {}

// ExitPartitionValue is called when production partitionValue is exited.
func (s *BaseStarRocksParserListener) ExitPartitionValue(ctx *PartitionValueContext) {}

// EnterDistributionClause is called when production distributionClause is entered.
func (s *BaseStarRocksParserListener) EnterDistributionClause(ctx *DistributionClauseContext) {}

// ExitDistributionClause is called when production distributionClause is exited.
func (s *BaseStarRocksParserListener) ExitDistributionClause(ctx *DistributionClauseContext) {}

// EnterDistributionDesc is called when production distributionDesc is entered.
func (s *BaseStarRocksParserListener) EnterDistributionDesc(ctx *DistributionDescContext) {}

// ExitDistributionDesc is called when production distributionDesc is exited.
func (s *BaseStarRocksParserListener) ExitDistributionDesc(ctx *DistributionDescContext) {}

// EnterRefreshSchemeDesc is called when production refreshSchemeDesc is entered.
func (s *BaseStarRocksParserListener) EnterRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// ExitRefreshSchemeDesc is called when production refreshSchemeDesc is exited.
func (s *BaseStarRocksParserListener) ExitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// EnterStatusDesc is called when production statusDesc is entered.
func (s *BaseStarRocksParserListener) EnterStatusDesc(ctx *StatusDescContext) {}

// ExitStatusDesc is called when production statusDesc is exited.
func (s *BaseStarRocksParserListener) ExitStatusDesc(ctx *StatusDescContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseStarRocksParserListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseStarRocksParserListener) ExitProperties(ctx *PropertiesContext) {}

// EnterExtProperties is called when production extProperties is entered.
func (s *BaseStarRocksParserListener) EnterExtProperties(ctx *ExtPropertiesContext) {}

// ExitExtProperties is called when production extProperties is exited.
func (s *BaseStarRocksParserListener) ExitExtProperties(ctx *ExtPropertiesContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseStarRocksParserListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseStarRocksParserListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterUserPropertyList is called when production userPropertyList is entered.
func (s *BaseStarRocksParserListener) EnterUserPropertyList(ctx *UserPropertyListContext) {}

// ExitUserPropertyList is called when production userPropertyList is exited.
func (s *BaseStarRocksParserListener) ExitUserPropertyList(ctx *UserPropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseStarRocksParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseStarRocksParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterInlineProperties is called when production inlineProperties is entered.
func (s *BaseStarRocksParserListener) EnterInlineProperties(ctx *InlinePropertiesContext) {}

// ExitInlineProperties is called when production inlineProperties is exited.
func (s *BaseStarRocksParserListener) ExitInlineProperties(ctx *InlinePropertiesContext) {}

// EnterInlineProperty is called when production inlineProperty is entered.
func (s *BaseStarRocksParserListener) EnterInlineProperty(ctx *InlinePropertyContext) {}

// ExitInlineProperty is called when production inlineProperty is exited.
func (s *BaseStarRocksParserListener) ExitInlineProperty(ctx *InlinePropertyContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseStarRocksParserListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseStarRocksParserListener) ExitVarType(ctx *VarTypeContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseStarRocksParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseStarRocksParserListener) ExitComment(ctx *CommentContext) {}

// EnterOutfile is called when production outfile is entered.
func (s *BaseStarRocksParserListener) EnterOutfile(ctx *OutfileContext) {}

// ExitOutfile is called when production outfile is exited.
func (s *BaseStarRocksParserListener) ExitOutfile(ctx *OutfileContext) {}

// EnterFileFormat is called when production fileFormat is entered.
func (s *BaseStarRocksParserListener) EnterFileFormat(ctx *FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *BaseStarRocksParserListener) ExitFileFormat(ctx *FileFormatContext) {}

// EnterString is called when production string is entered.
func (s *BaseStarRocksParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseStarRocksParserListener) ExitString(ctx *StringContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseStarRocksParserListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseStarRocksParserListener) ExitBinary(ctx *BinaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseStarRocksParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseStarRocksParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseStarRocksParserListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseStarRocksParserListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseStarRocksParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseStarRocksParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterTaskInterval is called when production taskInterval is entered.
func (s *BaseStarRocksParserListener) EnterTaskInterval(ctx *TaskIntervalContext) {}

// ExitTaskInterval is called when production taskInterval is exited.
func (s *BaseStarRocksParserListener) ExitTaskInterval(ctx *TaskIntervalContext) {}

// EnterTaskUnitIdentifier is called when production taskUnitIdentifier is entered.
func (s *BaseStarRocksParserListener) EnterTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// ExitTaskUnitIdentifier is called when production taskUnitIdentifier is exited.
func (s *BaseStarRocksParserListener) ExitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// EnterUnitIdentifier is called when production unitIdentifier is entered.
func (s *BaseStarRocksParserListener) EnterUnitIdentifier(ctx *UnitIdentifierContext) {}

// ExitUnitIdentifier is called when production unitIdentifier is exited.
func (s *BaseStarRocksParserListener) ExitUnitIdentifier(ctx *UnitIdentifierContext) {}

// EnterUnitBoundary is called when production unitBoundary is entered.
func (s *BaseStarRocksParserListener) EnterUnitBoundary(ctx *UnitBoundaryContext) {}

// ExitUnitBoundary is called when production unitBoundary is exited.
func (s *BaseStarRocksParserListener) ExitUnitBoundary(ctx *UnitBoundaryContext) {}

// EnterType is called when production type is entered.
func (s *BaseStarRocksParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseStarRocksParserListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseStarRocksParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseStarRocksParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseStarRocksParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseStarRocksParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterSubfieldDesc is called when production subfieldDesc is entered.
func (s *BaseStarRocksParserListener) EnterSubfieldDesc(ctx *SubfieldDescContext) {}

// ExitSubfieldDesc is called when production subfieldDesc is exited.
func (s *BaseStarRocksParserListener) ExitSubfieldDesc(ctx *SubfieldDescContext) {}

// EnterSubfieldDescs is called when production subfieldDescs is entered.
func (s *BaseStarRocksParserListener) EnterSubfieldDescs(ctx *SubfieldDescsContext) {}

// ExitSubfieldDescs is called when production subfieldDescs is exited.
func (s *BaseStarRocksParserListener) ExitSubfieldDescs(ctx *SubfieldDescsContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseStarRocksParserListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseStarRocksParserListener) ExitStructType(ctx *StructTypeContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseStarRocksParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseStarRocksParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseStarRocksParserListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseStarRocksParserListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterDecimalType is called when production decimalType is entered.
func (s *BaseStarRocksParserListener) EnterDecimalType(ctx *DecimalTypeContext) {}

// ExitDecimalType is called when production decimalType is exited.
func (s *BaseStarRocksParserListener) ExitDecimalType(ctx *DecimalTypeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseStarRocksParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseStarRocksParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseStarRocksParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseStarRocksParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterWriteBranch is called when production writeBranch is entered.
func (s *BaseStarRocksParserListener) EnterWriteBranch(ctx *WriteBranchContext) {}

// ExitWriteBranch is called when production writeBranch is exited.
func (s *BaseStarRocksParserListener) ExitWriteBranch(ctx *WriteBranchContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseStarRocksParserListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseStarRocksParserListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseStarRocksParserListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseStarRocksParserListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseStarRocksParserListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseStarRocksParserListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterIdentifierWithAlias is called when production identifierWithAlias is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// ExitIdentifierWithAlias is called when production identifierWithAlias is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// EnterIdentifierWithAliasList is called when production identifierWithAliasList is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {
}

// ExitIdentifierWithAliasList is called when production identifierWithAliasList is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {
}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierOrString is called when production identifierOrString is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierOrString(ctx *IdentifierOrStringContext) {}

// ExitIdentifierOrString is called when production identifierOrString is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierOrString(ctx *IdentifierOrStringContext) {}

// EnterIdentifierOrStringList is called when production identifierOrStringList is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierOrStringList(ctx *IdentifierOrStringListContext) {
}

// ExitIdentifierOrStringList is called when production identifierOrStringList is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierOrStringList(ctx *IdentifierOrStringListContext) {
}

// EnterIdentifierOrStringOrStar is called when production identifierOrStringOrStar is entered.
func (s *BaseStarRocksParserListener) EnterIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// ExitIdentifierOrStringOrStar is called when production identifierOrStringOrStar is exited.
func (s *BaseStarRocksParserListener) ExitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// EnterUserWithoutHost is called when production userWithoutHost is entered.
func (s *BaseStarRocksParserListener) EnterUserWithoutHost(ctx *UserWithoutHostContext) {}

// ExitUserWithoutHost is called when production userWithoutHost is exited.
func (s *BaseStarRocksParserListener) ExitUserWithoutHost(ctx *UserWithoutHostContext) {}

// EnterUserWithHost is called when production userWithHost is entered.
func (s *BaseStarRocksParserListener) EnterUserWithHost(ctx *UserWithHostContext) {}

// ExitUserWithHost is called when production userWithHost is exited.
func (s *BaseStarRocksParserListener) ExitUserWithHost(ctx *UserWithHostContext) {}

// EnterUserWithHostAndBlanket is called when production userWithHostAndBlanket is entered.
func (s *BaseStarRocksParserListener) EnterUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {
}

// ExitUserWithHostAndBlanket is called when production userWithHostAndBlanket is exited.
func (s *BaseStarRocksParserListener) ExitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {
}

// EnterAssignment is called when production assignment is entered.
func (s *BaseStarRocksParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseStarRocksParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseStarRocksParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseStarRocksParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BaseStarRocksParserListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BaseStarRocksParserListener) ExitDecimalValue(ctx *DecimalValueContext) {}

// EnterDoubleValue is called when production doubleValue is entered.
func (s *BaseStarRocksParserListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production doubleValue is exited.
func (s *BaseStarRocksParserListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseStarRocksParserListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseStarRocksParserListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseStarRocksParserListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseStarRocksParserListener) ExitNonReserved(ctx *NonReservedContext) {}
