package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/core"
	"strconv"
)

type GetBatchUsecase struct {
	batchRepository     core.BatchRepository
	batchOutputBoundary BatchOutputBoundary
}

func NewGetBatchUsecase(br core.BatchRepository, bob BatchOutputBoundary) GetBatchUsecase {
	return GetBatchUsecase{
		batchRepository:     br,
		batchOutputBoundary: bob,
	}
}

func (gbu GetBatchUsecase) Execute(input string) {
	id, err := strconv.Atoi(input)
	if err != nil {
		gbu.batchOutputBoundary.SendNotFoundResponse()
		return
	}
	b, err := gbu.batchRepository.FindById(id)
	if err != nil {
		gbu.batchOutputBoundary.SendNotFoundResponse()
	} else {
		gbu.batchOutputBoundary.SendBatchResponse(*b)
	}
}
