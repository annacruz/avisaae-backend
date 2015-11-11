package main

import (
  "fmt"
  "log"
  "net/http"
  "flag"
)

var (
  port *int
)

func init(){
  port = flag.Int("p", 8888, "port")
  flag.Parse()
}

func ReturnData(){

}

func main() {
  http.HandleFunc('/api/incidents', ReturnIncidents)


  logar("Iniciando servidor na porta %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
