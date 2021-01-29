package main

import (
	"fmt"
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

	// web層
	accountsRepository := internal.NewAccountsRepository(conn)
	accountsHandler := internal.NewAccountsHandler(accountsRepository)

	e := gin.Default()

	e.POST("/accounts", accountsHandler.CreateAccountHandle)
	e.GET("/accounts/:accountID", accountsHandler.GetAccountHandle)
	e.PUT("/accounts", accountsHandler.UpdateAccountHandle)
	e.DELETE("/accounts/:accountID", accountsHandler.DeleteAccountHandle)

	_ = e.Run(":" + port)
}
