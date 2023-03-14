package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
)

type GetBatchUsecase struct {
	batchRepository core.BatchRepository
}

func NewGetBatchUsecase(br core.BatchRepository) GetBatchUsecase {
	return GetBatchUsecase{batchRepository: br}
}

func (gbu *GetBatchUsecase) Execute(id int) (*entity.Batch, error) {
	return gbu.batchRepository.FindById(id)
}
