package webapi

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"github.com/hiroki90/goringorin/schedule-tool/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: web層のCreate,Find新しく作る
func NewSchedulesHandler(schedulesService *service.SchedulesService) *SchedulesHandler {
	return &SchedulesHandler{schedulesService: schedulesService}
}

type SchedulesHandler struct {
	schedulesService *service.SchedulesService
}

func (s SchedulesHandler) FindScheduleHandle(ctx *gin.Context) {
	scheduleID := ctx.Param("scheduleID")
	if scheduleID == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	res, err := s.schedulesService.Find(ctx, scheduleID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s SchedulesHandler) CreateScheduleHandle(ctx *gin.Context) {
	var req model.Schedule

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = s.schedulesService.Create(ctx, &req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

//func (s SchedulesHandler) GetSchedulesByAccountIDHandle(ctx *gin.Context) {
//	scheduleID := ctx.Param("scheduleID")
//	if scheduleID == "" {
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//
//	res, err := s.schedulesService.FindByAccountID(ctx, scheduleID)
//	if err != nil {
//		ctx.Status(http.StatusInternalServerError)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, res)
//}

//func (a SchedulesHandler) UpdateScheduleHandle(ctx *gin.Context) {
//	var req model.Schedule
//
//	err := ctx.ShouldBindJSON(&req)
//	if err != nil {
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//
//	//err = req.Validate()
//	if err != nil {
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//
//	err = a.schedulesService.Update(ctx, &req)
//	if err != nil {
//		ctx.Status(http.StatusInternalServerError)
//		return
//	}
//
//	ctx.Status(http.StatusOK)
//}

//func (a SchedulesHandler) DeleteScheduleHandle(ctx *gin.Context) {
//	scheduleID := ctx.Param("scheduleID")
//	if scheduleID == "" {
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//
//	err := a.schedulesService.Delete(ctx, scheduleID)
//
//	if err != nil {
//		ctx.Status(http.StatusInternalServerError)
//		return
////	}
//
//	ctx.Status(http.StatusOK)
//}
