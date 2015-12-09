package models

import(
  "time"
)

type Incident struct {
  Id int  `json:"id"`
  Details string `json:"details"`
  Type string `json:"type"`
  DateTimeIncident time.Time `json:"dateTimeIncident"`
  CreationDateTime time.Time `json:"creationDateTime"`
}
