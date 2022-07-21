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

func TestNearbyStationDetailGetByNearbyStation(t *testing.T) {
	toLo := func(v float64) vo.Longitude {
		l, err := vo.NewLongitude(v)
		if err != nil {
			t.Fatal(err)
		}
		return *l
	}
	toLa := func(v float64) vo.Latitude {
		l, err := vo.NewLatitude(v)
		if err != nil {
			t.Fatal(err)
		}
		return *l
	}
	toName := func(v string) vo.StationName {
		sn, err := vo.NewStationName(v)
		if err != nil {
			t.Fatal(err)
		}
		return *sn
	}
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
		{
			Station: external.ExspertStation{
				Code:        "24850",
				StationName: "大森・金城学院前",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "136.999444",
				Latitude:  "35.203611",
			},
		},
		{
			Station: external.ExspertStation{
				Code:        "22567",
				StationName: "大森海岸",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "139.738666",
				Latitude:  "35.584444",
			},
		},
		{
			Station: external.ExspertStation{
				Code:        "22218",
				StationName: "大森台",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "140.152639",
				Latitude:  "35.580889",
			},
		},
		{
			Station: external.ExspertStation{
				Code:        "22568",
				StationName: "大森町",
			},
			GeoPoint: external.ExspertGeoPoint{
				Longitude: "139.735305",
				Latitude:  "35.569333",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	exspert := mock_external.NewMockExspert(ctrl)

	testsOK := map[string]struct {
		arg      model.NearbyStation
		expected *model.NearbyStation
	}{
		"[正常]期待通りのデータが取れる": {
			arg: model.NearbyStation{
				Name:      toName("大森"),
				Longitude: toLo(139.728079),
				Latitude:  toLa(35.588903),
			},
			expected: &model.NearbyStation{
				Code:      "22566",
				Name:      toName("大森(東京都)"),
				Longitude: toLo(139.731138),
				Latitude:  toLa(35.585139),
			},
		},
	}
	for name, tt := range testsOK {
		t.Run(name, func(t *testing.T) {
			exspert.
				EXPECT().
				GetByName(tt.arg.Name).
				Return(mockData, nil)

			r := repository.NewNearbyStationDetail(exspert)
			actual, err := r.GetByNearbyStation(tt.arg)
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
		Name:      toName("大森"),
		Longitude: toLo(139.728079),
		Latitude:  toLa(35.588903),
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
