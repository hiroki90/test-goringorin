package main

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/infra"
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
	accountsRepository := infra.NewAccountRepository(conn)
	accountsHandler := webapi.NewAccountsHandler(accountsRepository)

	e.POST("/accounts", accountsHandler.CreateAccountHandle)
	e.GET("/accounts/:accountID", accountsHandler.GetAccountHandle)
	e.PUT("/accounts", accountsHandler.UpdateAccountHandle)
	e.DELETE("/accounts/:accountID", accountsHandler.DeleteAccountHandle)

	// Schedule:12/30
	schedulesRepository := infra.NewScheduleRepository(conn)
	schedulesHandler := webapi.NewSchedulesHandler(schedulesRepository)

	e.POST("/schedules", schedulesHandler.CreateScheduleHandle)
	e.GET("/schedules/:scheduleID", schedulesHandler.GetScheduleHandle)
	e.PUT("/schedules", schedulesHandler.UpdateScheduleHandle)
	e.DELETE("/schedules/:scheduleID", schedulesHandler.DeleteScheduleHandle)

	_ = e.Run(":" + port)
}
