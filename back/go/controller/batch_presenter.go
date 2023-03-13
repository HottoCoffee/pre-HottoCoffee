package controller

import "github.com/gin-gonic/gin"

type BatchPresenter struct {
	c gin.Context
}

func (p BatchPresenter) SendResponse(responseBody map[string]interface{}) {
	p.c.JSON(200, responseBody)
}
