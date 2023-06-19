package controller

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/gin-gonic/gin"
)

type CalendarPresenter struct {
	context *gin.Context
}

func NewCalendarPresenter(c *gin.Context) CalendarPresenter {
	return CalendarPresenter{c}
}

func (cp CalendarPresenter) SendCalendarResponse(calendarItems []entity.CalendarItem) {
	if len(calendarItems) == 0 {
		cp.context.JSON(200, []interface{}{})
		return
	}

	var responseBody []map[string]interface{}
	for _, calendarItem := range calendarItems {
		responseBody = append(responseBody, map[string]interface{}{
			"history_id":      calendarItem.Id,
			"batch_id":        calendarItem.Batch.Id,
			"batch_name":      calendarItem.Batch.BatchName,
			"start_datetime":  calendarItem.StartDatetime,
			"finish_datetime": calendarItem.FinishDatetime,
			"status":          string(calendarItem.Status),
		})
	}

	cp.context.JSON(200, responseBody)
}

func (cp CalendarPresenter) SendInvalidRequestResponse(message string) {
	cp.context.JSON(400, map[string]interface{}{"status": 400, "message": message})
}

func (cp CalendarPresenter) SendInternalServerErrorResponse() {
	cp.context.JSON(500, map[string]interface{}{"status": 500, "message": "Internal Server Error"})
}
