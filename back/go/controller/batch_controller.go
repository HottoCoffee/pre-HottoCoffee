package controller

import (
	"gorm.io/gorm"
	"time"
)

type BatchController struct {
	p  BatchPresenter
	db gorm.DB
}

func (bc BatchController) GetBatch(id int) {
	var batch Batch
	tx := bc.db.First(&batch, id)
	if tx.RowsAffected == 0 {
		// presenter.notFound
		return
	}

  bc.p.SendResponse((map[string]interface{})batch)
}

type Batch struct {
	gorm.Model
	BatchName         string
	ServerName        string
	CronSetting       string
	InitialDate       time.Time
	TimeLimit         int
	EstimatedDuration int
}
