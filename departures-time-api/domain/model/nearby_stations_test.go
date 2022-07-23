package model_test

import (
	"math"
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNearbyStationWithinRange(t *testing.T) {
	nearbyStation := model.NearbyStation{
		Name:      "東京",
		Longitude: 139.728079,
		Latitude:  35.601616,
		Distance:  3,
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

func TestNearbyStationDifference(t *testing.T) {
	nearbyStation := model.NearbyStation{
		Name:      "東京",
		Longitude: 139.728079,
		Latitude:  35.601616,
		Distance:  3,
	}
	xDiff := float64(139.728079) - float64(130.0)
	yDiff := float64(35.601616) - float64(36.0)

	assert.Equal(t, math.Sqrt(xDiff*xDiff+yDiff*yDiff), nearbyStation.Difference(130, 36))

}
