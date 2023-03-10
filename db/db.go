package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	sqlc "github.com/jjimgo/go_gin_mysql/db/sqlc"
)

func MigrateDataBase() *sqlc.Queries  {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/userDB?multiStatements=true&parseTime=true")

	if err != nil {
		log.Fatal("sql Open Error : ",err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal("driver instance : Error",err)
	}

	// same Code
	// m, err := migrate.New(
	// 	"file://db/migration",
	// 	"mysql://root:root@tcp(localhost:3306)/userDB?multiStatements=true",
	// )

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal("db instance Error : ",err)
	}

	// for delete database
	// if err := m.Down(); err != nil {
	// 	log.Fatal("miration Drop Failed : ", err)
	// }

	m.Steps(2)

	query := sqlc.New(db)
	return query
}