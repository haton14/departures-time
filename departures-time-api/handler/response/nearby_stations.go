package response

import (
	"github.com/haton14/departures-time/departures-time-api/domain/model"
)

type NearbyStationsGet struct {
	Stations []NearbyStation `json:"stations"`
}

type NearbyStation struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Distance int    `json:"distance"`
}

func NewNearbyStationsGet(stations []model.NearbyStation) NearbyStationsGet {
	datas := make([]NearbyStation, 0, len(stations))
	for _, s := range stations {
		data := NearbyStation{
			Code:     s.Code.Value(),
			Name:     s.Name.Value(),
			Distance: s.Distance.Value(),
		}
		datas = append(datas, data)
	}
	return NearbyStationsGet{
		Stations: datas,
	}
}
