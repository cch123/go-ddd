# data

generate code from database schema

https://entgo.io/blog/2021/10/11/generating-ent-schemas-from-existing-sql-databases/

go run -mod=mod entgo.io/ent/cmd/ent init

go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:@tcp(localhost:3306)/shopping"


go generate ./ent

