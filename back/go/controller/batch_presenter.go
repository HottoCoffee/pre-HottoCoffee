package controller

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/gin-gonic/gin"
)

type BatchPresenter struct {
	context *gin.Context
}

func NewBatchPresenter(c *gin.Context) BatchPresenter {
	return BatchPresenter{context: c}
}

func (p *BatchPresenter) SendBatchResponse(b entity.Batch) {
	p.context.JSON(200, map[string]interface{}{
		"id":                 b.Id,
		"batch_name":         b.BatchName,
		"server_name":        b.ServerName,
		"cron_setting":       b.CronSetting.ToString(),
		"initial_date":       b.StartDate,
		"time_limit":         b.TimeLimit,
		"estimated_duration": b.EstimatedDuration,
	})
}

func (p *BatchPresenter) SendNotFoundResponse() {
	p.context.JSON(404, map[string]string{"status": "404", "message": "Not Found"})
}
