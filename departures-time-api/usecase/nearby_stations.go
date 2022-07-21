package usecase

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/repository"
)

type NearbyStations interface {
	GetByCoordinateAndDistance(lo vo.Longitude, la vo.Latitude, distance vo.Distance) ([]model.NearbyStation, error)
}

type nearbyStations struct {
	nearbyStations      repository.NearbyStations
	nearbyStationDetail repository.NearbyStationDetail
}

func NewNearbyStations(ns repository.NearbyStations, nsd repository.NearbyStationDetail) NearbyStations {
	return nearbyStations{
		nearbyStations:      ns,
		nearbyStationDetail: nsd,
	}
}

func (n nearbyStations) GetByCoordinateAndDistance(lo vo.Longitude, la vo.Latitude, distance vo.Distance) ([]model.NearbyStation, error) {
	staions, err := n.nearbyStations.GetByLongitudeAndLatitudeAndDistance(lo, la, distance)
	if err != nil {
		return nil, fmt.Errorf("NearbyStations.GetByLongitudeAndLatitudeAndDistance() lo %f la %f distance %d: %s", lo.Value(), la.Value(), distance.Value(), err)
	}
	for i, s := range staions {
		detail, err := n.nearbyStationDetail.GetByNearbyStation(s)
		if err != nil {
			return nil, fmt.Errorf("NearbyStationDetail.GetByNearbyStation(): %s", err)
		}
		s.Code = detail.Code
		s.Name = detail.Name
		staions[i] = s
	}
	return staions, nil
}
