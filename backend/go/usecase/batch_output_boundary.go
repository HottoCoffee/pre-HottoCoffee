package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type BatchOutputBoundary interface {
	SendBatchResponse(b entity.Batch)
	SendBatchListResponse(b []entity.Batch)
	SendNotFoundResponse()
	SendInternalServerErrorResponse()
}
