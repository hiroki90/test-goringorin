package internal

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAccountsHandler(accountsRepository *AccountsRepository) *AccountsHandler {
	return &AccountsHandler{accountsRepository: accountsRepository}
}

type AccountsHandler struct {
	accountsRepository *AccountsRepository
}

func (a AccountsHandler) CreateAccountHandle(ctx *gin.Context) {
	var req model.Account

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

    // NOTE: ozzo-validation を上手く go get できないため保留
	//err = req.Validate()
	//if err != nil {
	//	ctx.Status(http.StatusBadRequest)
	//	return
	//}

	err = a.accountsRepository.Create(ctx, &req)
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

func (a AccountsHandler) UpdateAccountHandle(ctx *gin.Context) {
	var req model.Account

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	//err = req.Validate()
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = a.accountsRepository.Update(ctx, &req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (a AccountsHandler) DeleteAccountHandle(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	if accountID == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := a.accountsRepository.Delete(ctx, accountID)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
