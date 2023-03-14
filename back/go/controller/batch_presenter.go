package controller

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/gin-gonic/gin"
)

type BatchPresenter struct {
	c *gin.Context
}

func NewBatchPresenter(c *gin.Context) BatchPresenter {
	return BatchPresenter{c: c}
}

func (p *BatchPresenter) SendBatchResponse(b entity.Batch) {
	p.c.JSON(200, map[string]interface{}{
		"id":                 b.Id,
		"batch_name":         b.BatchName,
		"server_name":        b.ServerName,
		"cron_setting":       b.CronSetting,
		"initial_date":       b.StartDate,
		"time_limit":         b.TimeLimit,
		"estimated_duration": b.EsitimatedDuration,
	})
}

func (p *BatchPresenter) SendNotFoundResponse() {
	p.c.JSON(404, map[string]string{"status": "404", "message": "Not Found"})
}
