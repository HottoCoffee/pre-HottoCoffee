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

func (hp HistoryPresenter) SendHistoryResponse(history entity.History) {
	hp.context.JSON(200, map[string]interface{}{
		"history_id":     history.Id,
		"batch_id":       history.Batch.Id,
		"batch_name":     history.Batch.BatchName,
		"start_datetime": history.ReportedDatetime,
		"status":         mapStatusToResponseField(history.Status),
	})
}

func (hp HistoryPresenter) SendHistoryListResponse(histories entity.Histories) {
	if len(histories.SimpleHistories) == 0 {
		hp.context.JSON(200, []interface{}{})
		return
	}

	var response []map[string]interface{}
	for _, history := range histories.SimpleHistories {
		response = append(response, map[string]interface{}{
			"history_id":     history.Id,
			"batch_id":       histories.Batch.Id,
			"batch_name":     histories.Batch.BatchName,
			"start_datetime": history.ReportedDatetime,
			"status":         mapStatusToResponseField(history.Status),
		})
	}

	hp.context.JSON(200, response)
}

func (hp HistoryPresenter) SendNotFoundResponse() {
	hp.context.JSON(404, map[string]interface{}{"status": 404, "message": "Not Found"})
}

func mapStatusToResponseField(s entity.Status) string {
	switch s {
	case entity.BeforeStarted:
		return "before_started"
	case entity.InProgress:
		return "in_progress"
	case entity.Success:
		return "success"
	case entity.Failed:
		return "failed"
	default:
		panic("cannot reach here")
	}
}
