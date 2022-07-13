package model_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNearbyStationWithinRange(t *testing.T) {
	name, err := vo.NewStationName("東京")
	if err != nil {
		t.Fatal(err)
	}
	lo, err := vo.NewLongitude(139.728079)
	if err != nil {
		t.Fatal(err)
	}
	la, err := vo.NewLatitude(35.601616)
	if err != nil {
		t.Fatal(err)
	}
	distance, err := vo.NewDistance(3)
	if err != nil {
		t.Fatal(err)
	}
	nearbyStation := model.NearbyStation{
		Name:      *name,
		Longitude: *lo,
		Latitude:  *la,
		Distance:  *distance,
	}
	t.Run("駅が範囲内である", func(t *testing.T) {
		rangeDistance, err := vo.NewDistance(5)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, nearbyStation.WithinRange(*rangeDistance))
	})
	t.Run("駅が範囲外である", func(t *testing.T) {
		rangeDistance, err := vo.NewDistance(1)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, nearbyStation.WithinRange(*rangeDistance))
	})

}
