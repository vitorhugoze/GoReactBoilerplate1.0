package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var db *sql.DB

const (
	Select QueryType = iota
	Insert
	Update
	Delete
)

type QueryType int

func InitDb() *sql.DB {

	var (
		name = os.Getenv("DB_NAME")
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		pass = os.Getenv("DB_PASS")
	)

	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, pass, name)

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal().Err(fmt.Errorf("error opening postgres connection %w", err))
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(fmt.Errorf("error opening postgres connection %w", err))
	}

	return db
}

func PgConnect() *sql.DB {
	return db
}

/*
	QueryType

0-Select 1-Insert 2-Update 3-Delete

Select is the only one that return rows
*/
func PgQuery(db *sql.DB, t QueryType, query string) (*sql.Rows, error) {

	if t == Select {

		rows, err := db.Query(query)

		if err != nil {
			return nil, fmt.Errorf("error running pg query %w", err)
		}

		return rows, err
	} else {

		if _, err := db.Exec(query); err != nil {
			log.Err(err)
			return nil, fmt.Errorf("error running pg query %w", err)
		}

		return nil, nil
	}

}
