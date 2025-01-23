package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
		"twitter_user",
		"twitter_password",
		"mysql",
		3306,
		"twitter_clone",
	)
	return sql.Open("mysql", dsn)
} 