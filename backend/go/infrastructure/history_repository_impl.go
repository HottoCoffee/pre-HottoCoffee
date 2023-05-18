package infrastructure

import (
	"errors"
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"gorm.io/gorm"
	"time"
)

type HistoryRepositoryImpl struct {
	db              gorm.DB
	batchRepository core.BatchRepository
}

func NewHistoryRepositoryImpl(db gorm.DB, batchRepository core.BatchRepository) HistoryRepositoryImpl {
	return HistoryRepositoryImpl{db: db, batchRepository: batchRepository}
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
	BatchDeletedAt    *time.Time
	HistoryId         *int
	Status            *string
	HistoryCreatedAt  *time.Time
}

func (hr HistoryRepositoryImpl) FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.History, error) {
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

	return entity.NewHistory(record.HistoryId, batch, *record.Status, batch.CronSetting.Prev(*record.HistoryCreatedAt), record.HistoryCreatedAt)
}

func (hr HistoryRepositoryImpl) FindAllByBatchId(batchId int) (*entity.Histories, error) {
	batch, err := hr.batchRepository.FindById(batchId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("no batch found for batch id: %v", batchId))
	}

	var records []historyRecord
	hr.db.Where("batch_id = ?", batchId).Find(&records)

	var simpleHistories []entity.SimpleHistory
	for _, record := range records {
		history, err := entity.NewSimpleHistory(
			int(record.ID),
			record.Status,
			record.CreatedAt,
		)
		if err != nil {
			return nil, errors.New("TODO")
		}
		simpleHistories = append(simpleHistories, *history)
	}

	histories := entity.NewHistories(*batch, simpleHistories)
	return &histories, nil
}

func (hr HistoryRepositoryImpl) FindAllDuring(start time.Time, endInclusive time.Time) ([]entity.Histories, error) {
	var records []batchAndHistoryRecord
	hr.db.Debug().Table("batch").
		Select("batch.id as batch_id, batch.batch_name, batch.server_name, batch.cron_setting, batch.initial_date, batch.time_limit, batch.estimated_duration, batch.created_at as batch_created_at, batch.deleted_at as batch_deleted_at, history.id as history_id, history.status, history.created_at as history_created_at").
		Joins("left join history on batch.id = history.batch_id").
		Where("batch.created_at >= ? AND batch.created_at <= ?", start, endInclusive).
		Where("batch.deleted_at is null OR batch.deleted_at > ?", start).
		Where("history.created_at >= ? AND history.created_at <= date_add(?, interval batch.time_limit minute)", start, endInclusive).
		Where("history.deleted_at is null").
		Find(&records)

	batchHistoriesMap := map[entity.Batch][]entity.SimpleHistory{}
	for _, record := range records {
		batch, err := entity.NewBatch(
			record.BatchId,
			record.BatchName,
			record.ServerName,
			record.CronSetting,
			record.TimeLimit,
			record.EstimatedDuration,
			record.BatchCreatedAt,
			record.BatchDeletedAt,
		)
		if err != nil {
			return nil, err
		}

		if !record.hasHistory() {
			batchHistoriesMap[*batch] = []entity.SimpleHistory{}
			continue
		}

		history, err := entity.NewSimpleHistory(
			*record.HistoryId,
			*record.Status,
			*record.HistoryCreatedAt,
		)
		if err != nil {
			return nil, err
		}
		batchHistoriesMap[*batch] = append(batchHistoriesMap[*batch], *history)
	}

	var histories []entity.Histories
	for batch, simpleHistories := range batchHistoriesMap {
		histories = append(histories, entity.NewHistories(batch, simpleHistories))
	}
	return histories, nil
}

func (r batchAndHistoryRecord) hasHistory() bool {
	return r.HistoryId != nil
}
