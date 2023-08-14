package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ConfigMySql struct {
	DB *sql.DB
}

func NewConfigMySql() *ConfigMySql {
	fmt.Println("Connecting to database...")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, password, database)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	runMigrate(db)

	fmt.Println("Connected to the database")

	return &ConfigMySql{
		DB: db,
	}
}

func runMigrate(db *sql.DB) {
	fmt.Println("Run migrations...")

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println(err)
	}

	m.Up()
}
