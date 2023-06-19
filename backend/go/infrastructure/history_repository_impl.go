package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"gorm.io/gorm"
)

type HistoryRepositoryImpl struct {
	db gorm.DB
}

func NewHistoryRepositoryImpl(db gorm.DB) HistoryRepositoryImpl {
	return HistoryRepositoryImpl{db: db}
}

type batchAndHistoryRecord struct {
	BatchId           int
	BatchName         string
	ServerName        string
	CronSetting       string
	InitialDate       time.Time
	TimeLimit         int
	EstimatedDuration int
	BatchCreatedAt    time.Time
	BatchDeletedAt    *time.Time
	HistoryId         *int
	Status            *string
	StartDatetime     *time.Time
	FinishDatetime    *time.Time
}

type HistoryRecord struct {
	gorm.Model
	BatchId        int
	Status         string
	StartDatetime  time.Time
	FinishDatetime time.Time
}

func (HistoryRecord) TableName() string {
	return "history"
}

func (hr HistoryRepositoryImpl) FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.BatchExecutionHistory, error) {
	record := batchAndHistoryRecord{}
	tx := hr.db.Table("batch").
		Select("batch.id as batch_id, batch.batch_name, batch.server_name, batch.cron_setting, batch.initial_date, batch.time_limit, batch.estimated_duration, batch.created_at as batch_created_at, batch.deleted_at as batch_deleted_at, history.id as history_id, history.status, history.start_datetime, history.finish_datetime").
		Joins("join history on batch.id = history.batch_id").
		Where("history.id", historyId).
		Where("batch.id", batchId).
		Where("history.deleted_at is null").
		Limit(1).
		Find(&record)
	if tx.RowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("no batch and history record for batch historyId: %v, history historyId: %v", batchId, historyId))
	}

	batch, err := entity.NewBatch(
		record.BatchId,
		record.BatchName,
		record.ServerName,
		record.CronSetting,
		record.TimeLimit,
		record.EstimatedDuration,
		record.InitialDate,
		nil,
	)
	if err != nil {
		return nil, err
	}

	history, err := entity.NewHistory(*record.HistoryId, *record.Status, *record.StartDatetime, *record.FinishDatetime)
	if err != nil {
		return nil, err
	}

	batchExecutionHistory, err := entity.NewBatchExecutionHistory(*batch, *history)
	if err != nil {
		return nil, err
	}
	return batchExecutionHistory, nil
}

func (hr HistoryRepositoryImpl) FindByBatchId(batchId int) (*entity.BatchExecutionHistories, error) {
	var b BatchRecord
	tx := hr.db.Find(&b, batchId)
	if tx.RowsAffected == 0 {
		return nil, errors.New("no record")
	}

	batch, err := mapRecordToBatch(b)
	if err != nil {
		return nil, err
	}

	var hrs []HistoryRecord
	hr.db.Where("batch_id = ?", batchId).Find(&hrs)

	var hs []entity.History
	for _, record := range hrs {
		h, err := entity.NewHistory(int(record.ID), record.Status, record.StartDatetime, record.FinishDatetime)
		if err != nil {
			return nil, err
		}
		hs = append(hs, *h)
	}

	beh := entity.NewBatchExecutionHistories(*batch, hs)
	return &beh, nil
}

func (hr HistoryRepositoryImpl) FindAllDuring(startDate time.Time, endDate time.Time) ([]entity.BatchExecutionHistories, error) {
	var records []batchAndHistoryRecord
	hr.db.Table("batch").
		Select("batch.id as batch_id, batch.batch_name, batch.server_name, batch.cron_setting, batch.initial_date, batch.time_limit, batch.estimated_duration, batch.created_at as batch_created_at, batch.deleted_at as batch_deleted_at, history.id as history_id, history.status, history.start_datetime, history.finish_datetime").
		Joins("left join history on batch.id = history.batch_id").
		Where("batch.initial_date < ?", endDate).
		Where("batch.deleted_at is null or batch.deleted_at >= ?", startDate).
		Where("history.start_datetime >= ? and history.start_datetime < ?", startDate, endDate).
		Where("history.deleted_at is null").
		Find(&records)

	batchHistoriesMap := map[entity.Batch][]entity.History{}
	for _, record := range records {
		batch, err := entity.NewBatch(
			record.BatchId,
			record.BatchName,
			record.ServerName,
			record.CronSetting,
			record.TimeLimit,
			record.EstimatedDuration,
			record.InitialDate,
			record.BatchDeletedAt,
		)
		if err != nil {
			return nil, err
		}

		if record.HistoryId == nil {
			batchHistoriesMap[*batch] = []entity.History{}
			continue
		}

		history, err := entity.NewHistory(*record.HistoryId, *record.Status, *record.StartDatetime, *record.FinishDatetime)
		if err != nil {
			return nil, err
		}

		batchHistoriesMap[*batch] = append(batchHistoriesMap[*batch], *history)
	}

	var batchExecutionHistories []entity.BatchExecutionHistories
	for batch, histories := range batchHistoriesMap {
		batchExecutionHistories = append(batchExecutionHistories, entity.NewBatchExecutionHistories(batch, histories))
	}

	return batchExecutionHistories, nil
}
