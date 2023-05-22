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
		"id":           b.Id,
		"batch_name":   b.BatchName,
		"server_name":  b.ServerName,
		"cron_setting": b.CronSetting.ToString(),
		"initial_date": b.StartDate,
		"time_limit":   b.TimeLimit,
	})
}

func (p *BatchPresenter) SendBatchListResponse(bs []entity.Batch) {
	var response []map[string]interface{}
	if len(bs) == 0 {
		response = []map[string]interface{}{}
	} else {
		for _, b := range bs {
			response = append(response, map[string]interface{}{
				"id":           b.Id,
				"batch_name":   b.BatchName,
				"server_name":  b.ServerName,
				"cron_setting": b.CronSetting.ToString(),
				"initial_date": b.StartDate,
				"time_limit":   b.TimeLimit,
			})
		}
	}

	p.context.JSON(200, response)
}

func (p *BatchPresenter) SendInvalidRequestResponse(message string) {
	p.context.JSON(400, map[string]interface{}{"status": 400, "message": message})
}

func (p *BatchPresenter) SendNotFoundResponse() {
	p.context.JSON(404, map[string]interface{}{"status": 404, "message": "Not Found"})
}

func (p *BatchPresenter) SendInternalServerErrorResponse() {
	p.context.JSON(500, map[string]interface{}{"status": 500, "message": "Internal Server Error"})
}
