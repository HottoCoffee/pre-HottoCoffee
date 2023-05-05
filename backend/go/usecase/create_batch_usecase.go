package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
)

type CreateBatchUsecase struct {
	batchRepository     core.BatchRepository
	batchOutputBoundary BatchOutputBoundary
}

func NewCreateBatchUsecase(bp core.BatchRepository, bob BatchOutputBoundary) CreateBatchUsecase {
	return CreateBatchUsecase{batchRepository: bp, batchOutputBoundary: bob}
}

func (cbu CreateBatchUsecase) Execute(input BatchInput) {
	err := validateBatchInput(input)
	if err != nil {
		cbu.batchOutputBoundary.SendInvalidRequestResponse(err.Error())
		return
	}

	createdBatch, err := cbu.batchRepository.Create(input.BatchName, input.ServerName, input.CronSetting, input.TimeLimit, input.InitialDate)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		cbu.batchOutputBoundary.SendInternalServerErrorResponse()
		return
	}

	cbu.batchOutputBoundary.SendBatchResponse(*createdBatch)
}
