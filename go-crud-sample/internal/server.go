package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAccountsHandler(accountsRepository *AccountsRepository) *AccountsHandler {
	return &AccountsHandler{accountsRepository: accountsRepository}
}

type AccountsHandler struct {
	accountsRepository *AccountsRepository
}

func (a AccountsHandler) CreateAccountHandle(ctx *gin.Context) {
	var req Account

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = a.accountsRepository.Create(c, &req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (a AccountsHandler) GetAccountHandle(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	if accountID == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	res, err := a.accountsRepository.FindByID(ctx, accountID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
