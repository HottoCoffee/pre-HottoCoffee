package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"strconv"
)

type GetHistoryListUsecase struct {
	historyRepository     core.HistoryRepository
	historyOutputBoundary HistoryOutputBoundary
}

func NewGetHistoryListUsecase(historyRepository core.HistoryRepository, historyOutputBoundary HistoryOutputBoundary) GetHistoryListUsecase {
	return GetHistoryListUsecase{historyRepository: historyRepository, historyOutputBoundary: historyOutputBoundary}
}

func (ghlu GetHistoryListUsecase) Execute(batchIdString string) {
	batchId, err := strconv.Atoi(batchIdString)
	if err != nil {
		ghlu.historyOutputBoundary.SendNotFoundResponse()
		return
	}

	histories, err := ghlu.historyRepository.FindByBatchId(batchId)
	if err != nil {
		if entity.IsDomainRuleViolationError(err) {
			fmt.Println(err.Error())
			ghlu.historyOutputBoundary.SendInternalServerErrorResponse()
		} else {
			ghlu.historyOutputBoundary.SendNotFoundResponse()
		}
		return
	}

	ghlu.historyOutputBoundary.SendHistoryListResponse(*histories)
}
