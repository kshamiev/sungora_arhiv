package typ

import (
	"time"

	"sungora/lib/null"
)

type ReqSample struct {
	Id      string           `json:"id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"`
	Name    string           `json:"name"`
	Status  string           `json:"status"`
	Date    time.Time        `json:"date" example:"2006-01-02T15:04:05Z"`
	Flag    bool             `json:"flag"`
	Hobbit  null.StringArray `json:"hobbit"`
	Weight1 float32          `json:"weight1" example:"0.1"`
	Weight2 float64          `json:"weight2" example:"0.1"`
	Any     null.JSON        `json:"any" swaggertype:"string"`
	Number  int              `json:"number"`
}

type ResSample struct {
	Id      string           `json:"id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"`
	Name    string           `json:"name"`
	Status  string           `json:"status"`
	Date    time.Time        `json:"date" example:"2006-01-02T15:04:05Z"`
	Flag    bool             `json:"flag"`
	Hobbit  null.StringArray `json:"hobbit"`
	Weight1 float32          `json:"weight1" example:"0.1"`
	Weight2 float64          `json:"weight2" example:"0.1"`
	Any     null.JSON        `json:"any" swaggertype:"string"`
	Number  int              `json:"number"`
}
