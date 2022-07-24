package external_test

import (
	"os"
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
	"github.com/stretchr/testify/assert"
)

func TestExspertGetByName(t *testing.T) {
	tests := map[string]struct {
		arg      vo.StationName
		expected []external.ExspertDTO
	}{
		"[正常]:本物の公開APIにアクセスして期待通りデータが複数取れる": {
			arg: "大森",
			expected: []external.ExspertDTO{
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
			},
		},
		"[正常]:本物の公開APIにアクセスして期待通りデータが1件取れる": {
			arg: "大森海岸",
			expected: []external.ExspertDTO{
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
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			url := "http://api.ekispert.jp/v1/json/station"
			key := os.Getenv("EKISPERT_API_KEY")
			e := external.NewExspert(url, key)
			actual, err := e.GetByName(tt.arg)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	t.Run("[エラー]:本物の公開APIにアクセスしてデタラメなErrNotFoundで返ってくる", func(t *testing.T) {
		url := "http://api.ekispert.jp/v1/json/station"
		key := os.Getenv("EKISPERT_API_KEY")
		e := external.NewExspert(url, key)
		actual, err := e.GetByName("大森海岸XYZ")
		assert.ErrorIs(t, err, vo.ErrNotFound)
		assert.Nil(t, actual)
	})
}
