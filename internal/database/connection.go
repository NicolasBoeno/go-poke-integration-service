package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(config *Config) (*PostgresDB, error) {
	db, err := sql.Open("postgres", config.GetDSN())

	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// config connections pool
	maxConns, _ := strconv.Atoi(config.MaxConns)
	maxIdleConns, _ := strconv.Atoi(config.MaxIdleConns)
	maxLifetime, _ := strconv.Atoi(config.MaxLifetime)

	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}

func (p *PostgresDB) GetDb() *sql.DB {
	return p.db
}
