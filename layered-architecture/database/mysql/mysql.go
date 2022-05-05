package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/mmmommm/go-sample/layered-architecture/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func ProvideMysqlClient(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DBEngine, dataSource(config.DBUser, config.DBPass, config.DBHost, config.DBName, config.DBPort))
	if err != nil {
		return nil, err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations",
		config.DBName,
		driver,
	)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}
	log.Printf("database running on %s", config.DBAddr())
	return db, nil
}

func dataSource(user, password, host, name string, port int) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
}