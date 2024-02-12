package storage

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

func MysqlStorageConfig() *sql.DB {

	Db, _ := sql.Open("mysql", "root:root@/address")

	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	return Db

}
