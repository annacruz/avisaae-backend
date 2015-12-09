package main

import (
  "fmt"
  "log"
  "net/http"
  "flag"
  "encoding/json"
  "database/sql"
  "github.com/annacruz/avisaae-backend/db"
)

var (
  port *int
  database *sql.DB
  err error
)

type Headers map[string]string

func init(){
  port = flag.Int("p", 8888, "port")
  flag.Parse()
}

func responseWith(w http.ResponseWriter, status int, headers Headers){
  for k, v := range headers {
    w.Header().Set(k, v)
  }
  w.WriteHeader(status)
}

func responseWithJson(w http.ResponseWriter, response string){
  responseWith(w, http.StatusOK, Headers{"Content-Type": "application/json"})
  fmt.Fprintf(w, response)
}

func ReturnIncidents(w http.ResponseWriter, r *http.Request){
  incidentList, err := db.SelectAllIncidents(database)

  json, err := json.Marshal(incidentList)
  if err != nil {
    responseWith(w, http.StatusInternalServerError, nil)
    return
  }
  responseWithJson(w, string(json))

}

func main() {
  database, err = db.Initialize()
  defer database.Close()
  if err != nil {
    fmt.Printf("Um erro aconteceu!")
    return
  }

  http.HandleFunc("/api/incidents", ReturnIncidents)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
