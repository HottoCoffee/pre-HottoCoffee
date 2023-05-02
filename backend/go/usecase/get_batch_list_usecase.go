package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
)

type GetBatchListUsecase struct {
	batchRepository     core.BatchRepository
	batchOutputBoundary BatchOutputBoundary
}

func NewGetBatchListUsecase(br core.BatchRepository, bob BatchOutputBoundary) GetBatchListUsecase {
	return GetBatchListUsecase{
		batchRepository:     br,
		batchOutputBoundary: bob,
	}
}

func (gblu GetBatchListUsecase) Execute(query string) {
	var batches []entity.Batch
	var err error
	if query == "" {
		batches, err = gblu.batchRepository.FindAll()
	} else {
		batches, err = gblu.batchRepository.FindFilteredBy(query)
	}
	if err != nil {
		_ = fmt.Errorf(err.Error())
		gblu.batchOutputBoundary.SendInternalServerErrorResponse()
		return
	}
	gblu.batchOutputBoundary.SendBatchListResponse(batches)
}
