package controller

import (
    "github.com/HottoCoffee/HottoCoffee/usecase"
)

type BatchController struct {
	getBatchUsecase usecase.GetBatchUsecase
}

func NewBatchController(gbu usecase.GetBatchUsecase) BatchController {
	return BatchController{getBatchUsecase: gbu}
}

func (bc *BatchController) GetBatch(input string) {
	bc.getBatchUsecase.Execute(input)
}
