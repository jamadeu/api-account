package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/models"
)

type AccountHandler struct {
	ar models.AccountRepository
}

func NewAccountHandler(ar models.AccountRepository) *AccountHandler {
	return &AccountHandler{ar}
}

func (ah *AccountHandler) RegisterRoutes(router *gin.Engine, basePath string) {
	v1 := router.Group(basePath)
	{
		v1.GET("/account", ah.FindById)
	}
}

// @BasePath /api/v1

// @Summary Find account by id
// @Descriptaion find a account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id query string true "Account identification"
// @Success 200 {object} FindAccountByIdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /account [get]
func (h *AccountHandler) FindById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "id queryParameter is required")
		return
	}
	a, err := h.ar.FindById(id)
	if err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("account with id: %s not found", id))
		return
	}
	SendSuccess(ctx, http.StatusOK, "find-account-by-id", a)
}
