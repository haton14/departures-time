package request_test

import (
	"math"
	"net/http"
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/stretchr/testify/assert"
)

func TestNewNearbyStationsGet(t *testing.T) {
	lo, err := vo.NewLongitude(139.7274062)
	if err != nil {
		t.Fatal(err)
	}
	la, err := vo.NewLatitude(35.5920096)
	if err != nil {
		t.Fatal(err)
	}
	distance500, err := vo.NewDistance(500)
	if err != nil {
		t.Fatal(err)
	}
	distanceMaxInt, err := vo.NewDistance(math.MaxInt)
	if err != nil {
		t.Fatal(err)
	}

	testsOK := map[string]struct {
		query    map[string]string
		expected *request.NearbyStationsGet
	}{
		"[正常]:経度139.7274062,緯度35.5920096,距離500": {
			query: map[string]string{
				"longitude": "139.7274062",
				"latitude":  "35.5920096",
				"distance":  "500",
			},
			expected: &request.NearbyStationsGet{
				Longitude: *lo,
				Latitude:  *la,
				Distance:  *distance500,
			},
		}, "[正常]:経度139.7274062,緯度35.5920096,距離指定なし": {
			query: map[string]string{
				"longitude": "139.7274062",
				"latitude":  "35.5920096",
			},
			expected: &request.NearbyStationsGet{
				Longitude: *lo,
				Latitude:  *la,
				Distance:  *distanceMaxInt,
			},
		},
	}

	for name, tt := range testsOK {
		t.Run(name, func(t *testing.T) {
			// setup
			c := testHelper.createTestContext(
				http.MethodGet,
				"/v1/nearby-stations",
				tt.query,
			)

			// check
			actual, err := request.NewNearbyStationsGet(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	testsNG := map[string]struct {
		query map[string]string
	}{
		"[エラー]:経度指定なし,緯度35.5920096,距離500": {
			query: map[string]string{
				"latitude": "35.5920096",
				"distance": "500",
			},
		},
		"[エラー]:経度指定139.7274062,緯度指定なし,距離500": {
			query: map[string]string{
				"longitude": "139.7274062",
				"distance":  "500",
			},
		},
		"[エラー]:経度指定181,緯度指定35.5920096,距離500": {
			query: map[string]string{
				"longitude": "181",
				"latitude":  "35.5920096",
				"distance":  "500",
			},
		},
		"[エラー]:経度指定139.7274062,緯度指定181,距離500": {
			query: map[string]string{
				"longitude": "139.7274062",
				"latitude":  "181",
				"distance":  "500",
			},
		},
		"[エラー]:経度指定139.7274062,緯度指定35.5920096,距離-1": {
			query: map[string]string{
				"longitude": "139.7274062",
				"latitude":  "35.5920096",
				"distance":  "-1",
			},
		},
	}

	for name, tt := range testsNG {
		t.Run(name, func(t *testing.T) {
			// setup
			c := testHelper.createTestContext(
				http.MethodGet,
				"/v1/nearby-stations",
				tt.query,
			)

			// check
			actual, err := request.NewNearbyStationsGet(c)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	}

	t.Run("【エラー】bindエラー", func(t *testing.T) {
		// setup
		c := testHelper.createTestContextBindError(
			http.MethodPatch,
			"/v3/settings/shops/reservation-codes",
			nil,
		)
		// check
		actual, err := request.NewNearbyStationsGet(c)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}