package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 5
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute


func ConnectPostgres(dsn string) (*DB, error) {

	// connect to the postgres db using the driver and the dsn (data source name).
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil,err
	}

	dbConn.SQL = d
	return dbConn, nil

}

//test the DB connection and establish a new one if the connection is down.
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("***Pinged database successfully")
	return nil
}
