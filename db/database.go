package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/annacruz/avisaae-backend/models"
  "fmt"
)

var db *sql.DB
var err error

func Initialize(){
  connstring := "user=postgres password=1234 dbname=avisaae"
  db, err := sql.Open("postgres", connstring)

  if err != nil {
    fmt.Println(err)
  }

}

func Close() {
  db.Close()
}

func SelectAllIncidents() (returnedRows []incident, err error) {
  rows, err := db.Query("SELECT * From incidents")
  defer rows.Close()

  if err != nil {
    return nil, err
  }

  for rows.Next(){
    inc := incident.GenerateList()
    err = rows.Scan()
    returnedRows = append(returnedRows, inc)
  }

  return returnedRows, nil
}
