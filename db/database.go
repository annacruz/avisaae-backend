package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/annacruz/avisaae-backend/models"
  "fmt"
)

func Initialize() (*sql.DB, error){
  connstring := "user=postgres password=1234 dbname=avisaae"
  database, err := sql.Open("postgres", connstring)

  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  return database, nil
}

func SelectAllIncidents(db *sql.DB) (returnedRows []models.Incident, err error) {
  rows, err := db.Query("SELECT * From incidents")

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next(){
    inc := models.Incident{}
    err = rows.Scan(&inc.Id, &inc.Details, &inc.Type, &inc.DateTimeIncident, &inc.CreationDateTime)
    returnedRows = append(returnedRows, inc)
  }

  return returnedRows, nil
}
