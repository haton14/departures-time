package repository

import (
	"fmt"
	"strconv"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
)

type Destination interface {
	GetByName(name vo.StationName) ([]model.Destination, error)
}

type destination struct {
	exspert external.Exspert
}

func NewDestination(ext external.Exspert) Destination {
	return destination{
		exspert: ext,
	}
}

func (s destination) GetByName(name vo.StationName) ([]model.Destination, error) {
	datas, err := s.exspert.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("Exspert.GetByName() name %s: %s", name.Value(), err)
	}
	models := make([]model.Destination, 0, len(datas))
	for _, d := range datas {
		m, err := s.toDestination(d)
		if err != nil {
			return nil, err
		}
		models = append(models, *m)
	}
	return models, nil
}

func (s destination) toDestination(data external.ExspertDTO) (*model.Destination, error) {
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
	return &model.Destination{
		Code:      data.Station.Code,
		Name:      *name,
		Longitude: *lo,
		Latitude:  *la,
	}, nil
}
