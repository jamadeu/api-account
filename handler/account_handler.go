package handler

import "github.com/jamadeu/api-account/models"

type AccountHandler struct {
	ar models.AccountRepository
}

func NewAccountHandler(ar models.AccountRepository) *AccountHandler {
	return &AccountHandler{ar}
}
