package "db"

import (
  "database/sql"
  _ "github.com/lib/pq"
  "models/incident"
)

func ConnectDb() []Incident{}{  
  db, err := sql.Open("postgres", "user=postgres password=1234 dbname=incidentes")

  if err != nil {
    // Do something
  }

  rows, err := db.Query("SELECT * From incidents")

  if err != nil {
    // Do something
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
