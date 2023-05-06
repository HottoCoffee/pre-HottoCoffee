package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type HistoryOutputBoundary interface {
	SendHistoryResponse(history entity.History)
	SendHistoryListResponse(histories entity.Histories)
	SendNotFoundResponse()
}
