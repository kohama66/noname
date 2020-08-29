package db

import (
	"database/sql"

	//mysqlインポート用
	_ "github.com/go-sql-driver/mysql"
	"github.com/myapp/noname/api/env"
)

type Conn struct {
	*sql.DB
}

func New() *Conn {
	db, err := sql.Open("mysql", env.SqlSource())
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)
	conn := &Conn{db}
	return conn
}
