package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSorce  = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// var testQueries *Queries

var (
	testQueries *Queries
	testDB      *sql.DB
)

//1st
// func TestMain(m *testing.M) {
// 	//connect to database
// 	conn, err := sql.Open(dbDriver, dbSorce)

// 	//error handling
// 	if err != nil {
// 		log.Fatal("Cannot connect to database", err)
// 	}

// 	testQueries = New(conn)
// 	m.Run()

// }

//2nd
func TestMain(m *testing.M) {
	//connect to database
	var err error
	testDB, err = sql.Open(dbDriver, dbSorce)

	//error handling
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	testQueries = New(testDB)
	m.Run()

}
