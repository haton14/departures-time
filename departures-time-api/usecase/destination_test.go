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

func TestDestinationGetByName(t *testing.T) {

	toName := func(v string) vo.StationName {
		sn, err := vo.NewStationName(v)
		if err != nil {
			t.Fatal(err)
		}
		return *sn
	}

	mockStations := []model.Destination{
		{Code: "22566", Name: "大森(東京都)", Longitude: 139.728079, Latitude: 35.588903},
	}

	expected := []model.Destination{
		{Code: "22566", Name: "大森(東京都)", Longitude: 139.728079, Latitude: 35.588903},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	destination := mock_repository.NewMockDestination(ctrl)

	t.Run("[正常]期待通りデータが取れること", func(t *testing.T) {
		destination.
			EXPECT().
			GetByName(toName("大森(東京都)")).
			Return(mockStations, nil)

		u := usecase.NewDistination(destination)
		actual, err := u.GetByName("大森(東京都)")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)

		expected := []model.Destination{
			{
				Code:      "22566",
				Name:      "大森(東京都)",
				Longitude: 139.728079,
				Latitude:  35.588903,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("[エラー]Destination.GetByNameでエラー", func(t *testing.T) {
		destination.
			EXPECT().
			GetByName(toName("大森(東京都)")).
			Return(nil, errors.New("other error"))

		u := usecase.NewDistination(destination)
		actual, err := u.GetByName("大森(東京都)")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
