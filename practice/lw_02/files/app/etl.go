package main

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func LoadToSQLite(csvFile, dbFile string) error {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// Удаляем таблицу если существует (чтобы не было конфликта)
	_, err = db.Exec(`DROP TABLE IF EXISTS clients`)
	if err != nil {
		return err
	}

	createTable := `
	CREATE TABLE clients (
		id INTEGER,
		age INTEGER,
		income REAL,
		debt REAL,
		credit_score REAL,
		delinquent INTEGER
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO clients VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i, row := range rows {
		if i == 0 {
			continue
		}

		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5])
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}