package db

import (

  "database/sql"
  "fmt"
  "log"

_ "github.com/mattn/go-sqlite3"
)

var ErrNoMatch = fmt.Errorf("match not found")
type Database struct {
  Conn *sql.DB
}

func Initialize(dbFilename string) (Database, error) {

  db := Database{}
  conn, err := sql.Open("sqlite3", dbFilename)
  if err != nil {
    return db, err
  } else {
    conn.Exec(initDB)
  }

  db.Conn = conn
  err = db.Conn.Ping()
  if err != nil {
    return db, err
  }
  log.Println("Database connected")
  return db, nil
}

