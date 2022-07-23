package usecase

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/repository"
)

type Destination interface {
	GetByName(name vo.StationName) ([]model.Destination, error)
}

type destination struct {
	destinationRepository repository.Destination
}

func NewDistination(d repository.Destination) Destination {
	return destination{destinationRepository: d}
}

func (d destination) GetByName(name vo.StationName) ([]model.Destination, error) {
	destinations, err := d.destinationRepository.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("Destination.GetByName(): %s", err)
	}
	return destinations, nil
}
