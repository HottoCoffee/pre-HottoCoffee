package controller

import (
	"github.com/HottoCoffee/HottoCoffee/usecase"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BatchController struct {
	getBatchUsecase usecase.GetBatchUsecase
	batchPresenter  *BatchPresenter
}

func NewBatchController(gbu usecase.GetBatchUsecase, bp *BatchPresenter) BatchController {
	return BatchController{getBatchUsecase: gbu, batchPresenter: bp}
}

func (bc *BatchController) GetBatch(input string) {
	id, err := strconv.Atoi(input)
	if err != nil {
		bc.batchPresenter.SendNotFoundResponse()
		return
	}
	b, err := bc.getBatchUsecase.Execute(id)
	if err != nil {
		bc.batchPresenter.SendNotFoundResponse()
	} else {
		bc.batchPresenter.SendBatchResponse(*b)
	}
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
