package data

import (
	"entgo.io/ent/dialect"
	"github.com/cch123/go-ddd/repo/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

// NewClient get an ent client
func NewClient(dsn string) (*ent.Client, error) {
	c, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, err
	}

	return c, nil
}
