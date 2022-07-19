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

func TestNearbyStationsGetByLongitudeAndLatitudeAndDistance(t *testing.T) {
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
	toDistance := func(v int) vo.Distance {
		d, err := vo.NewDistance(v)
		if err != nil {
			t.Fatal(err)
		}
		return *d
	}
	mockData := []external.NeaRestApiDTO{
		{StationName: "大森", Location: []float64{139.728079, 35.588903}, Distance: 351},
		{StationName: "大森海岸", Location: []float64{39.73544, 35.587576}, Distance: 879},
		{StationName: "西大井", Location: []float64{39.721729, 35.601616}, Distance: 1186},
		{StationName: "立会川", Location: []float64{39.738886, 35.598489}, Distance: 1265},
		{StationName: "馬込", Location: []float64{39.711772, 35.596435}, Distance: 1499},
		{StationName: "平和島", Location: []float64{39.73491, 35.57868}, Distance: 1632},
		{StationName: "大井町", Location: []float64{39.73485, 35.606257}, Distance: 1723},
		{StationName: "大井競馬場前", Location: []float64{39.74708, 35.595006}, Distance: 1812},
		{StationName: "下神明", Location: []float64{39.726256, 35.608704}, Distance: 1861},
		{StationName: "鮫洲", Location: []float64{39.742227, 35.604977}, Distance: 1971},
		{StationName: "西馬込", Location: []float64{39.705942, 35.586859}, Distance: 2026},
		{StationName: "中延", Location: []float64{39.712526, 35.60571}, Distance: 2035},
		{StationName: "戸越公園", Location: []float64{39.718159, 35.608856}, Distance: 2054},
		{StationName: "大森町", Location: []float64{39.732027, 35.572431}, Distance: 2219},
		{StationName: "荏原町", Location: []float64{39.707571, 35.60382}, Distance: 2225},
		{StationName: "流通センター", Location: []float64{39.748964, 35.58186}, Distance: 2255},
		{StationName: "青物横丁", Location: []float64{39.742958, 35.609351}, Distance: 2389},
		{StationName: "荏原中延", Location: []float64{39.712196, 35.610056}, Distance: 2435},
		{StationName: "旗の台", Location: []float64{39.702286, 35.604923}, Distance: 2690},
		{StationName: "品川シーサイド", Location: []float64{39.749549, 35.608524}, Distance: 2720},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	neaRestApi := mock_external.NewMockNeaRestApi(ctrl)
	lo, err := vo.NewLongitude(139.7274062)
	if err != nil {
		t.Fatal(err)
	}
	la, err := vo.NewLatitude(35.5920096)
	if err != nil {
		t.Fatal(err)
	}

	testsOK := map[string]struct {
		arg      vo.Distance
		expected []model.NearbyStation
	}{
		"[正常]:期待通りデータが取れる 500m": {
			arg: toDistance(500),
			expected: []model.NearbyStation{
				{Name: toName("大森"), Longitude: toLo(139.728079), Latitude: toLa(35.588903), Distance: toDistance(351)},
			},
		},
		"[正常]:期待通りデータが取れる 2000m": {
			arg: toDistance(2000),
			expected: []model.NearbyStation{
				{Name: toName("大森"), Longitude: toLo(139.728079), Latitude: toLa(35.588903), Distance: toDistance(351)},
				{Name: toName("大森海岸"), Longitude: toLo(39.73544), Latitude: toLa(35.587576), Distance: toDistance(879)},
				{Name: toName("西大井"), Longitude: toLo(39.721729), Latitude: toLa(35.601616), Distance: toDistance(1186)},
				{Name: toName("立会川"), Longitude: toLo(39.738886), Latitude: toLa(35.598489), Distance: toDistance(1265)},
				{Name: toName("馬込"), Longitude: toLo(39.711772), Latitude: toLa(35.596435), Distance: toDistance(1499)},
				{Name: toName("平和島"), Longitude: toLo(39.73491), Latitude: toLa(35.57868), Distance: toDistance(1632)},
				{Name: toName("大井町"), Longitude: toLo(39.73485), Latitude: toLa(35.606257), Distance: toDistance(1723)},
				{Name: toName("大井競馬場前"), Longitude: toLo(39.74708), Latitude: toLa(35.595006), Distance: toDistance(1812)},
				{Name: toName("下神明"), Longitude: toLo(39.726256), Latitude: toLa(35.608704), Distance: toDistance(1861)},
				{Name: toName("鮫洲"), Longitude: toLo(39.742227), Latitude: toLa(35.604977), Distance: toDistance(1971)},
			},
		},
	}

	for name, tt := range testsOK {
		t.Run(name, func(t *testing.T) {
			neaRestApi.
				EXPECT().
				GetNearbyStations(*lo, *la).
				Return(mockData, nil)

			r := repository.NewNearbyStations(neaRestApi)
			actual, err := r.GetByLongitudeAndLatitudeAndDistance(*lo, *la, tt.arg)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	testsNG := map[string]struct {
		mockData []external.NeaRestApiDTO
	}{
		"[エラー]:NewStationNameでエラー": {
			mockData: []external.NeaRestApiDTO{{StationName: "", Location: []float64{139.728079, 35.588903}, Distance: 351}},
		},
		"[エラー]:NewLongitudeでエラー": {
			mockData: []external.NeaRestApiDTO{{StationName: "大森", Location: []float64{-181, 35.588903}, Distance: 351}},
		},
		"[エラー]:NewLatitudeでエラー": {
			mockData: []external.NeaRestApiDTO{{StationName: "大森", Location: []float64{139.728079, -181}, Distance: 351}},
		},
		"[エラー]:NewDistanceでエラー": {
			mockData: []external.NeaRestApiDTO{{StationName: "大森", Location: []float64{139.728079, 35.588903}, Distance: -1}},
		},
	}

	for name, tt := range testsNG {
		t.Run(name, func(t *testing.T) {
			neaRestApi.
				EXPECT().
				GetNearbyStations(*lo, *la).
				Return(tt.mockData, nil)

			r := repository.NewNearbyStations(neaRestApi)
			actual, err := r.GetByLongitudeAndLatitudeAndDistance(*lo, *la, toDistance(500))
			assert.Nil(t, actual)
			assert.Error(t, err)
		})
	}

	t.Run("[エラー]:NeaRestApi.GetNearbyStationsでエラー", func(t *testing.T) {
		neaRestApi.
			EXPECT().
			GetNearbyStations(*lo, *la).
			Return(nil, errors.New("other error"))
		r := repository.NewNearbyStations(neaRestApi)
		actual, err := r.GetByLongitudeAndLatitudeAndDistance(*lo, *la, toDistance(500))
		assert.Nil(t, actual)
		assert.Error(t, err)
	})
}
