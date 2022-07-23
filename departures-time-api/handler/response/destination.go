package response

import (
	"github.com/haton14/departures-time/departures-time-api/domain/model"
)

type DestinationGet struct {
	Stations []Destination `json:"stations"`
}

type Destination struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewDestinationGet(stations []model.Destination) DestinationGet {
	datas := make([]Destination, 0, len(stations))
	for _, s := range stations {
		data := Destination{
			Code: s.Code.Value(),
			Name: s.Name.Value(),
		}
		datas = append(datas, data)
	}
	return DestinationGet{
		Stations: datas,
	}
}
