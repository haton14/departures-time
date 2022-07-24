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
			url := "http://api.ekispert.jp/v1/json"
			key := os.Getenv("EKISPERT_API_KEY")
			e := external.NewExspert(url, key)
			actual, err := e.GetByName(tt.arg)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	t.Run("[エラー]:本物の公開APIにアクセスしてデタラメなErrNotFoundで返ってくる", func(t *testing.T) {
		url := "http://api.ekispert.jp/v1/json"
		key := os.Getenv("EKISPERT_API_KEY")
		e := external.NewExspert(url, key)
		actual, err := e.GetByName("大森海岸XYZ")
		assert.ErrorIs(t, err, vo.ErrNotFound)
		assert.Nil(t, actual)
	})
}

func TestExspertGetRoutingURL(t *testing.T) {
	t.Run("[正常]:本物の公開APIにアクセスして期待通りの結果が返ってくる", func(t *testing.T) {
		url := "http://api.ekispert.jp/v1/json"
		key := os.Getenv("EKISPERT_API_KEY")
		e := external.NewExspert(url, key)
		actual, err := e.GetRoutingURL("22566", "22567")
		assert.NoError(t, err)
		expected := `https://roote.ekispert.net/result?arr=%E5%A4%A7%E6%A3%AE%E6%B5%B7%E5%B2%B8&arr_code=22567&connect=true&dep=%E5%A4%A7%E6%A3%AE(%E6%9D%B1%E4%BA%AC%E9%83%BD)&dep_code=22566&express=true&highway=true&hour&liner=true&local=true&minute&plane=true&shinkansen=true&ship=true&sleep=false&sort=time&surcharge=3&type=dep&via1=&via1_code=&via2=&via2_code=`
		assert.Equal(t, expected, actual)
	})

	t.Run("[エラー]:本物の公開APIにアクセスしてデタラメなリクエストを送る", func(t *testing.T) {
		url := "http://api.ekispert.jp/v1/json"
		key := os.Getenv("EKISPERT_API_KEY")
		e := external.NewExspert(url, key)
		actual, err := e.GetRoutingURL("22566", "22567dummy")
		assert.Equal(t, "", actual)
		assert.Error(t, err)
	})
}
