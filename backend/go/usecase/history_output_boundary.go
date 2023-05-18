package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type HistoryOutputBoundary interface {
	SendHistoryResponse(history entity.History)
	SendHistoryListResponse(histories entity.Histories)
	SendCalenderResponse(histories []entity.History)
	SendBadRequestResponse(msg string)
	SendNotFoundResponse()
	SendInternalServerErrorResponse()
}
