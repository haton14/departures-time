package repository_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
	mock_external "github.com/haton14/departures-time/departures-time-api/external/mock"
	"github.com/haton14/departures-time/departures-time-api/repository"
	"github.com/stretchr/testify/assert"
)

func TestDestinationGetByName(t *testing.T) {

	mockData := []external.ExspertDTO{
		{
			Station: external.ExspertStation{
				Code:        "22566",
				StationName: "大森(東京都)",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "139.731138",
				Latitude:  "35.585139",
			},
		},
		{
			Station: external.ExspertStation{
				Code:        "29668",
				StationName: "大森(静岡県)",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "137.5125",
				Latitude:  "34.734444",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	exspert := mock_external.NewMockExspert(ctrl)

	testsOK := map[string]struct {
		arg      vo.StationName
		expected []model.NearbyStation
	}{
		"[正常]期待通りのデータが取れる": {
			arg: "大森",
			expected: []model.NearbyStation{
				{Code: "22566", Name: "大森(東京都)", Longitude: 139.731138, Latitude: 35.585139},
				{Code: "29668", Name: "大森(静岡県)", Longitude: 137.5125, Latitude: 34.734444},
			},
		},
	}
	for name, tt := range testsOK {
		t.Run(name, func(t *testing.T) {
			exspert.
				EXPECT().
				GetByName(tt.arg).
				Return(mockData, nil)

			r := repository.NewDestination(exspert)
			actual, err := r.GetByName(tt.arg)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	testsNG := map[string]struct {
		mockData []external.ExspertDTO
	}{
		"[エラー]:NewStationNameでエラー": {
			mockData: []external.ExspertDTO{
				{
					Station: external.ExspertStation{
						Code:        "22566",
						StationName: "",
					},
					GeoPoint: external.ExspertGeoPoint{
						Longitude: "139.731138",
						Latitude:  "35.585139",
					},
				},
			},
		},
		"[エラー]:NewLongitudeでエラー": {
			mockData: []external.ExspertDTO{
				{
					Station: external.ExspertStation{
						Code:        "22566",
						StationName: "大森(東京都)",
					},
					GeoPoint: external.ExspertGeoPoint{
						Longitude: "181",
						Latitude:  "35.585139",
					},
				},
			},
		},
		"[エラー]:Longitudeに数値以外が入っている": {
			mockData: []external.ExspertDTO{
				{
					Station: external.ExspertStation{
						Code:        "22566",
						StationName: "大森(東京都)",
					},
					GeoPoint: external.ExspertGeoPoint{
						Longitude: "dummy",
						Latitude:  "35.585139",
					},
				},
			},
		},
		"[エラー]:NewLatitudeでエラー": {
			mockData: []external.ExspertDTO{
				{
					Station: external.ExspertStation{
						Code:        "22566",
						StationName: "大森(東京都)",
					},
					GeoPoint: external.ExspertGeoPoint{
						Longitude: "139.731138",
						Latitude:  "181",
					},
				},
			},
		},
		"[エラー]:Latitudeに数値以外が入っている": {
			mockData: []external.ExspertDTO{
				{
					Station: external.ExspertStation{
						Code:        "22566",
						StationName: "大森(東京都)",
					},
					GeoPoint: external.ExspertGeoPoint{
						Longitude: "139.731138",
						Latitude:  "dummy",
					},
				},
			},
		},
	}

	ngArg := model.NearbyStation{
		Name:      "大森",
		Longitude: 139.728079,
		Latitude:  35.588903,
	}

	for name, tt := range testsNG {
		t.Run(name, func(t *testing.T) {
			exspert.
				EXPECT().
				GetByName(ngArg.Name).
				Return(tt.mockData, nil)

			r := repository.NewNearbyStationDetail(exspert)
			actual, err := r.GetByNearbyStation(ngArg)
			assert.Nil(t, actual)
			assert.Error(t, err)
		})
	}

	t.Run("[エラー]:Exspert.GetByNameでエラー", func(t *testing.T) {
		exspert.
			EXPECT().
			GetByName(ngArg.Name).
			Return(nil, errors.New("other error"))
		r := repository.NewNearbyStationDetail(exspert)
		actual, err := r.GetByNearbyStation(ngArg)
		assert.Nil(t, actual)
		assert.Error(t, err)
	})
}
