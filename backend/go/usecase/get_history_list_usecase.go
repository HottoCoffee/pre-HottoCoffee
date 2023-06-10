package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/core"
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
		ghlu.historyOutputBoundary.SendNotFoundResponse()
		return
	}

	ghlu.historyOutputBoundary.SendHistoryListResponse(*histories)
}
