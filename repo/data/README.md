# data

generate code from database schema

https://entgo.io/blog/2021/10/11/generating-ent-schemas-from-existing-sql-databases/

go run -mod=mod entgo.io/ent/cmd/ent init

go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:@tcp(localhost:3306)/shopping"


go generate ./ent


TODO

```go
package main

import (
    "time"

    "<your_project>/ent"
    "entgo.io/ent/dialect/sql"
)

func Open() (*ent.Client, error) {
    drv, err := sql.Open("mysql", "<mysql-dsn>")
    if err != nil {
        return nil, err
    }
    // Get the underlying sql.DB object of the driver.
    db := drv.DB()
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    return ent.NewClient(ent.Driver(drv)), nil
}
```

https://entgo.io/docs/sql-integration/
