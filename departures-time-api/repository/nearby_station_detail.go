package repository

import (
	"fmt"
	"math"
	"strconv"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
)

type NearbyStationDetail interface {
	GetByNearbyStation(station model.NearbyStation) (*model.NearbyStation, error)
}

type nearbyStationDetail struct {
	exspert external.Exspert
}

func NewNearbyStationDetail(ext external.Exspert) NearbyStationDetail {
	return nearbyStationDetail{
		exspert: ext,
	}
}

func (s nearbyStationDetail) GetByNearbyStation(station model.NearbyStation) (*model.NearbyStation, error) {
	datas, err := s.exspert.GetByName(station.Name)
	if err != nil {
		return nil, fmt.Errorf("Exspert.GetByName() name %s: %s", station.Name.Value(), err)
	}
	var modelCandidate model.NearbyStation
	minDifference := math.MaxFloat64
	for _, d := range datas {
		m, err := s.toNearbyStations(d)
		if err != nil {
			return nil, err
		}
		// NeaRestApiから得た駅情報との座標差分を算出
		diff := m.Difference(station.Longitude, station.Latitude)
		if diff < minDifference {
			modelCandidate = *m
			minDifference = diff
		}
	}
	return &modelCandidate, nil
}

func (s nearbyStationDetail) toNearbyStations(data external.ExspertDTO) (*model.NearbyStation, error) {
	name, err := vo.NewStationName(data.Station.StationName)
	if err != nil {
		return nil, fmt.Errorf("NewStationName(): %s", err)
	}
	flo, err := strconv.ParseFloat(data.GeoPoint.Longitude, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat() Longitude %s: %s", data.GeoPoint.Longitude, err)
	}
	lo, err := vo.NewLongitude(flo)
	if err != nil {
		return nil, fmt.Errorf("NewLongitude(): %s", err)
	}
	fla, err := strconv.ParseFloat(data.GeoPoint.Latitude, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat() Latitude %s: %s", data.GeoPoint.Latitude, err)
	}
	la, err := vo.NewLatitude(fla)
	if err != nil {
		return nil, fmt.Errorf("NewLatitude(): %s", err)
	}
	return &model.NearbyStation{
		Code:      data.Station.Code,
		Name:      *name,
		Longitude: *lo,
		Latitude:  *la,
	}, nil
}
