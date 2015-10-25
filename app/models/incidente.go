package models
import (
  "time"
)
type Incidente struct {
  ID int
  Detalhes string
  Tipo string
  Data_hora time.Time
  Data_criacao time.Time
}
