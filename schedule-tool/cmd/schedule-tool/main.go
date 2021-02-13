package main

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/infra"
	"github.com/hiroki90/goringorin/schedule-tool/internal/service"
	"github.com/hiroki90/goringorin/schedule-tool/internal/webapi"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hiroki90/goringorin/schedule-tool/db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	conn, err := db.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	e := gin.Default()

	// Account
	//accountsRepository := infra.NewAccountRepository(conn)
	//accountsService := service.NewAccountService(accountsRepository)
	//accountsHandler := webapi.NewAccountsHandler(accountsService)
	// TODO: Service層を挟むので，ひとつずつ書き換え（メモ：膨大になったら別のとこにでもok）

	//e.POST("/accounts", accountsHandler.CreateAccountHandle)
	//e.GET("/accounts/:accountID", accountsHandler.GetAccountHandle)
	//e.PUT("/accounts", accountsHandler.UpdateAccountHandle)
	//e.DELETE("/accounts/:accountID", accountsHandler.DeleteAccountHandle)


	// Schedule:12/30
	schedulesRepository := infra.NewSchedulesRepository(conn)
	schedulesService := service.NewSchedulesService(schedulesRepository)
	schedulesHandler := webapi.NewSchedulesHandler(schedulesService)
	// 0212メモ: wireが便利，インジェクションを自動で定義してくれる
	// TODO: Service層を挟むので，ひとつずつ書き換え（メモ膨大になったら別のとこにでもok）

	e.GET("/schedules/:scheduleID", schedulesHandler.FindScheduleHandle)
	e.POST("/schedules", schedulesHandler.CreateScheduleHandle)
	//e.PUT("/schedules", schedulesHandler.UpdateScheduleHandle)
	//e.DELETE("/schedules/:scheduleID", schedulesHandler.DeleteScheduleHandle)

	// TODO: Event

	_ = e.Run(":" + port)
}
