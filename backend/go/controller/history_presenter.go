package controller

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/gin-gonic/gin"
)

type HistoryPresenter struct {
	context *gin.Context
}

func NewHistoryPresenter(context *gin.Context) HistoryPresenter {
	return HistoryPresenter{context: context}
}

func (hp HistoryPresenter) SendHistoryResponse(batchExecutionHistory entity.BatchExecutionHistory) {
	hp.context.JSON(200, map[string]interface{}{
		"history_id":      batchExecutionHistory.History.Id,
		"batch_id":        batchExecutionHistory.Batch.Id,
		"batch_name":      batchExecutionHistory.Batch.BatchName,
		"start_datetime":  batchExecutionHistory.History.StartDatetime,
		"finish_datetime": batchExecutionHistory.History.FinishDatetime,
		"status":          string(batchExecutionHistory.History.ExecutionResult),
	})
}

func (hp HistoryPresenter) SendHistoryListResponse(batchExecutionHistories entity.BatchExecutionHistories) {
	if len(batchExecutionHistories.Histories) == 0 {
		hp.context.JSON(200, []map[string]interface{}{})
		return
	}

	var responseBody []map[string]interface{}
	for _, history := range batchExecutionHistories.Histories {
		responseBody = append(responseBody, map[string]interface{}{
			"history_id":      history.Id,
			"batch_id":        batchExecutionHistories.Batch.Id,
			"batch_name":      batchExecutionHistories.Batch.BatchName,
			"start_datetime":  history.StartDatetime,
			"finish_datetime": history.FinishDatetime,
			"status":          string(history.ExecutionResult),
		})
	}
	hp.context.JSON(200, responseBody)
}

func (hp HistoryPresenter) SendNotFoundResponse() {
	hp.context.JSON(404, map[string]interface{}{"status": 404, "message": "Not Found"})
}

func (hp HistoryPresenter) SendInternalServerErrorResponse() {
	hp.context.JSON(500, map[string]interface{}{"status": 500, "message": "Internal Server Error"})
}
