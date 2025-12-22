// Code generated from StarRocksParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package starrocks // StarRocksParser
import "github.com/antlr4-go/antlr/v4"

type BaseStarRocksParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStarRocksParserVisitor) VisitSqlStatements(ctx *SqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropDbStatement(ctx *DropDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDataStmt(ctx *ShowDataStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnDesc(ctx *ColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDefaultDesc(ctx *DefaultDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIndexDesc(ctx *IndexDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitEngineDesc(ctx *EngineDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCharsetDesc(ctx *CharsetDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCollateDesc(ctx *CollateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitKeyDesc(ctx *KeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOrderByDesc(ctx *OrderByDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnNullable(ctx *ColumnNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAggStateDesc(ctx *AggStateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAggDesc(ctx *AggDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRollupDesc(ctx *RollupDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRollupItem(ctx *RollupItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDupKeys(ctx *DupKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFromRollup(ctx *FromRollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTableStatement(ctx *ShowTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDescTableStatement(ctx *DescTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropViewStatement(ctx *DropViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTaskClause(ctx *TaskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitKillStatement(ctx *KillStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSyncStatement(ctx *SyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTypeDesc(ctx *TypeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLocationsDesc(ctx *LocationsDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDictionaryName(ctx *DictionaryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterClause(ctx *AlterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddBackendClause(ctx *AddBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropBackendClause(ctx *DropBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateImageClause(ctx *CreateImageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropIndexClause(ctx *DropIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableRenameClause(ctx *TableRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSwapTableClause(ctx *SwapTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOptimizeRange(ctx *OptimizeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOptimizeClause(ctx *OptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddColumnClause(ctx *AddColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropColumnClause(ctx *DropColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCompactionClause(ctx *CompactionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubfieldName(ctx *SubfieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNestedFieldName(ctx *NestedFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddFieldClause(ctx *AddFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropFieldClause(ctx *DropFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropBranchClause(ctx *DropBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropTagClause(ctx *DropTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableOperationClause(ctx *TableOperationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTagOptions(ctx *TagOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBranchOptions(ctx *BranchOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRefRetain(ctx *RefRetainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSnapshotId(ctx *SnapshotIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTimeUnit(ctx *TimeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInteger_list(ctx *Integer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSplitTabletClause(ctx *SplitTabletClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataSource(ctx *DataSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLoadProperties(ctx *LoadPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitImportColumns(ctx *ImportColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnProperties(ctx *ColumnPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitJobProperties(ctx *JobPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRegularColumns(ctx *RegularColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAllColumns(ctx *AllColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPredicateColumns(ctx *PredicateColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitHistogramStatement(ctx *HistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDisableBaselinePlanStatement(ctx *DisableBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitEnableBaselinePlanStatement(ctx *EnableBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitClassifier(ctx *ClassifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInlineFunction(ctx *InlineFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataDescList(ctx *DataDescListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataDesc(ctx *DataDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFormatProps(ctx *FormatPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBrokerDesc(ctx *BrokerDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitResourceDesc(ctx *ResourceDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowProcStatement(ctx *ShowProcStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitHelpStatement(ctx *HelpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowUserStatement(ctx *ShowUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnUser(ctx *GrantOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGrantOnAll(ctx *GrantOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivObjectName(ctx *PrivObjectNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBackupStatement(ctx *BackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRestoreStatement(ctx *RestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAddComputeNodeBlackListStatement(ctx *AddComputeNodeBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDelComputeNodeBlackListStatement(ctx *DelComputeNodeBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowComputeNodeBlackListStatement(ctx *ShowComputeNodeBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowExportStatement(ctx *ShowExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropFileStatement(ctx *DropFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetUserVar(ctx *SetUserVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetSystemVar(ctx *SetSystemVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIsolation_level(ctx *Isolation_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIsolation_types(ctx *Isolation_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLock_item(ctx *Lock_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTranslateStatement(ctx *TranslateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDialect(ctx *DialectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTranslateSQL(ctx *TranslateSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryRelation(ctx *QueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryPeriod(ctx *QueryPeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPeriodType(ctx *PeriodTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLimitConstExpr(ctx *LimitConstExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLimitElement(ctx *LimitElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFrom(ctx *FromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDual(ctx *DualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRollup(ctx *RollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCube(ctx *CubeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExcludeClause(ctx *ExcludeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRelations(ctx *RelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableAtom(ctx *TableAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFileTableFunction(ctx *FileTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPivotValue(ctx *PivotValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNamedArguments(ctx *NamedArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBracketHint(ctx *BracketHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitHintMap(ctx *HintMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionNames(ctx *PartitionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTabletList(ctx *TabletListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPrepareSql(ctx *PrepareSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitReplicaList(ctx *ReplicaListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMapExpressionList(ctx *MapExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMapExpression(ctx *MapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIsNull(ctx *IsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLike(ctx *LikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitConvert(ctx *ConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArrowExpression(ctx *ArrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitConcat(ctx *ConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMapConstructor(ctx *MapConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArraySlice(ctx *ArraySliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInformationFunction(ctx *InformationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWindowFunction(ctx *WindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableDesc(ctx *TableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExplainDesc(ctx *ExplainDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionExpr(ctx *PartitionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionDesc(ctx *PartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitListPartitionValue(ctx *ListPartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionValueList(ctx *PartitionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitKeyPartition(ctx *KeyPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPartitionValue(ctx *PartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDistributionClause(ctx *DistributionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDistributionDesc(ctx *DistributionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStatusDesc(ctx *StatusDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitExtProperties(ctx *ExtPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserPropertyList(ctx *UserPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInlineProperties(ctx *InlinePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInlineProperty(ctx *InlinePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitOutfile(ctx *OutfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitFileFormat(ctx *FileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTaskInterval(ctx *TaskIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnitBoundary(ctx *UnitBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubfieldDesc(ctx *SubfieldDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDecimalType(ctx *DecimalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitWriteBranch(ctx *WriteBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserWithHost(ctx *UserWithHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDecimalValue(ctx *DecimalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitDoubleValue(ctx *DoubleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksParserVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
