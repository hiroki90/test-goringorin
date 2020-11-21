package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiroki90/goringorin/go-crud-sample/configs"
)


func NewDBConnection()(*sql.DB,error){
	cfg:=configs.NewConfig()
	return sql.Open("mysql",cfg.DBDial)
}
