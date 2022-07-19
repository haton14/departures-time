package external_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
	"github.com/stretchr/testify/assert"
)

func TestHeartRailsExpressGetNearbyStations(t *testing.T) {
	tests := map[string]struct {
		expected []external.NeaRestApiDTO
	}{
		"[正常]:本物の公開APIにアクセスして期待通りデータが取れる": {
			expected: []external.NeaRestApiDTO{
				{
					StationName: "大森",
					Location:    []float64{139.728079, 35.588903},
					Distance:    351,
				},
				{
					StationName: "大森海岸",
					Location:    []float64{139.73544, 35.587576},
					Distance:    879,
				},
				{
					StationName: "西大井",
					Location:    []float64{139.721729, 35.601616},
					Distance:    1186,
				},
				{
					StationName: "立会川",
					Location:    []float64{139.738886, 35.598489},
					Distance:    1265,
				},
				{
					StationName: "馬込",
					Location:    []float64{139.711772, 35.596435},
					Distance:    1499,
				},
				{
					StationName: "平和島",
					Location:    []float64{139.73491, 35.57868},
					Distance:    1632,
				},
				{
					StationName: "大井町",
					Location:    []float64{139.73485, 35.606257},
					Distance:    1723,
				},
				{
					StationName: "大井競馬場前",
					Location:    []float64{139.74708, 35.595006},
					Distance:    1812,
				},
				{
					StationName: "下神明",
					Location:    []float64{139.726256, 35.608704},
					Distance:    1861,
				},
				{
					StationName: "鮫洲",
					Location:    []float64{139.742227, 35.604977},
					Distance:    1971,
				},
				{
					StationName: "西馬込",
					Location:    []float64{139.705942, 35.586859},
					Distance:    2026,
				},
				{
					StationName: "中延",
					Location:    []float64{139.712526, 35.60571},
					Distance:    2035,
				},
				{
					StationName: "戸越公園",
					Location:    []float64{139.718159, 35.608856},
					Distance:    2054,
				},
				{
					StationName: "大森町",
					Location:    []float64{139.732027, 35.572431},
					Distance:    2219,
				},
				{
					StationName: "荏原町",
					Location:    []float64{139.707571, 35.60382},
					Distance:    2225,
				},
				{
					StationName: "流通センター",
					Location:    []float64{139.748964, 35.58186},
					Distance:    2255,
				},
				{
					StationName: "青物横丁",
					Location:    []float64{139.742958, 35.609351},
					Distance:    2389,
				},
				{
					StationName: "荏原中延",
					Location:    []float64{139.712196, 35.610056},
					Distance:    2435,
				},
				{
					StationName: "旗の台",
					Location:    []float64{139.702286, 35.604923},
					Distance:    2690,
				},
				{
					StationName: "品川シーサイド",
					Location:    []float64{139.749549, 35.608524},
					Distance:    2720,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			url := "https://station.ic731.net/api/nearest"
			e := external.NewNeaRestApi(url)
			lo, err := vo.NewLongitude(139.7274062)
			if err != nil {
				t.Fatal(err)
			}
			la, err := vo.NewLatitude(35.5920096)
			if err != nil {
				t.Fatal(err)
			}
			actual, err := e.GetNearbyStations(*lo, *la)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
