package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	// Users
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	// Create tables

	return s.db, nil
}
