package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
)

type CreateBatchUsecase struct {
	batchRepository     core.BatchRepository
	batchOutputBoundary BatchOutputBoundary
}

func NewCreateBatchUsecase(bp core.BatchRepository, bob BatchOutputBoundary) CreateBatchUsecase {
	return CreateBatchUsecase{batchRepository: bp, batchOutputBoundary: bob}
}

func (cbu CreateBatchUsecase) Execute(input CreateBatchInput) {
	b, err := entity.NewBatch(0, input.BatchName, input.ServerName, input.CronSetting, input.TimeLimit, 0, input.InitialDate, nil)
	if err != nil {
		cbu.batchOutputBoundary.SendInvalidRequestResponse(err.Error())
		return
	}

	createdBatch, err := cbu.batchRepository.Save(*b)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		cbu.batchOutputBoundary.SendInternalServerErrorResponse()
		return
	}

	cbu.batchOutputBoundary.SendBatchResponse(*createdBatch)
}
