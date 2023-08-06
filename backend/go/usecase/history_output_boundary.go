package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type HistoryOutputBoundary interface {
	SendHistoryResponse(history entity.BatchExecutionHistory)
	SendHistoryListResponse(histories entity.BatchExecutionHistories)
	SendNotFoundResponse()
	SendInternalServerErrorResponse()
}
