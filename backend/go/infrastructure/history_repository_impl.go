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
	HistoryId         int
	Status            string
	StartDatetime     time.Time
	FinishDatetime    time.Time
}

func (hr HistoryRepositoryImpl) FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.BatchExecutionHistory, error) {
	record := batchAndHistoryRecord{}
	tx := hr.db.Table("batch").
		Select("batch.id as batch_id, batch.batch_name, batch.server_name, batch.cron_setting, batch.initial_date, batch.time_limit, batch.estimated_duration, batch.created_at as batch_created_at, history.id as history_id, history.status, history.start_datetime, history.finish_datetime").
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
		record.BatchCreatedAt,
		nil,
	)
	if err != nil {
		return nil, err
	}

	history, err := entity.NewHistory(record.HistoryId, record.Status, record.StartDatetime, record.FinishDatetime)
	if err != nil {
		return nil, err
	}

	batchExecutionHistory, err := entity.NewBatchExecutionHistory(*batch, *history)
	if err != nil {
		return nil, err
	}
	return batchExecutionHistory, nil
}
