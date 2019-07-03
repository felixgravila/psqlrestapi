package dao

// Author @felixgravila

import (
	"database/sql"
	"fmt"
)

// ConnectionDetails is needed to initialise the DB instance
type ConnectionDetails struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// GetDB initialises and returns the DB
func GetDB(details ConnectionDetails) (db *sql.DB, err error) {
	db, err = initialiseDB(details)
	if err != nil {
		return
	}
	err = testDB(db)
	if err != nil {
		return
	}
	fmt.Println("Successfully connected!")
	err = initTables(db)
	fmt.Println("Tables written!")
	return
}

func initialiseDB(details ConnectionDetails) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		details.Host, details.Port, details.User, details.Dbname)
	if len(details.Password) > 0 {
		psqlInfo = psqlInfo + " password=" + details.Password
	}

	return sql.Open("postgres", psqlInfo)
}

func testDB(db *sql.DB) error {
	return db.Ping()
}

func initTables(db *sql.DB) (err error) {
	// Create author
	_, err = db.Exec(`create table if not exists authors (
                    id SERIAL PRIMARY KEY,
                    firstname VARCHAR(30),
                    lastname VARCHAR(30)
                 );
              `)
	if err != nil {
		return
	}

	// Create book
	_, err = db.Exec(`create table if not exists books (
                    id SERIAL,
                    author_id INTEGER,
                    isbn VARCHAR(30),
                    title VARCHAR(30),
                    FOREIGN KEY (author_id) REFERENCES authors (id)
                 );
              `)
	return
}
