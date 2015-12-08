package main

import (
  "fmt"
  "log"
  "net/http"
  "flag"
  "encoding/json"
  "github.com/annacruz/avisaae-backend/db"
  "github.com/annacruz/avisaae-backend/models"
)

var (
  port *int
)

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
  incidentList, err := accessing.SelectAllIncidents

  json, err := json.Marshal(incidentList)
  if err != nil {
    w.WriteHeader)http.StatusInternalServerError)
    return
  }
  responseWithJson(w, string(json))

}

func main() {
  database.Initialize()
  defer database.Close()

  http.HandleFunc('/api/incidents', ReturnIncidents)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
