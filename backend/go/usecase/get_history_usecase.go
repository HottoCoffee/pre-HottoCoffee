package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"strconv"
)

type GetHistoryUsecase struct {
	historyRepository     core.HistoryRepository
	historyOutputBoundary HistoryOutputBoundary
}

func NewGetHistoryUsecase(historyRepository core.HistoryRepository, historyOutputBoundary HistoryOutputBoundary) GetHistoryUsecase {
	return GetHistoryUsecase{
		historyRepository:     historyRepository,
		historyOutputBoundary: historyOutputBoundary,
	}
}

func (ghu GetHistoryUsecase) Execute(batchIdString string, historyIdString string) {
	batchId, err := strconv.Atoi(batchIdString)
	if err != nil {
		ghu.historyOutputBoundary.SendNotFoundResponse()
		return
	}

	historyId, err := strconv.Atoi(historyIdString)
	if err != nil {
		ghu.historyOutputBoundary.SendNotFoundResponse()
		return
	}

	history, err := ghu.historyRepository.FindByHistoryIdAndBatchId(historyId, batchId)
	if err != nil {
		if entity.IsDomainRuleViolationError(err) {
			fmt.Println(err.Error())
			ghu.historyOutputBoundary.SendInternalServerErrorResponse()
		} else {
			ghu.historyOutputBoundary.SendNotFoundResponse()
		}
		return
	}

	ghu.historyOutputBoundary.SendHistoryResponse(*history)
}
