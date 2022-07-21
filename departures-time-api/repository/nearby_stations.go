package repository

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
)

type NearbyStations interface {
	GetByLongitudeAndLatitudeAndDistance(lo vo.Longitude, la vo.Latitude, distance vo.Distance) ([]model.NearbyStation, error)
}

type nearbyStations struct {
	neaRestApi external.NeaRestApi
}

func NewNearbyStations(ext external.NeaRestApi) NearbyStations {
	return nearbyStations{
		neaRestApi: ext,
	}
}

func (n nearbyStations) GetByLongitudeAndLatitudeAndDistance(lo vo.Longitude, la vo.Latitude, distance vo.Distance) ([]model.NearbyStation, error) {
	datas, err := n.neaRestApi.GetNearbyStations(lo, la)
	if err != nil {
		return nil, fmt.Errorf("NeaRestApi.GetNearbyStations() longitude %f latitude %f: %s", lo.Value(), la.Value(), err)
	}

	models := make([]model.NearbyStation, 0, len(datas))
	for _, d := range datas {
		m, err := n.toNearbyStations(d)
		if err != nil {
			return nil, err
		}
		// 現在地から指定した範囲内であるか
		if !m.WithinRange(distance) {
			continue
		}
		models = append(models, *m)
	}
	return models, nil
}

func (n nearbyStations) toNearbyStations(data external.NeaRestApiDTO) (*model.NearbyStation, error) {
	name, err := vo.NewStationName(data.StationName)
	if err != nil {
		return nil, fmt.Errorf("NewStationName(): %s", err)
	}
	lo, err := vo.NewLongitude(data.Location[0])
	if err != nil {
		return nil, fmt.Errorf("NewLongitude(): %s", err)
	}
	la, err := vo.NewLatitude(data.Location[1])
	if err != nil {
		return nil, fmt.Errorf("NewLatitude(): %s", err)
	}
	distance, err := vo.NewDistance(data.Distance)
	if err != nil {
		return nil, fmt.Errorf("NewDistance(): %s", err)
	}
	return &model.NearbyStation{
		Name:      *name,
		Longitude: *lo,
		Latitude:  *la,
		Distance:  *distance,
	}, nil
}
