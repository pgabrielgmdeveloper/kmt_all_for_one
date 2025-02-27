package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(dbUser, dbPassoword, dbName, dbHost, DbPort string, test bool) (*gorm.DB, error) {

	if test {
		return gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassoword, dbName, DbPort)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return gormDB, nil
}

func NewSQLDB(dbUser, dbPassoword, dbName, dbHost, DbPort string, test bool) (*sql.DB, error) {
	if test {
		return sql.Open("sqlite3", ":memory:")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassoword, dbName, DbPort)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return sqlDB, nil
}
