package db

import (
	"context"
	"database/sql"

	//mysqlインポート用
	_ "github.com/go-sql-driver/mysql"
	"github.com/myapp/noname/api/env"
)

// Conn DB構造体
type Conn struct {
	*sql.DB
}

// New DB初期化関数
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

// Tx トランザクション構造体
type Tx struct {
	*sql.Tx
}

// RunInTx トランザクション処理
func (c *Conn) RunInTx(ctx context.Context, f func(tx *Tx) error) error {
	tx, err := c.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := f(&Tx{tx}); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
