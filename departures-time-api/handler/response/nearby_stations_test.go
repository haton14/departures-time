package response_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/stretchr/testify/assert"
)

func TestNewNearbyStationsGet(t *testing.T) {
	arg := []model.NearbyStation{
		{Code: "22567", Name: "大森海岸", Longitude: 139.73544, Latitude: 35.587576, Distance: 879},
		{Code: "22566", Name: "大森(東京都)", Longitude: 139.728079, Latitude: 35.588903, Distance: 351},
	}
	expected := response.NearbyStationsGet{
		Stations: []response.NearbyStation{
			{Code: "22566", Name: "大森(東京都)", Distance: 351},
			{Code: "22567", Name: "大森海岸", Distance: 879},
		},
	}
	actual := response.NewNearbyStationsGet(arg)
	assert.Equal(t, expected, actual)
}
