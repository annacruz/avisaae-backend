package "db"

import (
  "database/sql"
  _ "github.com/lib/pq"
  "models/incident"
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

func SelectAllIncidents() []Incident{} {
  rows, err := db.Query("SELECT * From incidents")

  if err != nil {
    return nil, err
  }

  defer rows.Close()
  returnedRows := []Incident{}
  for rows.Next(){
    inc := Incident{}
    err = rows.Scan()
    returnedRows = append(returnedRows, inc)
  }

  return returnedRows
}
