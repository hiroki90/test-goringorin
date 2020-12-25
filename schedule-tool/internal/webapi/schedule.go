package webapi

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/infra"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSchedulesHandler(schedulesRepository *infra.SchedulesRepository) *SchedulesHandler {
	return &SchedulesHandler{schedulesRepository: schedulesRepository}
}

type SchedulesHandler struct {
	schedulesRepository *infra.SchedulesRepository
}

func (a SchedulesHandler) CreateScheduleHandle(ctx *gin.Context) {
	var req model.Schedule

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

	err = a.schedulesRepository.Create(ctx, &req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (a SchedulesHandler) GetScheduleHandle(ctx *gin.Context) {
	scheduleID := ctx.Param("scheduleID")
	if scheduleID == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	res, err := a.schedulesRepository.FindByID(ctx, scheduleID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (a SchedulesHandler) UpdateScheduleHandle(ctx *gin.Context) {
	var req model.Schedule

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

	err = a.schedulesRepository.Update(ctx, &req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (a SchedulesHandler) DeleteScheduleHandle(ctx *gin.Context) {
	scheduleID := ctx.Param("scheduleID")
	if scheduleID == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := a.schedulesRepository.Delete(ctx, scheduleID)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
