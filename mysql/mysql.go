package mysql

import (
	"database/sql"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/env"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const maxBatchSize = 2000 // ADJUST THIS IF GET TOO MANY PLACEHOLDERS IN PREPARED STATEMENT ERROR

func init() {
	log.Info("Initalising db")
	var err error
	db, err = sql.Open("mysql", getMySqlUri())
	if err != nil {
		log.Panicf("err opening mysql connection: %v", err)
	}
	db.SetMaxOpenConns(0) // unlimited max open connections
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(0) // unlimited max life time
}

// Return result rows from select query statement
func query[T any](
	fieldPtrs func(*T) []interface{},
	s string, args ...interface{},
) ([]*T, error) {
	rows, err := db.Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*T

	for rows.Next() {
		var t T
		err = rows.Scan(fieldPtrs(&t)...)
		if err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}

// Return result row from select query statement
func queryRow[T any](
	fieldPtrs func(*T) []interface{},
	s string, args ...interface{},
) (*T, error) {
	var t T
	err := db.QueryRow(s, args...).Scan(fieldPtrs(&t)...)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Exec mysql statement, eg insert
func exec(query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	return err
}

func getMySqlUri() string {
	uri := env.MYSQL_URI
	if len(uri) == 0 {
		uri = "root@tcp(db:3306)/"
	}
	return uri
}

// Create place holder i.e (?, ?, ... ?) for mysql statement.
func returnPlaceHolderString(args []any) string {
	if len(args) == 0 {
		return ""
	}
	var str []byte = make([]byte, len(args)*2+1)
	str[0] = '('
	for i := range args {
		str[i*2+1] = '?'
		str[i*2+2] = ','
	}
	str[len(str)-1] = ')'
	return string(str)
}
