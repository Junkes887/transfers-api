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
	host := os.Getenv("DB_HOST")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

	fmt.Println("dataSourceName: " + dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println("Erro ping")
		panic(err)
	}

	err = db.Ping()

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
	fmt.Println("Running migrations...")

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
