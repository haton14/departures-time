package response_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/stretchr/testify/assert"
)

func TestNewDestinationGet(t *testing.T) {
	arg := []model.Destination{
		{Code: "22566", Name: "大森(東京都)", Longitude: 139.728079, Latitude: 35.588903},
		{Code: "22567", Name: "大森海岸", Longitude: 139.73544, Latitude: 35.587576},
	}
	expected := response.DestinationGet{
		Stations: []response.Destination{
			{Code: "22566", Name: "大森(東京都)"},
			{Code: "22567", Name: "大森海岸"},
		},
	}
	actual := response.NewDestinationGet(arg)
	assert.Equal(t, expected, actual)
}
