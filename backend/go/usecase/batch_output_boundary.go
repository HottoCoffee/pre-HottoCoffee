package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type BatchOutputBoundary interface {
	SendBatchResponse(b entity.Batch)
	SendBatchesResponse(b []entity.Batch)
	SendNotFoundResponse()
	SendInternalServerErrorResponse()
}
