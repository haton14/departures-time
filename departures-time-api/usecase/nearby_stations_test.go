package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/haton14/departures-time/departures-time-api/domain/model"
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	mock_repository "github.com/haton14/departures-time/departures-time-api/repository/mock"
	"github.com/haton14/departures-time/departures-time-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNearbyStationsGetByCoordinateAndDistance(t *testing.T) {
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

	mockStations := []model.NearbyStation{
		{Name: "大森", Longitude: 139.728079, Latitude: 35.588903, Distance: 351},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	nearbyStations := mock_repository.NewMockNearbyStations(ctrl)
	nearbyStationDetail := mock_repository.NewMockNearbyStationDetail(ctrl)

	t.Run("[正常]期待通りデータが取れること", func(t *testing.T) {
		f := nearbyStations.
			EXPECT().
			GetByLongitudeAndLatitudeAndDistance(toLo(139.728079), toLa(35.588903), toDistance(500)).
			Return(mockStations, nil)

		nearbyStationDetail.
			EXPECT().
			GetByNearbyStation(mockStations[0]).
			Return(&model.NearbyStation{
				Code:      "22566",
				Name:      toName("大森(東京都)"),
				Longitude: toLo(139.731138),
				Latitude:  toLa(35.585139),
			}, nil).
			After(f)

		u := usecase.NewNearbyStations(nearbyStations, nearbyStationDetail)
		actual, err := u.GetByCoordinateAndDistance(139.728079, 35.588903, 500)
		assert.NoError(t, err)

		expected := []model.NearbyStation{
			{
				Code:      "22566",
				Name:      "大森(東京都)",
				Longitude: 139.728079,
				Latitude: 35.588903,
				Distance:  351,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("[エラー]NearbyStations.GetByLongitudeAndLatitudeAndDistanceでエラー", func(t *testing.T) {
		nearbyStations.
			EXPECT().
			GetByLongitudeAndLatitudeAndDistance(toLo(139.728079), toLa(35.588903), toDistance(500)).
			Return(nil, errors.New("other error"))

		u := usecase.NewNearbyStations(nearbyStations, nearbyStationDetail)
		actual, err := u.GetByCoordinateAndDistance(139.728079, 35.588903, 500)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("[エラー]NearbyStationDetail.GetByNearbyStationでエラー", func(t *testing.T) {
		f := nearbyStations.
			EXPECT().
			GetByLongitudeAndLatitudeAndDistance(toLo(139.728079), toLa(35.588903), toDistance(500)).
			Return(mockStations, nil)

		nearbyStationDetail.
			EXPECT().
			GetByNearbyStation(mockStations[0]).
			Return(nil, errors.New("other error")).
			After(f)

		u := usecase.NewNearbyStations(nearbyStations, nearbyStationDetail)
		actual, err := u.GetByCoordinateAndDistance(139.728079, 35.588903, 500)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
