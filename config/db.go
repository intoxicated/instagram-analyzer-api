package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Redshift struct {
	*sqlx.DB
}

type DB struct {
	*sqlx.DB
}

func InitDB(c *Configuration) *DB {
	db := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
		c.DB.Charset,
		c.DB.Location,
	))

	db.SetMaxIdleConns(c.DB.MaxIdleConns)
	db.SetMaxOpenConns(c.DB.MaxOpenConns)
	// db.LogMode(c.DB.LogMode)

	return &DB{db}
}
