package infrastructure

import (
	"errors"
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"gorm.io/gorm"
	"time"
)

type HistoryRepositoryImpl struct {
	db gorm.DB
}

func NewHistoryRepositoryImpl(db gorm.DB) HistoryRepositoryImpl {
	return HistoryRepositoryImpl{db: db}
}

type historyRecord struct {
	gorm.Model
	BatchId uint
	Status  string
}

func (historyRecord) TableName() string {
	return "history"
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
	HistoryCreatedAt  time.Time
}

func (hr HistoryRepositoryImpl) FindByIdAndBatchId(historyId int, batchId int) (*entity.History, error) {
	record := batchAndHistoryRecord{}
	tx := hr.db.Debug().Table("batch").
		Select("batch.id as batch_id, batch.batch_name, batch.server_name, batch.cron_setting, batch.initial_date, batch.time_limit, batch.estimated_duration, batch.created_at as batch_created_at, history.id as history_id, history.status, history.created_at as history_created_at").
		Joins("join history on batch.id = history.batch_id").
		Where("history.id", historyId).
		Where("batch.id", batchId).
		Where("batch.deleted_at is null").
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
		return nil, errors.New("TODO")
	}

	return entity.NewHistory(&record.HistoryId, batch, record.Status, &record.HistoryCreatedAt)
}
