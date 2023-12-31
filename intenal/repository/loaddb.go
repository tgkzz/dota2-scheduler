package repository

import (
	"database/sql"
	"dota2scheduler/config"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.DriverName, config.DB.DataSourceName)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	_, err = db.Exec("PRAGMA foreign_keys=ON;")
	if err != nil {
		return db, err
	}

	if err = createTable(db, config); err != nil {
		return db, err
	}

	return db, err
}

func createTable(db *sql.DB, config config.Config) error {
	fileSql, err := os.ReadFile(config.DB.SQL)
	if err != nil {
		return err
	}
	requests := strings.Split(string(fileSql), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			return err
		}
	}
	return err
}
