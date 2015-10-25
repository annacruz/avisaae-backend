package controllers

import (
    "github.com/revel/revel"
    "time"
    "trying-rest/app/models"
)

type Incidentes struct {
    *revel.Controller
}

var incidentes = []models.Incidente{
    models.Incidente{1, "La Trappe Quadrupel Oak Aged", "iluminacao", time.Now(), time.Now()},
    models.Incidente{2, "Cocoa Psycho", "iluminacao", time.Now(), time.Now()},
    models.Incidente{3, "American Dream", "iluminacao", time.Now(), time.Now()},
}

func (c Incidentes) List() revel.Result {
    return c.RenderJson(incidentes)
}

func (c Incidentes) Show(incidenteId int) revel.Result {
    var res models.Incidente

    for _, incidente := range incidentes {
        if incidente.ID == incidenteId {
            res = incidente
        }
    }

    if res.ID == 0 {
        return c.NotFound("Não foram achados incidentes")
    }

    return c.RenderJson(res)
}
