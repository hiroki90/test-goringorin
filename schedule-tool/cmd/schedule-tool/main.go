package main

import (
	"github.com/hiroki90/goringorin/schedule-tool/internal/infra"
	"github.com/hiroki90/goringorin/schedule-tool/internal/webapi"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hiroki90/goringorin/go-crud-sample/db"
	"github.com/hiroki90/goringorin/go-crud-sample/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("hoge hoge")
	conn, err := db.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Account
	accountsRepository := infra.NewAccountRepository(conn)
	accountsHandler := webapi.NewAccountsHandler(accountsRepository)

	e.POST("/accounts", accountsHandler.CreateAccountHandle)
	e.GET("/accounts/:accountID", accountsHandler.GetAccountHandle)
	e.PUT("/accounts", accountsHandler.UpdateAccountHandle)
	e.DELETE("/accounts/:accountID", accountsHandler.DeleteAccountHandle)

	_ = e.Run(":" + port)
}
