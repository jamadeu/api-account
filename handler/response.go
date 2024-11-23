package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/models"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type FindAccountByIdResponse struct {
	Message string         `json:"message"`
	Data    models.Account `json:"data"`
}
