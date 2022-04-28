package data

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cch123/go-ddd/infra/config/static"
	"github.com/cch123/go-ddd/repo/data/ent"
	_ "github.com/go-sql-driver/mysql"
)

// this function should be in a provider set
func NewClient(conf static.Config) (*ent.Client, error) {

	c, err := sql.Open(dialect.MySQL, conf.MySQL.DSN)
	if err != nil {
		return nil, err
	}

	// Get the underlying sql.DB object of the driver.
	db := c.DB()
	db.SetMaxIdleConns(conf.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(conf.MySQL.MaxOpenConns)
	db.SetConnMaxLifetime(conf.MySQL.MaxConnLifeTime)

	return ent.NewClient(ent.Driver(c)), nil
}
