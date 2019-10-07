package datastore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	ErrNoRows = sql.ErrNoRows
)

type Client struct {
	*sql.DB
}

func MustNew() *Client {
	db, err := sql.Open("mysql", "sqlboiler:sqlboiler@/sqlboiler?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(8)
	boil.SetDB(db)
	return &Client{db}
}
